package main

type CustomHeapItem struct {
	Value   int
	content any
}

type CustomHeap []CustomHeapItem

func (h CustomHeap) Left(i int) int {
	return 2*i + 1
}

func (h CustomHeap) Right(i int) int {
	return 2*i + 2
}

func (h CustomHeap) Parent(i int) int {
	return (i - 1) / 2 // la division d'entier tronque tjrs en go
}

func (h CustomHeap) Len() int {
	return len(h)
}

func (h CustomHeap) Less(i CustomHeapItem, j CustomHeapItem) bool {
	return i.Value < j.Value
}

// maintient la propriété de max heap
// i c'est l'indice tel que tous les subtrees de i sont des heaps valides
// donc on fait on sait que i c'est l'endroit où il y a un problème
func (h CustomHeap) maxHeapify(i int) {

	left := h.Left(i)
	right := h.Right(i)

	// on regarde les enfants de i, est-ce qu'ils sont bons ?
	// si c'est le cas, on change rien !

	largest := i

	if left < h.Len() && h.Less(h[i], h[left]) {
		// oups ! notre root est plus petit que notre membre à gauche.
		// on va devoir le faire descendre
		largest = left
	}

	if right < h.Len() && h.Less(h[i], h[right]) {
		largest = right
	}

	if largest != i {
		// on swap le root avec le membre le plus grand
		h[i], h[largest] = h[largest], h[i]

		h.maxHeapify(largest) // largest c'est la position de notre root qu'on vient de swapper
	}

}

func (h CustomHeap) buildMaxHeap() {
	// on veut appeler max heapify à partir du bas sur chaque noeud

	for i := (h.Len() - 2) / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
}
