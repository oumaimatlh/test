package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EditedText(Text string) []string {
	//Auto-Correction => Supprimer les espaces et les tabulation inutiles
	//Diviser chaque suite de caracteres qui sont séprarer par Espace
	//RQ => Dans ce cas j'ai laissé /n malgré est parmi les espaces blanc ; mais j'ai le besoin a la fin a fin de retourner le text avec ces sautes de lignes

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

	//Editing :
	//Search a Tags
	Tags := []string{"(up)", "(low)", "(bin)", "(hex)", "(cap)"}

	for i := 0; i < len(Words); i++ {
		for j := 0; j < len(Tags); j++ {
			Tag := Tags[j]
			LenTag := len(Tag)
			Matching := ""

			for q := 0; q < len(Words[i]); q++ {

				if LenTag <= len(Words[i]) {
					Matching = string(Words[i][q:LenTag])
				} else {
					Matching = string(Words[i][q:])
				}

				if Matching == Tag {
					//chaque '(Tag)' va faire son travail selon  le type de ce tag => cases

					WordBefore := Words[i][:q] // On va stocker le mot qui juste avant le tag sans espace
					WordAfter := Words[i][q+len(Tag):]

					index := i - 1

					if len(WordBefore) == 0 {
						if i > 0 {
							WordBefore := Words[index]

							for (WordBefore == "\n" || WordBefore == "") && index >= 0 {
								index--
								WordBefore = Words[index]
							}
							

							switch Tag {
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
						}

					} else if WordBefore == "\n" {
						if i > 0 {
							WordBefore = Words[index]

							for WordBefore == "\n" {
								if i-1 <= 0 {
									break
								}
								i--
								WordBefore = Words[index]

							}

							switch Tag {
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
						// Words[i][:q] est une suite de carac et non espace blanc et non  len(Words[i][:q]) == 0
						switch Tag {
						// dans ce cas le mot avant est tjr Words[i][:q]
						case "(up)":
							Words[i] = strings.ToUpper(WordBefore)+ WordAfter
							fmt.Println(WordBefore)

						case "(low)":
							Words[i] = strings.ToLower(WordBefore) + WordAfter
						case "(bin)":
							t, _ := strconv.ParseInt(WordBefore, 2, 64)
							Words[i] = strconv.FormatInt(t, 10) + Words[i][q+len(Tag):]

						case "(hex)":
							t, _ := strconv.ParseInt(WordBefore, 16, 64)
							Words[i] = strconv.FormatInt(t, 10) +  Words[i][q+len(Tag):]

						case "(cap)":
							fmt.Println(WordBefore)
							if len(WordBefore) > 0 {
								fmt.Println(WordBefore)
								fmt.Println(Words[i][q+len(Tag):])
								Words[i] = strings.ToUpper(string(WordBefore[0])) + strings.ToLower(WordBefore[1:]) +  WordAfter
							}

						}
						
					
					
					}
					q = -1
					
						
				} else {
					LenTag++
				}
			}
		}
	}
	return Words
}

func main() {
	fmt.Println(EditedText("hello(low)world(up)"))
}
