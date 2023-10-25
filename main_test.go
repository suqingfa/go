package example

import "testing"

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind[int]()

	uf.union(1, 2)

	if uf.find(2) != 1 {
		t.Error()
	}

	if !uf.isConnected(1, 2) {
		t.Error()
	}

	if uf.isConnected(1, 3) {
		t.Error()
	}
}
