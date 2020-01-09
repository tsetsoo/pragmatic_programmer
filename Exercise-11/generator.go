package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	textFile, err := os.Open("./file.txt")
	check(err)
	defer textFile.Close()

	scanner := bufio.NewScanner(textFile)

	scanner.Scan()
	text := scanner.Text()

	cHeader, err1 := os.Create(fmt.Sprintf("./%s.h", text))
	check(err1)
	defer cHeader.Close()

	cSource, err2 := os.Create(fmt.Sprintf("./%s.c", text))
	check(err2)
	defer cSource.Close()

	_, err3 := cHeader.WriteString(fmt.Sprintf("extern const char* %s_names[];\ntypedef enum {\n", strings.ToUpper(text)))
	check(err3)

	_, err4 := cSource.WriteString(fmt.Sprintf("const char* %s_names[] = {\n", strings.ToUpper(text)))
	check(err4)

	for scanner.Scan() {
		text := scanner.Text()

		_, err1 := cHeader.WriteString(fmt.Sprintf("  %s,\n", text))
		check(err1)

		_, err2 := cSource.WriteString(fmt.Sprintf("  \"%s\",\n", text))
		check(err2)
	}
	_, err5 := cHeader.WriteString(fmt.Sprintf("} %s;", strings.ToUpper(text)))
	check(err5)

	_, err6 := cSource.WriteString("};")
	check(err6)
}
