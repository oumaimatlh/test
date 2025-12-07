package main

import (
	"strings"
	"unicode"
)

func isWordChar(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsDigit(char) {
		return true
	}
	return false
}

func isPartOfWord(runes []rune, index int) bool {
	hasCharBefore := index > 0 && isWordChar(runes[index-1])
	hasCharAfter := index < len(runes)-1 && isWordChar(runes[index+1])
	return hasCharBefore && hasCharAfter
}

// finds the matching closing quote
func findClosingQuote(runes []rune, startIndex int) int {
	for j := startIndex + 1; j < len(runes); j++ {
		if runes[j] == '\'' {

			if isPartOfWord(runes, j) {
				continue
			}
			return j
		}
	}
	return -1
}

// gets the text between quotes and trims whitespace
func extractQuoteContent(runes []rune, startIndex int, endIndex int) string {
	content := string(runes[startIndex+1 : endIndex])
	return strings.TrimSpace(content)
}

// handleQuotes ensures proper spacing around single quotes
func Quotes(tokens []string) []string {
	text := strings.Join(tokens, " ")
	result := ""
	quoteContent := ""
	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		if runes[i] != '\'' {
			result += string(runes[i])
			continue
		}

		if isPartOfWord(runes, i) {
			result += string(runes[i])
			continue
		}

		closingQuoteIndex := findClosingQuote(runes, i)
		if closingQuoteIndex < 0 {
			result += string(runes[i])
			continue
		}

		quoteContent = extractQuoteContent(runes, i, closingQuoteIndex)
		result += "'" + quoteContent + "'"

		i = closingQuoteIndex
	}

	return strings.Split(result, " ")
}
