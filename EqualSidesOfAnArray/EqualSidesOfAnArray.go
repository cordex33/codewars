package main

import (
	"fmt"
)

func FindEvenIndex(arr []int) int {
	test := make(map[int]int)
	test2 := make(map[int]int)
	for i, ok := range arr {
		switch i {
		case 0:
			test[i] = ok
		case len(arr) - 1:

		default:
			test[i] = ok + test[i-1]
		}

	}

	for i, x := 0, len(arr)-1; x >= 0; i, x = i+1, x-1 {

		switch i {
		case 0:
			test2[i] = arr[x]
		case len(arr) - 1:
		default:
			test2[i] = arr[x] + test2[i-1]
		}
	}
	return -1
}

func main() {
	test := FindEvenIndex([]int{1, 100, 50, -51, 1, 1})
	fmt.Println(test)
}
