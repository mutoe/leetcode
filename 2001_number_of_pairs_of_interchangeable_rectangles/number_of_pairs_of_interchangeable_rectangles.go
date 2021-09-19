package number_of_pairs_of_interchangeable_rectangles

// https://leetcode.com/problems/number-of-pairs-of-interchangeable-rectangles

// level: 2
// time: O(n) 280ms 98.28%
// space: O(n) 20.8M 86.21%

//leetcode submit region begin(Prohibit modification and deletion)
func interchangeableRectangles(rectangles [][]int) int64 {
	ratioMap := make(map[float64]int)
	for _, rectangle := range rectangles {
		ratio := float64(rectangle[0]) / float64(rectangle[1])
		ratioMap[ratio]++
	}

	var ret int64
	for _, n := range ratioMap {
		ret += int64(n * (n - 1) / 2)
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
