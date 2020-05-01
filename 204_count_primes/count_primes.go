package count_primes

// https://leetcode.com/problems/count-primes

// Eratosthenes sieve
// time: O(n*log(n)) 8ms 81.42%
// space: O(n) 4.9M

// leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) (ret int) {
	if n < 3 {
		return 0
	}
	primes := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		if !primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = true
			}
			if i != n {
				ret++
			}
		}
	}

	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
