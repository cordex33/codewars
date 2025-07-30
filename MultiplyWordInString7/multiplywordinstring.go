package main

import (
	"fmt"
	"strings"
)

func ModifyMultiply(str string, loc, num int) string {
	xd := strings.Split(str, " ")
	test := ""
	if num == 0 {
		return xd[loc]
	}
	test = strings.Repeat(xd[loc]+"-", num)
	test = test[:len(test)-1]
	return test
}

func main() {
	test := ModifyMultiply("EfgHOkfRJAEdldPQ wieWg UOJdDCjFThoHEMCDZTK aGuRYSBzeKzuzteuQ", 3, 0)
	fmt.Println(test)
}
