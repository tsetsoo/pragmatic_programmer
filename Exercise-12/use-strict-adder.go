package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var files []string

	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	check(err)
	for _, file := range files {
		if !strings.Contains(file, ".pl") {
			continue
		}

	}
	writeToFile("test2.pl")
}

func writeToFile(file string) {

	f, err := os.OpenFile(file, os.O_RDWR, os.ModeAppend)
	check(err)
	var offset int64 = 0
	offsetWrite := make([]byte, 0)
	appendToOffsetWrite := false
	shouldAppend := true

	r4 := bufio.NewReader(f)

	for {
		bytes, err := r4.ReadBytes('\n')

		// loop termination condition 1:  EOF.
		// this is the normal loop termination condition.
		if err == io.EOF {
			break
		}
		text := string(bytes)

		if strings.Contains(text, "\nuse strict;") {
			shouldAppend = false
			break
		}
		if !strings.HasPrefix(text, "#") {
			appendToOffsetWrite = true
		}
		if appendToOffsetWrite {
			offsetWrite = append(offsetWrite, bytes...)
		} else {
			offset += int64(len(bytes))
		}
	}
	if shouldAppend {
		toWrite := fmt.Sprintf("use strict;\n%s", string(offsetWrite))
		print(toWrite)
		_, err6 := f.WriteAt([]byte(toWrite), offset)
		check(err6)
	}

}
