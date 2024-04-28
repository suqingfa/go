package ds

type Queue[T any] []T

func NewQueue[T any](values ...T) *Queue[T] {
	res := Queue[T](values)
	return &res
}

func (s *Queue[T]) Size() int {
	return len(*s)
}

func (s *Queue[T]) Push(t T) {
	*s = append(*s, t)
}

func (s *Queue[T]) MustPop() T {
	if s.Size() == 0 {
		panic("queue is empty")
	}
	res := (*s)[0]
	*s = (*s)[1:]
	return res
}
