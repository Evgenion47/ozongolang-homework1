package buildword

import (
	"math"
	"strings"
)

var allcounts []int

func BuildWord(word string, fragments []string) int {
	var temp string
	result := math.MaxInt
	allcounts = nil
	BuildWordUtil(len(word), word, fragments, temp)
	if allcounts == nil {
		return 0
	}
	for i := range allcounts {
		if result > allcounts[i] {
			result = allcounts[i]
		}
	}
	return result
}
func BuildWordUtil(n int, s string, dic []string, temp string) {
	for i := 1; i <= n; i++ {
		prefix := s[:i]
		if Contains(dic, prefix) {
			if i == n {
				temp = strings.Join([]string{temp, prefix}, "")
				count := 1
				for j := range temp {
					if temp[j] == '*' {
						count++
					}
				}
				allcounts = append(allcounts, count)
				return
			}
			BuildWordUtil(n-i, s[i:], dic, strings.Join([]string{temp, prefix, "*"}, ""))
		}

	}
}
func Contains(s []string, a string) bool {
	for i := range s {
		if s[i] == a {
			return true
		}
	}
	return false
}
