package main

import (
	"fmt"
	"os"
)

/**
*Gestion du traitement du fichier d'entrée et d sortie
*/
func main() {
	// Acess to the Arguments of Program
	if len(os.Args[1:]) != 2 {
		fmt.Println("Two Arguments =>  <input file> <output file> !!!")
	}
	InputFile := os.Args[1:][0]
	OutputFile := os.Args[1:][1]

	// Read the content of Input File ex : sample.txt
	ContentInputFile, err := os.ReadFile(string(InputFile))
	if err != nil {
		fmt.Println(err)
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
