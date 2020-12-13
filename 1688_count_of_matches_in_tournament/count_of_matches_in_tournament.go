package count_of_matches_in_tournament

// https://leetcode.com/problems/count-of-matches-in-tournament

// level: 2
// time: O(log(n)) 0ms 100%
// space: O(1) 2M 100%

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfMatches(n int) int {
	ret := 0
	rest := 0
	for n > 1 {
		rest = n % 2
		n /= 2
		ret += n
		n += rest
	}
	return ret + rest
}

// leetcode submit region end(Prohibit modification and deletion)
