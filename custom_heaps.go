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
func (h *CustomHeap) maxHeapify(i int) {

	heap := *h

	left := heap.Left(i)
	right := heap.Right(i)

	// on regarde les enfants de i, est-ce qu'ils sont bons ?
	// si c'est le cas, on change rien !

	largest := i

	if left < h.Len() && h.Less(heap[i], heap[left]) {
		// oups ! notre root est plus petit que notre membre à gauche.
		// on va devoir le faire descendre
		largest = left
	}

	if right < heap.Len() && heap.Less(heap[largest], heap[right]) {
		largest = right
	}

	if largest != i {
		// on swap le root avec le membre le plus grand
		heap[i], heap[largest] = heap[largest], heap[i]

		heap.maxHeapify(largest) // largest c'est la position de notre root qu'on vient de swapper
	}

}

func (h *CustomHeap) buildMaxHeap() {
	// on veut appeler max heapify à partir du bas sur chaque noeud

	for i := (h.Len() - 2) / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
}

func (h *CustomHeap) heapMaximum() CustomHeapItem {
	return (*h)[0]
}

func (h *CustomHeap) heapExtractMax() CustomHeapItem {
	heap := *h

	if heap.Len() == 0 {
		panic("heap underflow")
	}

	max := heap.heapMaximum()
	heap[0], heap[heap.Len()-1] = heap[heap.Len()-1], heap[0]
	heap = heap[:heap.Len()-1]

	heap.maxHeapify(0)

	return max
}

func (h *CustomHeap) heapIncreaseKey(idx int, value int) {

	heap := *h

	heap[idx].Value = value

	// tant que le parent est plus petit, on swap pour faire remonter la nouvelle clef
	for idx > 0 && heap[heap.Parent(idx)].Value < heap[idx].Value {
		heap[idx], heap[heap.Parent(idx)] = heap[heap.Parent(idx)], heap[idx]
		idx = heap.Parent(idx)
	}

}

func (h *CustomHeap) maxHeapInsert(item CustomHeapItem) {
	heap := *h

	heap = append(heap, item)

	// à heap.Len() -1 on va avoir l'élément qu'on vient juste d'ajouter
	// on trigger un nouveau calcul de sa position
	heap.heapIncreaseKey(heap.Len()-1, item.Value)
}
