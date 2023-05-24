package leetcode

import (
	"github.com/emirpasic/gods/utils"
	"reflect"
	"testing"
)

func TestToString(t *testing.T) {
	println(utils.ToString([]string{"hello", "world"}))
}

func TestSliceEqual(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	c := []int{1, 3, 2, 4}

	println(reflect.DeepEqual(a, b))
	println(reflect.DeepEqual(a, c))
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.ToString([]string{"hello", "world"})
	}
}
