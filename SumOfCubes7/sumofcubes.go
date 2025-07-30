package main

import (
	"fmt"
)

func SumCubes(n int) int {
	result := 0
	for x := 0; x <= n; x++ {
		result = result + (x * x * x)
	}
	return result
}

func main() {
	test := SumCubes(1)
	fmt.Println(test)
}
