package ioutil

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

// ReadAll
func TestReadAll(t *testing.T) {
	str := "hello world."
	buf := bytes.NewBufferString(str)
	b, err := ioutil.ReadAll(buf)
	assert.True(t, err == nil)
	assert.True(t, string(b) == str)
}

// ReadFile
func TestReadFile(t *testing.T) {
	str := "hello world."
	filename := "temp"
	b, err := ioutil.ReadFile(filename)
	assert.True(t, err == nil)
	assert.True(t, string(b) == str)
}

// WriteFile
func TestWriteFile(t *testing.T) {
	str := "hello world."
	filename := "temp"
	err := ioutil.WriteFile(filename, []byte(str), 0777)
	assert.True(t, err == nil)
}

// ReadDir
func TestReadDir(t *testing.T) {
	dirname := "../ioutil"
	rangeDir(dirname)
}

func rangeDir(dirname string, prefixes ...string) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.Join(prefixes, "") + dirname)
	prefixes = append(prefixes, "\t")
	p := strings.Join(prefixes, "")
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			rangeDir(fileInfo.Name(), p)
		} else {
			fmt.Println(p + fileInfo.Name())
		}
	}
}
