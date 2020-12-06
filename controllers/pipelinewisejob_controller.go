/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/go-logr/logr"
	"github.com/spf13/viper"
	batchv1 "k8s.io/api/batch/v1"
	kbatchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	kresource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ktypes "k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	batchv1alpha1 "github.com/dirathea/pipelinewise-operator/api/v1alpha1"
)

type ExternalResourceID string

const (
	ConfigMapExternalResourceID    ExternalResourceID = "config"
	VolumeExternalResourceID       ExternalResourceID = "volume"
	JobMapExternalResourceID       ExternalResourceID = "job"
	ConfigScriptExternalResourceID ExternalResourceID = "config-script"
	configModResourceName          string             = "pw-config-script"
	scriptFileName                 string             = "configuration-mod.sh"
)

// PipelinewiseJobReconciler reconciles a PipelinewiseJob object
type PipelinewiseJobReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=batch.pipelinewise,resources=pipelinewisejobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.pipelinewise,resources=pipelinewisejobs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch,resources=cronjobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;delete

func (r *PipelinewiseJobReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("pipelinewisejob", req.NamespacedName)

	// Load Pipelinewise Job
	var pipelinewiseJob batchv1alpha1.PipelinewiseJob
	if err := r.Get(ctx, req.NamespacedName, &pipelinewiseJob); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
		return ctrl.Result{}, err
	}

	// Setup Finalizer
	finalizerID := "pipelinewise"

	if pipelinewiseJob.DeletionTimestamp.IsZero() {
		if !containsString(pipelinewiseJob.ObjectMeta.Finalizers, finalizerID) {
			pipelinewiseJob.ObjectMeta.Finalizers = append(pipelinewiseJob.ObjectMeta.Finalizers, finalizerID)
			if err := r.Update(ctx, &pipelinewiseJob); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// Deletion flow
		if containsString(pipelinewiseJob.ObjectMeta.Finalizers, finalizerID) {
			if err := r.deleteExternalResources(&pipelinewiseJob); err != nil {
				return ctrl.Result{}, err
			}

			// remove our finalizer from the list and update it.
			pipelinewiseJob.ObjectMeta.Finalizers = removeString(pipelinewiseJob.ObjectMeta.Finalizers, finalizerID)
			if err := r.Update(ctx, &pipelinewiseJob); err != nil {
				return ctrl.Result{}, err
			}
		}

		// Stop reconciliation for deleted resource
		return ctrl.Result{}, nil
	}

	identifiers := resourcesIdentifier(&pipelinewiseJob)

	pwConfigScriptID := identifiers[ConfigScriptExternalResourceID]
	var pwConfigScript corev1.ConfigMap
	if err := r.Get(ctx, pwConfigScriptID, &pwConfigScript); err != nil {
		pwConfigScript = getConfigScript(pwConfigScriptID)
		err = r.Create(ctx, &pwConfigScript)
		if err != nil {
			log.Error(err, "Failed to create pipelinewise configuration")
			return ctrl.Result{}, err
		}
	}

	pwConfigID := identifiers[ConfigMapExternalResourceID]
	var pwConfig corev1.ConfigMap
	updatedPWConfig, err := r.getConfig(&pipelinewiseJob, pwConfigID)
	if err != nil {
		return ctrl.Result{}, err
	}
	if err := r.Get(ctx, pwConfigID, &pwConfig); err == nil {
		// Update the content from the CRD
		// Create new configMap
		pwConfig.Data = updatedPWConfig.Data
		err = r.Update(ctx, &pwConfig)
		if err != nil {
			log.Error(err, "Failed to update pipelinewise configuration")
			return ctrl.Result{}, err
		}
	} else {
		err = r.Create(ctx, &updatedPWConfig)
		if err != nil {
			log.Error(err, "Failed to create pipelinewise configuration")
			return ctrl.Result{}, err
		}
	}

	// Create PVC
	var pwVolume corev1.PersistentVolumeClaim
	volumeIdentifier := identifiers[VolumeExternalResourceID]
	if err := r.Get(ctx, volumeIdentifier, &pwVolume); err != nil {
		pwVolume := defaultVolume(volumeIdentifier)
		err = r.Create(ctx, &pwVolume)
		if err != nil {
			log.Error(err, "Failed to create PVC")
			return ctrl.Result{}, err
		}
	}

	// Create actual kubernetes job to run
	jobIdentifier := identifiers[JobMapExternalResourceID]
	var executorJob kbatchv1beta1.CronJob
	updatedExecutorJob := getExecutorJob(&pipelinewiseJob, jobIdentifier, pwConfig, pwConfigScript, pwVolume)
	if err := r.Get(ctx, jobIdentifier, &executorJob); err != nil {
		err = r.Create(ctx, &updatedExecutorJob)
		if err != nil {
			log.Error(err, "Failed to create executor Job")
			return ctrl.Result{}, err
		}
	} else {
		executorJob.Spec = updatedExecutorJob.Spec
		err = r.Update(ctx, &executorJob)
		if err != nil {
			log.Error(err, "Failed to update executor Job")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *PipelinewiseJobReconciler) deleteExternalResources(pipelinewiseJob *batchv1alpha1.PipelinewiseJob) error {
	//
	// delete any external resources associated with the cronJob
	//
	// Ensure that delete implementation is idempotent and safe to invoke
	// multiple types for same object.
	identifiers := resourcesIdentifier(pipelinewiseJob)
	deleteCtx := context.Background()
	var executorJob kbatchv1beta1.CronJob
	if err := r.Get(deleteCtx, identifiers[JobMapExternalResourceID], &executorJob); err == nil {
		// Found external resource job
		err := r.Delete(deleteCtx, &executorJob)
		if err != nil {
			return err
		}

		var jobList batchv1.JobList
		opts := []client.ListOption{
			client.InNamespace(executorJob.Namespace),
			client.MatchingFields{".metadata.ownerReferences[0].name": executorJob.Name},
		}
		err = r.List(deleteCtx, &jobList, opts...)
		if err != nil {
			return err
		}
		for _, pod := range jobList.Items {
			err = r.Delete(deleteCtx, &pod)
			if err != nil {
				return err
			}
		}
	}

	var volume corev1.PersistentVolumeClaim
	if err := r.Get(deleteCtx, identifiers[VolumeExternalResourceID], &volume); err == nil {
		// Found external resource volume
		err := r.Delete(deleteCtx, &volume)
		if err != nil {
			return err
		}
	}

	var config corev1.ConfigMap
	if err := r.Get(deleteCtx, identifiers[ConfigMapExternalResourceID], &config); err == nil {
		// Found external resource volume
		err := r.Delete(deleteCtx, &config)
		if err != nil {
			return err
		}
	}

	return nil
}

func getExecutorJob(pwJob *batchv1alpha1.PipelinewiseJob, identifier ktypes.NamespacedName, pwConfig, pwConfigScript corev1.ConfigMap, pwVolume corev1.PersistentVolumeClaim) kbatchv1beta1.CronJob {
	imageName := fmt.Sprintf("dirathea/pipelinewise:%v-%v-%v", viper.GetString("PIPELINEWISE_VERSION"), batchv1alpha1.GetTapConnectorID(pwJob), batchv1alpha1.GetTargetID(pwJob))
	if pwJob.Spec.Image != nil {
		imageName = *pwJob.Spec.Image
	}

	volumes := []corev1.Volume{
		corev1.Volume{
			Name: "pipelinewise-configuration",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: pwConfig.Name,
					},
				},
			},
		},
		corev1.Volume{
			Name: "runtime-volume",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: pwVolume.Name,
				},
			},
		},
	}

	volumeMounts := []corev1.VolumeMount{
		corev1.VolumeMount{
			Name:      "pipelinewise-configuration",
			MountPath: "/configurations",
		},
		corev1.VolumeMount{
			Name:      "runtime-volume",
			MountPath: "/root/.pipelinewise",
		},
	}

	importArgs := []string{
		"-c",
		"/app/entrypoint.sh import --dir /configurations",
	}

	runnerArgs := []string{
		"run_tap",
		"--tap",
		batchv1alpha1.GetTapID(pwJob),
		"--target",
		batchv1alpha1.GetTargetID(pwJob),
		"--extra_log",
	}

	if pwJob.Spec.Secret != nil {
		// Add secret as volume
		configModeDefaultMode := int32(0755)
		volumes = append(volumes,
			corev1.Volume{
				Name: "pw-master-password",
				VolumeSource: corev1.VolumeSource{
					Secret: &corev1.SecretVolumeSource{
						SecretName: pwJob.Spec.Secret.Name,
						Items: []corev1.KeyToPath{
							corev1.KeyToPath{
								Key:  pwJob.Spec.Secret.Key,
								Path: "master-password",
							},
						},
					},
				},
			},
			corev1.Volume{
				Name: "pw-config-mod",
				VolumeSource: corev1.VolumeSource{
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: pwConfigScript.Name,
						},
						DefaultMode: &configModeDefaultMode,
					},
				},
			})

		// Add secret volume mount
		volumeMounts = append(volumeMounts,
			corev1.VolumeMount{
				Name:      "pw-master-password",
				MountPath: "/secrets",
			},
			corev1.VolumeMount{
				Name:      "pw-config-mod",
				MountPath: "/pw-scripts",
			})

		// Append master token params
		importArgs[1] = fmt.Sprintf("/pw-scripts/%v /configurations /config-mod && /app/entrypoint.sh import --dir /config-mod --secret /secrets/master-password", scriptFileName)
	}

	return kbatchv1beta1.CronJob{
		ObjectMeta: identifierToMeta(identifier),
		Spec: kbatchv1beta1.CronJobSpec{
			Schedule: pwJob.Spec.Schedule,
			JobTemplate: kbatchv1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							RestartPolicy: corev1.RestartPolicyNever,
							InitContainers: []corev1.Container{
								corev1.Container{
									Name:  "import",
									Image: imageName,
									Args:  importArgs,
									Command: []string{
										"/bin/bash",
									},
									VolumeMounts: volumeMounts,
								},
							},
							Containers: []corev1.Container{
								corev1.Container{
									Name:         "runner",
									Image:        imageName,
									Args:         runnerArgs,
									VolumeMounts: volumeMounts,
								},
							},
							Volumes: volumes,
						},
					},
				},
			},
		},
	}
}

