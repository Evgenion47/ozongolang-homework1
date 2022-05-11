package secretmessage

// Decode func
import "sort"

type tuple struct {
	symbol       rune
	appearsCount int
}

func Decode(encoded string) string {
	var dict []tuple
	var result []rune
	encodedRunes := []rune(encoded)
	sort.Slice(encodedRunes, func(i, j int) bool { return encodedRunes[i] < encodedRunes[j] })

	var count int

	for i := 0; i < len(encodedRunes); i++ {
		if i < len(encodedRunes)-1 {
			if encodedRunes[i] == encodedRunes[i+1] {
				count++
			} else {
				dict = append(dict, tuple{encodedRunes[i], count + 1})
				count = 0
			}
		} else {
			if encodedRunes[i-1] == encodedRunes[i] {
				count++
				dict = append(dict, tuple{encodedRunes[i], count})
			} else {
				dict = append(dict, tuple{encodedRunes[i], 1})
			}
		}
	}
	sort.Slice(dict, func(i, j int) bool { return dict[i].appearsCount > dict[j].appearsCount })
	for i := range dict {
		if dict[i].symbol == '_' {
			return string(result)
		}
		result = append(result, dict[i].symbol)
	}

	return ""
}
