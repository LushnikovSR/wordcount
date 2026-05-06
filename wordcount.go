package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	args := os.Args
	count := 0

	for _, arg := range args {
		elems := strings.Fields(arg)
		for _, elem := range elems {
			if isOnlyLetters(elem) {
				count++
			}
		}
	}
	fmt.Println("count: ", count)
}

func isOnlyLetters(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, char := range s {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
