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
	nameFile, err := os.Open("./template.txt")
	check(err)
	defer nameFile.Close()

	scanner := bufio.NewScanner(nameFile)

	messageStarted := false
	model := Model{comments: make([]string, 0), properties: make(map[string]string)}

	for scanner.Scan() {
		text := scanner.Text()
		splitInput := strings.Fields(text)
		if splitInput[0] == "M" {
			if messageStarted || len(splitInput) < 2 {
				panic(fmt.Errorf("Invalid document structure"))
			}
			model.name = splitInput[1]
			messageStarted = true
		} else if splitInput[0] == "F" {
			if !messageStarted || len(splitInput) < 3 {
				panic(fmt.Errorf("Invalid document structure"))
			}
			model.properties[splitInput[1]] = splitInput[2]
		} else if splitInput[0] == "E" {
			if !messageStarted && len(model.properties) > 0 {
				panic(fmt.Errorf("Invalid document structure"))
			}
			break
		} else if splitInput[0] == "#" {
			model.comments = append(model.comments, text[2:])
		} else {
			panic(fmt.Errorf("Invalid document structure"))
		}
	}
	name := "generated"
	javaFile, err := os.Create(fmt.Sprintf("./%s.java", name))
	check(err)
	defer javaFile.Close()

	goFile, err := os.Create(fmt.Sprintf("./%s.go", name))
	check(err)
	defer goFile.Close()
}

type Generator interface {
	AddComments()
	AddContent(name string, properties map[string]string)
}

type Model struct {
	name       string
	properties map[string]string
	comments   []string
}
