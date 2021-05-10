package filepath

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// filepath: 目录相关工具包 (基于不同平台兼容性,建议优先使用filepath包而不是path包)

// filepath.join(): 目录拼接
func TestJoin(t *testing.T) {
	type JoinTest struct {
		Dirs     []string
		Expected string
	}
	jts := []*JoinTest{
		{[]string{"/tmp", "log"}, "/tmp/log"},         // 绝对路径拼接
		{[]string{"tmp/data", "log"}, "tmp/data/log"}, // 相对路径拼接
		{[]string{"/tmp", ""}, "/tmp"},                // 拼接空字符串
		{[]string{".", "tmp"}, "tmp"},                 // 拼接'.'
	}
	for _, jt := range jts {
		assert.Equal(t, jt.Expected, filepath.Join(jt.Dirs...))
	}
}

// filepath.Split(): 分割为文件目录和文件名
func TestSplit(t *testing.T) {
	type SplitTest struct {
		Path         string
		ExpectedDir  string
		ExpectedFile string
	}
	sts := []*SplitTest{
		{"/tmp/log/info.log", "/tmp/log/", "info.log"},
		{"tmp/log/info.log", "tmp/log/", "info.log"},
		{"/tmp/log/", "/tmp/log/", ""},
		{"/tmp/log//info.log", "/tmp/log//", "info.log"},
	}
	for _, st := range sts {
		dir, file := filepath.Split(st.Path)
		assert.Equal(t, st.ExpectedDir, dir)
		assert.Equal(t, st.ExpectedFile, file)
	}
}

// filepath.SplitList(): 分割特定os的目录为子目录列表,常用于PATH,GOPATH
func TestSplitList(t *testing.T) {
	type SplitListTest struct {
		Path     string
		Expected []string
	}
	sts := []*SplitListTest{
		{"/tmp/log/info.log", []string{"/tmp/log/info.log"}},
		{"/a/b/c:/usr/bin", []string{"/a/b/c", "/usr/bin"}},
		{"", []string{}},
	}
	for _, st := range sts {
		fmt.Println(filepath.SplitList(st.Path))
		assert.Equal(t, st.Expected, filepath.SplitList(st.Path))
	}
}

// filepath.Dir(): 获取文件目录(去除目录的文件名)
func TestDir(t *testing.T) {
	type DirTest struct {
		Path     string
		Expected string
	}
	dts := []*DirTest{
		{"/tmp/log/info.log", "/tmp/log"}, // 绝对路径
		{"tmp/log/info.log", "tmp/log"},   // 相对路径
	}
	for _, dt := range dts {
		fmt.Println(filepath.Dir(dt.Path))
	}
}

// filepath.Base(): 获取文件目录的文件名(去除文件路径)
func TestBase(t *testing.T) {
	type BaseTest struct {
		Path     string
		Expected string
	}
	dts := []*BaseTest{
		{"/tmp/log/info.log", "info.log"}, // 绝对路径
		{"tmp/log/info.log", "info.log"},  // 相对路径
		{"/tmp/log/info", "info"},         // 文件没有扩展名
	}
	for _, dt := range dts {
		assert.Equal(t, dt.Expected, filepath.Base(dt.Path))
	}
}

// filepath.Ext(): 获取文件扩展名
func TestExt(t *testing.T) {
	type ExtTest struct {
		Path     string
		Expected string
	}
	ets := []*ExtTest{
		{"/tmp/log/info.log", ".log"},
		{"/tmp/log/info", ""},
	}
	for _, et := range ets {
		assert.Equal(t, et.Expected, filepath.Ext(et.Path))
	}
}

