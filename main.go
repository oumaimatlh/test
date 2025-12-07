package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Acess to the Arguments of Program
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

	ContentInputFile, err := os.ReadFile(InputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	
	// Completion/Editing/Auto-correction
	EditingContent := EditedText(string(ContentInputFile))

	// write the edited content  to the output file ex : result.txt
	err = os.WriteFile(OutputFile, []byte(EditingContent), 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
