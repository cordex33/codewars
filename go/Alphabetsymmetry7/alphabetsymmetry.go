package main

import (
	"fmt"
	"strings"
)

func solve(slice []string) []int {

	xd := []int{}

	for _, word := range slice {
		count := 0
		for i, letter := range strings.ToLower(word) {

			if int(letter)-96 == i+1 {
				count++
			}
		}
		xd = append(xd, count)
	}

	return xd
}

func main() {
	test := solve([]string{"IAMDEFANDJKL", "thedefgh", "xyzDEFghijabc"})
	fmt.Println(test)
}
