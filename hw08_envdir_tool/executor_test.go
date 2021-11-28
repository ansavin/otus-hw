package main

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExit(t *testing.T) {
	t.Run("external error", func(t *testing.T) {
		err := exec.ExitError{}
		expected := err.ExitCode()

		actual := Exit(&err)
		require.Equal(t, actual, expected)
	})

	t.Run("internal error", func(t *testing.T) {
		actual := Exit(fmt.Errorf("something went wrong"))
		require.Equal(t, defaultErrorCode, actual)
	})
}

func TestRunCmd(t *testing.T) {
	t.Run("positive test - exit code 2", func(t *testing.T) {
		actual := RunCmd([]string{"grep", "-v", "123", "/path/to/nowhere"}, Environment{})
		require.Equal(t, actual, 2)
	})

	t.Run("positive test - exit code 0", func(t *testing.T) {
		actual := RunCmd([]string{"grep", "-q", "bar", "testdata/env/BAR"}, Environment{})
		require.Equal(t, actual, 0)
	})

	t.Run("negative test", func(t *testing.T) {
		actual := RunCmd([]string{"not existing command"}, Environment{})
		require.Equal(t, defaultErrorCode, actual)
	})
}
