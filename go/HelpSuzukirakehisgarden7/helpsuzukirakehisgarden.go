package main

import (
	"fmt"
	"strings"
)

func RakeGarden(garden string) string {
	spl := strings.Split(garden, " ")
	for i, w := range spl {
		if string(w) != "gravel" && string(w) != "rock" {
			spl[i] = "gravel"
		}
	}
	return strings.Join(spl, " ")
}

func main() {
	test := RakeGarden("slug spider rock gravel gravel gravel gravel gravel gravel gravel")
	fmt.Println(test)
}
