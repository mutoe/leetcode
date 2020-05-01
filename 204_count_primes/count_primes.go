package count_primes

import "math"

// https://leetcode.com/problems/count-primes

// time: O(n*log(n)) 1460ms 18.18%
// space: O(1) 1.9M

// leetcode submit region begin(Prohibit modification and deletion)
func isPrime(n int) bool {
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) (ret int) {
	if n < 3 {
		return 0
	}
	for i := 3; i < n; i++ {
		if isPrime(i) {
			ret += 1
		}
	}
	return ret + 1
}

// leetcode submit region end(Prohibit modification and deletion)
