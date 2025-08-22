package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	// Implement me! :)
	if text == ""{
		return ""
	}
	spit := strings.Split(text, " ")
	xd := []string{}
	wordencript := ""
	for _, word := range spit {
		if len(string(word)) > 2 {
			wordencript = strconv.Itoa(int(rune(word[0]))) + string(word[len(word)-1]) + word[2:len(word)-1] + word[1:2]
			xd = append(xd, wordencript)
		} else if len(string(word)) > 1 {
			wordencript = strconv.Itoa(int(rune(word[0]))) + word[1:]
			xd = append(xd, wordencript)
		} else {
			wordencript = strconv.Itoa(int(rune(word[0])))
			xd = append(xd, wordencript)
		}

	}
	wordencript = strings.Join(xd, " ")
	return wordencript
}

func main() {
	test := EncryptThis("The more he saw the less he spoke")
	fmt.Println(test)
}
