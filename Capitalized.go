package main

import (
	//"strings"
	"unicode"
)
func capitalizeWord(s string) string {  
    letters := []rune(s)
    check := false
    
    for i, r := range letters {
        if unicode.IsLetter(r) {
            if !check {
				letters[i] = unicode.ToUpper(r)
                check = true
            } else {
                letters[i] = unicode.ToLower(r)
            }
        }
    }
    return string(letters)
}

func Capitalized(words *[]string, wordBefore, wordAfter string, i, j int) []string {
	if len(wordBefore) == 0 {
		if i > 0 {
			wordBefore = (*words)[j]
			for (wordBefore == "\n" || wordBefore == "") && j >= 0 || !CheckWord(wordBefore) {
				j--
				if j >= 0 {
					wordBefore = (*words)[j]
				}
			}
			(*words)[j] = capitalizeWord(wordBefore)
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
			if len(wordBefore) > 0 {
				(*words)[j] = capitalizeWord(wordBefore)

			}
			(*words)[i] = "\n" + wordAfter
		} else {
			(*words)[i] = "\n" + wordAfter
		}
		//
	} else {
		if CheckWord(wordBefore) {
			(*words)[i] = capitalizeWord(wordBefore)+wordAfter
		} else {
			if i > 0 {
				wordBefore = (*words)[j]
				for (wordBefore == "\n" || wordBefore == "") && j >= 0 || !CheckWord(wordBefore) {
					j--
					if j >= 0 {
						wordBefore = (*words)[j]
					}
				}
				(*words)[j] = capitalizeWord(wordBefore)
				(*words)[i] = wordAfter

			} else {
				(*words)[i] =capitalizeWord(wordBefore) + wordAfter

			}
		}

	}
	return *words
}
