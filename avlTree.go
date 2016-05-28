package geom

// type Comparer

type avlNode struct {
	Right, Left *avlNode
	// Value Comparer
}

func (node *avlNode) rotateLeft() (newRoot *avlNode) {

	newRoot = node.Right
	node.Right, (node.Right).Left = node.Right.Left, node
	return
}

func (node *avlNode) rotateRight() (newRoot *avlNode) {

	newRoot = node.Left
	node.Left, (node.Left).Right = node.Left.Right, node
	return
}
