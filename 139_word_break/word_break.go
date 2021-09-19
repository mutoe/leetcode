package word_break

// https://leetcode.com/problems/word-break

// level: 2
// time:
// space:

// TODO: unresolved problem

//leetcode submit region begin(Prohibit modification and deletion)

func include(dict []string, target string) bool {
	for _, s := range dict {
		if s == target {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && include(wordDict, s[:i]) {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

//leetcode submit region end(Prohibit modification and deletion)
