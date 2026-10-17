package bench

import (
	"strings"
	"testing"
)

func BenchmarkStringAdd(b *testing.B) {
	s := ""
	for range b.N {
		s += "abc"
	}
	b.Log(len(s))
}

func BenchmarkStringBuilder(b *testing.B) {
	builder := strings.Builder{}
	for range b.N {
		builder.WriteString("abc")
	}
	b.Log(builder.Len())
}
