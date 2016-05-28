package geom

import "sort"

func isNotRightTurn(points []Vertex) bool {
	return points[2].SideOfLine(points[0], points[1]) != 1
}

func convexUpper(points []Vertex, ch chan []Vertex) {
	var lUpper []Vertex

	lUpper = append(lUpper, points[:2]...)

	for i := 2; i < len(points); i++ {
		lUpper = append(lUpper, points[i])

		for len(lUpper) > 2 && isNotRightTurn(lUpper[len(lUpper)-3:]) {
			index := len(lUpper) - 2
			lUpper = append(lUpper[:index], lUpper[index+1:]...)
		}
	}
	ch <- lUpper
}

func convexLower(points []Vertex, ch chan []Vertex) {
	var lLower []Vertex
	lenPoints := len(points)
	lLower = append(lLower, points[lenPoints-1], points[lenPoints-2])

	for i := lenPoints - 3; i >= 0; i-- {
		lLower = append(lLower, points[i])

		for len(lLower) > 2 && isNotRightTurn(lLower[len(lLower)-3:]) {
			index := len(lLower) - 2
			lLower = append(lLower[:index], lLower[index+1:]...)
		}
	}

	ch <- lLower
}

// ConvexHull takes a slice of Vertex and returns a new slice with the bounding points in clockwise order
func ConvexHull(points []Vertex) []Vertex {

	sort.Sort(Lexographically(points))

	chUpper, chLower := make(chan []Vertex), make(chan []Vertex)
	go convexUpper(points, chUpper)
	go convexLower(points, chLower)

	upper := <-chUpper
	lower := <-chLower
	return append(upper, lower[1:len(lower)-1]...)
}
