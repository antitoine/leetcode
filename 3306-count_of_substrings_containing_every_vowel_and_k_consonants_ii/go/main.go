package main

import (
	"fmt"
)

func countOfSubstrings(word string, k int) int64 {
	return substringsWithAtMost(word, k) - substringsWithAtMost(word, k-1)
}

func substringsWithAtMost(word string, k int) int64 {
	if k == -1 {
		return 0
	}
	var res int64
	vowels := 0
	uniqueVowels := 0
	vowelLastSeen := make(map[byte]int)
	l := 0
	for r := 0; r < len(word); r++ {
		if isVowel(word[r]) {
			vowels++
			if last, found := vowelLastSeen[word[r]]; !found || last < l {
				uniqueVowels++
			}
			vowelLastSeen[word[r]] = r
		}
		for r-l+1-vowels > k {
			if isVowel(word[l]) {
				vowels--
				if vowelLastSeen[word[l]] == l {
					uniqueVowels--
				}
			}
			l++
		}
		if uniqueVowels == 5 {
			minVowelIndex := min(
				vowelLastSeen['a'], vowelLastSeen['e'], vowelLastSeen['i'],
				vowelLastSeen['o'], vowelLastSeen['u'],
			)
			res += int64(minVowelIndex - l + 1)
		}
	}
	return res
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	word := "ieaouqqieaouqq"
	k := 1
	result := countOfSubstrings(word, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
