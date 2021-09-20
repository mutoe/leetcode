package first_unique_character_in_a_string

// https://leetcode.com/problems/first-unique-character-in-a-string

// level: 1
// time: O(n) 4ms 97.53%
// space: O(1) 5.5M 43.83%

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	bucket := [26]int{}
	for i := 0; i < len(s); i++ {
		bucket[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if bucket[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
