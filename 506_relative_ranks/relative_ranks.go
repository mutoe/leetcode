package relative_ranks

import "strconv"

// https://leetcode.com/problems/relative-ranks

// level: 1
// time: O(n*log(n)) 12ms 93.94%
// space: O(log(n)) 6.7M 24.24%

//leetcode submit region begin(Prohibit modification and deletion)

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	i, j, x := l-1, r+1, arr[l]
	for i < j {
		for ok := true; ok; ok = arr[i] < x {
			i++
		}
		for ok := true; ok; ok = arr[j] > x {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	quickSort(arr, l, j)
	quickSort(arr, j+1, r)
}

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
	QuickSort(sortedList)

	result := make([]string, len(score))
	for i := 0; i < len(score); i++ {
		score := sortedList[len(score)-i-1]
		result[scoreIndexMap[score]] = convertRandToResult(i)
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)
