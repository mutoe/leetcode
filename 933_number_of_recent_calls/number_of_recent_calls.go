package number_of_recent_calls

// https://leetcode.com/problems/number-of-recent-calls

// level: 1
// time: O(n) 128ms 79.37%
// space: O(n) 10.8M 19.93%

// leetcode submit region begin(Prohibit modification and deletion)
type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	i := 0
	for i < len(this.queue) && this.queue[i] < t-3000 {
		i++
	}
	this.queue = append(this.queue[i:], t)
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// leetcode submit region end(Prohibit modification and deletion)
