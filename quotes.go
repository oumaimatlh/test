package main

import (
	"fmt"
	"slices"
	"strings"
)

func CheckChar(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '_' || c == '-'
}
func Quotes(words *[]string) []string {
	for i := 0; i <= len((*words))-1; i++ {

		r := (*words)[i]
		pos := strings.Index(r, "'")

		count := 0
		closeQuote := 0
		openQuote := 0
		j := i + 1

		if pos != -1 {

			if len((*words)[i][:pos]) > 0 && len((*words)[i][pos+len("'"):]) > 0 {
				if CheckChar((*words)[i][:pos][len((*words)[i][:pos])-1]) && CheckChar((*words)[i][pos+len("'"):][0]) {
					continue
				}
			}

			count++
			pos2 := 0

			// D'abord chercher dans le même mot (après la première quote)
			pos2 = strings.Index((*words)[i][pos+1:], "'")

			if pos2 != -1 {
				// Vérifier si la quote est entre 2 caractères (apostrophe interne)
				actualPos2 := pos + 1 + pos2 // Position réelle dans (*words)[i]

				hasCharBefore := actualPos2 > 0 && CheckChar((*words)[i][actualPos2-1])
				hasCharAfter := actualPos2 < len((*words)[i])-1 && CheckChar((*words)[i][actualPos2+1])

				if hasCharBefore && hasCharAfter {
					// C'est une apostrophe interne, chercher dans les mots suivants
					pos2 = -1
					for k := i + 1; k < len((*words)); k++ {
						r2 := (*words)[k]
						pos2 = strings.Index(r2, "'")
						if pos2 != -1 {
							closeQuote = k
							count++
							break
						}
					}
				} else {
					// Les deux quotes sont dans le même mot
					closeQuote = i
					pos2 = actualPos2 // Ajuster la position réelle
					count++
				}
			} else {
				// Chercher dans les mots suivants
				for k := i + 1; k < len((*words)); k++ {
					r2 := (*words)[k]
					pos2 = strings.Index(r2, "'")
					if pos2 != -1 {
						closeQuote = k
						count++
						break
					}
				}
			}

			openQuote = i
			if count == 2 {
				// Handl Open Quote:
				// 1) "word' uhiuh.." => "word 'uhiuh.."
				if len((*words)[openQuote][:pos]) > 0 && len((*words)[openQuote][pos+len("'"):]) == 0 {
					if i < len((*words))-1 {
						next := (*words)[j]

						for j < len((*words))-1 && (next == "\n" || next == "") {
							j++
							next = (*words)[j]
							fmt.Println(next)
						}

						(*words)[i] = (*words)[i][:pos]
						(*words)[j] = "'" + next

					}
				} else if len((*words)[openQuote][:pos]) == 0 && len((*words)[openQuote][pos+len("'"):]) == 0 {
					if i < len((*words))-1 {
						next := (*words)[j]

						for j < len((*words))-1 && (next == "\n" || next == "") {
							j++
							next = (*words)[j]
						}

						(*words)[i] = ""
						(*words)[j] = "'" + next

						//
					}
				} else if len((*words)[openQuote][:pos]) > 0 && len((*words)[openQuote][pos+len("'"):]) > 0 {

					// Sauvegarder la partie après l'apostrophe
					after := "'" + (*words)[openQuote][pos+len("'"):]

					(*words)[i] = (*words)[openQuote][:pos]

					*words = append(*words, "")
					copy((*words)[i+2:], (*words)[i+1:])
					(*words)[i+1] = after
				}

				// Handl Close Quote:
				//
				fmt.Println(closeQuote)
				if len((*words)[closeQuote][:pos2]) == 0 && len((*words)[closeQuote][pos2+len("'"):]) > 0 {
					fmt.Println("cas1")
					if closeQuote > 0 {
						prev := (*words)[j]
						j = closeQuote - 1
						for j > 0 && (prev == "\n" || prev == "") {
							j--
							prev = (*words)[j]
							fmt.Println(prev)
						}

						(*words)[closeQuote] = (*words)[closeQuote][pos2+len("'"):]
						(*words)[j] = (*words)[j] + "'"

						//
					}
				} else if len((*words)[closeQuote][:pos2]) == 0 && len((*words)[closeQuote][pos2+len("'"):]) == 0 {
					fmt.Println("cas2")

					if closeQuote > 0 {
						j = closeQuote - 1
						prev := (*words)[j]
						fmt.Println(prev)
						j = closeQuote - 1

						for j > 0 && (prev == "\n" || prev == "") {
							j--
							prev = (*words)[j]
							fmt.Println(prev)
						}

						(*words)[closeQuote] = ""
						(*words)[j] = prev + "'"

						//
					}
				} else if len((*words)[closeQuote][:pos2]) > 0 && len((*words)[closeQuote][pos2+len("'"):]) > 0 {

					fmt.Println("cas3")

					// Sauvegarder la partie après l'apos2trophe
					after := (*words)[closeQuote][pos2+len("'"):]

					(*words)[closeQuote] = (*words)[closeQuote][:pos2] + "'"

					*words = slices.Insert(*words, closeQuote+1, after)

				}

			} else {
				continue
			}

			i = closeQuote
		}

	}
	return *words
}
