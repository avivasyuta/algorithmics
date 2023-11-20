package main

type Trie struct {
	children   map[rune]*Trie
	isFullWord bool
}

func Constructor() Trie {
	return Trie{
		children: make(map[rune]*Trie),
	}
}

func (t *Trie) Insert(word string) {
	t.insert([]rune(word))
}

func (t *Trie) Search(word string) bool {
	node, exist := t.traverse([]rune(word))
	if !exist {
		return false
	}

	return node.isFullWord
}

func (t *Trie) StartsWith(prefix string) bool {
	_, exist := t.traverse([]rune(prefix))

	return exist
}

func (t *Trie) insert(word []rune) {
	if len(word) == 0 {
		t.isFullWord = true
		return
	}

	child, exist := t.children[word[0]]
	if !exist {
		child = &Trie{
			children: make(map[rune]*Trie),
		}

		t.children[word[0]] = child
	}

	child.insert(word[1:])
}

func (t *Trie) traverse(prefix []rune) (*Trie, bool) {
	if len(prefix) == 0 {
		return t, true
	}

	child, exist := t.children[prefix[0]]
	if !exist {
		return nil, false
	}

	return child.traverse(prefix[1:])
}
