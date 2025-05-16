package main

import "container/heap"

type PriorityItem struct {
	value    any
	priority int
	index    int
}

type PriorityQueue []*PriorityItem

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *PriorityQueue) Push(x any) {
	*h = append(*h, x.(*PriorityItem))
}

func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (pq *PriorityQueue) update(item *PriorityItem, value any, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// // ajoute un élément à la priority queue (à la fin) et appelle heapify
// func (pq *PriorityQueue) Push(item *PriorityItem) {
// 	pq.items = append(pq.items, item)
// 	pq.heapIncreaseKey(len(pq.items) - 1)
// }

// // renvoie l'élément avec la priority la plus faible (le root)
// // et le retire de la priority queue
// // il swap le root avec le dernier élément
// func (pq *PriorityQueue) Pop() *PriorityItem {
// 	if len(pq.items) == 0 {
// 		return nil
// 	}
// 	min := pq.items[0]
// 	index_of_last_element := len(pq.items) - 1
// 	pq.items[0] = pq.items[index_of_last_element]
// 	pq.items = pq.items[:index_of_last_element]
// 	pq.heapify(0)
// 	return min
// }

// // on a augmenté la priority d'un élément
// // on doit le faire remonter dans l'arbre
// func (pq *PriorityQueue) heapIncreaseKey(index int) {
// 	for index > 0 {
// 		parent := (index - 1) / 2
// 		if pq.items[index].priority >= pq.items[parent].priority {
// 			// l'item est à la bonne place, son parent est plus petit
// 			break
// 		}
// 		pq.items[index], pq.items[parent] = pq.items[parent], pq.items[index]
// 		// on swap l'élément avec son parent
// 		index = parent
// 	}
// }

// // on sait que le root doit descendre dans l'arbre
// // (généralement parce qu'on a swapé le root avec le dernier élément)
// func (pq *PriorityQueue) heapify(index int) {
// 	n := len(pq.items)
// 	for {
// 		smallest := index
// 		left := 2*index + 1
// 		right := 2*index + 2

// 		if left < n && pq.items[left].priority < pq.items[smallest].priority {
// 			smallest = left
// 		}
// 		if right < n && pq.items[right].priority < pq.items[smallest].priority {
// 			smallest = right
// 		}
// 		if smallest == index {
// 			break
// 		}
// 		pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
// 		// on swap l'élément avec le plus petit enfant (comme ça si cet enfant devient le root, il sera à la bonne place)
// 		// et on continue à descendre
// 		index = smallest
// 	}
// }
