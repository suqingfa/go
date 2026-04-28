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
		list[i] = rand.Int()
		treap.Insert(list[i])
	}

	for _, v := range list {
		cnt := 0
		for _, num := range list {
			if v >= num {
				cnt++
			}
		}
		assert.Equal(t, cnt, treap.Rank(v))
	}
}
