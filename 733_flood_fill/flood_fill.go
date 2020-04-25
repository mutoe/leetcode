package flood_fill

// https://leetcode.com/problems/flood-fill

// DFS Solution
// time: O(n*m) 4ms 99%
// space: O(min(n,m)) 4.2M

// leetcode submit region begin(Prohibit modification and deletion)

var g [][]int
var oc int
var nc int

func dfs(x, y int) {
	if x < 0 || y < 0 || x >= len(g) || y >= len(g[0]) || g[x][y] != oc {
		return
	}
	g[x][y] = nc
	dfs(x-1, y)
	dfs(x+1, y)
	dfs(x, y-1)
	dfs(x, y+1)
}

func floodFill(image [][]int, sr int, sc int, newColor int) (result [][]int) {
	g = image
	nc = newColor
	oc = image[sr][sc]
	if nc != oc {
		dfs(sr, sc)
	}
	return g
}

// leetcode submit region end(Prohibit modification and deletion)
