package implement_trie_prefix_tree

// https://leetcode.com/problems/implement-trie-prefix-tree

// level: 2
// time: O(n) 72ms 50.75%
// space: O(n) 16.1M 73.68%

//leetcode submit region begin(Prohibit modification and deletion)
type Node struct {
	isWord bool
	next   map[byte]*Node
}

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := &Node{false, make(map[byte]*Node)}
	return Trie{root}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if cur.next[char] == nil {
			cur.next[char] = &Node{false, make(map[byte]*Node)}
		}
		cur = cur.next[char]
	}
	cur.isWord = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if cur.next[char] == nil {
			return false
		}
		cur = cur.next[char]
	}
	return cur.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if cur.next[char] == nil {
			return false
		}
		cur = cur.next[char]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