func (r *PipelinewiseJobReconciler) getConfig(pwJob *batchv1alpha1.PipelinewiseJob, identifier ktypes.NamespacedName) (corev1.ConfigMap, error) {
	pwConfig := corev1.ConfigMap{}
	// Create Pipelinewise Configuration via ConfigMap
	tapYaml, err := batchv1alpha1.ConstructTapConfiguration(pwJob)
	if err != nil {
		r.Log.Error(err, "Failed to construct tap configuration")
		return pwConfig, err
	}
	targetYaml, err := batchv1alpha1.ConstructTargetConfiguration(pwJob)
	if err != nil {
		r.Log.Error(err, "Failed to construct target configuration")
		return pwConfig, err
	}

	pwConfig.ObjectMeta = identifierToMeta(identifier)

	tapKeyName := fmt.Sprintf("tap_%v.yaml", batchv1alpha1.GetTapID(pwJob))
	targetKeyName := fmt.Sprintf("target_%v.yaml", batchv1alpha1.GetTargetID(pwJob))
	pwConfig.Data = map[string]string{
		tapKeyName:    string(tapYaml),
		targetKeyName: string(targetYaml),
	}
	return pwConfig, nil
}

func defaultVolume(identifier ktypes.NamespacedName) corev1.PersistentVolumeClaim {
	return corev1.PersistentVolumeClaim{
		ObjectMeta: identifierToMeta(identifier),
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: kresource.MustParse("1Gi"),
				},
			},
		},
	}
}

