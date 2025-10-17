package main

import (
	"fmt"
	"strconv"
)

func MaxRot(n int64) int64 {

	intastr := strconv.FormatInt(n, 10)
	base := intastr
	i62, _ := strconv.ParseInt(string(base), 10, 64)
	list := []int64{i62}
	for i := range intastr {
		//esta fue craneada historica
		base = base[:i] + base[i+1:] + base[i:i+1]
		i64, _ := strconv.ParseInt(string(base), 10, 64)
		list = append(list, i64)
	}
	var result int64
	for _, n := range list {
		if n > result {
			result = n
		}
	}
	return result
}

func main() {
	test := MaxRot(56789)
	fmt.Println(test)
}
