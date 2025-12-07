package main

import (
	"unicode"
)

//------CheckWord-----//
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

//------Quote-----//
func isWordChar(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsDigit(char) {
		return true
	}
	return false
}


