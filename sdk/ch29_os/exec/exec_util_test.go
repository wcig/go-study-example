package exec

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestRunCmd(t *testing.T) {
	cmdStr := "sleep 1 && echo ok"
	output, err := RunCmd(cmdStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("output:", output) // output: ok\n
}

func TestRunCmdWithTimeout(t *testing.T) {
	{
		cmdStr := "sleep 3"
		_, err := RunCmdWithTimeout(cmdStr, 5*time.Second)
		if err != nil {
			panic(err)
		}
	}
	{
		cmdStr := "sleep 3"
		_, err := RunCmdWithTimeout(cmdStr, 1*time.Second)
		if err == nil {
			panic("cmd run no timeout err")
		}
	}

}

// ------------------------------------------------------------------ //

func RunCmd(cmdStr string) (output string, err error) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("cmd: %s, stdout: %s, stderr: %s, err: %v", cmd, stdout.String(), stderr.String(), err)
		return stdout.String(), err
	}
	return stdout.String(), nil
}

func RunCmdWithTimeout(cmdStr string, timeout time.Duration) (output string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", cmdStr)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("cmd: %s, stdout: %s, stderr: %s, err: %v", cmd, stdout.String(), stderr.String(), err)
		return stdout.String(), err
	}
	return stdout.String(), nil
}
