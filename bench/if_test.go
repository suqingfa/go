package bench

import (
	"math/rand"
	"testing"
)

func BenchmarkIfSorted(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}

	b.ResetTimer()

	res := 0
	for i := 0; i < b.N; i++ {
		for _, j := range arr {
			if j%10 < 1 {
				res++
			} else {
				res--
			}
		}
	}
}

func BenchmarkIfNotSorted(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	b.ResetTimer()

	res := 0
	for i := 0; i < b.N; i++ {
		for _, j := range arr {
			if j%10 >= 1 {
				res--
			} else {
				res++
			}
		}
	}
}
