package ds

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcd(t *testing.T) {
	assert.Equal(t, 1, Gcd(2, 3))
	assert.Equal(t, 2, Gcd(2, 4))
	assert.Equal(t, 3, Gcd(6, 9))
}

func TestModPower(t *testing.T) {
	const M = 1e9 + 7
	assert.Equal(t, 1, ModPower(2, 0, M))
	assert.Equal(t, 2, ModPower(2, 1, M))
	assert.Equal(t, 4, ModPower(2, 2, M))
	assert.Equal(t, 8, ModPower(2, 3, M))
	assert.Equal(t, 16, ModPower(2, 4, M))
	assert.Equal(t, 32, ModPower(2, 5, M))
	assert.Equal(t, 140625001, ModPower(2, 1e9, M))
}

func TestSumAbs(t *testing.T) {
	assert.Equal(t, 10, Sum(1, 2, 3, 4))
	assert.Equal(t, 1, Abs(1))
	assert.Equal(t, 1, Abs(-1))
}

func TestSymmetricNumberSeq(t *testing.T) {
	expected := 0
	for i := range SymmetricNumberSeq {
		if i > 1e6 {
			break
		}
		for expected++; ; expected++ {
			itoa := strconv.Itoa(expected)
			bytes := []byte(itoa)
			slices.Reverse(bytes)
			if itoa == string(bytes) {
				assert.Equal(t, expected, i)
				break
			}
		}
	}
}

func TestFactorization(t *testing.T) {
	assert.Equal(t, map[int]int{3: 1}, Factorization(3))
	assert.Equal(t, map[int]int{2: 2}, Factorization(4))
	assert.Equal(t, map[int]int{2: 1, 3: 1}, Factorization(6))
	assert.Equal(t, map[int]int{2: 3}, Factorization(8))
}

func TestPrime(t *testing.T) {
	assert.False(t, IsPrime(0))
	assert.False(t, IsPrime(1))
	assert.True(t, IsPrime(2))
	assert.True(t, IsPrime(3))
	assert.False(t, IsPrime(4))

	n := 100_000
	m := make(map[int]bool)
	for _, prime := range InitPrimes(n) {
		m[prime] = true
	}

	for i := 2; i <= n; i++ {
		assert.False(t, IsPrime(i) && !m[i] || !IsPrime(i) && m[i])
	}
}

func TestCNK(t *testing.T) {
	cnk := initCNK(5)
	assert.Equal(t, [][]int{
		{1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0},
		{1, 2, 1, 0, 0, 0},
		{1, 3, 3, 1, 0, 0},
		{1, 4, 6, 4, 1, 0},
		{1, 5, 10, 10, 5, 1},
	}, cnk)
}

func TestNextPermutation(t *testing.T) {
	source := []int{1, 2, 3}

	assert.False(t, NextPermutation(source[:1]))

	assert.True(t, NextPermutation(source))
	assert.Equal(t, []int{1, 3, 2}, source)

	assert.True(t, NextPermutation(source))
	assert.Equal(t, []int{2, 1, 3}, source)

	assert.True(t, NextPermutation(source))
	assert.Equal(t, []int{2, 3, 1}, source)

	assert.True(t, NextPermutation(source))
	assert.Equal(t, []int{3, 1, 2}, source)

	assert.True(t, NextPermutation(source))
	assert.Equal(t, []int{3, 2, 1}, source)

	assert.False(t, NextPermutation(source))
	assert.Equal(t, []int{1, 2, 3}, source)
}
