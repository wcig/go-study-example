package client_go

import (
	"fmt"
	"testing"
	"time"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

func TestInformer(t *testing.T) {
	// create clientSet
	clientSet, err := kubernetes.NewForConfig(RestConfig())
	checkErr(err)

	// create informer factory (重同步周期)
	informerFactory := informers.NewSharedInformerFactory(clientSet, time.Second*30)

	deployInformer := informerFactory.Apps().V1().Deployments()
	informer := deployInformer.Informer()
	lister := deployInformer.Lister()
	_, err = informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    addDeployment,
		UpdateFunc: updateDeployment,
		DeleteFunc: deleteDeployment,
	})
	checkErr(err)

	stopCh := make(chan struct{})
	defer close(stopCh)

	informerFactory.Start(stopCh)
	informerFactory.WaitForCacheSync(stopCh)

	deployments, err := lister.Deployments("").List(labels.Everything())
	checkErr(err)
	for i, d := range deployments {
		fmt.Printf("index %d, namespace: %v, name: %v, replicas: %d\n", i, d.Namespace, d.Name, *d.Spec.Replicas)
	}
	stopCh <- struct{}{}

	// Output:
	// add deployment: example-foo
	// add deployment: coredns
	// add deployment: local-path-provisioner
	// index 0, namespace: default, name: example-foo, replicas: 2
	// index 1, namespace: kube-system, name: coredns, replicas: 2
	// index 2, namespace: local-path-storage, name: local-path-provisioner, replicas: 1
}

func TestSimpleInformer(t *testing.T) {
	clientSet, err := kubernetes.NewForConfig(RestConfig())
	checkErr(err)

	factory := informers.NewSharedInformerFactory(clientSet, time.Second*30)
	informer := factory.Apps().V1().Deployments().Informer()
	_, err = informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    addDeployment,
		UpdateFunc: updateDeployment,
		DeleteFunc: deleteDeployment,
	})
	checkErr(err)

	stopCh := make(chan struct{})
	defer close(stopCh)

	factory.Start(stopCh)
	if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
		panic("Timed out waiting for caches to sync")
	}
	stopCh <- struct{}{}
}

func addDeployment(obj interface{}) {
	deploy, ok := obj.(*v1.Deployment)
	if !ok {
		return
	}
	fmt.Println("add deployment:", deploy.Name)
}

func updateDeployment(old, new interface{}) {
	oldDeploy, ok := old.(*v1.Deployment)
	if !ok {
		return
	}
	newDeploy, ok := new.(*v1.Deployment)
	if !ok {
		return
	}
	fmt.Println("update deployment:", oldDeploy.Name, newDeploy.Name)
}

func deleteDeployment(obj interface{}) {
	deploy, ok := obj.(*v1.Deployment)
	if !ok {
		return
	}
	fmt.Println("delete deployment:", deploy.Name)
}
