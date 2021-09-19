package map_sum_pairs

// https://leetcode.com/problems/map-sum-pairs

// level: 2
// time: O(n) 0ms 100%
// space: O(n) 3M 19.17%

//leetcode submit region begin(Prohibit modification and deletion)
type Node struct {
	weight int
	next   map[byte]*Node
}

type MapSum struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	root := &Node{0, make(map[byte]*Node)}
	return MapSum{root}
}

func (ms *MapSum) Insert(key string, val int) {
	cur := ms.root
	for i := 0; i < len(key); i++ {
		char := key[i]
		if cur.next[char] == nil {
			cur.next[char] = &Node{0, make(map[byte]*Node)}
		}
		cur = cur.next[char]
	}
	cur.weight = val
}

func (ms *MapSum) sum(node *Node) int {
	sum := node.weight
	for _, n := range node.next {
		sum += ms.sum(n)
	}
	return sum
}

func (ms *MapSum) Sum(prefix string) int {
	cur := ms.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if cur.next[char] == nil {
			return 0
		}
		cur = cur.next[char]
	}
	return ms.sum(cur)
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
