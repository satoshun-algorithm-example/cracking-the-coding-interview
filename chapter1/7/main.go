package main

import "fmt"

func maxtrixRotation(a [][]int) [][]int {
	// TODO NxN matrix
	b := [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	for i, y := range a {
		for j, x := range y {
			b[j][3-i] = x
		}
	}
	for _, bb := range b {
		fmt.Println(bb)
	}
	return b
}

func main() {
	{
		maxtrixRotation([][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		})
	}
}
