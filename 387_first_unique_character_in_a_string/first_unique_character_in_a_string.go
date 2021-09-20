package first_unique_character_in_a_string

import "math"

// https://leetcode.com/problems/first-unique-character-in-a-string

// level: 1
// time: O(n) 9ms 65.65%
// space: O(1) 5.7M 33.02%

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	bucket := [26]int{}
	for i := 0; i < len(s); i++ {
		if bucket[s[i]-'a'] == 0 {
			bucket[s[i]-'a'] = i + 1
		} else {
			bucket[s[i]-'a'] = -1
		}
	}
	min := math.MaxInt32
	for _, charIndex := range bucket {
		if charIndex > 0 && charIndex-1 < min {
			min = charIndex - 1
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
