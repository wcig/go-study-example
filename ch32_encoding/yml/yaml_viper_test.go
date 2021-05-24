package yml

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// github.com/spf13/viper
// 可用于yaml文件的编解码和合并

// 读取yaml配置文件
func TestViperRead(t *testing.T) {
	src := `app:
  name: go-app
  port: 80
  profile: dev`

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(strings.NewReader(src))
	assert.Nil(t, err)
	fmt.Println(viper.Get("app")) // map[name:go-app port:80 profile:dev]
}

// 读取并写入yaml文件
func TestViperWrite(t *testing.T) {
	src := `app:
  name: go-app
  port: 80
  profile: dev`

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(strings.NewReader(src))
	assert.Nil(t, err)

	err = viper.WriteConfigAs("tmp.yaml")
	assert.Nil(t, err)
}

// 合并yaml文件
func TestYamlMerge(t *testing.T) {
	base := `app:
  name: go-app
  port: 8080
  profile: dev`

	dev := `app:
  port: 80`

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(strings.NewReader(base))
	assert.Nil(t, err)

	err = viper.MergeConfig(strings.NewReader(dev))
	assert.Nil(t, err)

	_ = viper.WriteConfigAs("tmp.yaml")
	// tmp.yaml:
	// app:
	//  name: go-app
	//  port: 80
	//  profile: dev
}

// 解析为结构体
func TestViperUnmarshal(t *testing.T) {
	src := `app:
  name: go-app
  port: 80
  profile: dev`

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(strings.NewReader(src))
	assert.Nil(t, err)

	var cfg Config
	err = viper.Unmarshal(&cfg)
	assert.Nil(t, err)
	printJson(&cfg) // {"app":{"name":"go-app","port":80,"profile":"dev"}}
}

func printJson(i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
