package ds

import (
	"reflect"
	"slices"
	"testing"
)

func TestGcd(t *testing.T) {
	if Gcd(2, 3) != 1 {
		t.Error()
	}

	if Gcd(2, 4) != 2 {
		t.Error()
	}

	if Gcd(6, 9) != 3 {
		t.Error()
	}
}

func TestSumAbs(t *testing.T) {
	if Sum(1, 2, 3, 4) != 10 {
		t.Error()
	}

	if Abs(1) != 1 {
		t.Error()
	}

	if Abs(-1) != 1 {
		t.Error()
	}
}

func TestPrime(t *testing.T) {
	if !IsPrime(0) {
		t.Error()
	}

	if !IsPrime(1) {
		t.Error()
	}

	if !IsPrime(2) {
		t.Error()
	}

	if !IsPrime(3) {
		t.Error()
	}

	if IsPrime(4) {
		t.Error()
	}

	n := 1_000_000
	m := make(map[int]bool)
	for _, prime := range InitPrimes(n) {
		m[prime] = true
	}

	for i := 2; i <= n; i++ {
		if IsPrime(i) && !m[i] || !IsPrime(i) && m[i] {
			t.Error()
		}
	}
}

func TestCNK(t *testing.T) {
	cnk := initCNK(5)
	if !reflect.DeepEqual(cnk, [][]int{
		{1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0},
		{1, 2, 1, 0, 0, 0},
		{1, 3, 3, 1, 0, 0},
		{1, 4, 6, 4, 1, 0},
		{1, 5, 10, 10, 5, 1},
	}) {
		t.Error()
	}
}

func TestNextPermutation(t *testing.T) {
	source := []int{1, 2, 3}

	NextPermutation(source[:1])

	if !NextPermutation(source) || !slices.Equal(source, []int{1, 3, 2}) {
		t.Error()
	}

	if !NextPermutation(source) || !slices.Equal(source, []int{2, 1, 3}) {
		t.Error()
	}

	if !NextPermutation(source) || !slices.Equal(source, []int{2, 3, 1}) {
		t.Error()
	}

	if !NextPermutation(source) || !slices.Equal(source, []int{3, 1, 2}) {
		t.Error()
	}

	if !NextPermutation(source) || !slices.Equal(source, []int{3, 2, 1}) {
		t.Error()
	}

	if NextPermutation(source) || !slices.Equal(source, []int{1, 2, 3}) {
		t.Error()
	}
}
