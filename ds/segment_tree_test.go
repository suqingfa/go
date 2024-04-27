package ds

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSegmentTree_Sum(t *testing.T) {
	const N = 1000
	st := NewSumSegmentTree(N)

	for i := range N {
		st.Update(i, i)
	}

	for i := range N {
		for j, sum := i, 0; j < N; j++ {
			sum += j
			assert.Equal(t, sum, st.Range(i, j), "i: %d, j: %d", i, j)
		}
	}
}

func TestSegmentTree_UpdateRange(t *testing.T) {
	const N = 10
	nums := make([]int, N)
	st := NewSumSegmentTree(N)

	for i := range N {
		nums[i] = i
		st.Update(i, i)
	}

	check := func() {
		for i := range N {
			for j, sum := i, 0; j < N; j++ {
				sum += nums[j]
				assert.Equal(t, sum, st.Range(i, j), "i: %d, j: %d, nums: %v", i, j, nums[i:j+1])
			}
		}
	}

	check()

	update := make([][]int, 0)
	for i := range N {
		for j := i; j < N; j++ {
			update = append(update, []int{i, j, rand.Intn(100) - 50})
		}
	}

	for _, ints := range update {
		for i := ints[0]; i <= ints[1]; i++ {
			nums[i] += ints[2]
		}
		st.UpdateRange(ints[0], ints[1], ints[2])
		check()
	}
}

func TestSegmentTree_Max(t *testing.T) {
	nums := []int{3, -1, 2, 5, 4}
	st := NewMaxSegmentTree(len(nums))

	for i, num := range nums {
		st.Update(i, num)
	}

	check := func() {
		for i := range len(nums) {
			for j, mx := i, nums[i]; j < len(nums); j++ {
				mx = max(mx, nums[j])
				assert.Equal(t, mx, st.Range(i, j), "i: %d, j: %d, nums: %v", i, j, nums[i:j+1])
			}
		}
	}

	check()

	st.Update(3, 3)
	nums[3] = max(nums[3], 3)
	check()
	st.Update(3, 6)
	nums[3] = max(nums[3], 6)
	check()
}
