package leetcode

import (
	"github.com/emirpasic/gods/utils"
	"testing"
)

func TestToString(t *testing.T) {
	println(utils.ToString([]string{"hello", "world"}))
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.ToString([]string{"hello", "world"})
	}
}
