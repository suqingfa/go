package test

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedBlackTree(t *testing.T) {
	tree := redblacktree.New[int, int]()

	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Put(3, 3)

	assert.Equal(t, 2, tree.GetNode(2).Value)

	assert.Equal(t, 1, tree.Left().Value)
	assert.Equal(t, 3, tree.Right().Value)

	tree.Remove(3)
	assert.Equal(t, 2, tree.Right().Value)
}
