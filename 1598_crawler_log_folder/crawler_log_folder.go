package crawler_log_folder

// https://leetcode.com/problems/crawler-log-folder

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 3.1M 100%

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(logs []string) int {
	l := 0
	for _, log := range logs {
		if log == "../" {
			if l > 0 {
				l--
			}
		} else if log == "./" {
			continue
		} else {
			l++
		}
	}
	return l
}

// leetcode submit region end(Prohibit modification and deletion)
