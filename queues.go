package main

type CustomQueueItem struct {
	Content any
}

const MAX_CUSTOM_QUEUE_SIZE = 20

type CustomQueue struct {
	head int
	tail int
	// qqch d'assez important avec les queues c'est que comme on fait des queue, dequeue
	// notre tableau se "déplace" vers la droite
	// donc l'algo de gestion de la queue utilise head/tail
	// et ici on en fait une démo avec une queue fixe
	items [MAX_CUSTOM_QUEUE_SIZE]CustomQueueItem
}

func getEmptyCustomQueue() CustomQueue {
	return CustomQueue{
		head:  0,
		tail:  0,
		items: [MAX_CUSTOM_QUEUE_SIZE]CustomQueueItem{},
	}
}

func (q *CustomQueue) IsFull() bool {
	next := (q.tail + 1) % MAX_CUSTOM_QUEUE_SIZE
	return next == q.head
}

func (q *CustomQueue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *CustomQueue) Enqueue(x CustomQueueItem) {
	if q.IsFull() {
		panic("queue is full!")
	}

	q.items[q.tail] = x
	q.tail = (q.tail + 1) % MAX_CUSTOM_QUEUE_SIZE
}

func (q *CustomQueue) Dequeue() CustomQueueItem {
	if q.IsEmpty() {
		panic("queue is empty!")
	}

	item := q.items[q.head]
	q.head = (q.head + 1) % MAX_CUSTOM_QUEUE_SIZE
	return item
}
