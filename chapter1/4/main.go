package main

import (
	"fmt"
	"strings"
)

func palindrome(s string) bool {
	var permit = 0
	if len(s)%2 == 1 {
		permit = 1
	}
	for _, c := range s {
		if strings.Count(s, string(c))%2 == 1 {
			permit -= 1
			if permit < 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(palindrome("tactcoa") == true)
	fmt.Println(palindrome("tactco") == false)
}
