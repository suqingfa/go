package ds

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSegmentTree_Sum(t *testing.T) {
	const N = 100
	tree := NewSegmentTree(
		N,
		0,
		func(node *int, val int) { *node = val },
		func(l int, r int) int { return l + r },
	)

	for i := range N {
		tree.Update(i, i)
	}

	assert.Equal(t, (0+9)*10/2, tree.Range(0, 9))
	assert.Equal(t, (10+99)*90/2, tree.Range(10, 99))
	assert.Equal(t, (0+99)*100/2, tree.Range(0, 99))
}

func TestSegmentTree_Max(t *testing.T) {
	nums := []int{3, -1, 2, 5, 4}
	tree := NewSegmentTree(
		len(nums),
		math.MinInt32,
		func(node *int, val int) { *node = val },
		func(l int, r int) int { return max(l, r) },
	)

	for i, num := range nums {
		tree.Update(i, num)
	}

	assert.Equal(t, 3, tree.Range(0, 2))
	assert.Equal(t, -1, tree.Range(1, 1))
	assert.Equal(t, 2, tree.Range(1, 2))
	assert.Equal(t, 5, tree.Range(0, 4))
	assert.Equal(t, 5, tree.Range(3, 4))
}

func TestSegmentTree_Set(t *testing.T) {
	nums := []int{3, -1, 2, 5, 4}
	tree := NewSegmentTree(
		len(nums),
		math.MinInt32,
		func(node *int, val int) { *node = max(*node, val) },
		func(l int, r int) int { return max(l, r) },
	)

	for i, num := range nums {
		tree.Update(i, num)
	}

	assert.Equal(t, 5, tree.Range(0, 4))
	tree.Update(3, 3)
	assert.Equal(t, 5, tree.Range(0, 4))
	tree.Update(3, 6)
	assert.Equal(t, 6, tree.Range(0, 4))
}
