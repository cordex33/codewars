package main

import (
	"fmt"
)

func GetMiddle(s string) string {
	xd := len(s) / 2

	return s[xd : xd+1]
}

func main() {
	test := GetMiddle("A")
	fmt.Println(test)
}
