package apolloconfig_agollo

import (
	"fmt"
	. "go-app/third/apollo/common"
	"sync"
	"testing"
	"time"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

// 参考: https://github.com/zouyx/agollo_demo

// 快速开始
func TestQuickStart(t *testing.T) {
	c := &config.AppConfig{
		AppID:          AppID,
		Cluster:        Cluster,
		IP:             fmt.Sprintf("http://%s", ApolloAddr),
		NamespaceName:  Namespace,
		IsBackupConfig: true,
		Secret:         Secret,
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	fmt.Println("初始化Apollo配置成功")

	cache := client.GetConfigCache(c.NamespaceName)

	value, _ := cache.Get("foo")
	fmt.Println("foo:", value)

	s, ok := value.(string)
	fmt.Println(s, ok)

	value, _ = cache.Get("timeout")
	fmt.Println("timeout:", value)

	s, ok = value.(string)
	fmt.Println(s, ok)

	i, ok := value.(int)
	fmt.Println(i, ok)

	cache.Range(func(key, value interface{}) bool {
		fmt.Println("range:", key, value)
		return true
	})

	time.Sleep(time.Second)
	// Output:
	// 初始化Apollo配置成功
	// foo: bar
	// bar true
	// timeout: 100
	// 100 true
	// 0 false
	// range: timeout 100
	// range: foo bar

	// SampleApp-application.json:
	// {"appId":"SampleApp","cluster":"","namespaceName":"application","releaseKey":"","configurations":{"foo":"bar","timeout":"100"}}
}

// 简单 Client
func TestSimpleClient(t *testing.T) {
	c := &config.AppConfig{
		AppID:          AppID,
		Cluster:        Cluster,
		IP:             fmt.Sprintf("http://%s", ApolloAddr),
		NamespaceName:  Namespace,
		IsBackupConfig: true,
		Secret:         Secret,
	}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("初始化Apollo配置成功")

	fooStrVal := client.GetStringValue("foo", "")
	fmt.Println("foo:", fooStrVal)

	timeoutStrVal := client.GetStringValue("timeout", "")
	fmt.Println("timeout:", timeoutStrVal)

	timeoutIntVal := client.GetIntValue("timeout", 0)
	fmt.Println("timeout:", timeoutIntVal)
}

// 监听配置变化
func TestListener(t *testing.T) {
	c := &config.AppConfig{
		AppID:          AppID,
		Cluster:        Cluster,
		IP:             fmt.Sprintf("http://%s", ApolloAddr),
		NamespaceName:  Namespace,
		IsBackupConfig: true,
		Secret:         Secret,
	}
	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("初始化Apollo配置成功")

	listener := &CustomChangeListener{}
	listener.wg.Add(1)
	client.AddChangeListener(listener)

	listener.wg.Wait()

	// Output:
	// 初始化Apollo配置成功
	// map[timeout:0x1400010e780]
	// change key :  timeout , value : &{100 200 1}
	// application
}

type CustomChangeListener struct {
	wg sync.WaitGroup
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	// write your code here
	fmt.Println(changeEvent.Changes)
	for key, value := range changeEvent.Changes {
		fmt.Println("change key : ", key, ", value :", value)
	}
	fmt.Println(changeEvent.Namespace)
	c.wg.Done()
}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	// write your code here
}
