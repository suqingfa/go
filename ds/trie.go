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

func (s *Trie) findChild(c byte, create bool) *Trie {
	if _, ok := s.child[c]; !ok && create {
		s.child[c] = &Trie{root: false, child: map[byte]*Trie{}, key: string(c)}
	}

	return s.child[c]
}

func (s *Trie) Insert(word string) {
	if s.root {
		child := s.findChild(word[0], true)
		child.Insert(word)
		return
	}

	if len(word) == 1 {
		s.end = true
		return
	}

	child := s.findChild(word[1], true)
	child.Insert(word[1:])
}

func (s *Trie) find(word string, findWithPrefix bool) bool {
	if s.root {
		child := s.findChild(word[0], false)
		if child == nil {
			return false
		}
		return child.find(word, findWithPrefix)
	}

	if len(word) == 1 {
		return findWithPrefix || s.end
	}

	child := s.findChild(word[1], false)
	if child == nil {
		return false
	}
	return child.find(word[1:], findWithPrefix)
}

func (s *Trie) Search(word string) bool {
	return s.find(word, false)
}

// StartsWith 字典树中是否存在以 prefix 为前缀的词
func (s *Trie) StartsWith(prefix string) bool {
	return s.find(prefix, true)
}
