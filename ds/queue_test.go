package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	stack := NewQueue[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(5)

	assert.Equal(t, 1, stack.MustPop())
	assert.Equal(t, 2, stack.MustPop())
	assert.Equal(t, 3, stack.MustPop())
	assert.Equal(t, 5, stack.MustPop())

	assert.Equal(t, 0, stack.Size())

	assert.Panics(t, func() {
		stack.MustPop()
	})
}
