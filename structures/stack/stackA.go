package stack

type stackA struct {
	count int
	data []int
}

func (s *stackA) Push(data int) {
	s.data = append(s.data, data)
	s.count = s.count + 1
}

func (s *stackA) Pop() (int) {
	if s.count == 0 {
		panic("Stack is empty")
	}
	s.count = s.count - 1
	i := s.data[s.count]
	s.data = s.data[:s.count]
	return i
}

func (s *stackA) Peek() (int) {
	if s.count == 0 {
		panic("Stack is empty")
	}
	return s.data[s.count-1]
}

// Output each item in the stackA
func (s *stackA) Print() {
	if s.count == 0 {
		out("Stack is empty\n")
		return
	}
	for i:=s.count-1;i>=0;i-- {
		out("%d ", s.data[i])
	}
	out("\n")
}

// Create and return a stackA
func InitStackA() (*stackA) {
	s := new(stackA)
	s.count = 0
	return s
}
