package example

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const N = 1e5

func TestGoroutine(t *testing.T) {
	c := make(chan int, 10)
	defer close(c)

	for i := 0; i < N; i++ {
		go func(c chan int) {
			time.Sleep(time.Millisecond)
			c <- 1
		}(c)
	}

	value := 0
	for i := 0; i < N; i++ {
		value += <-c
	}

	if value != N {
		t.Error()
	}
	println("value: ", value)
}

func TestLock(t *testing.T) {
	mutex := sync.Mutex{}
	value := 0

	c := make(chan int)
	defer close(c)

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

	if value != N {
		t.Error()
	}
	println("value: ", value)
}

func TestAtomic(t *testing.T) {
	value := atomic.Int64{}

	c := make(chan int)
	defer close(c)

	for i := 0; i < N; i++ {
		go func(c chan int) {
			value.Add(1)
			c <- 0
		}(c)
	}

	for i := 0; i < N; i++ {
		<-c
	}

	if value.Load() != N {
		t.Error()
	}
	println("value: ", value.Load())
}
