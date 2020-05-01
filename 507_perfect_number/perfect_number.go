package perfect_number

// https://leetcode.com/problems/perfect-number

// time: O(log(n)) 1124ms 40.38%
// space: O(1) 2M

// leetcode submit region begin(Prohibit modification and deletion)
func checkPerfectNumber(num int) bool {
	if num < 6 {
		return false
	}
	largestDivisor := num / 2
	divisorSum := 1
	for i := 2; i < largestDivisor; i++ {
		if num%i == 0 {
			divisorSum += i + num/i
			largestDivisor = num / i
		}
	}
	return divisorSum == num
}

// leetcode submit region end(Prohibit modification and deletion)
