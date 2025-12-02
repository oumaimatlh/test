package main

import (
	"strconv"
	//"strings"
)


func ConvertHexaDecimal(s string) string {
	t, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s
	}
	return strconv.FormatInt(t, 10)
}

func HexaDecimal(words *[]string, wordBefore, wordAfter string, i, j, position int) []string {
	if len(wordBefore) == 0 {
		if i > 0 {
			wordBefore = (*words)[j]
			for (wordBefore == "\n" || wordBefore == "") && j >= 0 || !CheckWord(wordBefore) {
				j--
				if j >= 0 {
					wordBefore = (*words)[j]
				}
			}
			(*words)[j] = ConvertHexaDecimal(wordBefore)
			(*words)[i] = wordAfter

		} else {
			(*words)[i] = wordAfter
		}
		//
	} else {
		if CheckWord(wordBefore) {
			(*words)[i] = ConvertHexaDecimal(wordBefore) + wordAfter
		} else {
			if i > 0 {
				wordBefore = (*words)[j]
				for (wordBefore == "\n" || wordBefore == "") && j >= 0 || !CheckWord(wordBefore) {
					j--
					if j >= 0 {
						wordBefore = (*words)[j]
					}
				}
				(*words)[j] = ConvertHexaDecimal(wordBefore)
				(*words)[i] = (*words)[i][:position] + wordAfter
			} else {
				(*words)[i] = wordBefore + wordAfter
			}
		}

	}
	return *words
}
