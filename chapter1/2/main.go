package main

import (
	"fmt"
	"strings"
)

func checkCharacter(a1 string, a2 string) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i, _ := range a1 {
		if strings.Count(a1, string(a1[i])) != strings.Count(a2, string(a1[i])) {
			return false
		}
	}
	return true
}

func main() {
	a1 := "test"
	a2 := "tset"
	fmt.Println(checkCharacter(a1, a2) == true)

	b1 := "testt"
	b2 := "tset"
	fmt.Println(checkCharacter(b1, b2) == false)

	c1 := "tast"
	c2 := "tset"
	fmt.Println(checkCharacter(c1, c2) == false)

	d1 := "aabbcceedd"
	d2 := "ddccbbaaee"
	fmt.Println(checkCharacter(d1, d2) == true)
}
