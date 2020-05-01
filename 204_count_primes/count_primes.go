package count_primes

// https://leetcode.com/problems/count-primes

// Euler sieve
// time: O(n) 8ms 81.42%
// space: O(n) 4.9M

// leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) (ret int) {
	if n < 3 {
		return 0
	}
	primes := make([]int, n+1)
	vis := make([]bool, n+1)

	for i := 2; i < n; i++ {
		if !vis[i] {
			primes[ret] = i
			ret++
		}
		for j := 0; j < ret && i*primes[j] < n; j++ {
			vis[i*primes[j]] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}

	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
