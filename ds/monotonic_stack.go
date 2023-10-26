package ds

import "sort"

// MonotonicStack 单调栈 从栈顶到栈底的元素是单调递增（或者单调递减）
type MonotonicStack[T sort.Interface] struct {
	source   T
	stack    []int
	topIndex int
}

func NewMonotonicStack[T sort.Interface](source T) *MonotonicStack[T] {
	return &MonotonicStack[T]{source, make([]int, source.Len()), 0}
}

func (s *MonotonicStack[T]) Size() int {
	return s.topIndex
}

func (s *MonotonicStack[T]) TopN(n int) int {
	return s.stack[s.topIndex-n-1]
}

func (s *MonotonicStack[T]) Top() int {
	return s.TopN(0)
}

func (s *MonotonicStack[T]) Pop() int {
	res := s.Top()
	s.topIndex--
	return res
}

// Push 将 source[index] 的索引 index 压入栈
// 返回值 栈顶大于index的索引
func (s *MonotonicStack[T]) Push(index int) []int {
	res := make([]int, 0)
	for s.Size() > 0 && s.source.Less(s.Top(), index) {
		res = append(res, s.Pop())
	}

	s.stack[s.topIndex] = index
	s.topIndex++

	return res
}
