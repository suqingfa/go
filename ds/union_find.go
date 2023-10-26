package ds

// UnionFind 并查集
type UnionFind[T comparable] struct {
	leader map[T]T
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{make(map[T]T)}
}

func (s *UnionFind[T]) Find(t T) T {
	if _, ok := s.leader[t]; !ok {
		return t
	}
	s.leader[t] = s.Find(s.leader[t])

	return s.leader[t]
}

func (s *UnionFind[T]) Connected(a T, b T) bool {
	return s.Find(a) == s.Find(b)
}

func (s *UnionFind[T]) Union(a T, b T) bool {
	la := s.Find(a)
	lb := s.Find(b)
	if la == lb {
		return false
	}

	s.leader[lb] = la

	return true
}
