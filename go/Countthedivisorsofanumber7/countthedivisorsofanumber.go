package main

import (
	"fmt"
)

func Divisors(n int) int {

	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

func main() {
	test := Divisors(54)
	fmt.Println(test)
}
