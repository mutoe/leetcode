package design_add_and_search_words_data_structure

// https://leetcode.com/problems/design-add-and-search-words-data-structure

// level: 2
// time: O(n) 96ms 38.61%
// space: O(n) 13M 84.16%

//leetcode submit region begin(Prohibit modification and deletion)
type Node struct {
	isWord bool
	next   map[byte]*Node
}

type WordDictionary struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := &Node{false, make(map[byte]*Node)}
	return WordDictionary{root}
}

func (wd *WordDictionary) AddWord(word string) {
	cur := wd.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if cur.next[char] == nil {
			cur.next[char] = &Node{false, make(map[byte]*Node)}
		}
		cur = cur.next[char]
	}
	cur.isWord = true
}

func (wd *WordDictionary) Search(word string) bool {
	return wd.match(wd.root, word, 0)
}

func (wd *WordDictionary) match(node *Node, word string, index int) bool {
	if index == len(word) {
		return node.isWord
	}
	char := word[index]
	if char != '.' {
		if node.next[char] == nil {
			return false
		}
		return wd.match(node.next[char], word, index+1)
	}
	for _, n := range node.next {
		if wd.match(n, word, index+1) {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
//leetcode submit region end(Prohibit modification and deletion)
