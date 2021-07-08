package exec

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestRunCmd(t *testing.T) {
	cmdStr := "sleep 1 && echo ok"
	output, err := RunCmd(cmdStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("output:", output) // output: ok\n
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
