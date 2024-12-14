package main

import (
	"fmt"
)

func continuousSubarrays(nums []int) int64 {
	var ans int64 = 1
	left := nums[0] - 2
	right := nums[0] + 2
	l := 0

	for r := 1; r < len(nums); r++ {
		if left <= nums[r] && nums[r] <= right {
			left = max(left, nums[r]-2)
			right = min(right, nums[r]+2)
		} else {
			left = nums[r] - 2
			right = nums[r] + 2
			l = r

			for l >= 0 && nums[r]-2 <= nums[l] && nums[l] <= nums[r]+2 {
				left = max(left, nums[l]-2)
				right = min(right, nums[l]+2)
				l--
			}
			l++
		}

		ans += int64(r - l + 1)
	}

	return ans
}

func main() {
	nums := []int{5, 4, 2, 4}
	result := continuousSubarrays(nums)
	fmt.Printf("Result: %#v\n", result)
	if result == 8 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