// file.Abs(): 获取文件的绝对路径,如果不是绝对路径则将当前工作路径拼接得到绝对路径(注意硬链接不一定正确)
func TestAbs(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.Nil(t, err)

	type AbsTest struct {
		Path     string
		Expected string
	}
	ats := []*AbsTest{
		{"/tmp/log/info.log", "/tmp/log/info.log"},
		{"info.log", filepath.Join(currentDir, "info.log")},
		{"log/info.log", filepath.Join(currentDir, "log/info.log")},
	}
	for _, at := range ats {
		val, err := filepath.Abs(at.Path)
		assert.Nil(t, err)
		assert.Equal(t, at.Expected, val)
	}
}

// filepath.IsAbs(): 目录是否绝对路径
func TestIsAbs(t *testing.T) {
	type IsAbs struct {
		Path     string
		Expected bool
	}
	its := []*IsAbs{
		{"/tmp/log/info.log", true},
		{"info.log", false},
		{"../info.log", false},
	}
	for _, it := range its {
		assert.Equal(t, it.Expected, filepath.IsAbs(it.Path))
	}
}

// filepath.Clean(): 整理路径
// 1.去除多余的分隔符
// 2.剔除每一个'.'(当前目录)
// 3.剔除路径内的'..'(父目录)和前面的'..'
// 4.剔除开始于根路径的'..', 即替换'/..'为'/'
// 5.空字符串则转换为'.'
func TestClean(t *testing.T) {
	type CleanTest struct {
		Path     string
		Expected string
	}
	cts := []*CleanTest{
		{"/tmp/log/info.log", "/tmp/log/info.log"},
		{"/tmp/log//info.log", "/tmp/log/info.log"},
		{"/tmp/log/./info.log", "/tmp/log/info.log"},
		{"./info.log", "info.log"},
		{"/tmp/log/../log/info.log", "/tmp/log/info.log"},
		{"../info.log", "../info.log"},
		{"/../info.log", "/info.log"},
		{"", "."},
		{"/tmp/log/", "/tmp/log"},
	}
	for _, ct := range cts {
		assert.Equal(t, ct.Expected, filepath.Clean(ct.Path))
	}
}

// filepath.Rel(): 计算目标目录对基础目录的相对路径
func TestRel(t *testing.T) {
	type RelTest struct {
		BasePath   string
		TargetPath string
		Expected   string
	}
	rts := []*RelTest{
		// {"", "/tmp/log/info", ""}, // error
		// {"tmp", "/tmp/log/info.log", ""}, // error
		{"/tmp", "/tmp/log/info.log", "log/info.log"},
		{"tmp", "tmp/log/info.log", "log/info.log"},
		{"/a", "/b/c", "../b/c"},
	}
	for _, rt := range rts {
		val, err := filepath.Rel(rt.BasePath, rt.TargetPath)
		fmt.Println("val:", val)
		fmt.Println("err:", err)
		assert.Equal(t, rt.Expected, val)
	}
}

