package main

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewOsFs()

	code := 1
	defer func() {
		os.Exit(code)
	}()

	args := os.Args

	if len(args) < 3 {
		fmt.Println("expects 2 or more args")
		return
	}

	envDirPath := args[1]

	envs, err := ReadDir(fs, envDirPath)
	if err != nil {
		fmt.Println("error reading env dir", err)
		return
	}

	code = RunCmd(args[2:], envs)
}
