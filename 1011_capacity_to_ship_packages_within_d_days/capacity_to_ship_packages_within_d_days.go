package capacity_to_ship_packages_within_d_days

import "math"

// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days

// level: 2
// time: O(n*log(n)) 48ms 29.27%
// space: O(1) 6.6M 80.49%

// TODO: Need to review

// leetcode submit region begin(Prohibit modification and deletion)
func shipWithinDays(weights []int, D int) int {
	l, r := 1, 0
	// O(n)
	for _, weight := range weights {
		r += weight
	}
	if D == 1 {
		return r
	}
	// O(log(n))
	// l == r ?
	for l <= r {
		mid := l + (r-l)>>1
		// O(n)
		if getDays(weights, mid) > D {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func getDays(weights []int, weight int) int {
	ret := 1
	rest := weight
	for _, w := range weights {
		if w > weight {
			return math.MaxInt32
		}
		if rest >= w {
			rest -= w
		} else {
			ret++
			rest = weight - w
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
