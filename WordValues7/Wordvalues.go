package main

import (
	"fmt"
)

func NameValue(my_list []string) []int {

	total := 0
	xd := []int{}
	for multi, x := range my_list {
		for _, xd := range x {
			if string(xd) != " " {
				total = total + int(xd) - 96
			}
		}
		if multi > 0 {
			xd = append(xd, total*(multi+1))
		} else {
			xd = append(xd, total)
		}

		total = 0
	}

	return xd
}

func main() {
	test := NameValue([]string{"abc", "abc", "abc", "abc"})
	fmt.Println(test)
}
