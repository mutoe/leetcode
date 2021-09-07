package relative_ranks

import (
	"sort"
	"strconv"
)

// https://leetcode.com/problems/relative-ranks

// level: 1
// time: O(n*log(n)) 12ms 93.94%
// space: O(log(n)) 6.7M 24.24%

//leetcode submit region begin(Prohibit modification and deletion)

func convertRandToResult(rank int) string {
	switch rank + 1 {
	case 1:
		return "Gold Medal"
	case 2:
		return "Silver Medal"
	case 3:
		return "Bronze Medal"
	default:
		return strconv.Itoa(rank + 1)
	}
}

func findRelativeRanks(score []int) []string {
	scoreIndexMap := make(map[int]int)
	for i := 0; i < len(score); i++ {
		scoreIndexMap[score[i]] = i
	}
	sortedList := score[:]
	sort.Ints(sortedList)

	result := make([]string, len(score))
	for i := 0; i < len(score); i++ {
		score := sortedList[len(score)-i-1]
		result[scoreIndexMap[score]] = convertRandToResult(i)
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)
