package leetcode

import (
	"encoding/json"
	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"reflect"
	"sync"
	"testing"
	"time"
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

func TestJson(t *testing.T) {
	bytes, _ := json.Marshal(map[string]int{"a": 1, "b": 2})

	m := make(map[string]int)
	_ = json.Unmarshal(bytes, &m)

	println(string(bytes), utils.ToString(m))

	if len(m) != 2 || m["a"] != 1 || m["b"] != 2 {
		t.Error()
	}
}

func TestQueue(t *testing.T) {
	q := linkedlistqueue.New()
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
		if value.(int) != i {
			t.Error()
		}
	}

	pq := priorityqueue.NewWith(utils.IntComparator)
	pq.Enqueue(3)
	pq.Enqueue(2)
	pq.Enqueue(1)
	pq.Enqueue(0)

	for i := 0; i < 4; i++ {
		value, ok := pq.Dequeue()
		if !ok {
			t.Error()
		}
		if value.(int) != i {
			t.Error()
		}
	}
}

func TestGoroutine(t *testing.T) {
	f := func(i int, c chan int) {
		time.Sleep(200_000_000)
		c <- i
	}

	const N = 1000

	c := make(chan int, 10)

	for i := 0; i < N; i++ {
		go f(i, c)
	}

	sum := 0
	for i := 0; i < N; i++ {
		sum += <-c
	}

	println("sum: ", sum)
	close(c)
}

func TestLock(t *testing.T) {
	mutex := sync.Mutex{}
	value := 0

	const N = 100_000
	c := make(chan int)

	for i := 0; i < N; i++ {
		go func(c chan int) {
			mutex.Lock()
			defer mutex.Unlock()
			value++
			c <- 0
		}(c)
	}

	for i := 0; i < N; i++ {
		<-c
	}

	println("value: ", value)
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.ToString([]string{"hello", "world"})
	}
}
