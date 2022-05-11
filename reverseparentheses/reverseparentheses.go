package reverseparentheses

import "strings"

func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Reverse(s string) string {
	var start int
	var end int
	for i := range s {
		if s[i] == '(' {
			start = i
		}
		if s[i] == ')' {
			end = i
			return Reverse(strings.Join([]string{s[:start], ReverseStr(s[1+start : end]), s[end+1:]}, ""))
		}
	}
	return s
}
