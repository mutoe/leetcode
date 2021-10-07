package reformat_date

import "strings"

// https://leetcode.com/problems/reformat-date

// level: 1
// time: O(1) 0ms 100%
// space: O(1) 2.3M 16%

//leetcode submit region begin(Prohibit modification and deletion)

func reformatDate(date string) string {
	DAY := map[string]string{
		"1st":  "01",
		"2nd":  "02",
		"3rd":  "03",
		"4th":  "04",
		"5th":  "05",
		"6th":  "06",
		"7th":  "07",
		"8th":  "08",
		"9th":  "09",
		"10th": "10",
		"11th": "11",
		"12th": "12",
		"13th": "13",
		"14th": "14",
		"15th": "15",
		"16th": "16",
		"17th": "17",
		"18th": "18",
		"19th": "19",
		"20th": "20",
		"21st": "21",
		"22nd": "22",
		"23rd": "23",
		"24th": "24",
		"25th": "25",
		"26th": "26",
		"27th": "27",
		"28th": "28",
		"29th": "29",
		"30th": "30",
		"31st": "31",
	}
	MONTH := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	dmy := strings.Split(date, " ")
	return strings.Join([]string{dmy[2], MONTH[dmy[1]], DAY[dmy[0]]}, "-")
}

//leetcode submit region end(Prohibit modification and deletion)
