package coins

func Piles(n int) int {
	n++
	p := make([]int, n)
	p[0] = 1
	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			p[j] = p[j] + p[j-i]
		}
	}
	return p[n-1]
}
