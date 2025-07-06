package client_go

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/types"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ClientSet 示例
func TestClientSet(t *testing.T) {
	clientSet, err := kubernetes.NewForConfig(RestConfig())
	checkErr(err)

	// node
	nodeList, err := clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	checkErr(err)
	for _, d := range nodeList.Items {
		fmt.Printf("[nodeList] name: %s\n", d.Name)
	}

	// namespace
	namespaceList, err := clientSet.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	checkErr(err)
	for _, d := range namespaceList.Items {
		fmt.Printf("[namespaceList] name: %s\n", d.Name)
	}

	// deployment
	deploymentList, err := clientSet.AppsV1().Deployments("kube-system").List(context.Background(), metav1.ListOptions{})
	checkErr(err)
	for _, d := range deploymentList.Items {
		fmt.Printf("[deploymentList] namespace: %v, name: %v, replicas: %d\n", d.Namespace, d.Name, *d.Spec.Replicas)
	}

	// pod
	podList, err := clientSet.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{})
	checkErr(err)
	for _, d := range podList.Items {
		fmt.Printf("[podList] namespace: %v, name: %v, status: %v\n", d.Namespace, d.Name, d.Status.Phase)
	}

	// Output:
	// [nodeList] name: 192.168.0.183
	// [namespaceList] name: default
	// [namespaceList] name: kube-node-lease
	// [namespaceList] name: kube-public
	// [namespaceList] name: kube-system
	// [namespaceList] name: monitoring
	// [deploymentList] namespace: kube-system, name: coredns, replicas: 2
	// [deploymentList] namespace: kube-system, name: everest-csi-controller, replicas: 2
	// [deploymentList] namespace: kube-system, name: node-local-dns-admission-controller, replicas: 2
	// [deploymentList] namespace: kube-system, name: node-problem-controller, replicas: 2
	// [podList] namespace: kube-system, name: cceaddon-npd-njh8f, status: Running
	// [podList] namespace: kube-system, name: coredns-dff6f9b-799nw, status: Pending
	// [podList] namespace: kube-system, name: coredns-dff6f9b-kf7zf, status: Running
	// [podList] namespace: kube-system, name: everest-csi-controller-86f5475978-968j4, status: Pending
	// [podList] namespace: kube-system, name: everest-csi-controller-86f5475978-bbb2g, status: Running
	// [podList] namespace: kube-system, name: everest-csi-driver-s2tfp, status: Running
	// [podList] namespace: kube-system, name: icagent-xkt5d, status: Running
	// [podList] namespace: kube-system, name: node-local-dns-578nl, status: Running
	// [podList] namespace: kube-system, name: node-local-dns-admission-controller-858d7bb84d-h2flk, status: Pending
	// [podList] namespace: kube-system, name: node-local-dns-admission-controller-858d7bb84d-jfzbs, status: Running
	// [podList] namespace: kube-system, name: node-problem-controller-7466df6b9-rrrhv, status: Running
	// [podList] namespace: kube-system, name: node-problem-controller-7466df6b9-wb425, status: Pending
}

func TestCRUD(t *testing.T) {
	clientSet, err := kubernetes.NewForConfig(RestConfig())
	checkErr(err)

	// create deploy
	log.Println("create deploy start")
	ctx := context.Background()
	namespace, name, appKey := "default", fmt.Sprintf("nginx-%s", time.Now().Format(time.DateOnly)), "app"
	deploy := newDeployment(namespace, name, appKey)
	createDeploy, err := clientSet.AppsV1().Deployments(namespace).Create(ctx, deploy, metav1.CreateOptions{})
	checkErr(err)
	log.Println("create deploy end")

	// update deploy
	log.Println("update deploy start")
	updateDeploy := createDeploy.DeepCopy()
	updateDeploy.Spec.Replicas = int32Ptr(2)
	updateDeploy, err = clientSet.AppsV1().Deployments(namespace).Update(ctx, updateDeploy, metav1.UpdateOptions{})
	checkErr(err)
	log.Println("update deploy end")

	// update deploy with patch
	log.Println("update deploy with patch start")
	patchBody := map[string]interface{}{
		"metadata": map[string]interface{}{
			"annotations": map[string]string{
				"app": name,
			},
		},
	}
	patchBytes, err := json.Marshal(patchBody)
	checkErr(err)
	updateDeploy, err = clientSet.AppsV1().Deployments(namespace).Patch(ctx, name, types.MergePatchType, patchBytes, metav1.PatchOptions{})
	checkErr(err)
	log.Println("update deploy with patch end")

	// query deploy
	log.Println("query deploy start")
	getDeploy, err := clientSet.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
	checkErr(err)
	log.Printf("query deploy end, body:\n%s\n", toJsonStr(getDeploy))

	// delete deploy
	log.Println("delete deploy start")
	err = clientSet.AppsV1().Deployments(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	checkErr(err)
	log.Println("delete deploy end")
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

func toJsonStr(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	checkErr(err)
	return string(data)
}
