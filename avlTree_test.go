package geom

import "testing"

func TestLeftRotate(t *testing.T) {
	a, b, c, p, q := new(avlNode), new(avlNode), new(avlNode), new(avlNode), new(avlNode)

	root := p
	p.Left, p.Right = a, q
	q.Left, q.Right = b, c

	root = root.rotateLeft()
	if root != q {
		t.Error("Tree root should be q")
	}
	if root.Left != p {
		t.Error("Root left should be p")
	}
	if root.Right != c {
		t.Error("Root right should be c")
	}
	if p.Left != a {
		t.Error("p left should be a")
	}
	if p.Right != b {
		t.Error("p right should be b")
	}
}

func TestRightRotate(t *testing.T) {
	a, b, c, p, q := new(avlNode), new(avlNode), new(avlNode), new(avlNode), new(avlNode)

	root := q
	p.Left, p.Right = a, b
	q.Left, q.Right = p, c

	root = root.rotateRight()
	if root != p {
		t.Error("Tree root should be p")
	}
	if root.Left != a {
		t.Error("Root left should be a")
	}
	if root.Right != q {
		t.Error("Root right should be q")
	}
	if q.Left != b {
		t.Error("q left should be b")
	}
	if q.Right != c {
		t.Error("q right should be c")
	}
}

func TestNilLeafRotate(t *testing.T) {
	p, q := new(avlNode), new(avlNode)

	root := p

	p.Right = q
	root = root.rotateLeft()
	if root != q {
		t.Error("rotateLeft failed")
	}

	root = root.rotateRight()
	if root != p {
		t.Error("rotateRight failed")
	}
}
