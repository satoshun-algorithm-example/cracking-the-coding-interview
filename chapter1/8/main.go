package main

import "fmt"

func fillZero(matrix [][]int) {
	// copy two dimensions slice
	m := [][]int{}
	for i := range matrix {
		mm := make([]int, len(matrix[i]))
		copy(mm, matrix[i])
		m = append(m, mm)
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == 0 {
				// fill x direction
				for xd := 0; xd < len(matrix[y]); xd++ {
					m[y][xd] = 0
				}

				// fill y direction
				for yd := 0; yd < len(matrix); yd++ {
					m[yd][x] = 0
				}
			}
		}
	}

	for i, _ := range m {
		fmt.Println(m[i])
	}
}

func fillZero2(matrix [][]int) {
	column := len(matrix)
	row := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = -1
			}
		}
	}
	for i := 0; i < column; i++ {
		for j := 0; j < row; j++ {
			if matrix[j][i] == 0 {
				matrix[0][i] = -1
			}
		}
	}
	for i := 0; i < row; i++ {
		if matrix[i][0] == -1 {
			for j := 0; j < column; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < column; i++ {
		if matrix[0][i] == -1 {
			for j := 0; j < row; j++ {
				matrix[j][i] = 0
			}
		}
	}

	for i, _ := range matrix {
		fmt.Println(matrix[i])
	}
}

func main() {
	{
		fillZero([][]int{
			{1, 2, 3, 4},
			{5, 6, 0, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		})
	}
	fmt.Println()
	{
		fillZero2([][]int{
			{1, 2, 3, 4},
			{5, 6, 0, 8},
			{9, 0, 11, 12},
			{13, 14, 15, 16},
		})
	}
}
