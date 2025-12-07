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

// gets the text between quotes and handles newlines specially
func extractQuoteContent(runes []rune, startIndex int, endIndex int) (string, string, string) {
	content := string(runes[startIndex+1 : endIndex])
	
	// Check for leading newlines after opening quote
	leadingNewlines := ""
	content = strings.TrimLeft(content, " ")
	for strings.HasPrefix(content, "\n") {
		leadingNewlines += "\n"
		content = strings.TrimPrefix(content, "\n")
		content = strings.TrimLeft(content, " ")
	}
	
	// Check for trailing newlines before closing quote
	trailingNewlines := ""
	content = strings.TrimRight(content, " ")
	for strings.HasSuffix(content, "\n") {
		trailingNewlines += "\n"
		content = strings.TrimSuffix(content, "\n")
		content = strings.TrimRight(content, " ")
	}
	
	return leadingNewlines, strings.TrimSpace(content), trailingNewlines
}

// handleQuotes ensures proper spacing around single quotes
func Quotes(tokens []string) []string {
	text := strings.Join(tokens, " ")
	result := ""
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

		leadingNewlines, quoteContent, trailingNewlines := extractQuoteContent(runes, i, closingQuoteIndex)
		
		// Only remove trailing space if there are leading newlines
		// (meaning the quote should move to the next line)
		if leadingNewlines != "" && len(result) > 0 && result[len(result)-1] == ' ' {
			result = result[:len(result)-1]
		}
		
		// Add: leading newlines + quote + content + quote + trailing newlines
		result += leadingNewlines + "'" + quoteContent + "'" + trailingNewlines

		i = closingQuoteIndex
	}

	return strings.Split(result, " ")
}