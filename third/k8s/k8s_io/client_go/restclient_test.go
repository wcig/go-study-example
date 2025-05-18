package client_go

import (
	"context"
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// RESTClient 示例
func TestRestClient(t *testing.T) {
	// 加载Config
	config := RestConfig()

	// 设置API路径 (这里查询的pod为无组名资源组使用/api而不是/apis)
	config.APIPath = "api"

	// 设置资源组和版本, 对应GVR中的GV
	config.GroupVersion = &corev1.SchemeGroupVersion

	// 设置编解码器
	config.NegotiatedSerializer = scheme.Codecs

	// 初始化Client
	restClient, err := rest.RESTClientFor(config)
	checkErr(err)

	// 构造接收pod列表对象
	result := &corev1.PodList{}

	// 查询pod列表
	err = restClient.Get().
		Namespace("kube-system").                                                // 命名空间
		Resource("pods").                                                        // 资源对象, 对应GVR的R
		VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec). // 参数及序列化工具
		Do(context.Background()).                                                // 发送请求
		Into(result)                                                             // 写入返回值
	checkErr(err)

	// 列出pod列表
	for _, d := range result.Items {
		fmt.Printf("namespace: %v, name: %v, status: %v\n", d.Namespace, d.Name, d.Status.Phase)
	}

	// Output:
	// namespace: kube-system, name: cceaddon-npd-njh8f, status: Running
	// namespace: kube-system, name: coredns-dff6f9b-799nw, status: Pending
	// namespace: kube-system, name: coredns-dff6f9b-kf7zf, status: Running
	// namespace: kube-system, name: everest-csi-controller-86f5475978-968j4, status: Pending
	// namespace: kube-system, name: everest-csi-controller-86f5475978-bbb2g, status: Running
	// namespace: kube-system, name: everest-csi-driver-s2tfp, status: Running
	// namespace: kube-system, name: icagent-xkt5d, status: Running
	// namespace: kube-system, name: node-local-dns-578nl, status: Running
	// namespace: kube-system, name: node-local-dns-admission-controller-858d7bb84d-h2flk, status: Pending
	// namespace: kube-system, name: node-local-dns-admission-controller-858d7bb84d-jfzbs, status: Running
	// namespace: kube-system, name: node-problem-controller-7466df6b9-rrrhv, status: Running
	// namespace: kube-system, name: node-problem-controller-7466df6b9-wb425, status: Pending
}
