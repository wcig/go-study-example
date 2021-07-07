package exec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
