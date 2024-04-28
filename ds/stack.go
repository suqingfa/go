package ds

type Stack[T any] []T

func NewStack[T any](values ...T) *Stack[T] {
	res := Stack[T](values)
	return &res
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) Push(t T) {
	*s = append(*s, t)
}

func (s *Stack[T]) MustPop() T {
	if s.Size() == 0 {
		panic("stack is empty")
	}
	res := (*s)[s.Size()-1]
	*s = (*s)[:s.Size()-1]
	return res
}
