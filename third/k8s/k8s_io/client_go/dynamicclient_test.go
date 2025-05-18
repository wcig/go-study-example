package client_go

import (
	"context"
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

// DynamicClient 示例
func TestDynamicClient(t *testing.T) {
	dynamicClient, err := dynamic.NewForConfig(RestConfig())
	checkErr(err)

	// pod
	podGVR := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
	unstructuredPodList, err := dynamicClient.Resource(podGVR).
		Namespace("kube-system").
		List(context.Background(), metav1.ListOptions{})
	checkErr(err)
	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredPodList.UnstructuredContent(), podList)
	checkErr(err)
	for _, d := range podList.Items {
		fmt.Printf("[podList] namespace: %v, name: %v, status: %v\n", d.Namespace, d.Name, d.Status.Phase)
	}

	// deployment
	deploymentGVR := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}
	unstructuredDeploymentList, err := dynamicClient.Resource(deploymentGVR).
		Namespace("kube-system").
		List(context.Background(), metav1.ListOptions{})
	checkErr(err)
	for _, d := range unstructuredDeploymentList.Items {
		replicas, found, err2 := unstructured.NestedInt64(d.Object, "spec", "replicas")
		if err2 != nil || !found {
			fmt.Printf("Replicas not found for deployment %s: error=%s", d.GetName(), err2)
			continue
		}
		fmt.Printf("[deploymentList] namespace: %v, name: %v, replicas: %d\n", d.GetNamespace(), d.GetName(), replicas)
	}

	// Output:
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
	// [deploymentList] namespace: kube-system, name: coredns, replicas: 2
	// [deploymentList] namespace: kube-system, name: everest-csi-controller, replicas: 2
	// [deploymentList] namespace: kube-system, name: node-local-dns-admission-controller, replicas: 2
	// [deploymentList] namespace: kube-system, name: node-problem-controller, replicas: 2
}
