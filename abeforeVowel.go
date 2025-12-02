package main

import (
	"fmt"
	"strings"
)

func ABeforeVowel(words *[]string, i int) []string {
	j := i + 1

	for j < len(*words) {
		if (*words)[j] == "\n" || (*words)[j] == "" || !CheckWord((*words)[j]) {
			j++
			continue
		}

		nextWord := (*words)[j]
		firstLetter := strings.ToLower(string(nextWord[0]))
		if strings.ContainsAny(firstLetter, "aeiouh") {
			if (*words)[i] == "A" {
				(*words)[i] = "An"
			}else {
				(*words)[i] = "an"
			}
		}
		break
	}

	fmt.Println((*words)[i])
	return *words
}
