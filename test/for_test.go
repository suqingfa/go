package test

import (
	"iter"
	"testing"
)

func TestFor(t *testing.T) {
	var seq iter.Seq[int] = func(yield func(int) bool) {
		for i := range 10 {
			yield(i)
		}
	}

	for i := range seq {
		t.Log(i)
	}

	var seq2 iter.Seq2[int, string] = func(yield func(int, string) bool) {
		arr := []string{"a", "b", "c"}
		for i, s := range arr {
			yield(i, s)
		}
	}

	for k, v := range seq2 {
		t.Log(k, v)
	}
}
