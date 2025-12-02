package main

import "fmt"

func Ponctuations(words *[]string, wordBefore, wordAfter, tag string, i, j, position int) []string {
	
	fmt.Println((*words)[i])
	fmt.Println(tag)

	
	if len(wordBefore) == 0 {
		if i > 0 {
			wordBefore = (*words)[j]

			for j >= 0 && (wordBefore == "\n" || wordBefore == "") {
				j--
				if j >= 0 {
					wordBefore = (*words)[j]
				}
			}
			if j >= 0 {
				(*words)[j] = (*words)[j] + tag
			}
			(*words)[i] = wordAfter
			fmt.Println((*words)[j] + (*words)[i])

		} else {
			(*words)[i] = wordBefore + tag + wordAfter
		}

		//
	} else if len(wordBefore) > 0 && len(wordAfter) > 0 {
		(*words)[i] = wordBefore + tag
		*words = append((*words)[:i+1], append([]string{wordAfter}, (*words)[i+1:]...)...)
	}
		
	return (*words)
}
