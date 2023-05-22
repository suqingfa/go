package leetcode

import (
	"github.com/emirpasic/gods/utils"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	println(utils.ToString([]string{"hello", "world"}))
}

func BenchmarkStringAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
