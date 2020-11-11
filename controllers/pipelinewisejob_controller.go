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
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	batchv1alpha1 "github.com/dirathea/pipelinewise-operator/api/v1alpha1"
)

// PipelinewiseJobReconciler reconciles a PipelinewiseJob object
type PipelinewiseJobReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=batch.pipelinewise,resources=pipelinewisejobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.pipelinewise,resources=pipelinewisejobs/status,verbs=get;update;patch

func (r *PipelinewiseJobReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("pipelinewisejob", req.NamespacedName)

	// your logic here

	// Load Pipelinewise Job
	var pipelinewiseJob batchv1alpha1.PipelinewiseJob
	if err := r.Get(ctx, req.NamespacedName, &pipelinewiseJob); err != nil {
		log.Error(err, "unable to fetch PipelinewiseJob")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

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

	configurationName := fmt.Sprintf("%v-pipelinewise-configuration", pipelinewiseJob.Name)
	persistenceName := fmt.Sprintf("%v-runtime-volume", pipelinewiseJob.Name)

	pipelinewiseConfigurationConfigMap := corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configurationName,
			Namespace: pipelinewiseJob.Namespace,
		},
		Data: map[string]string{
			"tap.yaml":    tapYaml,
			"target.yaml": targetYaml,
		},
	}
	err = r.Create(ctx, &pipelinewiseConfigurationConfigMap)
	if err != nil {
		log.Error(err, "Failed to create pipelinewise configuration")
		return ctrl.Result{}, err
	}

	// Create PVC
	constructPersistentLayer := func(piplinewiseJob *batchv1alpha1.PipelinewiseJob) (corev1.PersistentVolumeClaim, error) {
		pvc := corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      persistenceName,
				Namespace: pipelinewiseJob.Namespace,
			},
		}

		return pvc, nil
	}
	pvc, err := constructPersistentLayer(&pipelinewiseJob)
	if err != nil {
		log.Error(err, "Failed to construct persistence layer for executor")
		return ctrl.Result{}, err
	}
	err = r.Create(ctx, &pvc)
	if err != nil {
		log.Error(err, "Failed to create PVC")
		return ctrl.Result{}, err
	}

	// Create actual kubernetes job to run
	constructExecutorJob := func(pipelinewiseJob *batchv1alpha1.PipelinewiseJob) (batchv1.Job, error) {
		job := batchv1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Name:      configurationName,
				Namespace: pipelinewiseJob.Namespace,
			},
			Spec: batchv1.JobSpec{
				Template: corev1.PodTemplateSpec{
					Spec: corev1.PodSpec{
						InitContainers: []corev1.Container{
							corev1.Container{
								Command: []string{
									"/app/run.sh",
								},
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
								Command: []string{
									"/app/run.sh",
								},
								Args: []string{
									"run_tap",
									"--tap",
									"/configurations/tap.yaml",
									"--target",
									"/configurations/target.yaml",
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
											Name: configurationName,
										},
									},
								},
							},
							corev1.Volume{
								Name: "runtime-volume",
								VolumeSource: corev1.VolumeSource{
									PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
										ClaimName: persistenceName,
									},
								},
							},
						},
					},
				},
			},
		}
		return job, nil
	}

	executorJob, err := constructExecutorJob(&pipelinewiseJob)
	if err != nil {
		log.Error(err, "Failed to construct pipelinewise executor")
		return ctrl.Result{}, err
	}

	err = r.Create(ctx, &executorJob)
	if err != nil {
		log.Error(err, "Failed to create pipelinewise executor")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *PipelinewiseJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1alpha1.PipelinewiseJob{}).
		Complete(r)
}
