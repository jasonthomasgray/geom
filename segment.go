package geom

import "sort"

// Segment represents a line between two Vertex
type Segment struct {
	a, b Vertex
}

// OrderTop makes sure Vertex a of the segment is above Vertex b.
// If they have the same Y coord then X is sorted.
func (seg *Segment) OrderTop() {

	if seg.a.Y < seg.b.Y || (seg.a.Y == seg.b.Y && seg.a.X > seg.b.X) {
		// swap points
		seg.a, seg.b = seg.b, seg.a
	}
}

type orderSegments []Segment

func (s orderSegments) Len() int      { return len(s) }
func (s orderSegments) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s orderSegments) Less(i, j int) bool {
	if s[i].a.Y == s[j].a.Y {
		return s[i].a.X < s[j].a.X
	}
	return s[i].a.Y > s[j].a.Y
}

// Intersection is a Vertex representing an intersection point with a list of segments that contain that point.
type Intersection struct {
	Vertex
	segments []*Segment
}

// FindIntersections takes a slice of Segments an gives back a slice of Intersections
func FindIntersections(segments []Segment) {

	// Ensure segment points are ordered correctly.
	for i := range segments {
		segments[i].OrderTop()
	}
	sort.Sort(orderSegments(segments))

}
