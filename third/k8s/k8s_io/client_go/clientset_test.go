package client_go

import (
	"context"
	"fmt"
	"testing"

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
