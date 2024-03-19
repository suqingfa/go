package test

import (
	"encoding/json"
	"github.com/emirpasic/gods/v2/queues/linkedlistqueue"
	"github.com/emirpasic/gods/v2/queues/priorityqueue"
	"github.com/emirpasic/gods/v2/utils"
	"slices"
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	t.Log(utils.ToString([]string{"hello", "world"}))
}

func TestSliceEqual(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	c := []int{1, 3, 2, 4}

	if !slices.Equal(a, b) {
		t.Fail()
	}

	if slices.Equal(a, c) {
		t.Fail()
	}
}

func TestJson(t *testing.T) {
	bytes, _ := json.Marshal(map[string]int{"a": 1, "b": 2})

	m := make(map[string]int)
	_ = json.Unmarshal(bytes, &m)

	t.Log(string(bytes), m)

	if len(m) != 2 || m["a"] != 1 || m["b"] != 2 {
		t.Error()
	}
}

func TestQueue(t *testing.T) {
	q := linkedlistqueue.New[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Empty() {
		t.Error()
	}

	for i := 0; i < 4; i++ {
		value, ok := q.Dequeue()
		if !ok {
			t.Error()
		}
		if value != i {
			t.Error()
		}
	}

	pq := priorityqueue.New[int]()
	pq.Enqueue(3)
	pq.Enqueue(2)
	pq.Enqueue(1)
	pq.Enqueue(0)

	for i := 0; i < 4; i++ {
		value, ok := pq.Dequeue()
		if !ok {
			t.Error()
		}
		if value != i {
			t.Error()
		}
	}
}

func TestTimeParse(t *testing.T) {
	parse, err := time.Parse("2006:01:02:15:04:05", "2000:01:02:03:04:05")
	if err != nil {
		t.Error()
		return
	}
	t.Log(parse)
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.ToString([]string{"hello", "world"})
	}
}
