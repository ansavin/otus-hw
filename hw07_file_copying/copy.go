package main

import (
	"fmt"
	"io"

	"github.com/spf13/afero"
)

func Copy(fs afero.Fs, fromPath, toPath string, limit, offset, chunkSize int64) (finError error) {
	if limit < 0 {
		limit = 0
	}

	if offset < 0 {
		offset = 0
	}

	src, err := fs.Open(fromPath)
	defer func() {
		err := src.Close()
		if err != nil {
			finError = fmt.Errorf("can`t close file %s: %s", fromPath, err)
		}
	}()
	if err != nil {
		return fmt.Errorf("can`t open file %s: %s", fromPath, err)
	}

	fileStats, err := src.Stat()
	if err != nil {
		return fmt.Errorf("can`t stat file %s: %s", fromPath, err)
	}

	fileSize := fileStats.Size()
	if fileSize < offset {
		return fmt.Errorf("offset exceeds file size (%d bytes)", fileSize)
	}
	if limit > fileSize {
		limit = fileSize
	}

	dst, err := fs.Create(toPath)
	defer func() {
		err := dst.Close()
		if err != nil {
			finError = fmt.Errorf("can`t close file %s: %s", fromPath, err)
		}
	}()
	if err != nil {
		return fmt.Errorf("can`t create file %s: %s", toPath, err)
	}

	buf := make([]byte, chunkSize)
	isLastChank := false
	var initialOffset int64 = offset
	var totalRead int64 = 0
	var percent int64

	for offset < fileSize {
		read, err := src.ReadAt(buf, offset)
		if err != io.EOF && err != nil {
			return fmt.Errorf("can`t read from file %s: %s", fromPath, err)
		}
		if err == io.EOF {
			isLastChank = true
		}

		fmt.Printf("!!!! %#v, %#v, %d, %d, %d\n", buf, buf[:read], read, totalRead, limit)
		if limit > 0 && totalRead+int64(read) > limit {
			fmt.Println("foo")
			_, err = dst.WriteAt(buf[:(limit-totalRead)], offset-initialOffset)
			isLastChank = true
		} else {
			fmt.Println("bar")
			newbuf := buf[:read]
			_, err = dst.WriteAt(newbuf, offset-initialOffset)
		}

		if err != nil {
			return fmt.Errorf("can`t write to file %s: %s", toPath, err)
		}

		if limit > 0 && limit < fileSize {
			percent = 100 * totalRead / limit
		} else {
			percent = 100 * totalRead / (fileSize - initialOffset)
		}
		fmt.Printf("Copying: %.1d%% complete\n", percent)
		//time.Sleep(1 * time.Second)

		offset += int64(read)
		totalRead += int64(read)

		if isLastChank {
			break
		}
	}

	return nil

}
