package test

import (
	"runtime"
	"testing"
	"weak"

	"github.com/stretchr/testify/assert"
)

type T struct {
	// N.B. This must contain a pointer, otherwise the weak handle might get placed
	// in a tiny block making the tests in this package flaky.
	t *T
	a int
}

func TestWeak(t *testing.T) {
	p := weak.Make(new(T))
	assert.NotNil(t, p.Value())

	runtime.GC()
	assert.Nil(t, p.Value())
}
