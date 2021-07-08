package exec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestLookPath(t *testing.T) {
	path, err := exec.LookPath("gofmt")
	fmt.Println(path, err) // /usr/local/go/bin/gofmt <nil>
}

func TestCommand(t *testing.T) {
	cmd := exec.Command("echo", "ok")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String()) // ok
}

func TestCommandDir(t *testing.T) {
	cmd := exec.Command("ls", "-l")
	cmd.Dir = "/" // 设置工作路径

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
	// output:
	// total 9
	// drwxrwxr-x+ 70 root  admin  2240 Jul  3 11:21 Applications
	// drwxr-xr-x  65 root  wheel  2080 Aug  9  2020 Library
	// drwxr-xr-x@  8 root  wheel   256 Sep 30  2019 System
	// drwxr-xr-x   6 root  admin   192 Feb 22  2020 Users
	// drwxr-xr-x   5 root  wheel   160 Jul  8 20:29 Volumes
	// drwxr-xr-x@ 38 root  wheel  1216 Sep 30  2019 bin
	// drwxr-xr-x   2 root  wheel    64 Aug 25  2019 cores
	// dr-xr-xr-x   3 root  wheel  4545 Jul  8 20:25 dev
	// lrwxr-xr-x@  1 root  admin    11 Oct 20  2019 etc -> private/etc
	// lrwxr-xr-x   1 root  wheel    25 Jul  8 20:25 home -> /System/Volumes/Data/home
	// drwxr-xr-x   2 root  wheel    64 Aug 25  2019 opt
	// drwxr-xr-x   6 root  wheel   192 Sep 30  2019 private
	// drwxr-xr-x@ 64 root  wheel  2048 Oct 20  2019 sbin
	// lrwxr-xr-x@  1 root  admin    11 Oct 20  2019 tmp -> private/tmp
	// drwxr-xr-x@ 11 root  wheel   352 Oct 20  2019 usr
	// lrwxr-xr-x@  1 root  admin    11 Oct 20  2019 var -> private/var
	//
}

func TestCommandPath(t *testing.T) {
	cmd := exec.Command("/usr/local/go/bin/go", "version")
	fmt.Println("cmd path:", cmd.Path)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("output:", string(output))
	// output:
	// cmd path: /usr/local/go/bin/go
	// output: go version go1.16.4 darwin/amd64
	//
}

func TestCommandEnv(t *testing.T) {
	cmd := exec.Command("/bin/sh", "-c", "echo $my_var")
	cmd.Env = []string{"my_var=ok"}

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("output:", string(output))
	// output:
	// output: ok
	//
}

func TestCommandProcess(t *testing.T) {
	cmd := exec.Command("/bin/sh", "-c", "sleep 10 && echo ok")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("cmd process pid:", cmd.Process.Pid)
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("cmd process exit code:", cmd.ProcessState.ExitCode())
	// output:
	// cmd process pid: 6349
	// cmd process exit code: 0
}

func TestCommandContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sleep", "5")
	err := cmd.Run()
	fmt.Println(err)
	// output: 100毫秒后失败，睡五秒被中断
	// signal: killed
}

func TestCombinedOutput(t *testing.T) {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
	// output:
	// stdout
	// stderr
	//
}

func TestCommandStdin(t *testing.T) {
	stdin, err := os.Open("exec_example_test.go")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("wc", "-l")
	cmd.Stdin = stdin

	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("output:", string(b)) // output:      247\n
}

func TestOutput(t *testing.T) {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	// output:
	// Wed Jul  7 23:26:41 CST 2021
	//
}

func TestRun(t *testing.T) {
	cmd := exec.Command("sleep", "1")
	err := cmd.Run()
	fmt.Println(err) // <nil>
}

func TestStart(t *testing.T) {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	// output:
	// 2021/07/07 23:28:23 Waiting for command to finish...
	// 2021/07/07 23:28:28 Command finished with error: <nil>
}

func TestStderrPipe(t *testing.T) {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("stderr: %s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	// output:
	// stderr: stderr
	//
}

func TestStdinPipe(t *testing.T) {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
	// output:
	// values written to stdin are passed to cmd's standard input
}

func TestStdoutPipe(t *testing.T) {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	// output:
	// Bob is 32 years old
}

func TestProcessStatus(t *testing.T) {
	cmd := exec.Command("curl", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")
	var stdoutProcessStatus bytes.Buffer
	cmd.Stdout = io.MultiWriter(ioutil.Discard, &stdoutProcessStatus)
	done := make(chan struct{})
	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		for {
			select {
			case <-done:
				return
			case <-tick.C:
				log.Printf("downloaded: %d", stdoutProcessStatus.Len())
			}
		}
	}()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call Run(): %v", err)
	}
	close(done)
}
