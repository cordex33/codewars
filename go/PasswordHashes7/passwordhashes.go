package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// función encontrada en stackoverflow
func PassHash(str string) string {
	hash := md5.Sum([]byte(str))
	test := hex.EncodeToString(hash[:])
	return string(test)
}

func main() {
	test := PassHash("password")
	fmt.Println(test)
}
