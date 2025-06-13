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
	queue := *q

	if queue.IsFull() {
		panic("queue is full!")
	}

	queue.items[queue.tail] = x
	if queue.tail == MAX_CUSTOM_QUEUE_SIZE {
		queue.tail = 0 // on revient à zéro
	} else {
		queue.tail++
	}
}

func (q *CustomQueue) Dequeue() CustomQueueItem {
	queue := *q

	if queue.IsEmpty() {
		panic("queue is empty!")
	}

	item := queue.items[queue.head]
	if queue.head == MAX_CUSTOM_QUEUE_SIZE {
		queue.head = 0
	} else {
		queue.head++
	}

	return item
}
