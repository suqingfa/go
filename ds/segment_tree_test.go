package ds

import (
	"math/rand"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	const N = 100
	segmentTree := NewSegmentTree(0, N-1)

	ints := make([]int, N*10)
	for i := range ints {
		ints[i] = rand.Intn(N)
	}

	for _, i := range ints {
		segmentTree.Insert(i)
	}

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			cnt := 0
			for _, v := range ints {
				if i <= v && v <= j {
					cnt++
				}
			}

			if segmentTree.Search(i, j) != cnt {
				t.Error()
			}
		}
	}
}
