package kata

import (
	"sort"
)

func TwoToOne(a, b string) string {
	// put all chars of a in map
	// put all chars of b in map
	charsMap := make(map[rune]bool)
	for _, char := range a + b {
		charsMap[char] = true
	}

	// make slice out of map
	var chars []rune
	for char := range charsMap {
		chars = append(chars, char)
	}

	// sort slice
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	// join into str
	return string(chars)
}
