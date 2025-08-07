package test

import (
	"encoding/json"
	"slices"
	"testing"
	"time"

	"github.com/emirpasic/gods/v2/queues/linkedlistqueue"
	"github.com/emirpasic/gods/v2/queues/priorityqueue"
	"github.com/emirpasic/gods/v2/utils"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	t.Log(utils.ToString([]string{"hello", "world"}))
}

func TestSliceEqual(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	c := []int{1, 3, 2, 4}

	assert.True(t, slices.Equal(a, b))
	assert.False(t, slices.Equal(a, c))
}

func TestJson(t *testing.T) {
	bytes, _ := json.Marshal(map[string]int{"a": 1, "b": 2})

	m := make(map[string]int)
	_ = json.Unmarshal(bytes, &m)

	t.Log(string(bytes), m)

	assert.Equal(t, 2, len(m))
	assert.Equal(t, 1, m["a"])
	assert.Equal(t, 2, m["b"])
}

func TestQueue(t *testing.T) {
	q := linkedlistqueue.New[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.False(t, q.Empty())

	for i := 0; i < 4; i++ {
		value, ok := q.Dequeue()
		assert.True(t, ok)
		assert.Equal(t, i, value)
	}

	pq := priorityqueue.New[int]()
	pq.Enqueue(3)
	pq.Enqueue(2)
	pq.Enqueue(1)
	pq.Enqueue(0)

	for i := 0; i < 4; i++ {
		value, ok := pq.Dequeue()
		assert.True(t, ok)
		assert.Equal(t, i, value)
	}
}

func TestTimeParse(t *testing.T) {
	parse, err := time.Parse("2006:01:02:15:04:05", "2000:01:02:03:04:05")
	assert.Nil(t, err)
	t.Log(parse)
}
