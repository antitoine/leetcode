package main

import (
	"fmt"
)

func abs(v int) int64 {
	if v > 0 {
		return int64(v)
	}
	return int64(-v)
}

func maxMatrixSum(matrix [][]int) int64 {
	var nbNegative int64
	minValAbs := int64(10e5)
	var sumAbs int64
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix); c++ {
			v := matrix[r][c]
			if v < 0 {
				nbNegative++
			}
			vAbs := abs(v)
			sumAbs += vAbs
			if vAbs < minValAbs {
				minValAbs = vAbs
			}
		}
	}
	if nbNegative == 0 || nbNegative%2 == 0 {
		return sumAbs
	}
	return sumAbs - (2 * minValAbs)
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{-1, -2, -3},
		{1, 2, 3},
	}
	result := maxMatrixSum(matrix)
	fmt.Printf("Result: %#v\n", result)
	if result == 16 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
