package main

import (
	"fmt"
)

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	result := make([][]int, 0, len(nums1)+len(nums2))
	var i, j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] < nums2[j][0] {
			result = append(result, []int{nums1[i][0], nums1[i][1]})
			i++
		} else if nums1[i][0] > nums2[j][0] {
			result = append(result, []int{nums2[j][0], nums2[j][1]})
			j++
		} else {
			result = append(result, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		}
	}
	for i < len(nums1) {
		result = append(result, []int{nums1[i][0], nums1[i][1]})
		i++
	}
	for j < len(nums2) {
		result = append(result, []int{nums2[j][0], nums2[j][1]})
		j++
	}
	return result
}

func main() {
	nums1 := [][]int{{1, 2}, {2, 3}, {4, 5}}
	nums2 := [][]int{{1, 4}, {3, 2}, {4, 1}}
	result := mergeArrays(nums1, nums2)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[[1 6] [2 3] [3 2] [4 6]]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
