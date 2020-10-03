package ch15_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* 单元测试: package testing */

// 单元测试简单样例
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

// 表组测试
func TestTableDriven(t *testing.T) {
	dataList := []struct {
		Input  int
		Except int
	}{
		{
			-1,
			1,
		},
		{
			1,
			1,
		},
		{
			0,
			0,
		},
	}

	for _, data := range dataList {
		result := Abs(data.Input)
		if result != data.Except {
			t.Errorf("Abs(%d) = %d; want %d", data.Input, result, data.Except)
		}
	}
}

// Benchmark: 基准测试
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

func BenchmarkRandInt2(b *testing.B) {
	// long time program execute...
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

// Examples: 结果校验
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}

// assert: "github.com/stretchr/testify/assert"
func TestHello(t *testing.T) {
	val := "hello"
	except := "hello"
	assert.True(t, val == except)
}

// Skipping: 跳过测试
func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	fmt.Println("长时间测试任务...")
}

// 子测试
func TestFoo(t *testing.T) {
	// <setup code>
	t.Run("A=1", func(t *testing.T) { fmt.Println("a1 test") })
	t.Run("A=2", func(t *testing.T) { fmt.Println("a2 test") })
	t.Run("B=1", func(t *testing.T) { fmt.Println("b1 test") })
	// <tear-down code>
}

// Main
// func TestMain(m *testing.M) {
// 	fmt.Println("start...")
// 	existCode := m.Run()
// 	fmt.Println("end...")
// 	os.Exit(existCode)
// }

// Log
func TestLog(t *testing.T) {
	if testing.Short() {
		t.Skip("testing skip")
	}
	t.Log("testing log")
	t.Error("testing error")
	t.Fatal("testing fatal")
	fmt.Println("over...")
}

// Fail
func TestFail(t *testing.T) {
	defer func() {
		fmt.Println(t.Failed())
	}()

	fmt.Println("before...")
	if testing.Short() {
		t.Fail()
	}
	fmt.Println("after...")
}

// FailNow
func TestFailNow(t *testing.T) {
	defer func() {
		fmt.Println(t.Failed())
	}()

	fmt.Println("before...")
	t.FailNow()
	fmt.Println("after...")
}

// Skip
func TestSkip(t *testing.T) {
	defer func() {
		fmt.Println(t.Skipped())
	}()

	fmt.Println("before...")
	if testing.Short() {
		t.SkipNow()
	}
	fmt.Println("after...")
}

// Other
func TestOther(t *testing.T) {
	fmt.Println(t.Name()) // TestOther
}

// Mock
func TestMock(t *testing.T) {
	type result struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}

	fmt.Println("mock test start.")
	server := mockServer()
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("mock test failed, request err:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("mock test failed, check http response code err:", resp.StatusCode)
	}

	var mockResult result
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("mock test failed, read response body err:", err)
	}
	if err := json.Unmarshal(bytes, &mockResult); err != nil {
		t.Fatal("mock test failed, convert response body to result struct err:", err)
	}
	if mockResult.Code != 0 {
		t.Fatal("mock test failed, check result code err:", mockResult)
	}
	fmt.Println("mock test end.")
}

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"code":0}`)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}
