package main

import "github.com/emirpasic/gods/utils"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ////////////////////////////////////////////////

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primes(n int) []bool {
	res := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		res[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if res[i] {
			for j := 2 * i; j <= n; j += i {
				res[j] = false
			}
		}
	}

	return res
}

// ////////////////////////////////////////////

func main() {
	println(utils.ToString("hello world"))
}
