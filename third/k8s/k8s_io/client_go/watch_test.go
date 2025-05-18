package client_go

import (
	"context"
	"fmt"
	"testing"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Watch 示例
func TestWatch(t *testing.T) {
	clientSet, err := kubernetes.NewForConfig(RestConfig())
	checkErr(err)

	go modifyDeployment(clientSet)

	watch, err := clientSet.AppsV1().Deployments("default").Watch(context.Background(), metav1.ListOptions{})
	checkErr(err)

	fmt.Println(">> start watch deployments")
	for event := range watch.ResultChan() {
		deployment, ok := event.Object.(*appsv1.Deployment)
		if ok {
			fmt.Printf(">> watch event type: %8s, deployment name: %s, replicas: %d, generation: %d, resourceVersion: %s\n",
				event.Type, deployment.GetObjectMeta().GetName(), *deployment.Spec.Replicas,
				deployment.GetObjectMeta().GetGeneration(), deployment.GetObjectMeta().GetResourceVersion())
		}
	}

	// Output:
	// >> start watch deployments
	// Creating deployment...
	// Created deployment "nginx-deployment".
	// >> watch event type:    ADDED, deployment name: nginx-deployment, replicas: 1, generation: 1, resourceVersion: 205461
	// >> watch event type: MODIFIED, deployment name: nginx-deployment, replicas: 1, generation: 1, resourceVersion: 205463
	// >> watch event type: MODIFIED, deployment name: nginx-deployment, replicas: 1, generation: 1, resourceVersion: 205466
	// >> watch event type: MODIFIED, deployment name: nginx-deployment, replicas: 1, generation: 1, resourceVersion: 205472
	// >> watch event type: MODIFIED, deployment name: nginx-deployment, replicas: 1, generation: 1, resourceVersion: 205481
	// Updating deployment...
	// >> watch event type: MODIFIED, deployment name: nginx-deployment, replicas: 2, generation: 2, resourceVersion: 205533
	// Updated deployment "nginx-deployment".
	// >> watch event type: MODIFIED, deployment name: nginx-deployment, replicas: 2, generation: 2, resourceVersion: 205534
	// >> watch event type: MODIFIED, deployment name: nginx-deployment, replicas: 2, generation: 2, resourceVersion: 205539
	// >> watch event type: MODIFIED, deployment name: nginx-deployment, replicas: 2, generation: 2, resourceVersion: 205544
	// >> watch event type: MODIFIED, deployment name: nginx-deployment, replicas: 2, generation: 2, resourceVersion: 205553
	// Deleting deployment...
	// Deleted deployment "nginx-deployment".
	// >> watch event type:  DELETED, deployment name: nginx-deployment, replicas: 2, generation: 2, resourceVersion: 205605
}

func modifyDeployment(clientSet *kubernetes.Clientset) {
	deploymentsClient := clientSet.AppsV1().Deployments(corev1.NamespaceDefault)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx-deployment",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "nginx-deployment",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "web",
							Image: "nginx:alpine",
							Env: []corev1.EnvVar{
								{
									Name:  "PAAS_APP_NAME",
									Value: "nginx-deployment",
								},
								{
									Name:  "PAAS_NAMESPACE",
									Value: "default",
								},
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("250m"),
									corev1.ResourceMemory: resource.MustParse("512Mi"),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("250m"),
									corev1.ResourceMemory: resource.MustParse("512Mi"),
								},
							},
						},
					},
				},
			},
		},
	}

	// create deployment
	time.Sleep(10 * time.Second)
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	checkErr(err)
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	// modify deployment replicas
	time.Sleep(10 * time.Second)
	fmt.Println("Updating deployment...")
	deployment.Spec.Replicas = int32Ptr(2)
	result, err = deploymentsClient.Update(context.Background(), deployment, metav1.UpdateOptions{})
	checkErr(err)
	fmt.Printf("Updated deployment %q.\n", result.GetObjectMeta().GetName())

	// delete deployment
	time.Sleep(10 * time.Second)
	fmt.Println("Deleting deployment...")
	err = deploymentsClient.Delete(context.Background(), deployment.GetObjectMeta().GetName(), metav1.DeleteOptions{})
	checkErr(err)
	fmt.Printf("Deleted deployment %q.\n", deployment.GetObjectMeta().GetName())
}

func int32Ptr(i int32) *int32 { return &i }
