package example

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

func (s *UnionFind[T]) Find(t T) T {
	if _, ok := s.leader[t]; !ok {
		return t
	}
	s.leader[t] = s.Find(s.leader[t])

	return s.leader[t]
}

func (s *UnionFind[T]) Connected(a T, b T) bool {
	return s.Find(a) == s.Find(b)
}

func (s *UnionFind[T]) Union(a T, b T) bool {
	la := s.Find(a)
	lb := s.Find(b)
	if la == lb {
		return false
	}

	s.leader[lb] = la

	return true
}

// MonotonicStack 单调栈 从栈顶到栈底的元素是单调递增（或者单调递减）
type MonotonicStack[T sort.Interface] struct {
	source   T
	stack    []int
	topIndex int
}

func NewMonotonicStack[T sort.Interface](source T) *MonotonicStack[T] {
	return &MonotonicStack[T]{source, make([]int, source.Len()), 0}
}

func (s *MonotonicStack[T]) Size() int {
	return s.topIndex
}

func (s *MonotonicStack[T]) TopN(n int) int {
	return s.stack[s.topIndex-n-1]
}

func (s *MonotonicStack[T]) Top() int {
	return s.TopN(0)
}

func (s *MonotonicStack[T]) Pop() int {
	res := s.Top()
	s.topIndex--
	return res
}

// Push 将 source[index] 的索引 index 压入栈
// 返回值 栈顶大于index的索引
func (s *MonotonicStack[T]) Push(index int) []int {
	res := make([]int, 0)
	for s.Size() > 0 && s.source.Less(s.Top(), index) {
		res = append(res, s.Pop())
	}

	s.stack[s.topIndex] = index
	s.topIndex++

	return res
}

// MonotonicQueue 单调队列 队列中的数据单调递增 主要用来辅助解决滑动窗口相关的问题
type MonotonicQueue struct {
	queue []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{nil}
}

func (s *MonotonicQueue) Size() int {
	return len(s.queue)
}

func (s *MonotonicQueue) Peek() int {
	return s.queue[0]
}

func (s *MonotonicQueue) Dequeue(head int) bool {
	res := s.Size() != 0 && s.queue[0] == head
	if res {
		s.queue = s.queue[1:]
	}
	return res
}

func (s *MonotonicQueue) Enqueue(value int) {
	for s.Size() != 0 && value > s.queue[s.Size()-1] {
		s.queue = s.queue[:s.Size()-1]
	}
	s.queue = append(s.queue, value)
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

func (s *Trie) findChild(c byte, create bool) *Trie {
	if _, ok := s.child[c]; !ok && create {
		s.child[c] = &Trie{root: false, child: map[byte]*Trie{}, key: string(c)}
	}

	return s.child[c]
}

func (s *Trie) Insert(word string) {
	if s.root {
		child := s.findChild(word[0], true)
		child.Insert(word)
		return
	}

	if len(word) == 1 {
		s.end = true
		return
	}

	child := s.findChild(word[1], true)
	child.Insert(word[1:])
}

func (s *Trie) find(word string, findWithPrefix bool) bool {
	if s.root {
		child := s.findChild(word[0], false)
		if child == nil {
			return false
		}
		return child.find(word, findWithPrefix)
	}

	if len(word) == 1 {
		return findWithPrefix || s.end
	}

	child := s.findChild(word[1], false)
	if child == nil {
		return false
	}
	return child.find(word[1:], findWithPrefix)
}

func (s *Trie) Search(word string) bool {
	return s.find(word, false)
}

// StartsWith 字典树中是否存在以 prefix 为前缀的词
func (s *Trie) StartsWith(prefix string) bool {
	return s.find(prefix, true)
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
func (s *SegmentTree) mid() int {
	return (s.start + s.end) / 2
}

func (s *SegmentTree) getLeft() *SegmentTree {
	if s.left == nil {
		s.left = &SegmentTree{start: s.start, end: s.mid()}
	}
	return s.left
}

func (s *SegmentTree) getRight() *SegmentTree {
	if s.right == nil {
		s.right = &SegmentTree{start: s.mid() + 1, end: s.end}
	}
	return s.right
}

func (s *SegmentTree) Insert(node int) {
	s.value++
	if node == s.start && node == s.end {
		return
	}

	if node <= s.mid() {
		s.getLeft().Insert(node)
	} else {
		s.getRight().Insert(node)
	}
}

func (s *SegmentTree) Search(start int, end int) int {
	if start == s.start && end == s.end {
		return s.value
	}

	if end <= s.mid() {
		return s.getLeft().Search(start, end)
	} else if s.mid() < start {
		return s.getRight().Search(start, end)
	} else {
		return s.getLeft().Search(start, s.mid()) + s.getRight().Search(s.mid()+1, end)
	}
}

// // ////////////////////////////////////////////////

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

func Reverse[T any](arr []T) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
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
