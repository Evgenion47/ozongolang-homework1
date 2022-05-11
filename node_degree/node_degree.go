package nodedegree

// Degree func
import (
	"fmt"
)

func Degree(nodes int, graph [][2]int, node int) (int, error) {
	var result int
	for i := range graph {
		if graph[i][0] == node || graph[i][1] == node {
			result++
		}
	}
	if result == 0 {
		return 0, fmt.Errorf("node %d not found in the graph", node)
	}
	return result, nil
}
