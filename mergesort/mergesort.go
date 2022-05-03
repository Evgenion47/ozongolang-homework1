package mergesort

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	first := MergeSort(input[:len(input)/2])
	second := MergeSort(input[len(input)/2:])
	return Merge(first, second)
}
func Merge(a []int, b []int) []int {
	var result []int
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	for j < len(b) {
		result = append(result, b[j])
		j++
	}
	return result
}
