package main

import (
	"fmt"
	"math"
)

// Metodo robado para ver si son primos
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Metodo robado y adaptado para ver si una lista tiene un númnero dado.
// Usaría slices.contains pero codewars al parecer tiene una versión antigua de go
func stringInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Step(g, m, n int) []int {
	result := []int{}
	for x := m; x < n; x++ {
		if isPrime(x) {
			result = append(result, x)
		}

	}

	for _, n := range result {
		cont := n + g
		if stringInSlice(cont, result) {
			return []int{n, n + g}
		}
	}

	return nil
}

func main() {
	test := Step(6, 100, 110)
	fmt.Println(test)
}
