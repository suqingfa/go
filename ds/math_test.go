package ds

import (
	"reflect"
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
	if !IsPrime(2) {
		t.Error()
	}

	if !IsPrime(3) {
		t.Error()
	}

	if IsPrime(4) {
		t.Error()
	}

	primes := initPrimes(11)
	if !reflect.DeepEqual(primes, []int{2, 3, 5, 7, 11}) {
		t.Error()
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

func TestReverse(t *testing.T) {
	source := []int{1, 2, 3, 4}
	Reverse(source)
	if !reflect.DeepEqual(source, []int{4, 3, 2, 1}) {
		t.Error()
	}

	source = []int{1, 2, 3, 4, 5}
	Reverse(source)
	if !reflect.DeepEqual(source, []int{5, 4, 3, 2, 1}) {
		t.Error()
	}
}

func TestNextPermutation(t *testing.T) {
	source := []int{1, 2, 3}

	if !NextPermutation(source) || !reflect.DeepEqual(source, []int{1, 3, 2}) {
		t.Error()
	}

	if !NextPermutation(source) || !reflect.DeepEqual(source, []int{2, 1, 3}) {
		t.Error()
	}

	if !NextPermutation(source) || !reflect.DeepEqual(source, []int{2, 3, 1}) {
		t.Error()
	}

	if !NextPermutation(source) || !reflect.DeepEqual(source, []int{3, 1, 2}) {
		t.Error()
	}

	if !NextPermutation(source) || !reflect.DeepEqual(source, []int{3, 2, 1}) {
		t.Error()
	}

	if NextPermutation(source) || !reflect.DeepEqual(source, []int{1, 2, 3}) {
		t.Error()
	}
}