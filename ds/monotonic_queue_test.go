package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonotonicQueue(t *testing.T) {
	source := []int{3, 1, 2, 3, 2, 1, 4}
	queue := NewMonotonicQueue()

	except := [][]int{
		{3},
		{3, 1},
		{3, 2},
		{3, 3},
		{3, 3, 2},
		{3, 3, 2, 1},
		{4},
	}

	for i, v := range source {
		queue.Enqueue(v)
		assert.Equal(t, except[i], queue.queue)
	}

	assert.Equal(t, 4, queue.Peek())

	assert.False(t, queue.Dequeue(3))
	assert.True(t, queue.Dequeue(4))
	assert.False(t, queue.Dequeue(4))
}
