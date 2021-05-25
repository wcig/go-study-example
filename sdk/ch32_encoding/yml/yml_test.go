package yml

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

// yml文件编解码
// "gopkg.in/yaml.v2"

type Config struct {
	App *App `yaml:"app,omitempty" json:"app,omitempty"`
}

type App struct {
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Port    int    `yaml:"port,omitempty" json:"port,omitempty"`
	Profile string `yaml:"profile,omitempty" json:"profile,omitempty"`
}

// func Marshal(in interface{}) (out []byte, err error)
// yaml编码
func TestMarshal(t *testing.T) {
	cfg := &Config{
		App: &App{
			Name:    "go-app",
			Profile: "dev",
			Port:    80,
		},
	}

	b, err := yaml.Marshal(cfg)
	assert.Nil(t, err)
	fmt.Println(string(b))
	// output:
	// app:
	//  name: go-app
	//  port: 80

	err = os.WriteFile("tmp.yml", b, os.ModePerm)
	assert.Nil(t, err)
}

// func Unmarshal(in []byte, out interface{}) (err error)
// yaml解码
func TestUnmarshal(t *testing.T) {
	b, err := os.ReadFile("tmp.yml")
	assert.Nil(t, err)
	fmt.Println(string(b))
	// output:
	// app:
	//  name: go-app
	//  port: 80

	cfg := &Config{}
	err = yaml.Unmarshal(b, cfg)
	assert.Nil(t, err)
	fmt.Println(cfg.App)
	// output:
	// &{go-app 80}
}

// yaml编解码的第二种方式: 使用编解码器
// func NewEncoder(w io.Writer) *Encoder: NewEncoder返回一个写入w的新编码器。 使用后，应关闭编码器以将所有数据刷新到w。
// func NewDecoder(r io.Reader) *Decoder: NewDecoder返回一个新的解码器，该解码器从r读取。解码器引入自己的缓冲，并且可以从r中读取超出请求的YAML值的数据。
