package main

import (
	"fmt"
)

func GetMiddle(s string) string {
	xd := len(s) / 2

	if len(s)%2 == 0 {
		return s[xd-1 : xd+1]
	} else {
		return s[xd : xd+1]
	}
}

func main() {
	test := GetMiddle("test")
	fmt.Println(test)
}
