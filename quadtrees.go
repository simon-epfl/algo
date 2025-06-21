package main

import "math"

type CustomQuadTreeNode struct {
	X, Y        int // coordonnées du centre du noeud
	TopLeft     *CustomQuadTree
	TopRight    *CustomQuadTree
	BottomLeft  *CustomQuadTree
	BottomRight *CustomQuadTree
	Content     any
}

type CustomQuadTree struct {
	Root *CustomQuadTreeNode
}

func (cq *CustomQuadTree) Search(x, y int) any {
	q := cq.Root
	if q == nil {
		return nil
	}

	if x == q.X && y == q.Y {
		return q.Content
	}

	if x < 0 && y < 0 {
		return q.TopLeft.Search(x, y)
	} else if x >= 0 && y < 0 {
		return q.TopRight.Search(x, y)
	} else if x < 0 && y >= 0 {
		return q.BottomLeft.Search(x, y)
	} else {
		return q.BottomRight.Search(x, y)
	}
}

func (cq *CustomQuadTree) Southmost() int { // renvoie la coordonnée Y du point le plus bas
	if cq.Root == nil {
		return math.MaxInt
	}
	node := cq.Root

	w, e := math.MaxInt, math.MaxInt

	if node.TopLeft != nil {
		w = node.TopLeft.Southmost()
	}

	if node.TopRight != nil {
		e = node.TopRight.Southmost()
	}

	if w < e && w < node.Y {
		return w
	}

	if e < w && e < node.Y {
		return e
	}

	if node.Y < w && node.Y < e {
		return node.Y
	}

	return math.MaxInt

}
