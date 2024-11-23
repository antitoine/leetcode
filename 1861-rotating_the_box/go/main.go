package main

import (
	"fmt"
)

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	rotatedBox := make([][]byte, n)
	for i := 0; i < n; i++ {
		rotatedBox[i] = make([]byte, m)
	}
	for rowIdx := 0; rowIdx < m; rowIdx++ {
		rotatedColIdx := m - 1 - rowIdx
		for colIdx := n - 1; colIdx >= 0; colIdx-- {
			rotatedRowIdx := colIdx
			rotatedBox[rotatedRowIdx][rotatedColIdx] = box[rowIdx][colIdx]
			if string(rotatedBox[rotatedRowIdx][rotatedColIdx]) == "#" {
				i := rotatedRowIdx + 1
				for i < n && string(rotatedBox[i][rotatedColIdx]) == "." {
					i++
				}
				if i != rotatedRowIdx+1 {
					rotatedBox[i-1][rotatedColIdx], rotatedBox[rotatedRowIdx][rotatedColIdx] = rotatedBox[rotatedRowIdx][rotatedColIdx], rotatedBox[i-1][rotatedColIdx]
				}
			}
		}
	}
	return rotatedBox
}

func main() {
	box := [][]byte{
		{'#', '#', '*', '.', '*', '.'},
		{'#', '#', '#', '*', '.', '.'},
		{'#', '#', '#', '.', '#', '.'},
	}
	result := rotateTheBox(box)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[[46 35 35] [46 35 35] [35 35 42] [35 42 46] [35 46 42] [35 46 46]]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
