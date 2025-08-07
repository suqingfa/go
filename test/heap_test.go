package test

import (
	"container/heap"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type H []int

func (h *H) Len() int           { return len(*h) }
func (h *H) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *H) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *H) Push(v any)         { *h = append(*h, v.(int)) }
func (h *H) Pop() any           { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func TestHeap(t *testing.T) {
	h := H([]int{3, 4, 5, 3, 2, 1, 1})

	heap.Init(&h)
	heap.Push(&h, 10)

	h[2] = 20
	heap.Fix(&h, 2)

	// [1 2 3 3 4 5 20 10]
	log.Println(h)

	// 移除第 i 个元素并重新整理堆
	assert.Equal(t, 3, heap.Remove(&h, 2))

	assert.Equal(t, 1, heap.Pop(&h))
	assert.Equal(t, 2, heap.Pop(&h))
	assert.Equal(t, 3, heap.Pop(&h))
	assert.Equal(t, 4, heap.Pop(&h))
	assert.Equal(t, 5, heap.Pop(&h))
	assert.Equal(t, 10, heap.Pop(&h))
	assert.Equal(t, 20, heap.Pop(&h))

	assert.Equal(t, 0, len(h))
}
