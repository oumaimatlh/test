package main

import (
	"fmt"
	//"regexp"
	"strconv"
	"strings"
)

func EditedText(Text string) string {
	words := []string{}
	word := ""
	check := false

	for _, r := range Text {
		if r != '\t' && r != ' ' {
			word += string(r)
			check = true
		} else if check {
			words = append(words, word)
			word = ""
			check = false
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

		tags := []string{"(up)", "(low)", "(bin)", "(hex)", "(cap)", "(up,", "(low,", "(cap,"}

		for i, word := range words {
			for q, t := range tags {

				pos := strings.Index(word, t)

				//TagNumber (<tag>, <number>)
				if (q == 5 || q == 6 || q == 7) && pos != -1 {
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
								tagNumber = number
								fmt.Println(tagNumber)
							}
						} else {
							continue
						}
					} else {
						continue
					}

				}
				if pos != -1 {

					if wordIndex == -1 || i < wordIndex || (i == wordIndex && pos < position) {
						wordIndex = i
						tag = t
						position = pos
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
		case "(up)":
			UpperCase(&words, wordBefore, wordAfter, tag, i, j, position, tagNumber)
			//
		case "(low)":
			LowerCase(&words, wordBefore, wordAfter, i, j, position)
			//
		case "(cap)":
			Capitalized(&words, wordBefore, wordAfter, i, j, position)
			//
		case "(bin)":
			BinaryDecimal(&words, wordBefore, wordAfter, i, j, position)
			//
		case "(hex)":
			HexaDecimal(&words, wordBefore, wordAfter, i, j, position)

		case "(up,":
			UpperCase(&words, wordBefore, wordAfter, tag, i, j, position, tagNumber)

		}

		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
		}
	}
	newText := ""
	for _, r := range words {
		newText += r + " "
	}
	return newText
}
