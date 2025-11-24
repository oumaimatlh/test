package main

import (
	"fmt"
	"strings"
)

func EditFile(content string) {
	words := []string{}
	word := ""
	b := false
	for _, r := range content {
		if r != '\n' && r != '\t' && r != ' ' {
			word += string(r)
			b = true
		} else if b {
			words = append(words, word)
			word = ""
			b = false
		}
	}

	if word != "" {
		words = append(words, word)
	}

	//newText := ""
	// Traitement word by word
	//WordBefore := ""
	for j := 0; j < len(words); j++ {
		substring := "(up)"
		k := len(substring)
		match := ""
		h := ""
		for i := 0; i < len(words[j]); i++ {
			if k <= len(words[j]) {
				match = string(words[j][i:k])
			} else {
				match = string(words[j][i:])

			}

			if match == substring {
				WordBefore := words[j][:i]
				if len(WordBefore) == 0 {
					WordBefore = words[j-1]
					h = strings.ToUpper(WordBefore)

				} else {
					h = strings.ToUpper(WordBefore)
				}
				fmt.Println(h)

			} 
			k++
		}
	}

		for j := 0; j < len(words); j++ {
		substring := "(low)"
		k := len(substring)
		match := ""
		h := ""
		for i := 0; i < len(words[j]); i++ {
			if k <= len(words[j]) {
				match = string(words[j][i:k])
			} else {
				match = string(words[j][i:])

			}

			if match == substring {
				WordBefore := words[j][:i]
				if len(WordBefore) == 0 {
					WordBefore = words[j-1]
					h = strings.ToLower(WordBefore)

				} else {
					h = strings.ToLower(WordBefore)
				}
				fmt.Println(h)

			} 
			k++
		}
	}

}
func main() {
	EditFile("it (cap) was the best of times, it was the worst of times    (up)) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS ILOVEYY      (low) winter of despair.\nSimply add 42 (hex) and 10 (bin) and you will see the result is 68.")
}
 