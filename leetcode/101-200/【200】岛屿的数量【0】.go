package _01_200

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
// Related Topics 深度优先搜索 广度优先搜索 并查集
// 👍 968 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func NumIslands_DFS(grid [][]byte) int {
	var res int
	yl := len(grid)
	xl := len(grid[0])
	var dfs func(x,y int)
	dfs = func(x, y int) {
		grid[y][x] = '0'
		for _, l := range [][2]int{{ x, y-1},{ x,y+1},{ x-1,y},{ x+1,y}}{
			x, y = l[0], l[1]
			if (x >=0 && x<xl )&& (y >=0 && y<yl ) && grid[y][x] == '1'{
				dfs(x, y)
			}
		}

	}


	for y := range grid{
		for x := range grid[y]{
			if grid[y][x] == '1'{
				res += 1
				dfs(x,y)
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region begin(Prohibit modification and deletion)
func NumIslands_BFS(grid [][]byte) int {
	var res int
	yl := len(grid)
	if yl == 0{
		return 0
	}
	xl := len(grid[0])

	for y := range grid{
		for x := range grid[y]{
			if grid[y][x] == '1'{
				res += 1
				grid[y][x] = '0'
				var queue = [][2]int{{x, y}}
				for len(queue)!=0{
					node := queue[0]
					queue = queue[1:]
					col, row := node[0], node[1]
					for _, l := range [][2]int{{ col, row-1},{ col,row+1},{ col-1,row},{ col+1,row}}{
						xx, yy := l[0], l[1]
						if (xx >=0 && xx<xl )&& (yy >=0 && yy<yl ) && grid[yy][xx] == '1'{
							grid[yy][xx] = '0'
							queue = append(queue,[2]int{xx, yy})
						}
					}

				}
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
