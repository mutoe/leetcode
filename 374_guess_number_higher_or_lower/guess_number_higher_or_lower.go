package guess_number_higher_or_lower

// https://leetcode.com/problems/guess-number-higher-or-lower

// level: 1
// time: O(log(n)) 0ms 100%
// space: O(1) 1.9M 100%

var pick int

func guess(num int) int {
	if num == pick {
		return 0
	} else if num > pick {
		return -1
	} else {
		return 1
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)>>1
		diff := guess(mid)
		if diff == 0 {
			return mid
		} else if diff > 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 0
}

// leetcode submit region end(Prohibit modification and deletion)
