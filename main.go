package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) (string, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return "", err
	}
	const maxSize = 100 * 1024
	if info.Size() > maxSize {
		return "", fmt.Errorf("input file is too large")
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: go run .  <input file> <output file> !!!")
		return
	}

	InputFile := os.Args[1:][0]
	OutputFile := os.Args[1:][1]

	if !strings.HasSuffix(strings.ToLower(InputFile), ".txt") {
		fmt.Println("Error: input file must have .txt extension")
		return
	}

	if !strings.HasSuffix(strings.ToLower(OutputFile), ".txt") {
		fmt.Println("Error: output file must have .txt extension")
		return
	}

	ContentInputFile, err := readFile(InputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	EditingContent := EditedText(string(ContentInputFile))

	err = os.WriteFile(OutputFile, []byte(EditingContent), 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
