package ds

// Trie 字典树
type Trie struct {
	children map[rune]*Trie
	isLeaf   bool
}

func NewTrie() *Trie {
	return &Trie{map[rune]*Trie{}, false}
}

func (t *Trie) Insert(word string) {
	cur := t
	for _, r := range word {
		next, ok := cur.children[r]
		if !ok {
			next = NewTrie()
			cur.children[r] = next
		}
		cur = next
	}
	cur.isLeaf = true
}

func (t *Trie) find(word string, findWithPrefix bool) bool {
	next, ok := t, false
	for _, r := range word {
		next, ok = next.children[r]
		if !ok {
			return false
		}
	}

	if findWithPrefix {
		return true
	}

	return next.isLeaf
}

func (t *Trie) Find(word string) bool {
	return t.find(word, false)
}

// HasPrefixString 字典树中是否存在以 prefix 为前缀的词
func (t *Trie) HasPrefixString(prefix string) bool {
	return t.find(prefix, true)
}
