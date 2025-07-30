package main

import (
	"fmt"
	"strings"
)

func Capitalize(st string) []string {
	xd := []string{}
	a := ""
	for i, x := range st {
		if i%2 == 0 {
			a += strings.ToUpper(string(x))
		} else {
			a += string(x)
		}
	}
	xd = append(xd, a)
	a = ""
	for i, x := range st {
		if i%2 == 0 {
			a += string(x)
		} else {
			a += strings.ToUpper(string(x))
		}
	}
	xd = append(xd, a)
	return xd
}

func main() {
	test := Capitalize("codewars")
	fmt.Println(test)
}
