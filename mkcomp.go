package main

import (
	"fmt"
	"mkcomp/constants"
	"os"
	"strings"
)

func FmtTemplate(template string, word string) string {
	wordCount := strings.Count(template, "%s")
	wordList := make([]interface{}, wordCount)
	for i := 0; i < wordCount; i++ {
		wordList[i] = word
	}
	return fmt.Sprintf(template, wordList...)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass a component name")
		return
	}
	component := strings.Title(os.Args[1])
	err := os.Mkdir(component, 0777)
	if err != nil {
		fmt.Printf("Couldn't create %s\n%s\n", component, err.Error())
	}
	for name, template := range constants.Files {
		f, err := os.Create(fmt.Sprintf("%s/%s", component, FmtTemplate(name, component)))
		if err != nil {
			fmt.Printf("Couldn't create %s/%s\n%s\n", component, name, err.Error())
		}
		f.WriteString(FmtTemplate(template, component))
	}
}
