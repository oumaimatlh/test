package main

import (
	//"regexp"
	//"fmt"
	//"fmt"
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
	Quotes(&words)

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

				//Manipulation d (tag, <Number>):
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
					Aftertag := words[i][pos+len(t):]

					if len(words[i][:pos]) > 0 && (len(Aftertag) == 0 ||
						(Aftertag[0] == '?' || Aftertag[0] == '!' ||
							Aftertag[0] == '.' || Aftertag[0] == ',' ||
							Aftertag[0] == ';' || Aftertag[0] == ':')) ||
						(i == 0 && len(words[i][:pos]) == 0 && len(Aftertag) >= 0) {
						continue
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
