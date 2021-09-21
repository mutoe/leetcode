package longest_happy_prefix

// https://leetcode.com/problems/longest-happy-prefix

// level: 3
// time: O(n) 16ms 72.22%
// space: O(n) 7.6M 50%

//leetcode submit region begin(Prohibit modification and deletion)
const B = 26
const P = int64(10e9 + 7)

func longestPrefix(s string) string {
	N := len(s)
	if N == 1 {
		return ""
	}
	pow26 := make([]int64, N+1)
	pow26[0] = 1
	for i := 0; i < N; i++ {
		pow26[i+1] = pow26[i] * B % P
	}

	length := 0
	leftHash, rightHash := int64(0), int64(0)
	for i := 0; i < N-1; i++ {
		leftChar, rightChar := int64(s[i]-'a'), int64(s[N-i-1]-'a')
		leftHash = (leftHash*B + leftChar) % P
		pow := pow26[i]
		rightHash = (rightChar*pow + rightHash) % P
		if leftHash == rightHash {
			length = i + 1
		}
	}
	if length == 0 {
		return ""
	}
	return s[:length]
}

//leetcode submit region end(Prohibit modification and deletion)
