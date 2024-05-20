package shima_park_agollo

import (
	"fmt"
	. "go-app/third/apollo/common"
	"testing"

	"github.com/shima-park/agollo"
	remote "github.com/shima-park/agollo/viper-remote"
	"github.com/spf13/viper"
)

// 快速开始
func TestQuickStart(t *testing.T) {
	options := []agollo.Option{
		agollo.Cluster(Cluster),
		agollo.PreloadNamespaces(Namespace),
		agollo.AutoFetchOnCacheMiss(),
		agollo.WithClientOptions(agollo.WithAccessKey(Secret)),
	}
	a, err := agollo.New(ApolloAddr, AppID, options...)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%s], [%s], [%s], [%s]\n",

		// 默认读取Namespace：application下key: foo的value
		a.Get("foo"),

		// 获取namespace为test.json的所有配置项
		a.GetNameSpace(Namespace),

		// 当key：foo不存在时，提供一个默认值bar
		a.Get("foo2", agollo.WithDefault("bar")),

		// 读取Namespace为other_namespace, key: foo的value
		a.Get("foo", agollo.WithNamespace("other_namespace")),
	)
	// Output:
	// [bar], [map[foo:bar timeout:100]], [bar], []

	// .agollo:
	// {"application":{"foo":"bar","timeout":"100"}}
}

// 监听配置变化
func TestListener(t *testing.T) {
	options := []agollo.Option{
		agollo.Cluster(Cluster),
		agollo.PreloadNamespaces(Namespace),
		agollo.AutoFetchOnCacheMiss(),
		agollo.WithClientOptions(agollo.WithAccessKey(Secret)),
	}
	a, err := agollo.New(ApolloAddr, AppID, options...)
	if err != nil {
		panic(err)
	}

	// Start后会启动goroutine监听变化，并更新agollo对象内的配置cache
	// 或者忽略错误处理直接 a.Start()
	errorCh := a.Start()
	watchCh := a.Watch()
	for {
		select {
		case pollErr := <-errorCh:
			fmt.Println("err:", pollErr)
		case resp := <-watchCh:
			fmt.Println(
				"Namespace:", resp.Namespace,
				"OldValue:", resp.OldValue,
				"NewValue:", resp.NewValue,
				"Error:", resp.Error,
			)
		}
	}

	// Output:
	// Namespace: application OldValue: map[foo:bar timeout:100] NewValue: map[foo:bar timeout:200] Error: <nil>
}

// 配合 viper
func TestWithViper(t *testing.T) {
	type Config struct {
		Timeout int    `mapstructure:"timeout" json:"timeout"`
		Foo     string `mapstructure:"foo" json:"foo"`
	}

	remote.SetAppID(AppID)
	remote.SetAgolloOptions(
		agollo.Cluster(Cluster),
		agollo.WithClientOptions(agollo.WithAccessKey(Secret)),
		agollo.AutoFetchOnCacheMiss(),
		agollo.FailTolerantOnBackupExists(),
	)

	app := viper.New()
	app.SetConfigType("prop")
	if err := app.AddRemoteProvider("apollo", ApolloAddr, Namespace); err != nil {
		panic(err)
	}
	if err := app.ReadRemoteConfig(); err != nil {
		panic(err)
	}
	var cfg Config
	if err := app.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	fmt.Println("cfg:", cfg)

	// Output:
	// cfg: {100 bar}
}
