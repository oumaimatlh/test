package main

import (
	"fmt"
	"strings"
)

func EditedText(Text string) []string {
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

		tags := []string{"(up)", "(low)", "(bin)", "(hex)", "(cap)"}

		for i, word := range words {
			for _, t := range tags {
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
			UpperCase(&words, wordBefore, wordAfter, i, j)
			//
		case "(low)":
			LowerCase(&words, wordBefore, wordAfter, i, j)
			//
		case "(cap)":
			Capitalized(&words, wordBefore, wordAfter, i, j)
			//
		case "(bin)":
			BinaryDecimal(&words, wordBefore, wordAfter, i, j)
			//
		case "(hex)":
			HexaDecimal(&words, wordBefore, wordAfter, i, j)

		}

		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func main() {
	result := EditedText("AD\n1E(cap)")
	fmt.Println(result)
}