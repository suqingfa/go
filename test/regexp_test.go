package test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	matched, err := regexp.MatchString(`^Go\w*123$`, "Golang123")
	assert.Nil(t, err)
	assert.True(t, matched)
}

func TestFindString(t *testing.T) {
	reg := regexp.MustCompile(`[+-]?\d*x?`)
	assert.Equal(t, []string{"x", "+5", "-3", "+x"}, reg.FindAllString("x+5-3+x", -1))
}
