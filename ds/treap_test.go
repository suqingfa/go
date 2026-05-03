package ds

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreap(t *testing.T) {
	treap := Treap{}

	list := make([]int, 100)
	for i := range list {
		list[i] = rand.Intn(10)

		cnt := 0
		for _, num := range list[:i] {
			if num < list[i] {
				cnt++
			}
		}

		rank := treap.Rank(list[i])
		assert.Equal(t, cnt+1, rank)

		treap.Insert(list[i])
	}
}
