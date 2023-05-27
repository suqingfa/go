package leetcode

import (
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
