package main

import (
	"algo/shared"
	"container/heap"
	"fmt"
)

// c'est pas exactement l'implémentation de heapsort classique,
// mais ça utilise une priority queue pour trier les entiers
// en O(n log n) temps

func HeapSortInts(input []int) []int {
	// On créé une priority queue vide
	pq := make(shared.PriorityQueue, 0, len(input))
	heap.Init(&pq)

	// On envoie toutes nos valeurs dans la queue
	for _, v := range input {
		item := &shared.PriorityItem{
			Value:    v,
			Priority: v, // la prio de l'item c'est sa valeur elle-même
		}
		heap.Push(&pq, item)
	}

	// tout est déjà trié ici !

	sorted := make([]int, 0, len(input))
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*shared.PriorityItem)
		sorted = append(sorted, item.Value.(int))
	}

	return sorted
}

func runheapsort() {
	unsorted := []int{5, 1, 7, 3, 2, 9, 4}
	sorted := HeapSortInts(unsorted)
	fmt.Println("Non trié:", unsorted)
	fmt.Println("Trié:  ", sorted)
}
