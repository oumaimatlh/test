package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EditedTex(Text string) []string {
	// Auto-Correction => Supprimer les espaces et les tabulation inutiles
	Words := []string{}
	Word := ""
	Check := false

	for _, r := range Text {
		if r != '\t' && r != ' ' {
			Word += string(r)
			Check = true
		} else if Check {
			Words = append(Words, Word)
			Word = ""
			Check = false
		}
	}
	if Word != "" {
		Words = append(Words, Word)
	}

	for {
		foundIndex := -1
		foundTag := ""
		foundPos := -1
		foundWordIndex := -1

		tags := []string{"(up)", "(low)", "(bin)", "(hex)", "(cap)"}

		for wordIndex, word := range Words {
			for _, tag := range tags {
				position := strings.Index(word, tag)
				if position != -1 {
					if foundIndex == -1 || wordIndex < foundWordIndex || (wordIndex == foundWordIndex && position < foundPos) {
						foundIndex = wordIndex
						foundTag = tag
						foundPos = position
						foundWordIndex = wordIndex
					}
				}
			}
		}

		// Sortir si aucun tag n'a été trouvé
		if foundIndex == -1 {
			break
		}
		// Traitement du tag trouvé
		i := foundIndex
		tag := foundTag
		pos := foundPos

		WordBefore := Words[i][:pos]
		
		WordAfter := Words[i][pos+len(tag):]


		index := i - 1

		if len(WordBefore) == 0 {
			if i > 0 {
				WordBefore := Words[index]

				for (WordBefore == "\n" || WordBefore == "") && index >= 0 {
					index--
					if index >= 0 {
						WordBefore = Words[index]
					}
				}

				if index >= 0 {
					switch tag {
					case "(up)":
						Words[index] = strings.ToUpper(WordBefore)
						Words[i] = WordAfter

					case "(low)":
						Words[index] = strings.ToLower(WordBefore)
						Words[i] = WordAfter

					case "(bin)":
						t, _ := strconv.ParseInt(WordBefore, 2, 64)
						Words[index] = strconv.FormatInt(t, 10)
						Words[i] = WordAfter

					case "(hex)":
						t, _ := strconv.ParseInt(WordBefore, 16, 64)
						Words[index] = strconv.FormatInt(t, 10)
						Words[i] = WordAfter

					case "(cap)":
						Words[index] = strings.ToUpper(string(WordBefore[0])) + strings.ToLower(WordBefore[1:])
						Words[i] = WordAfter
					}
				} else {
					Words[i] = WordAfter
				} // Devrait donner ["HELLOWORLD"]
			} else {
				Words[i] = WordAfter
			}

		} else if WordBefore == "\n" {
			if i > 0 {
				WordBefore = Words[index]
				// Devrait donner ["HELLOWORLD"]
				for WordBefore == "\n" && index > 0 {
					index--
					WordBefore = Words[index]
				}

				switch tag {
				case "(up)":
					Words[index] = strings.ToUpper(WordBefore)
					Words[i] = "\n" + WordAfter

				case "(low)":
					Words[index] = strings.ToLower(WordBefore)
					Words[i] = "\n" + WordAfter

				case "(bin)":
					t, _ := strconv.ParseInt(WordBefore, 2, 64)
					Words[index] = strconv.FormatInt(t, 10)
					Words[i] = "\n" + WordAfter

				case "(hex)":
					t, _ := strconv.ParseInt(WordBefore, 16, 64)
					Words[index] = strconv.FormatInt(t, 10)
					Words[i] = "\n" + WordAfter

				case "(cap)":
					if len(WordBefore) > 0 {
						Words[index] = strings.ToUpper(string(WordBefore[0])) + strings.ToLower(WordBefore[1:])
					}
					Words[i] = "\n" + WordAfter
				}
			} else {
				Words[i] = WordAfter
			}

		} else {
			switch tag {
			case "(up)":
				Words[i] = strings.ToUpper(WordBefore) + WordAfter

			case "(low)":
				Words[i] = strings.ToLower(WordBefore) + WordAfter

			case "(bin)":

				t, err := strconv.ParseInt(WordBefore, 2, 64)
				if err == nil {
					Words[i] = strconv.FormatInt(t, 10) + WordAfter
				} else {
					Words[i] = WordBefore + WordAfter
				}

			case "(hex)":
				t, _ := strconv.ParseInt(WordBefore, 16, 64)
				Words[i] = strconv.FormatInt(t, 10) + WordAfter

			case "(cap)":
				if len(WordBefore) > 0 {
					Words[i] = strings.ToUpper(string(WordBefore[0])) + strings.ToLower(WordBefore[1:]) + WordAfter
				}
			}
		}

		if Words[i] == "" {
			Words = append(Words[:i], Words[i+1:]...)
		}
	}

	return Words
}

func main() {
	result := EditedTex("1zdjhbfjhbf0\n(low)(up)kjnfkjnf(up)")
	fmt.Println(result)

}
