package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// defaultErrorCode represents return code when tool meets an internal error.
const defaultErrorCode = 146

// Exit stops tool and returns correct return code.
func Exit(err error) int {
	exitError := &exec.ExitError{}

	if errors.As(err, &exitError) {
		return exitError.ExitCode()
	}
	return defaultErrorCode
}

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for k, e := range env {
		if e.NeedRemove {
			os.Unsetenv(k)
		}
		os.Setenv(k, e.Value)
	}

	c := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	stdout, err := c.StdoutPipe()
	if err != nil {
		fmt.Println("cant read stdout:", err)
		return Exit(err)
	}

	stderr, err := c.StderrPipe()
	if err != nil {
		fmt.Println("cant read stderr:", err)
		return Exit(err)
	}

	err = c.Start()
	if err != nil {
		fmt.Println(err)
		return Exit(err)
	}

	errorBuf, _ := io.ReadAll(stderr)
	os.Stderr.WriteString(string(errorBuf))

	outBuf, _ := io.ReadAll(stdout)
	fmt.Print(string(outBuf))

	if err := c.Wait(); err != nil {
		return Exit(err)
	}
	return 0
}
