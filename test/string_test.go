package test

import (
	"strings"
	"testing"
)

func BenchmarkStringAdd(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		s += "abc"
	}
	b.Log(len(s))
}

func BenchmarkStringBuilder(b *testing.B) {
	builder := strings.Builder{}
	for i := 0; i < b.N; i++ {
		builder.WriteString("abc")
	}
	b.Log(builder.Len())
}
