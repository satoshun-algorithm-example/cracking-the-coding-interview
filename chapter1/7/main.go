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

func maxtrixRotation2(a [][]int) [][]int {
	n := len(a)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := a[first][i]

			a[first][i] = a[last-offset][first]
			a[last-offset][first] = a[last][last-offset]
			a[last][last-offset] = a[i][last]
			a[i][last] = top
		}
	}

	for _, bb := range a {
		fmt.Println(bb)
	}
	return a
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
	fmt.Println()
	{
		maxtrixRotation2([][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		})
	}
}
