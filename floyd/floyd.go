package floyd

// Triangle makes a Floyd's triangle matrix with rows count.
func Triangle(rows int) [][]int {
	var result [][]int
	for i := 0; i < rows; i++ {
		result = append(result, nil)
	}
	lowerLeftCorner := rows*(rows-1)/2 + 1
	lastInColumn := lowerLeftCorner
	lastInRow := 1
	for i, row := 1, 1; row <= rows; i++ {
		if i < lastInRow {
			result[row-1] = append(result[row-1], i)
			lastInColumn++
		} else {
			result[row-1] = append(result[row-1], i)
			row++
			lastInRow += row
			lastInColumn = lowerLeftCorner
		}
	}
	return result
}
