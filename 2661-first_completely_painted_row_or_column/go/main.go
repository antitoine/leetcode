package main

import (
	"fmt"
	"image"
)

func firstCompleteIndex(arr []int, mat [][]int) int {
	positions := make([]image.Point, len(arr))
	for y := 0; y < len(mat); y++ {
		for x := 0; x < len(mat[y]); x++ {
			positions[mat[y][x]-1] = image.Pt(x, y)
		}
	}
	lineLength := len(mat)
	columnLength := len(mat[0])
	lines := make([]int, lineLength)
	columns := make([]int, columnLength)
	for i, val := range arr {
		pos := positions[val-1]
		lines[pos.Y]++
		if lines[pos.Y] == columnLength {
			return i
		}
		columns[pos.X]++
		if columns[pos.X] == lineLength {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 4, 5, 2, 6, 3}
	mat := [][]int{{4, 3, 5}, {1, 2, 6}}
	result := firstCompleteIndex(arr, mat)
	fmt.Printf("Result: %#v\n", result)
	if result == 1 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
