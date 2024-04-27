package ds

import "math"

// SegmentTree 线段树
type SegmentTree struct {
	tree  []int
	lazy  []int
	zero  int
	set   func(*int, int)
	merge func(int, int) int
}

func NewSegmentTree(n int, zero int, set func(node *int, val int), merge func(l int, r int) int) *SegmentTree {
	tree := make([]int, 4*n)
	for i := range tree {
		tree[i] = zero
	}
	return &SegmentTree{tree, make([]int, 4*n), zero, set, merge}
}

func NewSumSegmentTree(n int) *SegmentTree {
	return NewSegmentTree(
		n,
		0,
		func(node *int, val int) { *node = val },
		func(l int, r int) int { return l + r },
	)
}

func NewMaxSegmentTree(n int) *SegmentTree {
	return NewSegmentTree(
		n,
		math.MinInt,
		func(node *int, val int) { *node = max(*node, val) },
		func(l int, r int) int { return max(l, r) },
	)
}

func (st *SegmentTree) Update(index int, val int) {
	var dfs func(node, nodeLeft, nodeRight int)
	dfs = func(node, nodeLeft, nodeRight int) {
		if nodeLeft == nodeRight {
			st.set(&st.tree[node], val)
			return
		}
		nodeMid := (nodeLeft + nodeRight) / 2
		if index <= nodeMid {
			dfs(2*node, nodeLeft, nodeMid)
		} else {
			dfs(2*node+1, nodeMid+1, nodeRight)
		}
		st.tree[node] = st.merge(st.tree[2*node], st.tree[2*node+1])
	}
	dfs(1, 0, len(st.tree)/4-1)
}

func (st *SegmentTree) UpdateRange(left int, right int, delta int) {
	var dfs func(node, nodeLeft, nodeRight int)
	dfs = func(node, nodeLeft, nodeRight int) {
		if nodeLeft > right || nodeRight < left {
		} else if nodeLeft >= left && nodeRight <= right {
			st.lazy[node] += delta
		} else {
			nodeMid := (nodeLeft + nodeRight) / 2
			dfs(2*node, nodeLeft, nodeMid)
			dfs(2*node+1, nodeMid+1, nodeRight)
		}
	}
	dfs(1, 0, len(st.tree)/4-1)
}

func (st *SegmentTree) pushDown(node, nodeLeft, nodeRight int) {
	if st.lazy[node] != 0 {
		if nodeLeft != nodeRight {
			st.lazy[2*node] += st.lazy[node]
			st.lazy[2*node+1] += st.lazy[node]
		}
		st.tree[node] += (nodeRight - nodeLeft + 1) * st.lazy[node]
		st.lazy[node] = 0

		// push up
		for ; node > 0; node /= 2 {
			st.tree[node/2] = st.merge(st.tree[node/2*2], st.tree[node/2*2+1])
		}
	}
}

func (st *SegmentTree) Range(left int, right int) int {
	var dfs func(node, nodeLeft, nodeRight int) int
	dfs = func(node, nodeLeft, nodeRight int) int {
		st.pushDown(node, nodeLeft, nodeRight)

		if nodeLeft > right || nodeRight < left {
			return st.zero
		} else if nodeLeft >= left && nodeRight <= right {
			return st.tree[node]
		} else {
			nodeMid := (nodeLeft + nodeRight) / 2
			l := dfs(2*node, nodeLeft, nodeMid)
			r := dfs(2*node+1, nodeMid+1, nodeRight)
			return st.merge(l, r)
		}
	}
	return dfs(1, 0, len(st.tree)/4-1)
}
