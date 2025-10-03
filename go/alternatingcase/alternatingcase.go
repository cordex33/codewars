package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ToAlternatingCase(str string) string {
	result := ""
	for _, l := range str {
		switch {
		case unicode.IsNumber(l):
			result += string(l)
		case unicode.IsUpper(l):
			result += strings.ToLower(string(l))
		case unicode.IsLower(l):
			result += strings.ToUpper(string(l))
		default:
			result += string(l)
		}
	}
	return result
}

func main() {
	a := ToAlternatingCase("1A2b3c4d5e")

	fmt.Println(a)
}
