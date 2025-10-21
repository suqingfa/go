package test

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = i
	}

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	assert.False(t, slices.IsSorted(arr))

	slices.Sort(arr)
	assert.True(t, slices.IsSorted(arr))
}

func TestSearch(t *testing.T) {
	arr := []int{1, 2, 2, 3, 6, 6, 7}

	// 找到时，返回最小索引
	i, _ := slices.BinarySearch(arr, 2)
	assert.Equal(t, 1, i)

	// 未找到时，返回目标应该放置的位置
	i, _ = slices.BinarySearch(arr, 4)
	assert.Equal(t, 4, i)
}
