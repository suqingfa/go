package ds

import (
	"cmp"
	"iter"
	"slices"
	"strconv"
)

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// ModInverse 当 m 是素数时, 求 ax mod m == 1 的解
func ModInverse(a, m int) int {
	return ModPower(a, m-2, m)
}

func ModPower(base, n, mod int) int {
	if n == 0 {
		return 1
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

// SymmetricNumberSeq 回文数字序列
func SymmetricNumberSeq(yield func(int) bool) {
	for i := range 9 {
		if !yield(i + 1) {
			return
		}
	}

	for start := 1; ; start *= 10 {
		for prefix := start; prefix < 10*start; prefix++ {
			itoa := strconv.Itoa(prefix)
			bytes := []byte(itoa)
			slices.Reverse(bytes)
			atoi, _ := strconv.Atoi(itoa + string(bytes))
			if !yield(atoi) {
				return
			}
		}

		for prefix := start; prefix < 10*start; prefix++ {
			itoa := strconv.Itoa(prefix)
			bytes := []byte(itoa)
			slices.Reverse(bytes)
			for mid := 0; mid < 10; mid++ {
				atoi, _ := strconv.Atoi(itoa + strconv.Itoa(mid) + string(bytes))
				if !yield(atoi) {
					return
				}
			}
		}
	}
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
			for j := i * i; j <= n; j += i {
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

// CNK 组合数
type CNK struct {
	m   int
	fac []int
}

func NewCNK(n, m int) *CNK {
	fac := make([]int, n+1)
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = (fac[i-1] * i) % m
	}
	return &CNK{m: m, fac: fac}
}

// c(n, k) = fac(n) / (fac(k)*fac(n-k))
func (cnk *CNK) cnk(n, k int) int {
	res := cnk.fac[n]
	res = (res * ModInverse(cnk.fac[k], cnk.m)) % cnk.m
	res = (res * ModInverse(cnk.fac[n-k], cnk.m)) % cnk.m
	return res
}

func Subset(mask int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for s := mask; yield(s) && s != 0; s = (s - 1) & mask {
		}
	}
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
