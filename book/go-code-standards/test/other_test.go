package test

import (
	"fmt"
	"testing"
	"time"
)

// TestTimeSince
func TestTimeSince(t *testing.T) {
	tt := time.Date(2019, 5, 25, 0, 0, 0, 0, time.Local)

	t1 := time.Now().Sub(tt)

	t2 := time.Since(tt)

	fmt.Println(tt.Unix())
	fmt.Println(t1.Seconds())
	fmt.Println(t2.Seconds())
}

// 多返回值函数定义
func getSize01(videoDir string) (int, int) {
	//...
	return 0, 0
}

func getSize02(videoDir string) (width, height int) {
	//...
	return 0, 0
}
