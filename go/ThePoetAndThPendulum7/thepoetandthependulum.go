package main

import (
	"fmt"
	"slices"
)

func Pendulum(values []int) []int {
	slices.Sort(values)
	test := []int{}
	qwe := true
	for _, n := range values {
		if !qwe {
			test = append(test, n)
			qwe = true
		} else {
			test = append([]int{n}, test...)
			qwe = false
		}
	}
	return test
}

func main() {
	test := Pendulum([]int{22, 26, 21, 27, 24, 21, 15, 26, 25})
	fmt.Println(test)
}
