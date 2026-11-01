type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)

	if len(s.min) == 0 || val < s.min[len(s.min)-1] {
		s.min = append(s.min, val)
	} else {
		s.min = append(s.min, s.min[len(s.min)-1])
	}
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.min = s.min[:len(s.min)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.min[len(s.min)-1]
}
