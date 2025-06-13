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
