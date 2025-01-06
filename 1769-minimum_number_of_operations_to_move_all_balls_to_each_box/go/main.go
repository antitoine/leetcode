package main

import (
	"fmt"
)

func minOperations(boxes string) []int {
	answer := make([]int, len(boxes))

	ascBalls := make([]int, len(boxes))
	if boxes[0] == '1' {
		ascBalls[0] = 1
	}
	ascAnswer := make([]int, len(boxes))
	for i := 1; i < len(boxes); i++ {
		ascBalls[i] = ascBalls[i-1]
		ascAnswer[i] = ascBalls[i-1] + ascAnswer[i-1]
		if boxes[i] == '1' {
			ascBalls[i]++
		}
		answer[i] = ascAnswer[i]
	}

	descBalls := make([]int, len(boxes))
	if boxes[len(boxes)-1] == '1' {
		descBalls[len(boxes)-1] = 1
	}
	descAnswer := make([]int, len(boxes))
	for i := len(boxes) - 2; i >= 0; i-- {
		descBalls[i] = descBalls[i+1]
		descAnswer[i] = descAnswer[i+1] + descBalls[i+1]
		if boxes[i] == '1' {
			descBalls[i]++
		}
		answer[i] += descAnswer[i]
	}

	return answer
}

func main() {
	boxes := "001011"
	result := minOperations(boxes)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[11 8 5 4 3 4]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
