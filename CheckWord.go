
package main

func CheckWord(word string) bool {
	for i := 0; i < len(word); i++ {
		char := word[i]
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			return true
		}
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}