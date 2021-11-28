package main

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	prefix := "test"
	simpleFile := "FOO"
	simpleFileContent := []byte("123")

	emptyFile := "BAR"
	emptyFileContent := []byte{}

	multilineFile := "BAZ"
	multilineFileContent := []byte("123\n231\n312")

	tralingSpaceFile := "FIZZ"
	tralingSpaceFileContent := []byte("123   ")

	multulineVarFile := "BUZZ"
	multulineVarFileContent := []byte{0x31, 0x00, 0x32, 0x00, 0x33}

	fakeFs := afero.NewMemMapFs()
	fakeFs.MkdirAll(prefix+"/", 0o775)
	afero.WriteFile(fakeFs, prefix+"/"+simpleFile, simpleFileContent, 0o660)
	afero.WriteFile(fakeFs, prefix+"/"+emptyFile, emptyFileContent, 0o660)
	afero.WriteFile(fakeFs, prefix+"/"+multilineFile, multilineFileContent, 0o660)
	afero.WriteFile(fakeFs, prefix+"/"+tralingSpaceFile, tralingSpaceFileContent, 0o660)
	afero.WriteFile(fakeFs, prefix+"/"+multulineVarFile, multulineVarFileContent, 0o660)

	t.Run("positive test", func(t *testing.T) {
		expected := make(Environment)

		expected[simpleFile] = EnvValue{
			Value:      string(simpleFileContent),
			NeedRemove: false,
		}
		expected[emptyFile] = EnvValue{
			Value:      "",
			NeedRemove: true,
		}
		expected[multilineFile] = EnvValue{
			Value:      "123",
			NeedRemove: false,
		}
		expected[tralingSpaceFile] = EnvValue{
			Value:      "123",
			NeedRemove: false,
		}
		expected[multulineVarFile] = EnvValue{
			Value:      "1\n2\n3",
			NeedRemove: false,
		}

		actual, err := ReadDir(fakeFs, prefix)
		require.NoError(t, err)

		require.Equal(t, expected, actual)
	})

	t.Run("negative test", func(t *testing.T) {
		fakeFs := afero.NewMemMapFs()

		_, err := ReadDir(fakeFs, "/path/to/nowhere")
		require.Error(t, err)
	})
}
