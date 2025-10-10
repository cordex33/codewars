package main

import (
	"fmt"
	"strconv"
)

func Encode(str string, key int) []int {
	result := []int{}
	xd := strconv.Itoa(key)
	for _, l := range str {
		result = append(result, int(l)-96)
	}
	a := 0
	for i := range result {
		if a >= len(xd) {
			a = 0
		}
		test := int(xd[a]) - 48
		result[i] = result[i] + test
		a++
	}
	return result
}

func main() {
	test := Encode("masterpiece", 1939)
	fmt.Println(test)
}
