package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	trie.Insert("a")
	trie.Insert("ab")
	trie.Insert("abc")
	trie.Insert("abd")

	findExcept := map[string]bool{
		"a":    true,
		"ab":   true,
		"abc":  true,
		"abcd": false,
		"x":    false,
		"xy":   false,
	}

	for s, b := range findExcept {
		assert.Equal(t, b, trie.Search(s))
	}

	prefixExcept := map[string]bool{
		"a":    true,
		"ab":   true,
		"abc":  true,
		"abcd": false,
		"x":    false,
		"xy":   false,
	}

	for s, b := range prefixExcept {
		assert.Equal(t, b, trie.StartWith(s))
	}
}
