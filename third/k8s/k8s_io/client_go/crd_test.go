package client_go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCRD(t *testing.T) {
	// create crd clientSet
	cs, err := clientset.NewForConfig(RestConfig())
	assert.Nil(t, err)

	// create crd
	result, err := cs.ApiextensionsV1().CustomResourceDefinitions().Create(context.Background(), crd, metav1.CreateOptions{})
	assert.Nil(t, err)
	assert.Equal(t, result.GetName(), crd.GetName())

	// delete crd
	err = cs.ApiextensionsV1().CustomResourceDefinitions().Delete(context.Background(), crd.GetName(), metav1.DeleteOptions{})
	assert.Nil(t, err)
}

var crd = &v1.CustomResourceDefinition{
	ObjectMeta: metav1.ObjectMeta{
		Name: "examples.mygroup.mydomain",
	},
	Spec: v1.CustomResourceDefinitionSpec{
		Group: "mygroup.mydomain",
		Names: v1.CustomResourceDefinitionNames{
			Plural:   "examples",
			Singular: "example",
			Kind:     "Example",
		},
		Scope: v1.NamespaceScoped,
		Versions: []v1.CustomResourceDefinitionVersion{
			{
				Name:    "v1alpha1",
				Served:  true,
				Storage: true,
				Schema: &v1.CustomResourceValidation{
					OpenAPIV3Schema: &v1.JSONSchemaProps{
						Type: "object",
						Properties: map[string]v1.JSONSchemaProps{
							"spec": {
								Type: "object",
								Properties: map[string]v1.JSONSchemaProps{
									"deploymentName": {
										Type: "string",
									},
									"replicas": {
										Type:    "integer",
										Minimum: float64Ptr(1),
										Maximum: float64Ptr(10),
									},
								},
							},
							"status": {
								Type: "object",
								Properties: map[string]v1.JSONSchemaProps{
									"availableReplicas": {
										Type: "integer",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func float64Ptr(f float64) *float64 {
	return &f
}
