package main

import (
	"fmt"
	"strconv"
)

func PrevMultOfThree(n int) interface{} {
	if n%3 == 0 {
		return n
	}

	a := strconv.Itoa(n)
	for x := 1; x < len(a); x++ {
		test := a[:len(a)-x]
		i, _ := strconv.Atoi(test)
		if i%3 == 0 {
			return i
		}
	}
	return nil
}

func main() {
	test := PrevMultOfThree(782566)
	fmt.Println(test)
}
