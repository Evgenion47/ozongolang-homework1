package warriors

import (
	"math"
	"strings"
)

func dfs(img [][]bool, i int, j int) [][]bool {
	img[i][j] = false
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			if img[k][l] == true {
				dfs(img, k, l)
			}
		}
	}
	return img
}

func Count(image string) int {
	var result int
	imgStrSl := strings.Split(image, "\n")
	n := int(math.Sqrt(float64(len(image))))
	slImage := make([][]bool, n+2)
	slImage[0] = make([]bool, n+2)
	slImage[n+1] = make([]bool, n+2)
	for i := range imgStrSl {
		slImage[i+1] = make([]bool, n+2)
		for j, r := range imgStrSl[i] {
			slImage[i+1][j+1] = r == '1'
		}
	}
	for i := 0; i < len(slImage); i++ {
		for j := 0; j < len(slImage[i]); j++ {
			if slImage[i][j] == true {
				dfs(slImage, i, j)
				result++
			}
		}
	}

	return result
}
