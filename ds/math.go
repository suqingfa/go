package ds

import "sort"

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return Gcd(b, a%b)
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

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func initPrimes(n int) []int {
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

func NextPermutation(arr []int) bool {
	n := len(arr)

	i := n - 2
	for ; i > 0 && arr[i] >= arr[i+1]; i-- {
	}

	for j := n - 1; j > i; j-- {
		if arr[j] > arr[i] {
			arr[i], arr[j] = arr[j], arr[i]
			sort.Ints(arr[i+1:])
			return true
		}
	}

	sort.Ints(arr)
	return false
}

func Reverse[T any](arr []T) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}
