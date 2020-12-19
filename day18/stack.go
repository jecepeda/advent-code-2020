package day18

// Stack represents a LIFO (Last In, First Out)
// data structure with basic operations to manage data
type Stack struct {
	Data []Token
}

// NewStack creates a new stack
func NewStack() *Stack {
	return &Stack{}
}

// Push pushes data into the stack
func (s *Stack) Push(value Token) {
	s.Data = append(s.Data, value)
}

// Pop gets the last element from the stack, removing the last
// element inserted
func (s *Stack) Pop() Token {
	if len(s.Data) == 0 {
		return Token{}
	}
	v := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return v
}

// Peek gets the last element from the stack without removing it
func (s *Stack) Peek() Token {
	return s.Data[len(s.Data)-1]
}

func (s *Stack) Len() int {
	return len(s.Data)
}
