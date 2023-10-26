package ds

import "testing"

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
