package ds

// UnionFind 并查集
type UnionFind[T comparable] interface {

	// Find 查找 t 的代表
	Find(t T) T

	// Union 将 a 与 b 合并，返回 true 为当前操作将两个集合合并在一起
	Union(a, b T) bool

	// IsConnected 测试 a 与 b 是否在同一个集合
	IsConnected(a, b T) bool
}

func NewUnionFind[T comparable]() UnionFind[T] {
	return &unionFindUseMap[T]{make(map[T]T)}
}

func NewUnionFindInt(size int) UnionFind[int] {
	arr := make([]int, size)
	for i := range size {
		arr[i] = i
	}
	return &unionFindUseArray{arr}
}

type unionFindUseMap[T comparable] struct {
	leader map[T]T
}

func (uf *unionFindUseMap[T]) Find(t T) T {
	if _, ok := uf.leader[t]; !ok {
		return t
	}
	uf.leader[t] = uf.Find(uf.leader[t])

	return uf.leader[t]
}

func (uf *unionFindUseMap[T]) IsConnected(a T, b T) bool {
	return uf.Find(a) == uf.Find(b)
}

func (uf *unionFindUseMap[T]) Union(a T, b T) bool {
	la, lb := uf.Find(a), uf.Find(b)
	if la == lb {
		return false
	}

	uf.leader[lb] = la
	return true
}

type unionFindUseArray struct {
	arr []int
}

func (uf *unionFindUseArray) Find(t int) int {
	res := uf.arr[t]
	for res != uf.arr[res] {
		res = uf.arr[res]
	}

	uf.arr[t] = res

	return res
}

func (uf *unionFindUseArray) Union(a, b int) bool {
	la, lb := uf.Find(a), uf.Find(b)
	if la == lb {
		return false
	}

	uf.arr[lb] = la
	return true
}

func (uf *unionFindUseArray) IsConnected(a, b int) bool {
	return uf.Find(a) == uf.Find(b)
}
