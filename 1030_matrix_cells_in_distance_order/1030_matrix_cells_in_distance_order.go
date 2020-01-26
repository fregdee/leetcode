import (
	"math"
	"sort"
)


func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	distanceMap := make(map[float64][][]int)

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			// the Manhattan distance ( |r1 - r2| + |c1 - c2| )
			distance := math.Abs(float64(r - r0)) + math.Abs(float64(c - c0))
			distanceMap[distance] = append(distanceMap[distance], []int{r, c})
		}
	}

	var distances []float64

	for k, _ := range distanceMap {
		distances = append(distances, k)
	}

	sort.Float64s(distances)

	var result [][]int

	for _, i := range distances {
		for _, v := range distanceMap[i] {
			result = append(result, v)
		}
	}

	return result
}
