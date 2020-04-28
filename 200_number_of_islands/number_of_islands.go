package number_of_islands

// https://leetcode.com/problems/number-of-islands

// DFS Solution
// time: O(n*m) 0ms 100%
// space: O(min(n,m)) 2.4M

// leetcode submit region begin(Prohibit modification and deletion)

var g [][]byte

func dfs(i, j int) {
	if i >= 0 && j >= 0 && i < len(g) && j < len(g[0]) && g[i][j] == '1' {
		g[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
}

func numIslands(grid [][]byte) (islandsCount int) {
	g = grid
	for i, row := range g {
		for j := range row {
			if g[i][j] == '1' {
				islandsCount++
				dfs(i, j)
			}
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
