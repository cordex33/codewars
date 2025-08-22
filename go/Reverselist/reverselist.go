package main

import (
	"fmt"
)

func ReverseList(lst []int) []int {
	xd := []int{}
	for _, n := range lst {
		xd = append([]int{n}, xd...)
	}
	return xd
}

func main() {
	test := ReverseList([]int{2, 4, 5, 6})
	fmt.Println(test)
}
