package leetcode

import (
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
// 数据结构

// UnionFind 并查集
type UnionFind[T comparable] struct {
	father map[T]T
	count  map[T]int
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{make(map[T]T), make(map[T]int)}
}

func (this *UnionFind[T]) size(t T) int {
	return this.count[this.find(t)] + 1
}

func (this *UnionFind[T]) find(t T) T {
	if _, ok := this.father[t]; !ok {
		return t
	}
	old := this.father[t]
	this.father[t] = this.find(this.father[t])

	if old != this.father[t] {
		this.count[this.father[t]]++
		delete(this.count, old)
	}

	return this.father[t]
}

func (this *UnionFind[T]) isConnected(a T, b T) bool {
	return this.find(a) == this.find(b)
}

func (this *UnionFind[T]) union(a T, b T) bool {
	pa := this.find(a)
	pb := this.find(b)
	if pa == pb {
		return false
	}

	this.father[pb] = pa

	this.count[pa]++
	delete(this.count, pb)
	this.find(a)
	this.find(b)

	return true
}

// MonotonicStack 单调栈
type MonotonicStack[T sort.Interface] struct {
	source   T
	stack    []int
	topIndex int
}

func NewMonotonicStack[T sort.Interface](source T) *MonotonicStack[T] {
	return &MonotonicStack[T]{source, make([]int, source.Len()), 0}
}

func (this *MonotonicStack[T]) size() int {
	return this.topIndex
}

func (this *MonotonicStack[T]) isEmpty() bool {
	return this.size() == 0
}

func (this *MonotonicStack[T]) top() int {
	return this.stack[this.topIndex-1]
}

func (this *MonotonicStack[T]) pop() int {
	res := this.top()
	this.topIndex--
	return res
}

func (this *MonotonicStack[T]) push(index int) []int {
	res := make([]int, 0)
	for this.size() > 0 && this.source.Less(this.top(), index) {
		res = append(res, this.pop())
	}

	if this.topIndex == len(this.stack) {
		this.stack = append(this.stack, index)
	} else {
		this.stack[this.topIndex] = index
	}
	this.topIndex++

	return res
}

// MonotonicQueue 单调队列
type MonotonicQueue struct {
	queue []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{nil}
}

func (this *MonotonicQueue) size() int {
	return len(this.queue)
}

func (this *MonotonicQueue) isEmpty() bool {
	return this.size() == 0
}

func (this *MonotonicQueue) peek() int {
	return this.queue[0]
}

func (this *MonotonicQueue) dequeue(head int) {
	if !this.isEmpty() && this.queue[0] == head {
		this.queue = this.queue[1:]
	}
}

func (this *MonotonicQueue) enqueue(tail int) {
	for !this.isEmpty() && tail > this.queue[this.size()-1] {
		this.queue = this.queue[:this.size()-1]
	}
	this.queue = append(this.queue, tail)
}

// // ////////////////////////////////////////////////

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

func reverse[T any](arr []T) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
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
