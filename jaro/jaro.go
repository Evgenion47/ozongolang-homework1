package jaro

import (
	"math"
	"strings"
)

func Distance(word1 string, word2 string) float64 {
	word1, word2 = strings.ToLower(word1), strings.ToLower(word2)
	if word1 == word2 {
		return 1
	}
	len1 := len(word1)
	len2 := len(word2)
	maxDist := math.Floor((math.Max(float64(len1), float64(len2)) / 2) - 1)
	match := 0
	hs1 := make([]int, len1)
	hs2 := make([]int, len2)

	for i := 0; i < len1; i++ {
		for j := int(math.Max(0, float64(i)-maxDist)); j < int(math.Min(float64(len2), float64(i)+maxDist+1)); j++ {
			if word1[i] == word2[j] && hs2[j] == 0 {
				hs1[i] = 1
				hs2[j] = 1
				match++
				break
			}
		}
	}
	if match == 0 {
		return 0
	}
	t := 0.0
	point := 0
	for i := 0; i < len1; i++ {
		if hs1[i] == 1 {
			for hs2[point] == 0 {
				point++
			}
			if word1[i] != word2[point] {
				t++
			}
			point++
		}
	}
	t /= 2
	return ((float64(match) / float64(len1)) + (float64(match) / float64(len2)) + ((float64(match) - t) / float64(match))) / 3.0
}
