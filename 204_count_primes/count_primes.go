package count_primes

// https://leetcode.com/problems/count-primes

// Eratosthenes sieve
// time: O(n*log(n)) 12ms 57.31%
// space: O(n) 4.9M

// leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) (ret int) {
	if n < 3 {
		return 0
	}
	primes := make([]bool, n+1)

	primes[0] = true
	primes[1] = true

	for i := 2; i <= n; i++ {
		if !primes[i] {
			for j := i * 2; j <= n; j += i {
				primes[j] = true
			}
		}
	}

	for i := 2; i < len(primes)-1; i++ {
		if !primes[i] {
			ret++
		}
	}

	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
