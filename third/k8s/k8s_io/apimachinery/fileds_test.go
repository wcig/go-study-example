package apimachinery

import (
	"context"
	"log"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

func TestField(t *testing.T) {
	c := NewClientSet()
	const namespace, name, appKey = "default", "nginx", "app"

	log.Println("query deploy nginx")
	_, err := c.AppsV1().Deployments(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		log.Println("query not found deploy nginx")
		log.Println("create deploy nginx")
		newDeploy := newDeployment(namespace, name, appKey)
		_, err = c.AppsV1().Deployments("default").Create(context.Background(), newDeploy, metav1.CreateOptions{})
		if err != nil {
			log.Println("create deploy nginx success")
		}
	}

	log.Println("list deploy with nginx label")
	requirement, err := labels.NewRequirement(appKey, selection.Equals, []string{name})
	checkErr(err)
	selector := labels.NewSelector().Add(*requirement)
	listOptions := metav1.ListOptions{LabelSelector: selector.String()}
	deployList, err := c.AppsV1().Deployments(namespace).List(context.Background(), listOptions)
	if err != nil {
		log.Printf("list deploy with nginx label err: %v\n", err)
		return
	}
	log.Printf("list deploy with nginx label success, size: %d\n", len(deployList.Items))

	log.Println("delete deploy nginx")
	err = c.AppsV1().Deployments("default").Delete(context.Background(), "nginx", metav1.DeleteOptions{})
	if err != nil {
		log.Printf("delete deploy nginx err: %v\n", err)
		return
	}
	log.Println("delete deploy nginx success")

	// Output:
	// 2025/07/06 20:21:58 query deploy nginx
	// 2025/07/06 20:21:58 query not found deploy nginx
	// 2025/07/06 20:21:58 create deploy nginx
	// 2025/07/06 20:21:58 list deploy with nginx label
	// 2025/07/06 20:21:58 list deploy with nginx label success, size: 1
	// 2025/07/06 20:21:58 delete deploy nginx
	// 2025/07/06 20:21:58 delete deploy nginx success
}

func newDeployment(namespace, name, appKey string) *appsv1.Deployment {
	deployLabels := map[string]string{
		appKey: name,
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    deployLabels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: deployLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: deployLabels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:latest",
						},
					},
				},
			},
		},
	}
}

func int32Ptr(i int32) *int32 { return &i }
