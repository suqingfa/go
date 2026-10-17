package bench

import "testing"

func BenchmarkOp(b *testing.B) {
	res := 0
	for range b.N {
		res++
	}
	b.Log(res)
}

func BenchmarkSwitchBS(b *testing.B) {
	res := 0
	for i := range b.N {
		switch i % 7 {
		case 1:
			res++
		case 2:
			res++
		case 3:
			res++
		case 4:
			res++
		case 5:
			res++
		case 6:
			res++
		case 7:
			res++
		}
	}
	b.Log(res)
}

func BenchmarkSwitchJT(b *testing.B) {
	res := 0
	for i := range b.N {
		switch i % 7 {
		case 1:
			res++
		case 2:
			res++
		case 3:
			res++
		case 4:
			res++
		case 5:
			res++
		case 6:
			res++
		case 7:
			res++
		case 8:
			res++
		}
	}
	b.Log(res)
}
