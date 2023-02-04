package main

import (
	"github.com/emirpasic/gods/utils"
	"sort"
)

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

func nextPermutation(nums []int) bool {
	n := len(nums)

	i := n - 2
	for ; i > 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i < 0 {
		i = 0
	}

	for j := n - 1; j > i; j-- {
		if nums[j] > nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
			sort.Ints(nums[i+1:])
			return true
		}
	}

	sort.Ints(nums)
	return false
}

// ////////////////////////////////////////////

func main() {
	println(utils.ToString(
		"",
	))
}
