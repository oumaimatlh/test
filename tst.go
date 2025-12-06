package main

import (
	"fmt"
	"strings"
)

//"fmt"
//"strings"

func Quotess(words *[]string) []string {
	for i := 0; i <= len((*words))-1; i++ {

		count := 0
		r := (*words)[i]
		pos := strings.Index(r, "'")
		j := i + 1
		closeQuote := 0

		if r == "'" {
			count++
			for k := i + 1; k < len((*words)); k++ {
				r2 := (*words)[k]
				if r2 == "'" {
					closeQuote = k
					count++
					break
				}
			}

			if count == 2 {
				// Handl Open Quote
				if i < len((*words))-1 {
					next := (*words)[j]

					for j < len((*words))-1 && (next == "\n" || next == "") {
						j++
						next = (*words)[j]
						fmt.Println(next)
					}

					(*words)[i] = ""
					(*words)[j] = "'" + next

					//
				}

				// Handl Close
				j = closeQuote - 1
				if closeQuote > 0 {
					prev := (*words)[j]
					for j > 0 && (prev == "\n" || prev == "") {
						j--
						prev = (*words)[j]
						fmt.Println(prev)
					}

					(*words)[closeQuote] = ""
					(*words)[j] = prev + "'"

					//
				}
				i = closeQuote + 1

			}

			//

		} else if pos != -1 && r != "'" {
			count++
			for k := i + 1; k < len((*words)); k++ {
				r2 := (*words)[k]
				if r2 == "'" {
					closeQuote = k
					count++
					break
				}
			}
			if count == 2 {
				// Handl Open Quote
				if len((*words)[i][:pos]) > 0 && len((*words)[i][pos+len("'"):]) == 0 {
					if i < len((*words))-1 {
						next := (*words)[j]

						for j < len((*words))-1 && (next == "\n" || next == "") {
							j++
							next = (*words)[j]
							fmt.Println(next)
						}

						(*words)[i] = (*words)[i][:pos]
						(*words)[j] = "'" + next

						//
					}

					// Handl Close
					j = closeQuote - 1
					if closeQuote > 0 {
						prev := (*words)[j]
						for j > 0 && (prev == "\n" || prev == "") {
							j--
							prev = (*words)[j]
							fmt.Println(prev)
						}

						(*words)[closeQuote] = ""
						(*words)[j] = prev + "'"

						//
					}
				} else if len((*words)[i][:pos]) == 0 && len((*words)[i][pos+len("'"):]) >= 0 {

					// Handl Close
					j = closeQuote - 1
					if closeQuote > 0 {
						prev := (*words)[j]
						for j > 0 && (prev == "\n" || prev == "") {
							j--
							prev = (*words)[j]
							fmt.Println(prev)
						}

						(*words)[closeQuote] = ""
						(*words)[j] = prev + "'"

						//
					}
				}
			}
		}
		
	}
	return *words
}
