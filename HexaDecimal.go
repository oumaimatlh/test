package main

import "strconv"

//import "strings"

func ConvertHexaDecimal(s string) string {

	if s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}

	t, err := strconv.ParseInt(s, 16,64)

	if err != nil {
		return s
	}
	return strconv.FormatInt(t, 10)
}

func HexaDecimal(words *[]string, wordBefore,wordAfter string, i,j int) []string {
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
	} else if wordBefore == "\n" {
		if i > 0 {
			wordBefore = (*words)[j]
			for (wordBefore == "\n" && j > 0) || !CheckWord(wordBefore) {
				j--
				wordBefore = (*words)[j]
			}
			(*words)[j] = ConvertHexaDecimal(wordBefore)
			(*words)[i] = "\n" + wordAfter
		} else {
			(*words)[i] = "\n" + wordAfter
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
				(*words)[i] = wordAfter

			} else {
				(*words)[i] =wordBefore + wordAfter
			}
		}

	}
	return  *words
}
