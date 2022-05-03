package anagram

import (
	"sort"
	"strings"
)

func convertToComparable(input string) string {
	runeArray := []rune(strings.ToLower(strings.Replace(input, " ", "", -1)))
	sort.Slice(runeArray, func(i, j int) bool { return runeArray[i] < runeArray[j] })

	return string(runeArray)
}

func FindAnagrams(dictionary []string, word string) (result []string) {
	var answer []string
	convertedWord := convertToComparable(word)
	if convertedWord == "" {
		return []string{}
	}
	for i := range dictionary {
		if convertedWord == convertToComparable(dictionary[i]) && strings.ToLower(dictionary[i]) != strings.ToLower(word) {
			answer = append(answer, dictionary[i])
		}
	}

	return answer
}
