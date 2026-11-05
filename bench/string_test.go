package bench

import (
	"strings"
	"testing"
)

func BenchmarkStringAdd(b *testing.B) {
	s := ""
	for b.Loop() {
		s += "abc"
	}
	b.Log(len(s))
}

func BenchmarkStringBuilder(b *testing.B) {
	builder := strings.Builder{}
	for b.Loop() {
		builder.WriteString("abc")
	}
	b.Log(builder.Len())
}
