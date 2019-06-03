package main

import (
	"fmt"
	"strings"
)

func isRotationString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return isSubString(s1+s1, s2)
}

func isSubString(s1 string, s2 string) bool {
	return strings.Contains(s1, s2)
}

func main() {
	{
		fmt.Println(isRotationString("waterbottle", "erbottlewat") == true)
		fmt.Println()
		fmt.Println(isRotationString("waterbottlea", "erbottlewata") == false)
	}
}
