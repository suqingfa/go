package test

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type hp []int

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(int)) }
func (h *hp) Pop() any           { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func TestHeap(t *testing.T) {
	h := hp([]int{3, 4, 5, 3, 2, 1, 1})
	heap.Init(&h)
	log.Println(h)

	assert.Equal(t, 1, heap.Pop(&h))
	assert.Equal(t, 1, heap.Pop(&h))
	assert.Equal(t, 2, heap.Pop(&h))
	assert.Equal(t, 3, heap.Pop(&h))
	assert.Equal(t, 3, heap.Pop(&h))
	assert.Equal(t, 4, heap.Pop(&h))
	assert.Equal(t, 5, heap.Pop(&h))
	assert.Equal(t, 0, len(h))
}
