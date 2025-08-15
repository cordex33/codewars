package main

import (
	"fmt"
	"strconv"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func Thirt(n int) int {
	list := []int{1, 10, 9, 12, 3, 4}
	num := reverse(strconv.Itoa(n))
	devolver := 0
	usado := 0
	for x := 0; len(num) > 2; x++ {
		devolver = 0
		usado = 0
		for _, xd := range num {
			a, _ := strconv.Atoi(string(xd))
			devolver = devolver + (a * list[usado])
			usado++
			if usado == 6 {
				usado = 0
			}
		}
		num = reverse(strconv.Itoa(devolver))
	}

	return devolver
}

func main() {
	test := Thirt(1234567)
	fmt.Println(test)
}
