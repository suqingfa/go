package ds

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
