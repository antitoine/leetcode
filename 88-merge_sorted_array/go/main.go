package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	mIdx := m - 1
	nIdx := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if (mIdx >= 0 && nIdx >= 0 && nums1[mIdx] > nums2[nIdx]) || nIdx < 0 {
			nums1[i] = nums1[mIdx]
			mIdx--
		} else {
			nums1[i] = nums2[nIdx]
			nIdx--
		}
	}
}

func main() {
	nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
	merge(nums1, m, nums2, n)
	fmt.Printf("Result: %#v\n", nums1)
	if fmt.Sprintf("%v", nums1) == "[1 2 2 3 5 6]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
