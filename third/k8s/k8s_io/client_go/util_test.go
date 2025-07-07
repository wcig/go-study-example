package client_go

import (
	"context"
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
)

func TestRetry(t *testing.T) {
	// 加载 kubeconfig 文件
	config, err := clientcmd.BuildConfigFromFlags("", "/path/to/kubeconfig")
	if err != nil {
		panic(err.Error())
	}

	// 创建一个新的 Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 重试更新 Pod
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// 获取 Pod
		pod, getErr := clientset.CoreV1().Pods("default").Get(context.TODO(), "example-pod", metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		// 更新 Pod 的镜像
		pod.Spec.Containers[0].Image = "nginx:latest"
		_, updateErr := clientset.CoreV1().Pods("default").Update(context.TODO(), pod, metav1.UpdateOptions{})
		return updateErr
	})

	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}

	fmt.Println("Updated Pod example-pod.")
}
