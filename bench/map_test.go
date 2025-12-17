package bench

import (
	"testing"
)

func BenchmarkSliceAccess(b *testing.B) {
	arr := make([]int, 1e5)
	for i := 0; i < b.N; i++ {
		arr[i%1e5] = i
	}
}

func BenchmarkMapFast64(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	b.Log(len(m))
}

func BenchmarkMap(b *testing.B) {
	m := make(map[float64]int)
	for i := 0; i < b.N; i++ {
		m[float64(i)] = i
	}
	b.Log(len(m))
}
