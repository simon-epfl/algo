package main

import (
	"fmt"
)

func customHeapSortInts(input []int) []int {
	var heap CustomHeap
	for i := range input {
		heap = append(heap, CustomHeapItem{Value: i})
	}
	heap.buildMaxHeap()

	fmt.Println(heap)

	var sorted []int

	for heap.Len() > 0 {
		sorted = append(sorted, heap[0].Value)
		heap[0] = heap[heap.Len()-1] // on prend le dernier élément et on le met à la place du root
		heap = heap[:heap.Len()-1]   // on enlève le dernier élément

		heap.maxHeapify(0)
	}

	return sorted
}

func runHeapsort() {
	unsorted := []int{5, 1, 7, 3, 2, 9, 4}
	sorted := customHeapSortInts(unsorted)
	fmt.Println("Non trié:", unsorted)
	fmt.Println("Trié:  ", sorted)
}
