package main

import "strings"

func UpperCase(words *[]string, wordBefore,wordAfter string, i,j int) []string {
	if len(wordBefore) == 0 {
		if i > 0 {
			wordBefore = (*words)[j]
			for (wordBefore == "\n" || wordBefore == "") && j >= 0 || !CheckWord(wordBefore) {
				j--
				if j >= 0 {
					wordBefore = (*words)[j]
				}
			}
			(*words)[j] = strings.ToUpper(wordBefore)
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
			(*words)[j] = strings.ToUpper(wordBefore)
			(*words)[i] = "\n" + wordAfter
		} else {
			(*words)[i] = "\n" + wordAfter
		}
		//
	} else {
		if CheckWord(wordBefore) {
			(*words)[i] = strings.ToUpper(wordBefore) + wordAfter
		} else {
			if i > 0 {
				wordBefore = (*words)[j]
				for (wordBefore == "\n" || wordBefore == "") && j >= 0 || !CheckWord(wordBefore) {
					j--
					if j >= 0 {
						wordBefore = (*words)[j]
					}
				}
				(*words)[j] = strings.ToUpper(wordBefore)
				(*words)[i] = wordAfter

			} else {
				(*words)[i] =wordBefore + wordAfter
			}
		}

	}
	return  *words
}
