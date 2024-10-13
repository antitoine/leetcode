package main

import (
	"fmt"
)

type Frequency struct {
	Occurrences    uint8
	DistinctValues map[int]bool
}

func maxFrequencyElements(nums []int) int {
	frequencyByNum := make(map[int]*Frequency)
	maxFrequencies := make(map[uint8]*Frequency)
	var maxFrequency *Frequency
	for _, num := range nums {
		var newFrequency *Frequency
		if frequency, found := frequencyByNum[num]; !found {
			if existingFrequency, foundFreq := maxFrequencies[1]; !foundFreq {
				freq := &Frequency{
					Occurrences:    1,
					DistinctValues: map[int]bool{num: true},
				}
				frequencyByNum[num] = freq
				maxFrequencies[1] = freq
				newFrequency = freq
			} else {
				existingFrequency.DistinctValues[num] = true
				frequencyByNum[num] = existingFrequency
				newFrequency = existingFrequency
			}
		} else {
			delete(frequency.DistinctValues, num)
			newOccurrences := frequency.Occurrences + 1
			if existingFrequency, foundFreq := maxFrequencies[newOccurrences]; !foundFreq {
				freq := &Frequency{
					Occurrences:    newOccurrences,
					DistinctValues: map[int]bool{num: true},
				}
				frequencyByNum[num] = freq
				maxFrequencies[newOccurrences] = freq
				newFrequency = freq
			} else {
				existingFrequency.DistinctValues[num] = true
				frequencyByNum[num] = existingFrequency
				newFrequency = existingFrequency
			}
		}
		if maxFrequency == nil || newFrequency.Occurrences > maxFrequency.Occurrences {
			maxFrequency = newFrequency
		}
	}
	return int(maxFrequency.Occurrences) * len(maxFrequency.DistinctValues)
}

func main() {
	input := []int{1, 2, 2, 3, 1, 4}
	result := maxFrequencyElements(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
