package trie

type Node struct {
	isWord bool
	next   map[byte]*Node
}

type Trie struct {
	root *Node
	size int
}

func (t *Trie) Add(word string) {
	if t.root == nil {
		t.root = &Node{}
	}
	cur := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if cur.next == nil {
			cur.next = make(map[byte]*Node)
		}
		if _, exist := cur.next[char]; !exist {
			cur.next[char] = &Node{}
		}
		cur = cur.next[char]
	}
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exist := cur.next[char]; !exist {
			return false
		}
		cur = cur.next[char]
	}
	return cur.isWord
}
