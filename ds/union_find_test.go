package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind[int]()

	assert.True(t, uf.Union(1, 2))
	assert.False(t, uf.Union(1, 2))

	assert.Equal(t, 1, uf.Find(2))

	assert.True(t, uf.IsConnected(1, 2))
	assert.False(t, uf.IsConnected(1, 3))
}
