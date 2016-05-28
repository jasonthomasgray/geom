package geom

import "math"

// Vertex represents a point in 2 dimensions
type Vertex struct {
	X, Y float64
}

// Vertices is a slice of Vertex
type Vertices []Vertex

// Convert takes Vertices and flattens them into a float32 slice
func (l Vertices) Convert() (data []float32) {

	data = make([]float32, len(l)*2)
	for i, v := range l {
		index := i * 2
		data[index] = float32(v.X)
		data[index+1] = float32(v.Y)
	}
	return
}

// SideOfLine checks which side of the line ab the Vertex is on
func (c Vertex) SideOfLine(a Vertex, b Vertex) int {

	cross := float64((b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X))

	if cross == 0 {
		return 0
	}
	if math.Signbit(cross) {
		return 1
	}
	return -1

}

// Lexographically is a sorter to sort Vertexs by X coordinate
type Lexographically []Vertex

func (a Lexographically) Len() int      { return len(a) }
func (a Lexographically) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Lexographically) Less(i, j int) bool {
	if a[i].X == a[j].X {
		return a[i].Y < a[j].Y
	}
	return a[i].X < a[j].X
}
