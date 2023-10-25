package example

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind[int]()

	if !uf.Union(1, 2) {
		t.Error()
	}

	if uf.Union(1, 2) {
		t.Error()
	}

	if uf.Find(2) != 1 {
		t.Error()
	}

	if !uf.Connected(1, 2) {
		t.Error()
	}

	if uf.Connected(1, 3) {
		t.Error()
	}
}

func TestMonotonicStack(t *testing.T) {
	source := []int{3, 1, 2, 3, 2, 1, 4}
	stack := NewMonotonicStack[sort.IntSlice](source)

	except := [][]int{
		{},           // [3]
		{},           // [3,1]
		{1},          // [3,2]
		{2},          // [3,3]
		{},           // [3,3,2]
		{},           // [3,3,2,1]
		{5, 4, 3, 0}, // [4]
	}

	for i := 0; i < len(source); i++ {
		push := stack.Push(i)
		if !reflect.DeepEqual(push, except[i]) {
			t.Error(push)
		}
	}
}

func TestMonotonicQueue(t *testing.T) {
	source := []int{3, 1, 2, 3, 2, 1, 4}
	queue := NewMonotonicQueue()

	except := [][]int{
		{3},
		{3, 1},
		{3, 2},
		{3, 3},
		{3, 3, 2},
		{3, 3, 2, 1},
		{4},
	}

	for i, v := range source {
		queue.Enqueue(v)
		if !reflect.DeepEqual(queue.queue, except[i]) {
			t.Error()
		}
	}

	if queue.Peek() != 4 {
		t.Error()
	}

	if queue.Dequeue(3) {
		t.Error()
	}

	if !queue.Dequeue(4) {
		t.Error()
	}

	if queue.Dequeue(4) {
		t.Error()
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()

	trie.Insert("a")
	trie.Insert("ab")
	trie.Insert("abc")
	trie.Insert("abd")

	searchExcept := map[string]bool{
		"a":    true,
		"ab":   true,
		"abc":  true,
		"abcd": false,
		"x":    false,
		"xy":   false,
	}

	for s, b := range searchExcept {
		if trie.Search(s) != b {
			t.Error(s, b)
		}
	}

	startWithExcept := map[string]bool{
		"a":    true,
		"ab":   true,
		"abc":  true,
		"abcd": false,
		"x":    false,
		"xy":   false,
	}

	for s, b := range startWithExcept {
		if trie.StartsWith(s) != b {
			t.Error(s, b)
		}
	}
}

func TestSegmentTree(t *testing.T) {
	const N = 100
	segmentTree := NewSegmentTree(0, N-1)

	ints := make([]int, N*10)
	for i := range ints {
		ints[i] = rand.Intn(N)
	}

	for _, i := range ints {
		segmentTree.Insert(i)
	}

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			cnt := 0
			for _, v := range ints {
				if i <= v && v <= j {
					cnt++
				}
			}

			if segmentTree.Search(i, j) != cnt {
				t.Error()
			}
		}
	}
}

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
