package main

type PriorityItem struct {
	value    string
	priority int
}

type PriorityQueue struct {
	items []*PriorityItem
}

// ajoute un élément à la priority queue (à la fin) et appelle heapify
func (pq *PriorityQueue) Push(item *PriorityItem) {
	pq.items = append(pq.items, item)
	pq.heapIncreaseKey(len(pq.items) - 1)
}

// renvoie l'élément avec la priority la plus faible (le root)
// et le retire de la priority queue
// il swap le root avec le dernier élément
func (pq *PriorityQueue) Pop() *PriorityItem {
	if len(pq.items) == 0 {
		return nil
	}
	min := pq.items[0]
	index_of_last_element := len(pq.items) - 1
	pq.items[0] = pq.items[index_of_last_element]
	pq.items = pq.items[:index_of_last_element]
	pq.heapify(0)
	return min
}

// on a augmenté la priority d'un élément
// on doit le faire remonter dans l'arbre
func (pq *PriorityQueue) heapIncreaseKey(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if pq.items[index].priority >= pq.items[parent].priority {
			// l'item est à la bonne place, son parent est plus petit
			break
		}
		pq.items[index], pq.items[parent] = pq.items[parent], pq.items[index]
		// on swap l'élément avec son parent
		index = parent
	}
}

// on sait que le root doit descendre dans l'arbre
// (généralement parce qu'on a swapé le root avec le dernier élément)
func (pq *PriorityQueue) heapify(index int) {
	n := len(pq.items)
	for {
		smallest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < n && pq.items[left].priority < pq.items[smallest].priority {
			smallest = left
		}
		if right < n && pq.items[right].priority < pq.items[smallest].priority {
			smallest = right
		}
		if smallest == index {
			break
		}
		pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
		// on swap l'élément avec le plus petit enfant (comme ça si cet enfant devient le root, il sera à la bonne place)
		// et on continue à descendre
		index = smallest
	}
}
