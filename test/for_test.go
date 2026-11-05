package test

import (
	"iter"
	"testing"
)

func TestFor(t *testing.T) {
	var seq iter.Seq[int] = func(yield func(int) bool) {
		if !yield(1) {
			return
		}
		if !yield(2) {
			return
		}
		if !yield(3) {
			return
		}
		yield(4)
	}

	for i := range seq {
		t.Log(i)
	}

	var seq2 iter.Seq2[int, string] = func(yield func(int, string) bool) {
		if !yield(1, "a") {
			return
		}
		if !yield(2, "b") {
			return
		}
		yield(3, "c")
	}

	for k, v := range seq2 {
		t.Log(k, v)
	}
}
