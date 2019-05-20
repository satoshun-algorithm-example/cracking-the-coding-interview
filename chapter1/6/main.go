package main

import (
	"fmt"
	"strconv"
)

func stringCompression(s string) string {
	var cs string
	var b rune
	var i int = 0

	for _, c := range s {
		if b != c {
			if i == 0 {
				b = c
				i = 1
				continue
			}
			cs += string(b) + strconv.Itoa(i)

			b = c
			i = 1
			continue
		}
		i++
	}

	cs += string(b) + strconv.Itoa(i)
	return cs
}

func main() {
	fmt.Println(stringCompression("aabcccccaaa") == "a2b1c5a3")
}
