package missingnumbers

import "sort"

func Missing(numbers []int) []int {
	sort.Ints(numbers)
	var result []int
	for i := 0; i < len(numbers); i++ {
		if i+1 != numbers[i] {
			result = append(result, i+1)
			numbers = append(numbers, i+1)
			sort.Ints(numbers)
			i--
		}
	}
	for len(result) < 2 {
		if len(result) == 0 {
			result = append(result, len(numbers)+len(result)+1)
			result = append(result, len(numbers)+len(result)+1)
			continue
		}
		result = append(result, len(numbers)+len(result))
	}
	return result
}
