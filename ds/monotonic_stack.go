package ds

import (
	"cmp"
)

// MonotonicStack 单调栈 从栈顶到栈底的元素是单调递增
type MonotonicStack[T cmp.Ordered] struct {
	source   []T
	stack    []int
	topIndex int
}

func NewMonotonicStack[T cmp.Ordered](source []T) *MonotonicStack[T] {
	return &MonotonicStack[T]{source, make([]int, len(source)), 0}
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
	for s.Size() > 0 && s.source[s.Top()] < s.source[index] {
		res = append(res, s.Pop())
	}

	s.stack[s.topIndex] = index
	s.topIndex++

	return res
}
