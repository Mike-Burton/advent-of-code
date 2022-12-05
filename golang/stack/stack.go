package stack

// TODO using slices may be faster than LL on modern hardware.

type Stack[T any] struct {
	head *node[T]
	size int
}

type node[T any] struct {
	data T
	next *node[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Push(data T) {
	newHead := node[T]{data, s.head}
	s.head = &newHead
	s.size++
}

func (s *Stack[T]) Pop() (t T) {
	if s.head == nil {
		return t
	}

	r := s.head.data
	s.head = s.head.next
	s.size--

	return r
}

func (s *Stack[T]) Peek() (t T) {
	if s.head == nil {
		return t
	}
	return s.head.data
}
