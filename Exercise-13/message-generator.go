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
	fmt.Println("Template parsing completed successfully.")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter target file (END if finished):\n")
		text, err := reader.ReadString('\n')
		check(err)
		text = strings.TrimSpace(text)
		if text == "END" {
			break
		}
		splitInput := strings.Split(text, ".")
		extension := splitInput[len(splitInput)-1]
		file, err := os.Create(fmt.Sprintf("./%s", text))
		check(err)
		defer file.Close()
		generator, err := contentGenerator(extension)
		check(err)
		wrapper := GeneratorWrapper{
			generator: generator,
			model:     model,
		}
		fmt.Println(wrapper.Content())
		_, err = file.WriteString(wrapper.Content())
		check(err)
	}
	fmt.Println("Finished")
}

func contentGenerator(extension string) (Generator, error) {
	switch extension {
	case "go":
		return GoGenerator{}, nil
	case "java":
		return JavaGenerator{}, nil
	case "c":
		return CGenerator{}, nil
	case "pp":
		return PascalGenerator{}, nil
	default:
		return nil, fmt.Errorf("Does not support extension: %s", extension)
	}
}

type Generator interface {
	Comments(comments []string) string
	Struct(name string, fields map[string]string) string
}

type GoGenerator struct {
}

func (g GoGenerator) Comments(comments []string) string {
	return cStyleComments(comments)
}

func (g GoGenerator) Struct(name string, fields map[string]string) string {
	return ""
}

type JavaGenerator struct {
}

func (g JavaGenerator) Comments(comments []string) string {
	return cStyleComments(comments)
}
func (g JavaGenerator) Struct(name string, fields map[string]string) string {
	return ""
}

type CGenerator struct {
}

func (g CGenerator) Comments(comments []string) string {
	return cStyleComments(comments)
}
func (g CGenerator) Struct(name string, fields map[string]string) string {
	return ""
}

type PascalGenerator struct {
}

func (g PascalGenerator) Comments(comments []string) string {
	return ""
}
func (g PascalGenerator) Struct(name string, fields map[string]string) string {
	return ""
}

type GeneratorWrapper struct {
	generator Generator
	model     Model
}

func (g GeneratorWrapper) Content() string {
	return g.generator.Comments(g.model.comments) + g.generator.Struct(g.model.name, g.model.properties)
}

func cStyleComments(comments []string) string {
	var sb strings.Builder
	for _, str := range comments {
		sb.WriteString(fmt.Sprintf("// %s\n", str))
	}
	return sb.String()
}

type Model struct {
	name       string
	properties map[string]string
	comments   []string
}
