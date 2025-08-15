package main

import (
	"fmt"
	"unicode"
)

func Solve(s string) []int {
	retornar := []int{0, 0, 0, 0}
	for _, l := range s {
		if unicode.IsUpper(l) {
			retornar[0]++
		} else if unicode.IsLower(l) {
			retornar[1]++
		} else if unicode.IsNumber(l) {
			retornar[2]++
		} else {
			retornar[3]++
		}
	}
	return retornar
}

func main() {
	test := Solve("bgA5<1d-tOwUZTS8yQ")
	fmt.Println(test)
}
