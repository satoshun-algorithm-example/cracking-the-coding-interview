package main

import "fmt"

func oneShotTranslate(s1 string, s2 string) bool {
	if len(s1) < len(s2) {
		tmp := s2
		s2 = s1
		s1 = tmp
	}
	if len(s1)-len(s2) == 0 {
		var c = 0
		for i, _ := range s1 {
			if s1[i] != s2[i] {
				c++
			}
		}
		return c <= 1
	}
	if len(s1)-len(s2) == 1 {
		var c = 0
		for i, _ := range s1 {
			if len(s2) == (i - c) {
				c++
				break
			}
			if s1[i] != s2[i-c] {
				c++
				if c == 2 {
					break
				}
			}
		}
		return c == 1
	}
	return false
}

func main() {
	fmt.Println(oneShotTranslate("pale", "ple") == true)
	fmt.Println(oneShotTranslate("pales", "pale") == true)
	fmt.Println(oneShotTranslate("pale", "bale") == true)
	fmt.Println(oneShotTranslate("pale", "bake") == false)
}
