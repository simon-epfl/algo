package main

import "fmt"

type BinarySearchTree struct {
	Parent  *BinarySearchTree
	Left    *BinarySearchTree
	Right   *BinarySearchTree
	Content int // juste pour que ce soit facile à comparer...
}

func (t *BinarySearchTree) Search(x int) *BinarySearchTree {
	if t == nil {
		return nil
	}
	if t.Content == x {
		return t
	}
	if x < t.Content {
		return t.Left.Search(x) // même si t.Left est nul ça va appeler la fonction avec nil!
	}
	return t.Right.Search(x)
}

func (t *BinarySearchTree) Minimum() *BinarySearchTree {
	for t.Left != nil {
		t = t.Left
	}
	return t
}

func (t *BinarySearchTree) Maximum() *BinarySearchTree {
	for t.Right != nil {
		t = t.Right
	}
	return t
}

func (t *BinarySearchTree) Successor() *BinarySearchTree {

	if t.Right != nil {
		return t.Right.Minimum()
	}

	// on remonte jusqu'au root
	// OU on remonte jusqu'au noeud telle que notre x était un enfant à gauche (donc plus petit) du parent
	currentTree := t
	currentParent := t.Parent

	for currentParent != nil && currentTree == currentParent.Right {
		currentTree = currentParent
		currentParent = currentParent.Parent
	}

	return currentParent
}

func (t *BinarySearchTree) PrintInorder() {
	if t.Left != nil {
		t.Left.PrintInorder()
	}

	fmt.Println(t.Content)

	if t.Right != nil {
		t.Right.PrintInorder()
	}
}

func (t *BinarySearchTree) Insert(content int) {

	// STEP1 : on cherche l'endroit d'insertion
	var insertionPosition *BinarySearchTree
	currentTree := t
	for currentTree != nil {
		insertionPosition = currentTree
		if content < currentTree.Content {
			currentTree = currentTree.Left
		} else {
			currentTree = currentTree.Right
		}
	}

	// STEP2 : insertion
	newTree := &BinarySearchTree{
		Parent:  insertionPosition, // on ajoute le nouveau noeud à notre arbre
		Content: content,
	}

	if newTree.Content < newTree.Parent.Content {
		newTree.Parent.Left = newTree
	} else {
		newTree.Parent.Right = newTree
	}
}

// Transplant remplace l'arbre U par l'arbre V
func (t *BinarySearchTree) Transplant(u, v *BinarySearchTree) {
	if u.Parent == nil { // si u n'a pas de parent, c'était le root
		*t = *v
	} else if u == u.Parent.Left { // si u était l'enfant à gauche
		u.Parent.Left = v // on met v à la place dans le parent
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent // enfin, on set le parent de v à u
	}
}

func (t *BinarySearchTree) Delete(z *BinarySearchTree) {

	// Cas 1 : pas d'enfant à gauche
	if z.Left == nil {
		t.Transplant(z, z.Right)
		return
	}

	// Cas 2 : pas d'enfant à droite
	if z.Right == nil {
		t.Transplant(z, z.Left)
		return
	}

	// Case 3 : deux enfants, on cherche le successor
	successor := z.Successor()

	if successor.Parent != z {
		t.Transplant(successor, successor.Right) // on fait monter l'enfant droit du successor
		// par définition, le successor n'a pas d'enfant à gauche!! donc on ne s'en préoccupe pas
		successor.Right = z.Right
		successor.Right.Parent = successor
	}

	t.Transplant(z, successor)
	successor.Left = z.Left
	successor.Left.Parent = successor

}
