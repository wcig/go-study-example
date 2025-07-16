package client_go

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestFakeClient(t *testing.T) {
	const (
		pod1Name = "test-pod-1"
		pod2Name = "test-pod-2"
		ns       = "default"
	)

	pod1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pod1Name,
			Namespace: ns,
		},
	}
	pod2 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pod2Name,
			Namespace: ns,
		},
	}

	// create fake client with pod1
	cs := fake.NewClientset(pod1)

	// create pod2
	_, err := cs.CoreV1().Pods(ns).Create(context.Background(), pod2, metav1.CreateOptions{})
	assert.Nil(t, err)

	// list pod
	podList, err := cs.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(podList.Items))
	var podNames []string
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.GetName())
	}
	sort.StringSlice(podNames).Sort()
	assert.Equal(t, podNames, []string{pod1Name, pod2Name})
}
