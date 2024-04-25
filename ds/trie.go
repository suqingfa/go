package ds

// Trie 字典树
type Trie struct {
	child map[byte]*Trie
	key   string
	root  bool
	end   bool
}

func NewTrie() *Trie {
	return &Trie{root: true, child: map[byte]*Trie{}}
}

func (t *Trie) findChild(c byte, create bool) *Trie {
	if _, ok := t.child[c]; !ok && create {
		t.child[c] = &Trie{root: false, child: map[byte]*Trie{}, key: string(c)}
	}

	return t.child[c]
}

func (t *Trie) Insert(word string) {
	if t.root {
		child := t.findChild(word[0], true)
		child.Insert(word)
		return
	}

	if len(word) == 1 {
		t.end = true
		return
	}

	child := t.findChild(word[1], true)
	child.Insert(word[1:])
}

func (t *Trie) find(word string, findWithPrefix bool) bool {
	if t.root {
		child := t.findChild(word[0], false)
		if child == nil {
			return false
		}
		return child.find(word, findWithPrefix)
	}

	if len(word) == 1 {
		return findWithPrefix || t.end
	}

	child := t.findChild(word[1], false)
	if child == nil {
		return false
	}
	return child.find(word[1:], findWithPrefix)
}

func (t *Trie) Search(word string) bool {
	return t.find(word, false)
}

// StartsWith 字典树中是否存在以 prefix 为前缀的词
func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix, true)
}
