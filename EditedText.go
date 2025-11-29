package main

import (
	"fmt"
	"regexp"
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

		tags := []string{"(up)", "(low)", "(bin)", "(hex)", "(cap)", `^\((up|low|cap), [0-9]+\)$`}

		for i, word := range words {
			for q, t := range tags {
				if q == len(tags)-1 {
					if regexp.MustCompile(t).MatchString(word) {
						fmt.Println("Match !")
					}
				}
				pos := strings.Index(word, t)

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
			UpperCase(&words, wordBefore, wordAfter, i, j, position)
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

		}

		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
		}
	}
	newText := ""
	for _, r:= range words {
		newText += r + " "
	}
	return newText
}

