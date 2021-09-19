package word_search_ii

// https://leetcode.com/problems/word-search-ii

// level: 3
// time: unknown 1432ms 26.06%
// space: O(n) 2.6M 62.42%

//leetcode submit region begin(Prohibit modification and deletion)

type Node struct {
	char byte
	word string
	next map[byte]*Node
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	root := &Node{byte(0), "", make(map[byte]*Node)}
	return Trie{root}
}

func (t *Trie) Add(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if cur.next[char] == nil {
			cur.next[char] = &Node{char, "", make(map[byte]*Node)}
		}
		cur = cur.next[char]
	}
	cur.word = word
}

type Result struct {
	board   [][]byte
	result  map[string]bool
	visited map[[2]int]bool
}

func (r *Result) search(x, y int, node *Node) {
	if x < 0 || y < 0 || y >= len(r.board) || x >= len(r.board[0]) || r.visited[[2]int{x, y}] {
		return
	}
	r.visited[[2]int{x, y}] = true
	for char, n := range node.next {
		if char != r.board[y][x] {
			continue
		}
		r.search(x+1, y, n)
		r.search(x-1, y, n)
		r.search(x, y+1, n)
		r.search(x, y-1, n)

		if n.word != "" {
			r.result[n.word] = true
			continue
		}
	}
	r.visited[[2]int{x, y}] = false
}

func findWords(board [][]byte, words []string) []string {
	t := Constructor()
	for _, word := range words {
		t.Add(word)
	}
	r := &Result{board, map[string]bool{}, map[[2]int]bool{}}
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[j]); i++ {
			r.search(i, j, t.root)
		}
	}
	ret, i := make([]string, len(r.result)), 0
	for s := range r.result {
		ret[i] = s
		i++
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
