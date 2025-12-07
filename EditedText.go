package main

import (
	//"regexp"
	//"fmt"
	//"fmt"
	//"fmt"
	"fmt"
	"strconv"
	"strings"
)

func EditedText(Text string) string {
	words := []string{}
	word := ""
	check := false

	for _, r := range Text {
		if r != '\t' && r != ' ' && r != '\n' && r != '\r' {
			word += string(r)
			check = true
		} else {
			if check {
				words = append(words, word)
				word = ""
				check = false
			}
			if r == '\n' {
				words = append(words, "\n")
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}

	for {
		wordIndex := -1
		tag := ""
		position := -1
		tagNumber := 0

		tags := []string{"(up)", "(low)", "(bin)", "(hex)", "(cap)", "(up,", "(low,", "(cap,", ".", ",", "!", "?", ":", ";"}
		for i, word := range words {
			n := 0

			for q, t := range tags {
				pos := strings.Index(word, t)

				// Manipulation d (tag, <Number>):
				if (q == 5 || q == 6 || q == 7) && pos != -1 {
					if len(words[i][pos+len(t):]) == 0 {
						if i+1 < len(words) {
							s := ""
							for _, r := range words[i+1] {
								if r != ')' {
									s += string(r)
								} else {
									break
								}
							}

							if len(s) < len(words[1+i]) {
								if words[i+1][len(s)] == ')' {
									number, err := strconv.Atoi(s)
									if err != nil {
										continue
									} else {
										n = number
									}
									//
								} else {
									continue
								}
							} else {
								continue
							}

						} else {
							continue
						}
					} else {
						continue
					}
				}
				// Manipulation Ponctuation
				if (q == 8 || q == 9 || q == 10 || q == 11 || q == 12 || q == 13) && pos != -1 {
					beforeTag := words[i][:pos]
					Aftertag := words[i][pos+len(t):]

					fmt.Println(Aftertag)

					// Skip si ponctuation à la fin
					if len(beforeTag) > 0 && len(Aftertag) == 0 {
						continue
					}

					// Skip si au début ET (première ligne OU pas de mot valide avant)
					if len(beforeTag) == 0 {
						// Vérifier s'il y a un mot valide avant
						hasValidPrevious := false
						for k := i - 1; k >= 0; k-- {
							if words[k] != "\n" && words[k] != "" {
								hasValidPrevious = true
								break
							}
						}

						// Skip si pas de mot valide avant (début ou que des \n)
						if !hasValidPrevious {
							continue
						}
					}

					// Skip si après il y a SEULEMENT des ponctuations
					if len(beforeTag) > 0 && len(Aftertag) > 0 {
						allPunct := true
						for _, r := range Aftertag {
							if r != '?' && r != '!' && r != '.' && r != ',' && r != ';' && r != ':' {
								allPunct = false
								break
							}
						}
						if allPunct {
							continue
						}
					}
				}
				if pos != -1 {
					if wordIndex == -1 || i < wordIndex || (i == wordIndex && pos < position) {
						wordIndex = i
						tag = t
						position = pos
						tagNumber = n
					}
				}
			}
		}

		if wordIndex == -1 {
			break
		}
		i := wordIndex

		wordBefore := words[i][:position]
		wordAfter := words[i][position+len(tag):]

		j := i - 1

		switch tag {
		case "(up)", "(up,":
			UpperCase(&words, wordBefore, wordAfter, tag, i, j, position, tagNumber)

		case "(low)", "(low,":
			LowerCase(&words, wordBefore, wordAfter, tag, i, j, position, tagNumber)
			//
		case "(cap)", "(cap,":
			Capitalized(&words, wordBefore, wordAfter, tag, i, j, position, tagNumber)
			//
		case "(bin)":
			BinaryDecimal(&words, wordBefore, wordAfter, i, j, position)
			//
		case "(hex)":
			HexaDecimal(&words, wordBefore, wordAfter, i, j, position)

		case ".", ",", "!", "?", ":", ";":
			Ponctuations(&words, wordBefore, wordAfter, tag, i, j, position)
			//

		}

		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
		}

	}

	// Manipulation = > vowels "a" et "A"
	for i := 0; i < len(words); i++ {
		if words[i] == "a" || words[i] == "A" {
			ABeforeVowel(&words, i)
		}
	}

	//Manipulation des Quotes
	words = Quotes(words)

	newText := ""
	for i, r := range words {
		if r == "\n" {
			newText += r
			continue
		}
		newText += r
		if i+1 < len(words) && words[i+1] != "\n" {
			newText += " "
		}
	}

	newText = strings.TrimSpace(newText)
	return newText
}
