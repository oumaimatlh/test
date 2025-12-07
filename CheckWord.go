package main

import (
	"unicode"
)

func CheckWord(word string) bool {
	for _, char := range word {
		
		if unicode.IsLetter(char) {
			return true
		}
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}
func CheckChar(c byte) bool {
    return unicode.IsLetter(rune(c))
}
