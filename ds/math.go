package ds

import (
	"cmp"
	"slices"
)

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return Gcd(b, a%b)
}

func ModPower(base, n, mod int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return base % mod
	}

	res := ModPower(base, n/2, mod)
	res = (res * res) % mod
	if n%2 == 0 {
		return res
	} else {
		return (res * base) % mod
	}
}

func Sum[T int | int64 | byte | rune | float64](arr ...T) T {
	res := T(0)
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func Abs[T int | int64 | byte | rune | float64](n T) T {
	if n < T(0) {
		return T(0) - n
	}

	return n
}

func Factorization(n int) map[int]int {
	m := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for ; n%i == 0; n /= i {
			m[i]++
		}
	}

	if n != 1 {
		m[n]++
	}

	return m
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func InitPrimes(n int) []int {
	m := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		m[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if m[i] {
			for j := 2 * i; j <= n; j += i {
				m[j] = false
			}
		}
	}

	res := make([]int, 0)
	for i := 0; i <= n; i++ {
		if m[i] {
			res = append(res, i)
		}
	}

	return res
}

// c(n, k) 组合数
func initCNK(n int) [][]int {
	c := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		c[i] = make([]int, n+1)
	}

	c[0][0] = 1
	for i := 1; i <= n; i++ {
		c[i][0] = 1
		for j := 1; j <= i; j++ {
			c[i][j] = c[i-1][j] + c[i-1][j-1]
		}
	}

	return c
}

func NextPermutation[S ~[]E, E cmp.Ordered](arr S) bool {
	n := len(arr)
	if n <= 1 {
		return false
	}

	i := n - 2
	for ; i > 0 && arr[i] >= arr[i+1]; i-- {
	}

	for j := n - 1; j > i; j-- {
		if arr[j] > arr[i] {
			arr[i], arr[j] = arr[j], arr[i]
			slices.Sort(arr[i+1:])
			return true
		}
	}

	slices.Sort(arr)
	return false
}
