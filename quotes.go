package main

import (
	"fmt"
	"strings"
)

func Quotes(words *[]string) []string {
	for i := 0; i <= len((*words))-1; i++ {
	if i < len(*words)-1 {
    if (*words)[i] == "'" && (*words)[i+1] == "'" {
        (*words)[i] = "''"
        *words = append((*words)[:i+1], (*words)[i+2:]...)
        continue 
    }
}

		r := (*words)[i]
		pos := strings.Index(r, "'")

		count := 0
		closeQuote := 0
		openQuote := 0
		j := i + 1

		if pos != -1 {

			if CheckWord((*words)[i][:pos]) && CheckWord((*words)[i][pos+len("'"):]) {
				continue
			}

			count++
			pos2 := 0
			for k := i + 1; k < len((*words)); k++ {
				r2 := (*words)[k]
				pos2 = strings.Index(r2, "'")

				if pos2 != -1 {
					closeQuote = k
					fmt.Print(closeQuote)

					count++
					break
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
				fmt.Print(closeQuote)
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
					after := "'" + (*words)[closeQuote][pos2+len("'"):]

					(*words)[closeQuote] = (*words)[closeQuote][:pos2]

					*words = append(*words, "")
					copy((*words)[closeQuote+2:], (*words)[closeQuote+1:])
					(*words)[closeQuote+1] = after
				}

			} else {
				continue
			}

			i = closeQuote + 1
		}

	}
	return *words
}
