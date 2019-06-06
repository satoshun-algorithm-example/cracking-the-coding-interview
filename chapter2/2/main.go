package main

import "fmt"

func exploreBehindK(k int, v []interface{}) interface{} {
	n := 0
	for {
		a := _getOrNil(v, n)
		if a == nil {
			break
		}
		n++
	}
	return v[n-k-1]
}

// simulates a Singly Linked Lists
func _getOrNil(v []interface{}, index int) interface{} {
	if len(v) <= index {
		return nil
	}
	return v[index]
}

func main() {
	{
		fmt.Println(
			exploreBehindK(
				1,
				[]interface{}{
					"a",
					"b",
					"c",
					"d",
					"e",
				},
			),
		)

		fmt.Println(
			exploreBehindK(
				3,
				[]interface{}{
					"a",
					"b",
					"c",
					"d",
					"e",
				},
			),
		)
	}
}
