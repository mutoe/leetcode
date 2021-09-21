package longest_chunked_palindrome_decomposition

// https://leetcode.com/problems/longest-chunked-palindrome-decomposition

// level: 3
// time: O(n) 0ms 100%
// space: O(n) 2.2M 59.52%

//leetcode submit region begin(Prohibit modification and deletion)
const P = int64(1e9 + 7)

const B = 26

func longestDecomposition(text string) int {
	count := 0
	N := len(text)
	pow26 := make([]int64, N+1)
	pow26[0] = 1
	for i := 1; i < len(pow26); i++ {
		pow26[i] = pow26[i-1] * B % P
	}

	leftHash, rightHash, length := int64(0), int64(0), 1
	i, j := 0, N-1
	for ; i <= j; i, j = i+1, j-1 {
		if i == j && length == 1 {
			count++
			break
		}

		leftHash = (leftHash*B + int64(text[i]-'a')) % P
		pow := pow26[length-1]
		rightHash = (int64(text[j]-'a')*pow + rightHash) % P
		if leftHash == rightHash && text[i+1-length:i+1] == text[j:j+length] {
			count += 2
			leftHash, rightHash, length = 0, 0, 1
			continue
		}
		length++
	}
	if leftHash != 0 || rightHash != 0 {
		count++
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
