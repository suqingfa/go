package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, add(1, 2))
	assert.Equal(t, 1.5, add(1.0, 0.5))
	assert.Equal(t, "abc", add("ab", "c"))
}

func TestIf(t *testing.T) {
	assert.Equal(t, 1, If(true, 1, 2))
	assert.Equal(t, 2, If(false, 1, 2))
}
