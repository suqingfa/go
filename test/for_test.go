package test

import (
	"iter"
	"testing"
)

func TestFor(t *testing.T) {
	var seq iter.Seq[int] = func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
		yield(4)
	}

	for i := range seq {
		t.Log(i)
	}

	var seq2 iter.Seq2[int, string] = func(yield func(int, string) bool) {
		yield(1, "a")
		yield(2, "b")
		yield(3, "c")
	}

	for k, v := range seq2 {
		t.Log(k, v)
	}
}
