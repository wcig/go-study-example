package client_go

import (
	"fmt"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery/cached/disk"
)

// CachedDiscoveryClient 示例
func TestCachedDiscoveryClient(t *testing.T) {
	cachedDiscoveryClient, err := disk.NewCachedDiscoveryClientForConfig(
		RestConfig(),
		".cache/discovery",
		".cache/http",
		time.Hour)
	checkErr(err)

	apiGroupList, apiResourceList, err := cachedDiscoveryClient.ServerGroupsAndResources()
	checkErr(err)

	for _, d := range apiGroupList {
		fmt.Printf("[apiGroupList] name: %s, versions: %v\n", d.Name, d.Versions)
	}

	for _, d := range apiResourceList {
		gv, err2 := schema.ParseGroupVersion(d.GroupVersion)
		checkErr(err2)
		for _, apiResource := range d.APIResources {
			fmt.Printf("[apiResourceList] name: %s, group: %s, version: %s, kind: %s\n",
				apiResource.Name, gv.Group, gv.Version, apiResource.Kind)
		}
	}
}
