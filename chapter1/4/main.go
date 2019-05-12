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

func palindrome2(s string) bool {
	var index = 0
	for _, c := range s {
		var i = uint8(c - 'a')
		var j = 1 << i
		index ^= j
	}

	var permit = 0
	if len(s)%2 == 1 {
		permit = 1
	}

	for {
		if index <= 0 {
			break
		}

		var b = index & 1
		if b == 1 {
			if permit == 0 {
				return false
			}
			permit -= 1
		}
		index = index >> 1
	}

	return true
}

func main() {
	{
		fmt.Println(palindrome("tactcoa") == true)
		fmt.Println(palindrome("tactco") == false)
	}
	{
		fmt.Println(palindrome2("tactcoa") == true)
		fmt.Println(palindrome2("tactco") == false)
	}
}
