package apimachinery

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/restmapper"
)

func TestRestMapper(t *testing.T) {
	// create discovery client
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(RestConfig())
	checkErr(err)

	// api-resource list
	apiGroupResourcesList, err := restmapper.GetAPIGroupResources(discoveryClient)
	checkErr(err)

	// create restMapper
	restMapper := restmapper.NewDiscoveryRESTMapper(apiGroupResourcesList)

	// convert GVK to GVR
	gvk := schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Pod",
	}
	mapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	checkErr(err)
	fmt.Printf("gvr: |%s|%s|%s|, gvk: |%s|%s|%s|, scope: |%s|\n",
		mapping.Resource.Group, mapping.Resource.Version, mapping.Resource.Resource,
		mapping.GroupVersionKind.Group, mapping.GroupVersionKind.Version, mapping.GroupVersionKind.Kind,
		mapping.Scope.Name(),
	)
	// Output:
	// gvr: ||v1|pods|, gvk: ||v1|Pod|, scope: |namespace|
}
