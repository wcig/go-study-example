package ch29_os

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 目录获取
func TestGetDir(t *testing.T) {
	// TempDir
	fmt.Println(os.TempDir())

	// UserCacheDir
	cacheDir, err := os.UserCacheDir()
	fmt.Println(cacheDir, err)

	// UserConfigDir
	configDir, err := os.UserConfigDir()
	fmt.Println(configDir, err)

	// UserHomeDir
	homeDir, err := os.UserHomeDir()
	fmt.Println(homeDir, err)

	// Getwd: 当前工作目录，Chdir: 修改当前工作目录
	curDir, err := os.Getwd()
	fmt.Println("before:", curDir, err)
	err = os.Chdir("/Users/yangbo/Documents")
	assert.Nil(t, err)
	curDir, err = os.Getwd()
	fmt.Println("after:", curDir, err)

	// output:
	// /var/folders/vh/lks7z1qx6x90j10nwtm3njlw0000gn/T/
	// /Users/yangbo/Library/Caches <nil>
	// /Users/yangbo/Library/Application Support <nil>
	// /Users/yangbo <nil>
	// before: /Users/yangbo/Documents/MyGithub/go-study-example/sdk/ch29_os <nil>
	// after: /Users/yangbo/Documents <nil>
}

// 目录创建
func TestCreateDir(t *testing.T) {
	// Mkdir
	err := os.Mkdir("tmp.1", 0755)
	assert.Nil(t, err)

	// MkdirAll
	err = os.MkdirAll("tmp.2/a/b/c", 0755)
	assert.Nil(t, err)

	// MkdirTemp
	dir, err := os.MkdirTemp("", "tmp.3")
	fmt.Println(dir, err) // /var/folders/vh/lks7z1qx6x90j10nwtm3njlw0000gn/T/tmp.32077096737 <nil>
}

// 目录、文件删除
func TestRemoveDir(t *testing.T) {
	err := os.Remove("tmp.1")
	assert.Nil(t, err)

	err = os.RemoveAll("tmp.2")
	assert.Nil(t, err)
}

// 环境变量
func TestEnv(t *testing.T) {
	// Environ: 获取所有的环境变量
	envList := os.Environ()
	fmt.Println("env size:", len(envList))
	fmt.Println("last:", envList[len(envList)-1])

	// GetEnv: 获取指定环境变量
	goRoot := os.Getenv("GOROOT")
	fmt.Println("go root:", goRoot)

	// Clearenv: 清除所有环境变量
	os.Clearenv()
	envList = os.Environ()
	fmt.Println("env size:", len(envList))

	// Setenv: 设置环境变量
	err := os.Setenv("GOROOT", "/home/go")
	assert.Nil(t, err)
	goRoot = os.Getenv("GOROOT")
	fmt.Println("go root:", goRoot)

	// LookupEnv: 查询环境变量
	val, has := os.LookupEnv("GOROOT")
	fmt.Println(val, has)

	// : 取消设置环境变量
	err = os.Unsetenv("GOROOT")
	assert.Nil(t, err)
	val, has = os.LookupEnv("GOROOT")
	fmt.Println(val, has)

	// output:
	// env size: 46
	// last: CGO_ENABLED=1
	// go root: /usr/local/go
	// env size: 0
	// go root: /home/go
	// /home/go true
	//  false
}

// 字符串值替换
func TestExpand(t *testing.T) {
	// Expand: 以指定映射函数替换字符串文本的${var}或$var
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}
		return ""
	}
	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper))

	// Expand: 基于环境变量替换字符串文本的${var}或$var
	_ = os.Setenv("NAME", "gopher")
	_ = os.Setenv("BURROW", "/usr/gopher")
	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))
	// output:
	// Good morning, Gopher!
	// gopher lives in /usr/gopher.
}

// 用户uid，用户组gid
func TestUidGid(t *testing.T) {
	// Getuid
	uid := os.Getuid()
	fmt.Println("uid:", uid)

	// Geteuid
	euid := os.Geteuid()
	fmt.Println("uid:", euid)

	// Getgid
	gid := os.Getgid()
	fmt.Println("gid:", gid)

	// Getegid
	egid := os.Getegid()
	fmt.Println("gid:", egid)

	// output:
	// uid: 501
	// uid: 501
	// gid: 20
	// gid: 20
}
