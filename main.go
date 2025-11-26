package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EditedText(Text string) []string {
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

	// Nouvelle approche : traiter les tags dans l'ordre d'apparition
	for {
		// Trouver le premier tag dans tout le texte
		earliestIndex := -1
		earliestTag := ""
		earliestPos := -1 // position dans le mot
		earliestWordIndex := -1

		Tags := []string{"(up)", "(low)", "(bin)", "(hex)", "(cap)"}

		// Parcourir tous les mots pour trouver le tag le plus à gauche
		for i, word := range Words {
			for _, tag := range Tags {
				pos := strings.Index(word, tag)
				if pos != -1 {
					if earliestIndex == -1 || i < earliestWordIndex || (i == earliestWordIndex && pos < earliestPos) {
						earliestIndex = i
						earliestTag = tag
						earliestPos = pos
						earliestWordIndex = i
					}
				}
			}
		}

		// Si aucun tag trouvé, sortir de la boucle
		if earliestIndex == -1 {
			break
		}

		// Traiter le tag trouvé
		i := earliestIndex
		tag := earliestTag
		pos := earliestPos

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
				}// Devrait donner ["HELLOWORLD"]
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
			// Le tag est au milieu d'un mot
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
					// Si pas valide, on garde WordBefore tel quel
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

		// Si le mot devient vide après suppression du tag, le retirer
		if Words[i] == "" {
			Words = append(Words[:i], Words[i+1:]...)
		}
	}

	return Words
}

func main() {
	result := EditedText("10 (bin) \n 10 (bin) jh \npp")
	fmt.Println(result)

}
