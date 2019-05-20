package main

import (
	"fmt"
	"strconv"
	"strings"
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

func stringCompression2(s string) string {
	var cs strings.Builder
	var b rune
	var i int = 0

	for _, c := range s {
		if b != c {
			if i == 0 {
				b = c
				i = 1
				continue
			}
			cs.WriteRune(b)
			cs.WriteString(strconv.Itoa(i))

			b = c
			i = 1
			continue
		}
		i++
	}

	cs.WriteRune(b)
	cs.WriteString(strconv.Itoa(i))
	return cs.String()
}

func main() {
	{
		fmt.Println(stringCompression("aabcccccaaa") == "a2b1c5a3")
	}

	{
		fmt.Println(stringCompression2("aabcccccaaa") == "a2b1c5a3")
	}
}
