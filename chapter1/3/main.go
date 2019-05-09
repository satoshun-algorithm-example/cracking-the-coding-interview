package main

import "fmt"

func urlify(s string) string {
	space := 0
	for _, c := range s {
		if c == ' ' {
			space += 1
		}
	}

	var ns []rune
	ns = make([]rune, len(s)+space*2)
	i := 0
	for _, c := range s {
		if c == ' ' {
			ns[i] = '%'
			ns[i+1] = '2'
			ns[i+2] = '0'
			i += 3
		} else {
			ns[i] = c
			i += 1
		}
	}

	return string(ns)
}

func main() {
	fmt.Println(urlify("Mr John Smith") == "Mr%20John%20Smith")
	fmt.Println(urlify("MrJohn") == "MrJohn")
}
