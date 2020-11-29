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

	"github.com/go-logr/logr"
	"github.com/spf13/viper"
	batchv1 "k8s.io/api/batch/v1"
	kbatchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
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
	ConfigMapExternalResourceID ExternalResourceID = "config"
	VolumeExternalResourceID    ExternalResourceID = "volume"
	JobMapExternalResourceID    ExternalResourceID = "job"
)

// PipelinewiseJobReconciler reconciles a PipelinewiseJob object
type PipelinewiseJobReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=batch.pipelinewise,resources=pipelinewisejobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.pipelinewise,resources=pipelinewisejobs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch,resources=cronjobs,verbs=get;create;list;watch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;delete

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

	// Create Pipelinewise Configuration via ConfigMap
	tapYaml, err := batchv1alpha1.ConstructTapConfiguration(&pipelinewiseJob)
	if err != nil {
		log.Error(err, "Failed to construct tap configuration")
		return ctrl.Result{}, err
	}
	targetYaml, err := batchv1alpha1.ConstructTargetConfiguration(&pipelinewiseJob)
	if err != nil {
		log.Error(err, "Failed to construct target configuration")
		return ctrl.Result{}, err
	}

	identifierToMeta := func(identifier ktypes.NamespacedName) metav1.ObjectMeta {
		return metav1.ObjectMeta{
			Name:      identifier.Name,
			Namespace: identifier.Namespace,
		}
	}

	configIdentifier := identifiers[ConfigMapExternalResourceID]
	var pipelinewiseConfigurationConfigMap corev1.ConfigMap
	if err := r.Get(ctx, configIdentifier, &pipelinewiseConfigurationConfigMap); err == nil {
		// Update the content from the CRD
		// Create new configMap
		tapKeyName := fmt.Sprintf("tap_%v.yaml", batchv1alpha1.GetTapID(&pipelinewiseJob))
		pipelinewiseConfigurationConfigMap.Data[tapKeyName] = string(tapYaml)
		targetKeyName := fmt.Sprintf("target_%v.yaml", batchv1alpha1.GetTargetID(&pipelinewiseJob))
		pipelinewiseConfigurationConfigMap.Data[targetKeyName] = string(targetYaml)
		err = r.Update(ctx, &pipelinewiseConfigurationConfigMap)
		if err != nil {
			log.Error(err, "Failed to update pipelinewise configuration")
			return ctrl.Result{}, err
		}
	} else {
		pipelinewiseConfigurationConfigMap = corev1.ConfigMap{}
		pipelinewiseConfigurationConfigMap.Namespace = configIdentifier.Namespace
		pipelinewiseConfigurationConfigMap.Name = configIdentifier.Name
		pipelinewiseConfigurationConfigMap.Data = map[string]string{}
		tapKeyName := fmt.Sprintf("tap_%v.yaml", batchv1alpha1.GetTapID(&pipelinewiseJob))
		pipelinewiseConfigurationConfigMap.Data[tapKeyName] = string(tapYaml)
		targetKeyName := fmt.Sprintf("target_%v.yaml", batchv1alpha1.GetTargetID(&pipelinewiseJob))
		pipelinewiseConfigurationConfigMap.Data[targetKeyName] = string(targetYaml)
		err = r.Create(ctx, &pipelinewiseConfigurationConfigMap)
		if err != nil {
			log.Error(err, "Failed to create pipelinewise configuration")
			return ctrl.Result{}, err
		}
	}

	// Create PVC
	constructPersistentLayer := func(piplinewiseJob *batchv1alpha1.PipelinewiseJob, identifier ktypes.NamespacedName) corev1.PersistentVolumeClaim {
		pvc := corev1.PersistentVolumeClaim{
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

		return pvc
	}
	var pvc corev1.PersistentVolumeClaim
	volumeIdentifier := identifiers[VolumeExternalResourceID]
	if err := r.Get(ctx, volumeIdentifier, &pvc); err != nil {
		pvc = constructPersistentLayer(&pipelinewiseJob, volumeIdentifier)
		err = r.Create(ctx, &pvc)
		if err != nil {
			log.Error(err, "Failed to create PVC")
			return ctrl.Result{}, err
		}
	}

	// Create actual kubernetes job to run
	constructExecutorJob := func(pipelinewiseJob *batchv1alpha1.PipelinewiseJob, identifier ktypes.NamespacedName) kbatchv1beta1.CronJob {
		imageName := fmt.Sprintf("dirathea/pipelinewise:%v-%v-%v", viper.GetString("PIPELINEWISE_VERSION"), batchv1alpha1.GetTapConnectorID(pipelinewiseJob), batchv1alpha1.GetTargetID(pipelinewiseJob))
		if pipelinewiseJob.Spec.Image != nil {
			imageName = *pipelinewiseJob.Spec.Image
		}
		cronJob := kbatchv1beta1.CronJob{
			ObjectMeta: identifierToMeta(identifier),
			Spec: kbatchv1beta1.CronJobSpec{
				Schedule: pipelinewiseJob.Spec.Schedule,
				JobTemplate: kbatchv1beta1.JobTemplateSpec{
					Spec: batchv1.JobSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								RestartPolicy: corev1.RestartPolicyNever,
								InitContainers: []corev1.Container{
									corev1.Container{
										Name:  "configuration-import",
										Image: imageName,
										Args: []string{
											"import",
											"--dir",
											"/configurations",
										},
										VolumeMounts: []corev1.VolumeMount{
											corev1.VolumeMount{
												Name:      "pipelinewise-configuration",
												MountPath: "/configurations",
											},
											corev1.VolumeMount{
												Name:      "runtime-volume",
												MountPath: "/root/.pipelinewise",
											},
										},
									},
								},
								Containers: []corev1.Container{
									corev1.Container{
										Name:  "pipelinewise",
										Image: imageName,
										Args: []string{
											"run_tap",
											"--tap",
											batchv1alpha1.GetTapID(pipelinewiseJob),
											"--target",
											batchv1alpha1.GetTargetID(pipelinewiseJob),
											"--extra_log",
										},
										VolumeMounts: []corev1.VolumeMount{
											corev1.VolumeMount{
												Name:      "pipelinewise-configuration",
												MountPath: "/configurations",
											},
											corev1.VolumeMount{
												Name:      "runtime-volume",
												MountPath: "/root/.pipelinewise",
											},
										},
									},
								},
								Volumes: []corev1.Volume{
									corev1.Volume{
										Name: "pipelinewise-configuration",
										VolumeSource: corev1.VolumeSource{
											ConfigMap: &corev1.ConfigMapVolumeSource{
												LocalObjectReference: corev1.LocalObjectReference{
													Name: pipelinewiseConfigurationConfigMap.Name,
												},
											},
										},
									},
									corev1.Volume{
										Name: "runtime-volume",
										VolumeSource: corev1.VolumeSource{
											PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
												ClaimName: pvc.Name,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		return cronJob
	}
	jobIdentifier := identifiers[JobMapExternalResourceID]
	var executorJob kbatchv1beta1.CronJob
	if err := r.Get(ctx, jobIdentifier, &executorJob); err != nil {
		executorJob = constructExecutorJob(&pipelinewiseJob, jobIdentifier)
		err = r.Create(ctx, &executorJob)
		if err != nil {
			log.Error(err, "Failed to create kubernetes Job")
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

		var podList v1.PodList
		opts := []client.ListOption{
			client.InNamespace(executorJob.Namespace),
			client.MatchingLabels{"job-name": executorJob.Name},
		}
		err = r.List(deleteCtx, &podList, opts...)
		if err != nil {
			return err
		}
		for _, pod := range podList.Items {
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

// Helper function to get external resources identifier
func resourcesIdentifierGenerator(pipelinewiseJob *batchv1alpha1.PipelinewiseJob, prefix string) ktypes.NamespacedName {
	return ktypes.NamespacedName{
		Namespace: pipelinewiseJob.Namespace,
		Name:      fmt.Sprintf("%v-%v", prefix, pipelinewiseJob.Name),
	}
}

func resourcesIdentifier(pipelinewiseJob *batchv1alpha1.PipelinewiseJob) map[ExternalResourceID]ktypes.NamespacedName {
	return map[ExternalResourceID]ktypes.NamespacedName{
		ConfigMapExternalResourceID: resourcesIdentifierGenerator(pipelinewiseJob, "pipelinewise-config"),
		VolumeExternalResourceID:    resourcesIdentifierGenerator(pipelinewiseJob, "pipelinewise-volume"),
		JobMapExternalResourceID:    resourcesIdentifierGenerator(pipelinewiseJob, "pipelinewise-job"),
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