// filepath.EvalSymlinks(): 获取符号链接指定的文件名(符号链接文件需与源文件在同一目录)
func Test1(t *testing.T) {
	var err error
	srcFileName := "srcFile"
	err = ioutil.WriteFile(srcFileName, []byte("test"), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(srcFileName)

	symlinkDir := "symlink"
	err = os.Mkdir(symlinkDir, os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(symlinkDir)

	newFileName1 := filepath.Join(symlinkDir, "newFile")
	err = os.Symlink(srcFileName, newFileName1)
	assert.Nil(t, err)
	defer os.Remove(newFileName1)

	newFileName2 := "newFile2"
	err = os.Symlink(srcFileName, newFileName2)
	assert.Nil(t, err)
	defer os.Remove(newFileName2)

	fmt.Println(filepath.EvalSymlinks(newFileName1))
	fmt.Println(os.Readlink(newFileName1))

	fmt.Println(filepath.EvalSymlinks(newFileName2))
	fmt.Println(os.Readlink(newFileName2))
}

// output:
//  lstat symlink/srcFile: no such file or directory
// srcFile <nil>
// srcFile <nil>
// srcFile <nil>

// filepath.Match(): 文件是否与pattern匹配
func TestMatch(t *testing.T) {
	var err error
	dirs := []string{"a", "a/b", "a/c"}
	files := []string{"a/b/f1.txt", "a/c/f2.txt", "a/f3.txt"}
	for _, dir := range dirs {
		err = os.MkdirAll(dir, 0777)
		assert.Nil(t, err)
	}
	for _, file := range files {
		err = ioutil.WriteFile(file, []byte(""), 0755)
		assert.Nil(t, err)
	}
	baseDir := dirs[0]
	defer os.RemoveAll(baseDir)

	type MatchTest struct {
		Pattern  string
		Name     string
		Expected bool
	}
	mts := []*MatchTest{
		{"a/b/*.txt", "a/b/f1.txt", true},

		{"a/*/*.txt", "a/b/f1.txt", true},
		{"a/*/*.txt", "a/c/f2.txt", true},

		{"a/*", "a/b", true},
		{"a/*", "a/c", true},
		{"a/*", "a/f3.txt", true},
		{"a/*", "a/b/f1.txt", false},
	}

	for _, mt := range mts {
		matched, err := filepath.Match(mt.Pattern, mt.Name)
		assert.Nil(t, err)
		assert.Equal(t, mt.Expected, matched)
	}
}

// filepath.Glob(): 返回所有匹配文件的名称,没有则返回nil
func TestGlob(t *testing.T) {
	var err error
	dirs := []string{"a", "a/b", "a/c"}
	files := []string{"a/b/f1.txt", "a/c/f2.txt", "a/f3.txt"}
	for _, dir := range dirs {
		err = os.MkdirAll(dir, 0777)
		assert.Nil(t, err)
	}
	for _, file := range files {
		err = ioutil.WriteFile(file, []byte(""), 0755)
		assert.Nil(t, err)
	}
	baseDir := dirs[0]
	defer os.RemoveAll(baseDir)

	type GlobTest struct {
		Pattern  string
		Expected []string
	}
	gts := []*GlobTest{
		{"a/b/*.txt", []string{"a/b/f1.txt"}},
		{"a/*/*.txt", []string{"a/b/f1.txt", "a/c/f2.txt"}},
		{"a/*", []string{"a/b", "a/c", "a/f3.txt"}},
	}
	for _, gt := range gts {
		matches, err := filepath.Glob(gt.Pattern)
		assert.Nil(t, err)
		assert.Equal(t, gt.Expected, matches)
	}
}

// filepath.Walk(): 遍历目录下所有文件和目录
// a
// ├── b
// │   └── f1
// ├── c
// │   └── f2
// └── f3
func TestWalk(t *testing.T) {
	var err error
	dirs := []string{"a", "a/b", "a/c"}
	files := []string{"a/b/f1", "a/c/f2", "a/f3"}
	for _, dir := range dirs {
		err = os.MkdirAll(dir, 0777)
		assert.Nil(t, err)
	}
	for _, file := range files {
		err = ioutil.WriteFile(file, []byte(""), 0755)
		assert.Nil(t, err)
	}
	baseDir := dirs[0]
	defer os.RemoveAll(baseDir)

	err = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("path:%7s, isDir:%6t, name:%3s, err: %v\n", path, info.IsDir(), info.Name(), err)
		return err
	})
	assert.Nil(t, err)
}

// output:
// path:      a, isDir:  true, name:  a, err: <nil>
// path:    a/b, isDir:  true, name:  b, err: <nil>
// path: a/b/f1, isDir: false, name: f1, err: <nil>
// path:    a/c, isDir:  true, name:  c, err: <nil>
// path: a/c/f2, isDir: false, name: f2, err: <nil>
// path:   a/f3, isDir: false, name: f3, err: <nil>

// windows平台3函数
func TestWin(t *testing.T) {
	filepath.VolumeName(`C:\foo\bar`)
	filepath.FromSlash(`/foo/bar`)
	filepath.ToSlash(`C:\foo\bar`)
}
