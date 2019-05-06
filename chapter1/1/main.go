package main

import (
	"fmt"
)

// use array
func checkDuplicateChar(str string) bool {
	var checker [65536]bool
	for _, c := range str {
		if checker[c] == true {
			return false
		}
		checker[c] = true
	}
	return true
}

// use bit op, but only used from a-z
func checkDuplicateChar2(str string) bool {
	var checker uint32 = 0
	for _, c := range str {
		var b = uint8(c - 'a')
		var bb = uint32(1 << b)
		if (checker & bb) != 0 {
			return false
		}
		checker |= bb
	}
	return true
}

func main() {
	fmt.Println(checkDuplicateChar("abcdefg"))
	fmt.Println(checkDuplicateChar("abcdefgb"))
	fmt.Println(checkDuplicateChar2("aa"))

	fmt.Println("------------------------------------")

	fmt.Println(checkDuplicateChar2("abcdefg"))
	fmt.Println(checkDuplicateChar2("abcdefgb"))
	fmt.Println(checkDuplicateChar2("aa"))

	fmt.Println("------------------------------------")
}
