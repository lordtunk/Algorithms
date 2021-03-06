package stack

import "fmt"

type item struct {
	data int
	next *item
}
type stack struct {
	count int
	top *item
}

func main() {
	s := InitStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	// Should Print in reverse order (3 2 1) since it is a stack
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
}

func (s *stack) Push(data int) {
	i := initItem(data)
	i.next = s.top
	s.top = i
	s.count = s.count + 1
}

func (s *stack) Pop() (int) {
	if s.count == 0 {
		panic("Stack is empty")
	}
	i := s.top
	s.top = i.next
	s.count = s.count - 1
	return i.data
}

func (s *stack) Peek() (*item) {
	if s.count == 0 {
		return nil
	}
	return s.top
}

// Output each item in the stack
func (s *stack) Print() {
	if s.count == 0 {
		out("stack is empty\n")
		return
	}
	i := s.top
	for count := 0; i != nil && count < s.count; count++ {
		out("%d ", i.data)
		i = i.next
	}
	out("\n")
}

// Create and return a stack
func InitStack() (*stack) {
	s := new(stack)
	s.count = 0
	s.top = nil
	return s
}

// Create and return an item using the provided data
func initItem(data int) (*item) {
	i := new(item)
	i.data = data
	i.next = nil
	return i
}

// Shorthand function for fmt.Printf
func out(format string, v...interface{}) {
	fmt.Printf(format, v...)
}