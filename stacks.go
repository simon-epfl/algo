package main

type CustomStackItem struct {
	Value   int
	Content any
}

type CustomStack []CustomStackItem

func (s CustomStack) Empty() bool {
	return len(s) == 0
}

func (s *CustomStack) Push(item CustomStackItem) {
	stack := *s
	*s = append(stack, item)
}

func (s *CustomStack) Pop(item CustomStackItem) CustomStackItem {
	if s.Empty() {
		panic("underflow")
	} else {
		stack := *s
		*s = stack[:len(stack)-1]
		return stack[len(stack)-1]
	}
}
