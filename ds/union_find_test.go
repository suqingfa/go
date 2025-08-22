package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind[int]()

	assert.True(t, uf.Union(1, 2))
	assert.False(t, uf.Union(1, 2))

	assert.Equal(t, 1, uf.Find(2))

	assert.True(t, uf.IsConnected(1, 2))
	assert.False(t, uf.IsConnected(1, 3))
}

func TestUnionFindUseIntArray(t *testing.T) {
	uf := NewUnionFindInt(4)

	assert.True(t, uf.Union(1, 2))
	assert.False(t, uf.Union(1, 2))

	assert.Equal(t, 1, uf.Find(2))

	assert.True(t, uf.IsConnected(1, 2))
	assert.False(t, uf.IsConnected(1, 3))
}
