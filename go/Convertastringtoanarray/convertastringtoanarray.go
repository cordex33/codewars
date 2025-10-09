package main

import (
	"fmt"
	"strings"
)

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}

func main() {
	test := StringToArray("I love arrays they are my favorite")
	fmt.Println(test)
}
