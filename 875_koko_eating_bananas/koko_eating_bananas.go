package koko_eating_bananas

// https://leetcode.com/problems/koko-eating-bananas

// level: 2
// time: O(n*log(n)) 48ms 55.77%
// space: O(1) 6.3M 36.54%

// leetcode submit region begin(Prohibit modification and deletion)
func minEatingSpeed(piles []int, H int) int {
	min, max := 1, piles[0]
	// O(n)
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	// O(log(n))
	for min <= max {
		mid := min + (max-min)>>1
		// O(n)
		if getEatingCount(piles, mid) > H {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return min
}

func getEatingCount(piles []int, K int) int {
	ret := 0
	for _, pile := range piles {
		ret += (pile-1)/K + 1
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
