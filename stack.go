package main

// Stack is a simple LIFO stack.
type Stack struct {
	items []int
}

// Push adds an item to the stack.
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item.
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}
