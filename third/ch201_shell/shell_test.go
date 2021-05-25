package ch201_shell

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 执行shell命令
func TestShell(t *testing.T) {
	cmdStr := "echo ok"
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	bytes, err := cmd.Output()
	fmt.Println(string(bytes))
	fmt.Println(err)
}

// 输出:
// ok
//
// <nil>

// 记录stdout,stderr
func TestGetStd(t *testing.T) {
	cmdStr := "echo ok"
	cmd := exec.Command("/bin/sh", "-c", cmdStr)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()
	fmt.Println("stdout:", stdoutBuf.String())
	fmt.Println("stderr:", stderrBuf.String())
	fmt.Println("err:", err)
}

// 输出:
// stdout: ok
//
// stderr:
// err: <nil>

// 执行文件
func TestExecFile(t *testing.T) {
	fileName := "./demo.sh"
	cmdStr := "date"
	err := ioutil.WriteFile(fileName, []byte(cmdStr), 0777)
	assert.Nil(t, err)
	defer os.Remove(fileName)

	cmd := exec.Command("/bin/sh", "-c", fileName)
	b, err := cmd.Output()
	assert.Nil(t, err)
	fmt.Println(string(b))
}

// 添加超时时间
func TestExpire(t *testing.T) {
	fileName := "./demo.sh"
	cmdStr := "sleep 100s;\ndate"
	err := ioutil.WriteFile(fileName, []byte(cmdStr), 0777)
	assert.Nil(t, err)
	defer os.Remove(fileName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", fileName)
	b, err := cmd.Output()
	fmt.Println(string(b))
	fmt.Println(err)
	// err = cmd.Run()
	// fmt.Println(err)

	/*ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	go func() {
		cmd := exec.CommandContext(ctx, "/bin/sh", "-c", fileName)
		b, err := cmd.Output()
		fmt.Println(string(b))
		fmt.Println(err)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}*/
}
