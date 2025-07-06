package apimachinery

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// 加载配置文件生成Config对象
func RestConfig() *rest.Config {
	homeDir := os.Getenv("HOME")
	kubeConfigPath := filepath.Join(homeDir, ".kube/config")
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	checkErr(err)
	return restConfig
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewClientSet() *kubernetes.Clientset {
	clientSet, err := kubernetes.NewForConfig(RestConfig())
	checkErr(err)
	return clientSet
}
