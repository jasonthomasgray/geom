package geom

import "testing"

func TestSegmentOrderTop(t *testing.T) {

	seg := Segment{
		Vertex{1, 4},
		Vertex{3, 7},
	}
	seg.OrderTop()

	if seg.a.Y < seg.b.Y {
		t.Logf("expected Segment point a to have greater Y coord than b (%v, %v)", seg.a, seg.b)
		t.Fail()
	}

	seg = Segment{
		Vertex{3, 4},
		Vertex{1, 4},
	}
	seg.OrderTop()

	if seg.a.X > seg.b.X {
		t.Logf("expected Segment point a to have lower X coord than b (%v, %v)", seg.a, seg.b)
		t.Fail()
	}
}

func TestSegmentIntersection(t *testing.T) {

	segments := []Segment{
		{Vertex{0, 2}, Vertex{0, 4}},
		{Vertex{0, -4}, Vertex{0, 7}},
	}
	FindIntersections(segments)
}