func getConfigScript(identifier ktypes.NamespacedName) corev1.ConfigMap {
	configBytes, _ := ioutil.ReadFile(scriptFileName)
	return corev1.ConfigMap{
		ObjectMeta: identifierToMeta(identifier),
		Data: map[string]string{
			scriptFileName: string(configBytes),
		},
	}
}

// Helper function to get external resources identifier
func resourcesIdentifierGenerator(pwJob *batchv1alpha1.PipelinewiseJob, prefix string) ktypes.NamespacedName {
	return ktypes.NamespacedName{
		Namespace: pwJob.Namespace,
		Name:      fmt.Sprintf("%v-%v", prefix, pwJob.Name),
	}
}

func resourcesIdentifier(pwJob *batchv1alpha1.PipelinewiseJob) map[ExternalResourceID]ktypes.NamespacedName {
	return map[ExternalResourceID]ktypes.NamespacedName{
		ConfigMapExternalResourceID: resourcesIdentifierGenerator(pwJob, "pw-config"),
		VolumeExternalResourceID:    resourcesIdentifierGenerator(pwJob, "pw-volume"),
		JobMapExternalResourceID:    resourcesIdentifierGenerator(pwJob, "pw-job"),
		ConfigScriptExternalResourceID: ktypes.NamespacedName{
			Name:      configModResourceName,
			Namespace: pwJob.Namespace,
		},
	}
}

func identifierToMeta(identifier ktypes.NamespacedName) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      identifier.Name,
		Namespace: identifier.Namespace,
	}
}

// Helper functions to check and remove string from a slice of strings.
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func (r *PipelinewiseJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1alpha1.PipelinewiseJob{}).
		Complete(r)
}
