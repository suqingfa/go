package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapFunc(t *testing.T) {
	l := []int{1, 2, 3}

	assert.Equal(t, []int{2, 4, 6}, MapFunc(l, func(e int) int {
		return e * 2
	}))

	assert.Equal(t, []string{"1", "2", "3"}, MapFunc(l, func(e int) string {
		return fmt.Sprintf("%d", e)
	}))
}

func TestCount(t *testing.T) {
	assert.Equal(t, 3, Count([]int{1, 2, 2, 3, 3, 3}, 3))
}

func TestIndexes(t *testing.T) {
	assert.Equal(t, []int{0, 3, 6}, Indexes([]int{1, 2, 3, 1, 2, 3, 1}, 1))
}
