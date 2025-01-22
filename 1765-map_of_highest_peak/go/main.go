package main

import (
	"container/list"
	"fmt"
)

func highestPeak(isWater [][]int) [][]int {
	numRows, numCols := len(isWater), len(isWater[0])

	answerGrid := make([][]int, numRows)
	for i := range answerGrid {
		answerGrid[i] = make([]int, numCols)
		for j := range answerGrid[i] {
			answerGrid[i][j] = -1
		}
	}

	queue := list.New()
	for rowIndex, row := range isWater {
		for colIndex, value := range row {
			if value == 1 {
				queue.PushBack([2]int{rowIndex, colIndex})
				answerGrid[rowIndex][colIndex] = 0
			}
		}
	}

	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).([2]int)
		currentRow, currentCol := current[0], current[1]

		for _, direction := range directions {
			neighborRow, neighborCol := currentRow+direction[0], currentCol+direction[1]

			if neighborRow >= 0 && neighborRow < numRows &&
				neighborCol >= 0 && neighborCol < numCols &&
				answerGrid[neighborRow][neighborCol] == -1 {

				answerGrid[neighborRow][neighborCol] = answerGrid[currentRow][currentCol] + 1
				queue.PushBack([2]int{neighborRow, neighborCol})
			}
		}
	}

	return answerGrid
}

func main() {
	isWater := [][]int{{0, 1}, {0, 0}}
	result := highestPeak(isWater)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[[1 0] [2 1]]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
