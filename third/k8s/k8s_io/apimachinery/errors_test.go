package apimachinery

import (
	"context"
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestErrors(t *testing.T) {
	c := NewClientSet()
	_, err := c.AppsV1().Deployments("default").Get(context.Background(), "test-not-found", metav1.GetOptions{})
	if err != nil {
		handleErr(err)
	}
	// Output:
	// match apimachinery not found err: deployments.apps "test-not-found" not found
}

// 错误处理
func handleErr(err error) {
	if errors.IsNotFound(err) {
		fmt.Printf("match apimachinery not found err: %v\n", err)
	} else if errors.IsConflict(err) {
		fmt.Printf("match apimachinery conflict err: %v\n", err)
	} else if statusError, ok := err.(*errors.StatusError); ok {
		fmt.Printf("match apimachinery err: %v\n", statusError)
	} else {
		fmt.Printf("no match apimachinery err: %v\n", err)
	}
}
