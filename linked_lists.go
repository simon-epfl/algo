package main

type CustomLinkedListItem struct {
	Prev  *CustomLinkedListItem
	Value any
	Next  *CustomLinkedListItem
}

type CustomLinkedList struct {
	head *CustomLinkedListItem
}

// runtime O(n)
func (l *CustomLinkedList) Search(value any) any {
	x := l.head
	for x != nil && x.Value != value {
		x = x.Next
	}
	if x.Value == value {
		return x
	} else {
		return -1
	}
}

// runtime en O(1)
func (l *CustomLinkedList) Insert(x any) {
	item := CustomLinkedListItem{
		Next:  l.head,
		Value: x,
		Prev:  nil,
	}
	if l.head != nil {
		l.head.Prev = &item
	}
	l.head = &item
}

func (l *CustomLinkedList) Delete(x *CustomLinkedListItem) {
	prevX := x.Prev
	nextX := x.Next
	if prevX != nil {
		prevX.Next = nextX
	} else { // y'avait rien avant x!
		l.head = nextX
	}
	if nextX != nil {
		nextX.Prev = prevX
	}
}
