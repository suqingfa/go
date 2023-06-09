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
	leader map[T]T
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{make(map[T]T)}
}

func (this *UnionFind[T]) find(t T) T {
	if _, ok := this.leader[t]; !ok {
		return t
	}
	this.leader[t] = this.find(this.leader[t])

	return this.leader[t]
}

func (this *UnionFind[T]) isConnected(a T, b T) bool {
	return this.find(a) == this.find(b)
}

func (this *UnionFind[T]) union(a T, b T) bool {
	la := this.find(a)
	lb := this.find(b)
	if la == lb {
		return false
	}

	this.leader[lb] = la

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

func (this *MonotonicStack[T]) topN(n int) int {
	return this.stack[this.topIndex-n-1]
}

func (this *MonotonicStack[T]) top() int {
	return this.topN(0)
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

// Trie 字典树
type Trie struct {
	child map[byte]*Trie
	key   string
	root  bool
	end   bool
}

func NewTrie() *Trie {
	return &Trie{root: true, child: map[byte]*Trie{}}
}

func (this *Trie) findChild(c byte, create bool) *Trie {
	if _, ok := this.child[c]; !ok && create {
		this.child[c] = &Trie{root: false, child: map[byte]*Trie{}, key: string(c)}
	}

	return this.child[c]
}

func (this *Trie) Insert(word string) {
	if this.root {
		child := this.findChild(word[0], true)
		child.Insert(word)
		return
	}

	if len(word) == 1 {
		this.end = true
		return
	}

	child := this.findChild(word[1], true)
	child.Insert(word[1:])
}

func (this *Trie) find(word string, findWithPrefix bool) bool {
	if this.root {
		child := this.findChild(word[0], false)
		if child == nil {
			return false
		}
		return child.find(word, findWithPrefix)
	}

	if len(word) == 1 {
		return findWithPrefix || this.end
	}

	child := this.findChild(word[1], false)
	if child == nil {
		return false
	}
	return child.find(word[1:], findWithPrefix)
}

func (this *Trie) Search(word string) bool {
	return this.find(word, false)
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix, true)
}

func (this *Trie) Prefix(word string) (string, bool) {
	if this.root {
		child := this.findChild(word[0], false)
		if child == nil {
			return "", false
		}
		return child.Prefix(word)
	}

	if this.end {
		return word[:1], true
	}

	if len(word) == 1 {
		return "", false
	}

	child := this.findChild(word[1], false)
	if child == nil {
		return "", false
	}

	prefix, find := child.Prefix(word[1:])
	if !find {
		return "", false
	}

	return word[:1] + prefix, true
}

// SegmentTree 线段树
type SegmentTree struct {
	value int

	start int
	end   int

	// [start, mid]
	left *SegmentTree
	// [mid+1, end]
	right *SegmentTree
}

func NewSegmentTree(start int, end int) *SegmentTree {
	return &SegmentTree{start: start, end: end}
}
func (this *SegmentTree) mid() int {
	return (this.start + this.end) / 2
}

func (this *SegmentTree) getLeft() *SegmentTree {
	if this.left == nil {
		this.left = &SegmentTree{start: this.start, end: this.mid()}
	}
	return this.left
}

func (this *SegmentTree) getRight() *SegmentTree {
	if this.right == nil {
		this.right = &SegmentTree{start: this.mid() + 1, end: this.end}
	}
	return this.right
}

func (this *SegmentTree) Insert(node int) {
	this.value++
	if node == this.start && node == this.end {
		return
	}

	if node <= this.mid() {
		this.getLeft().Insert(node)
	} else {
		this.getRight().Insert(node)
	}
}

func (this *SegmentTree) Search(start int, end int) int {
	if start == this.start && end == this.end {
		return this.value
	}

	if end <= this.mid() {
		return this.getLeft().Search(start, end)
	} else if this.mid() < start {
		return this.getRight().Search(start, end)
	} else {
		return this.getLeft().Search(start, this.mid()) + this.getRight().Search(this.mid()+1, end)
	}
}

// // ////////////////////////////////////////////////

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func min(arr ...int) int {
	res := arr[0]
	for _, i := range arr {
		if res > i {
			res = i
		}
	}

	return res
}

func max(arr ...int) int {
	res := arr[0]
	for _, i := range arr {
		if res < i {
			res = i
		}
	}

	return res
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
