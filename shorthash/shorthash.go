package shorthash

import "strings"

var ShortHash []string

func GenerateShortHashes(dictionary string, len int) []string {
	if dictionary == "" || len == 0 {
		return []string{}
	}
	ShortHash = nil
	for i := 1; i < len+1; i++ {
		SetByLength(i, dictionary)
	}
	return ShortHash
}

func FoundNext(prevPointer *[]int, base int) bool {
	prev := *prevPointer
	for i := len(prev) - 1; i >= 0; i-- {
		if prev[i]+1 < base {
			prev[i]++
			return true
		} else {
			prev[i] = 0
		}
	}
	return false
}

func SetByLength(length int, dic string) {
	hNext := true
	runeDic := []rune(dic)
	alPower := len(runeDic)
	a := make([]int, length)
	builder := strings.Builder{}

	for hNext {
		for i := range a {
			builder.WriteRune(runeDic[a[i]])
		}
		ShortHash = append(ShortHash, builder.String())
		builder.Reset()
		hNext = FoundNext(&a, alPower)
	}
}
