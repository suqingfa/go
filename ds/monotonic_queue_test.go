package ds

import (
	"reflect"
	"testing"
)

func TestMonotonicQueue(t *testing.T) {
	source := []int{3, 1, 2, 3, 2, 1, 4}
	queue := NewMonotonicQueue()

	except := [][]int{
		{3},
		{3, 1},
		{3, 2},
		{3, 3},
		{3, 3, 2},
		{3, 3, 2, 1},
		{4},
	}

	for i, v := range source {
		queue.Enqueue(v)
		if !reflect.DeepEqual(queue.queue, except[i]) {
			t.Error()
		}
	}

	if queue.Peek() != 4 {
		t.Error()
	}

	if queue.Dequeue(3) {
		t.Error()
	}

	if !queue.Dequeue(4) {
		t.Error()
	}

	if queue.Dequeue(4) {
		t.Error()
	}
}
