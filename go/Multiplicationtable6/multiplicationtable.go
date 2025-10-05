package main

import (
	"fmt"
)

func MultiplicationTable(size int) [][]int {
	result := [][]int{}
	temp := []int{}

	for n := 1; n <= size; n++ {

		for n2 := 1; n2 <= size; n2++ {
			temp = append(temp, (n)*(n2))
		}
		result = append(result, temp)
		temp = []int{}
	}

	return result
}

func main() {
	test := MultiplicationTable(3)
	fmt.Println(test)
}
