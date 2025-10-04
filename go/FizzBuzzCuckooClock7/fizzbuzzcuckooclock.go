package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) string {
	//con las variable abcd, son las horas por ejemplo:
	// 12:00 = la a representa el 1, b el 2, c el 0 y d finalmente el último 0
	ab, _ := strconv.Atoi(time[len(time)-5 : len(time)-3])
	_ = ab
	cd, _ := strconv.Atoi(time[len(time)-2:])
	_ = cd
	if cd%3 != 0 && cd%5 != 0 {
		return "tick"
	} else if cd%3 == 0 && cd%5 == 0 && ab == 00 || ab == 12 && cd == 00 {
		return "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"
	} else if cd%3 == 0 && cd%5 == 0 && cd == 00 {
		if ab > 12 {
			test := strings.Repeat("Cuckoo ", ab-12)
			return test[:len(test)-1]
		} else if ab <= 12 {
			test := strings.Repeat("Cuckoo ", ab)
			return test[:len(test)-1]
		}
	} else if cd%3 == 0 && cd%5 == 0 && cd == 30 {
		return "Cuckoo"
	} else if cd%3 == 0 && cd%5 == 0 {
		return "Fizz Buzz"
	} else if cd%3 == 0 {
		return "Fizz"
	} else {
		return "Buzz"
	}

	return ""
}

func main() {
	test := FizzBuzzCuckooClock("12:45")
	fmt.Println(test)
}
