package ds

// SegmentTree 线段树
type SegmentTree struct {
	tree  []int
	n     int
	zero  int
	set   func(*int, int)
	merge func(int, int) int
}

func NewSegmentTree(n int, zero int, set func(node *int, val int), merge func(l int, r int) int) SegmentTree {
	return SegmentTree{make([]int, 4*n), n, zero, set, merge}
}

func (this *SegmentTree) Update(index int, val int) {
	var dfs func(i int, l int, r int, index int, val int)
	dfs = func(i int, l int, r int, index int, val int) {
		if l == r {
			this.set(&this.tree[i], val)
			return
		}

		mid := (l + r) / 2
		if index <= mid {
			dfs(2*i+1, l, mid, index, val)

		} else {
			dfs(2*i+2, mid+1, r, index, val)
		}
		this.tree[i] = this.merge(this.tree[2*i+1], this.tree[2*i+2])
	}

	dfs(0, 0, this.n-1, index, val)
}

func (this *SegmentTree) Range(left int, right int) int {
	var dfs func(i int, l int, r int, left int, right int) int
	dfs = func(i int, l int, r int, left int, right int) int {
		if l > right || r < left {
			return this.zero
		} else if l >= left && r <= right {
			return this.tree[i]
		}

		mid := (l + r) / 2
		return this.merge(
			dfs(2*i+1, l, mid, left, right),
			dfs(2*i+2, mid+1, r, left, right),
		)
	}

	return dfs(0, 0, this.n-1, left, right)
}
