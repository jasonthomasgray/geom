package geom

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
)

func TestPointSort(t *testing.T) {

	points := []Vertex{
		{1, 7}, {7, 1}, {1, 8}, {5, 0}, {4, 0}, {5, 1}, {1, 9}, {8, 4}, {8, 5}, {1, 6},
	}

	expected := []Vertex{
		{1, 6}, {1, 7}, {1, 8}, {1, 9}, {4, 0}, {5, 0}, {5, 1}, {7, 1}, {8, 4}, {8, 5},
	}
	sort.Sort(Lexographically(points))

	for i, v := range points {
		if v != expected[i] {
			t.Logf("Expected value %v at %v to equal %v.", v, i, expected[i])
			t.Fail()
		}
	}
}

func TestSideOfLine(t *testing.T) {

	p, a, b := Vertex{-1, 1}, Vertex{0, 0}, Vertex{0, 1}
	var result int

	if result = p.SideOfLine(a, b); result != -1 {
		t.Errorf("Expected point (%v) to be to left (-1) got [%v]", p, result)
	}

	if result = p.SideOfLine(b, a); result != 1 {
		t.Errorf("Expected point (%v) to be to right (1) got [%v]", p, result)
	}

	p = Vertex{0, 2}

	if result = p.SideOfLine(a, b); result != 0 {
		t.Errorf("Expected point (%v) to be to left (-1) got [%v]", p, result)
	}
}

func TestConvexHull(t *testing.T) {

	// points := []Vertex{
	// 	{1, 7}, {7, 1}, {1, 8}, {5, 0}, {4, 0}, {5, 1}, {1, 9}, {8, 4}, {8, 5}, {1, 6},
	// }

	points := make([]Vertex, 20)

	for i := range points {
		points[i].X = int(math.Floor(rand.NormFloat64()*2 + 5))
		points[i].Y = int(math.Floor(rand.NormFloat64()*2 + 5))
	}

	fmt.Println(points)
	fmt.Println(ConvexHull(points))
}
