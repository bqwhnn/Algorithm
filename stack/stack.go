package stack

// import "fmt"

// Stack Struct
type Stack struct {
	data []interface{}
	top int
} 

// NewStack is new stack
func NewStack() Stack {
	return Stack{
		data : make([]interface{}, 0, 100),
		top : -1,
	}
}

// IsEmpty check
func (s *Stack) IsEmpty() bool {
	return s.top < 0
}

// Push op
func (s *Stack) Push(x interface{}) {
	s.top++
	if s.top > len(s.data) - 1 {
		s.data = append(s.data, x)
	} else {
		s.data[s.top] = x
	}
}

// Pop op
func (s *Stack) Pop() interface{} {
	if s.top < 0 {
		return nil
	}
	res := s.data[s.top]
	s.top--
	return res
}

// Top get stack top val
func (s *Stack) Top() interface{} {
	if s.top < 0 {
		return nil
	}
	return s.data[s.top]
}