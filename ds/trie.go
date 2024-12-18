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

func (t *Trie) search(word string, findWithPrefix bool) bool {
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

func (t *Trie) Search(word string) bool {
	return t.search(word, false)
}

// StartWith 字典树中是否存在以 prefix 为前缀的词
func (t *Trie) StartWith(prefix string) bool {
	return t.search(prefix, true)
}
