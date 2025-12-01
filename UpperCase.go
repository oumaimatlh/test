package main

import (
	//"fmt"
	"strings"
)

func UpperCase(words *[]string, wordBefore, wordAfter, tag string, i, j, position, tagNumber int) []string {

	//Case = > "(up)"
	if tag == "(up)" {
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
					(*words)[i] = (*words)[i][:position] + wordAfter

				} else {
					(*words)[i] = wordBefore + wordAfter
				}
			}

		}
	} else {
		//Case = > "(up, <number>)"
		var pos int
		for p, r := range (*words)[i+1] {
			if r == ')' {
				pos = p + 1

				break
			}
		}
		if len(wordBefore) == 0 {
			if i > 0 {
				wordBefore = (*words)[j]
				for (wordBefore == "\n" || wordBefore == "") && j >= 0 || !CheckWord(wordBefore) {
					j--
					if j >= 0 {
						wordBefore = (*words)[j]
					}
				}

				for i := 1; i <= tagNumber && j >= 0; i++ {
					(*words)[j] = strings.ToUpper(wordBefore)
					if j > 0 {
						j--
						wordBefore = (*words)[j]
						for j >= 0 && ((wordBefore == "\n" || wordBefore == "") || !CheckWord(wordBefore)) {
							j--
							if j >= 0 {
								wordBefore = (*words)[j]
							} else {
								break
							}
						}
					} else {
						break
					}
				}
				(*words)[i] = wordAfter
				(*words)[i+1] = (*words)[i+1][pos:]

			} else {
				(*words)[i] = wordAfter
				(*words)[i+1] = (*words)[i+1][pos:]

			}
			//

		}else {
			

		}

	}
	return *words
}
