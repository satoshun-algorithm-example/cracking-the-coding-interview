package main

import (
	"fmt"
	"sort"
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

func checkCharacter2(a1 string, a2 string) bool {
	if len(a1) != len(a2) {
		return false
	}
	t1 := strings.Split(a1, "")
	sort.Strings(t1)
	a1 = strings.Join(t1, "")

	t2 := strings.Split(a2, "")
	sort.Strings(t2)
	a2 = strings.Join(t2, "")

	return a1 == a2
}

func main() {
	{
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

	{
		a1 := "test"
		a2 := "tset"
		fmt.Println(checkCharacter2(a1, a2) == true)

		b1 := "testt"
		b2 := "tset"
		fmt.Println(checkCharacter2(b1, b2) == false)

		c1 := "tast"
		c2 := "tset"
		fmt.Println(checkCharacter2(c1, c2) == false)

		d1 := "aabbcceedd"
		d2 := "ddccbbaaee"
		fmt.Println(checkCharacter2(d1, d2) == true)
	}
}
