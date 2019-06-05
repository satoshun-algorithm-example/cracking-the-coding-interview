package main

import "fmt"

func removeDuplcatesElement(elements []interface{}) {
	i := 0
	for {
		if i >= len(elements) {
			break
		}
		e := elements[i]
		for j := i + 1; ; {
			if j >= len(elements) {
				break
			}
			if elements[j] == e {
				fmt.Println("100")
				elements = remove(elements, j)
				continue
			}
			j++
		}
		i++
	}

	fmt.Println(elements)
}

func remove(s []interface{}, i int) []interface{} {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	{
		removeDuplcatesElement(
			[]interface{}{"1", "2", "1", "3"},
		)
	}
}
