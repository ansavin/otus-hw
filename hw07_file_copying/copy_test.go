package main

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	file := "test/file"
	emptyFile := "test/emptyFile"
	cantReadFile := "test/cantReadFile"
	cantWriteFile := "test/cantWriteFile"

	fakeFs := afero.NewMemMapFs()
	fakeFs.MkdirAll("test/", 0755)
	afero.WriteFile(fakeFs, file, []byte("simple_file"), 0660)
	afero.WriteFile(fakeFs, emptyFile, []byte(""), 0660)
	afero.WriteFile(fakeFs, cantReadFile, []byte("you_shall_not_pass"), 0110)
	afero.WriteFile(fakeFs, cantWriteFile, []byte("you_shall_not_pass"), 0220)

	negativeTests := []struct {
		name        string
		from        string
		to          string
		limit       int64
		offset      int64
		chunkSize   int64
		expectedErr error
	}{
		//{name: "no --from", from: "", to: emptyFile, limit: 1, offset: 1, chunkSize: 1},
		//{name: "no --to", from: file, to: "", limit: 1, offset: 1, chunkSize: 1},
		{name: "--from doesn't exist", from: "/path/to/nonexist/file", to: emptyFile, limit: 1, offset: 1, chunkSize: 1},
		{name: "--from read-protected", from: cantReadFile, to: emptyFile, limit: 1, offset: 1, chunkSize: 1},
		{name: "--to write-protected", from: file, to: cantWriteFile, limit: 1, offset: 1, chunkSize: 1},
	}
	for _, tc := range negativeTests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := Copy(fakeFs, tc.from, tc.to, tc.limit, tc.offset, tc.chunkSize)
			require.Error(t, err)
		})
	}
}
