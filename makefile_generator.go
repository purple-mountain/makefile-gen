package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput(prompt string) string {
	var s string
	fmt.Printf(prompt)
	fmt.Scanln(&s)
	return strings.TrimSpace(s)
}

func main() {
	programName := getInput("Program name: ")
	language := getInput("Language (cpp or c): ")
	outputPath := getInput("Output path (leave empty for current folder): ")
	warningFlags := getInput("Enable warning flags (y/n): ")
	additionalFlags := getInput("Additional flags (leave empty if none): ")

	if len(outputPath) != 0 && strings.Contains(outputPath[:len(outputPath)-1], "/") {
		log.Fatalln("Invalid outputh path")
	}

	if language != "cpp" && language != "c" {
		log.Fatalln("Invalid language")
	}

	if warningFlags == "y" {
		warningFlags = "-Wall -Weffc++ -Wextra -Wconversion -Wsign-conversion"
	} else {
		warningFlags = ""
	}

	if outputPath == "." || (!strings.Contains(outputPath, "/") && len(outputPath) != 0) {
		outputPath += "/"
	}

	output := fmt.Sprintf("%s: %s.%s\n\t\tg++ -o %s%s %s.%s %s %s", programName, programName, language, outputPath, programName, programName, language, warningFlags, additionalFlags)

	err := os.WriteFile("makefile", []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}
