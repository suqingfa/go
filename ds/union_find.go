package ds

// UnionFind 并查集
type UnionFind[T comparable] struct {
	leader map[T]T
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{make(map[T]T)}
}

func (uf *UnionFind[T]) Find(t T) T {
	if _, ok := uf.leader[t]; !ok {
		return t
	}
	uf.leader[t] = uf.Find(uf.leader[t])

	return uf.leader[t]
}

func (uf *UnionFind[T]) IsConnected(a T, b T) bool {
	return uf.Find(a) == uf.Find(b)
}

func (uf *UnionFind[T]) Union(a T, b T) bool {
	la, lb := uf.Find(a), uf.Find(b)
	if la == lb {
		return false
	}

	uf.leader[lb] = la
	return true
}
