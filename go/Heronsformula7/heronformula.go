package main

import (
	"fmt"
	"math"
)

func Heron(a, b, c int) float64 {
	q, w, e := float64(a), float64(b), float64(c)
	s := (q + w + e) / 2
	result := math.Sqrt(s * (s - q) * (s - w) * (s - e))
	return result
}

func main() {
	s := Heron(3, 4, 5)
	fmt.Println(s)
}
