package ds

import (
	"reflect"
	"sort"
	"testing"
)

func TestMonotonicStack(t *testing.T) {
	source := []int{3, 1, 2, 3, 2, 1, 4}
	stack := NewMonotonicStack[sort.IntSlice](source)

	except := [][]int{
		{},           // [3]
		{},           // [3,1]
		{1},          // [3,2]
		{2},          // [3,3]
		{},           // [3,3,2]
		{},           // [3,3,2,1]
		{5, 4, 3, 0}, // [4]
	}

	for i := 0; i < len(source); i++ {
		push := stack.Push(i)
		if !reflect.DeepEqual(push, except[i]) {
			t.Error(push)
		}
	}
}
