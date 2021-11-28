package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/afero"
)

// Environment represents environment variables as pair key => value.
type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(fs afero.Fs, dir string) (Environment, error) {
	isDir, _ := afero.DirExists(fs, dir)
	if !isDir {
		return nil, fmt.Errorf("directory %s doesn't exist", dir)
	}

	files, err := afero.ReadDir(fs, dir)
	if err != nil {
		return nil, fmt.Errorf("can't read files in %s", dir)
	}

	env := make(Environment)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.Contains(file.Name(), "=") {
			return nil, fmt.Errorf("files in env dir mustn't contain '=' sign")
		}

		filename := fmt.Sprintf("%s/%s", dir, file.Name())
		content, err := fs.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("can't open file %s", filename)
		}

		value, err := io.ReadAll(content)
		if err != nil {
			return nil, fmt.Errorf("can't read file %s", filename)
		}

		firstLine := []byte(strings.Split(string(value), "\n")[0])
		stringValue := strings.TrimRight(string(bytes.Replace(firstLine, []byte{0x00}, []byte{'\n'}, 1024)), " ")

		env[file.Name()] = EnvValue{Value: stringValue, NeedRemove: len(value) == 0}
	}

	return env, nil
}
