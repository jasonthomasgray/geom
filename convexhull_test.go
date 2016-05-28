package geom

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestConvexHull(t *testing.T) {

	// points := []Vertex{
	// 	{1, 7}, {7, 1}, {1, 8}, {5, 0}, {4, 0}, {5, 1}, {1, 9}, {8, 4}, {8, 5}, {1, 6},
	// }

	points := make([]Vertex, 20)

	for i := range points {
		points[i].X = math.Floor(rand.NormFloat64()*2 + 5)
		points[i].Y = math.Floor(rand.NormFloat64()*2 + 5)
	}

	fmt.Println(points)
	fmt.Println(ConvexHull(points))
}
