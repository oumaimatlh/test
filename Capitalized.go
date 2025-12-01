package main

import (
	"unicode"
)

func capitalizeWord(s string) string {  
    letters := []rune(s)
    check := false
    
    for i, r := range letters {
        if unicode.IsLetter(r) {
            if !check {
                letters[i] = unicode.ToUpper(r)
                check = true
            } else {
                letters[i] = unicode.ToLower(r)
            }
        }
    }
    return string(letters)
}

func Capitalized(words *[]string, wordBefore, wordAfter, tag string, i, j, position, tagNumber int) []string {
    //Case = > "(cap)"
    if tag == "(cap)" {
        if len(wordBefore) == 0 {
            if i > 0 {
                wordBefore = (*words)[j]
                for j >= 0 && ((wordBefore == "\n" || wordBefore == "") || !CheckWord(wordBefore)) {
                    j--
                    if j >= 0 {
                        wordBefore = (*words)[j]
                    }
                }
                if j >= 0 {
                    (*words)[j] = capitalizeWord(wordBefore)
                }
                (*words)[i] = wordAfter
            } else {
                (*words)[i] = wordAfter
            }
        } else {
            if CheckWord(wordBefore) {
                (*words)[i] = capitalizeWord(wordBefore) + wordAfter
            } else {
                if i > 0 {
                    wordBefore = (*words)[j]
                    for j >= 0 && ((wordBefore == "\n" || wordBefore == "") || !CheckWord(wordBefore)) {
                        j--
                        if j >= 0 {
                            wordBefore = (*words)[j]
                        }
                    }
                    if j >= 0 {
                        (*words)[j] = capitalizeWord(wordBefore)
                    }
                    (*words)[i] = (*words)[i][:position] + wordAfter
                } else {
                    (*words)[i] = wordBefore + wordAfter
                }
            }
        }
    } else {
        //Case = > "(cap, <number>)"
        var pos int
        for p, r := range (*words)[i+1] {
            if r == ')' {
                pos = p + 1
                break
            }
        }

        // Vérification pour les nombres <= 0
        if tagNumber <= 0 {
            if len(wordBefore) == 0 {
                (*words)[i] = wordAfter + (*words)[i+1][pos:]
            } else {
                (*words)[i] = wordBefore + (*words)[i+1][pos:]
            }
            (*words)[i+1] = ""
            return *words
        }

        if len(wordBefore) == 0 {
            if i > 0 {
                wordBefore = (*words)[j]
                for j >= 0 && ((wordBefore == "\n" || wordBefore == "") || !CheckWord(wordBefore)) {
                    j--
                    if j >= 0 {
                        wordBefore = (*words)[j]
                    }
                }

                for count := 1; count <= tagNumber && j >= 0; count++ {
                    (*words)[j] = capitalizeWord(wordBefore)
                    if j > 0 {
                        j--
                        if j >= 0 {
                            wordBefore = (*words)[j]
                            for j >= 0 && ((wordBefore == "\n" || wordBefore == "") || !CheckWord(wordBefore)) {
                                j--
                                if j >= 0 {
                                    wordBefore = (*words)[j]
                                } else {
                                    break
                                }
                            }
                        }
                    } else {
                        break
                    }
                }

                (*words)[i] = wordAfter
                (*words)[i+1] = (*words)[i+1][pos:]

            } else {
                (*words)[i] = wordAfter
                (*words)[i+1] = (*words)[i+1][pos:]
            }

        } else {
            if CheckWord(wordBefore) {
                (*words)[i] = capitalizeWord(wordBefore) + (*words)[i+1][pos:]

                if i > 0 {
                    wordBefore = (*words)[j]
                    for j >= 0 && ((wordBefore == "\n" || wordBefore == "") || !CheckWord(wordBefore)) {
                        j--
                        if j >= 0 {
                            wordBefore = (*words)[j]
                        }
                    }

                    for count := 1; count <= tagNumber-1 && j >= 0; count++ {
                        (*words)[j] = capitalizeWord(wordBefore)
                        if j > 0 {
                            j--
                            if j >= 0 {
                                wordBefore = (*words)[j]
                                for j >= 0 && ((wordBefore == "\n" || wordBefore == "") || !CheckWord(wordBefore)) {
                                    j--
                                    if j >= 0 {
                                        wordBefore = (*words)[j]
                                    } else {
                                        break
                                    }
                                }
                            }
                        } else {
                            break
                        }
                    }

                    (*words)[i+1] = ""

                } else {
                    (*words)[i] = capitalizeWord(wordBefore) + (*words)[i+1][pos:]
                    (*words)[i+1] = ""
                }
            } else {
                if i > 0 {
                    wordBefore = (*words)[j]
                    for j >= 0 && ((wordBefore == "\n" || wordBefore == "") || !CheckWord(wordBefore)) {
                        j--
                        if j >= 0 {
                            wordBefore = (*words)[j]
                        }
                    }

                    for count := 1; count <= tagNumber && j >= 0; count++ {
                        (*words)[j] = capitalizeWord(wordBefore)
                        if j > 0 {
                            j--
                            if j >= 0 {
                                wordBefore = (*words)[j]
                                for j >= 0 && ((wordBefore == "\n" || wordBefore == "") || !CheckWord(wordBefore)) {
                                    j--
                                    if j >= 0 {
                                        wordBefore = (*words)[j]
                                    } else {
                                        break
                                    }
                                }
                            }
                        } else {
                            break
                        }
                    }

                    (*words)[i] =  (*words)[i][:position]+(*words)[i+1][pos:]
                    (*words)[i+1] = ""

                } else {
                    (*words)[i] =  (*words)[i][:position]+(*words)[i+1][pos:]
                    (*words)[i+1] = ""
                }
            }
        }
    }
    return *words
}