package ds

import "testing"

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
		if trie.Search(s) != b {
			t.Error(s, b)
		}
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
		if trie.StartsWith(s) != b {
			t.Error(s, b)
		}
	}
}
