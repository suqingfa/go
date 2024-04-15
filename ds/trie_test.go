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

	searchExcept := map[string]bool{
		"a":    true,
		"ab":   true,
		"abc":  true,
		"abcd": false,
		"x":    false,
		"xy":   false,
	}

	for s, b := range searchExcept {
		assert.Equal(t, b, trie.Search(s))
	}

	startWithExcept := map[string]bool{
		"a":    true,
		"ab":   true,
		"abc":  true,
		"abcd": false,
		"x":    false,
		"xy":   false,
	}

	for s, b := range startWithExcept {
		assert.Equal(t, b, trie.StartsWith(s))
	}
}
