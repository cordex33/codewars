package main

import (
	"fmt"
)

func Dative(word string) string {
	front := []string{"e", "é", "i", "í", "ö", "ő", "ü", "ű"}
	back := []string{"a", "á", "o", "ó", "u", "ú"}

	reverword := ""

	for _, x := range word {
		reverword = reverword + string(x)
	}

	for i, l := range front {
		for _, l2 := range reverword {
			if l == string(l2) {
				return reverword + "nek"
			}
			if i < 6 {
				if back[i] == string(l2) {
					return reverword + "nak"
				}
			}
		}
	}

	return ""
}

func main() {
	test := Dative("tükör")
	fmt.Println(test)
}
