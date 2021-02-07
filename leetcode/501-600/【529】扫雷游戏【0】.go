package _01_600

//让我们一起来玩扫雷游戏！
//
// 给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B' 代表没有相邻（上，下，左，右，和所有4个对角线）
//地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。
//
// 现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：
//
//
// 如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
// 如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的未挖出方块都应该被递归地揭露。
// 如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
// 如果在此次点击中，若无更多方块可被揭露，则返回面板。
//
//
//
//
// 示例 1：
//
// 输入:
//
//[['E', 'E', 'E', 'E', 'E'],
// ['E', 'E', 'M', 'E', 'E'],
// ['E', 'E', 'E', 'E', 'E'],
// ['E', 'E', 'E', 'E', 'E']]
//
//Click : [3,0]
//
//输出:
//
//[['B', '1', 'E', '1', 'B'],
// ['B', '1', 'M', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
//
//解释:
//
//
//
// 示例 2：
//
// 输入:
//
//[['B', '1', 'E', '1', 'B'],
// ['B', '1', 'M', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
//
//Click : [1,2]
//
//输出:
//
//[['B', '1', 'E', '1', 'B'],
// ['B', '1', 'X', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
//
//解释:
//
//
//
//
//
// 注意：
//
//
// 输入矩阵的宽和高的范围为 [1,50]。
// 点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。
// 输入面板不会是游戏结束的状态（即有地雷已被挖出）。
// 简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。
//
// Related Topics 深度优先搜索 广度优先搜索
// 👍 213 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func UpdateBoard_DFS(board [][]byte, click []int) [][]byte {
	ly := len(board)
	if ly == 0 || len(click) != 2 {
		return nil
	}
	lx := len(board[0])
	y, x := click[0], click[1]
	if y < ly && y >= 0 && x < lx && x >= 0 && board[y][x] == 'M' {
		board[y][x] = 'X'
		return board
	}
	var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
	var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		var count byte
		for i, v := range dirX {
			tx, ty := x+v, y+dirY[i]
			if tx < lx && tx >= 0 && ty < ly && ty >= 0 && board[ty][tx] == 'M' {
				count++
			}
		}
		if count > 0 {
			board[y][x] = count + '0'
		} else {
			board[y][x] = 'B'
			for i, v := range dirX {
				tx, ty := x+v, y+dirY[i]
				if tx < lx && tx >= 0 && ty < ly && ty >= 0 && board[ty][tx] == 'E' {
					dfs(tx, ty)
				}
			}

		}
	}
	dfs(x, y)
	return board

}
func UpdateBoard_BFS(board [][]byte, click []int) [][]byte {
	ly := len(board)
	if ly == 0 || len(click) != 2 {
		return nil
	}
	lx := len(board[0])
	y, x := click[0], click[1]
	if y < ly && y >= 0 && x < lx && x >= 0 && board[y][x] == 'M' {
		board[y][x] = 'X'
		return board
	}
	var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
	var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}


	vis := make([][]bool, len(board))
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, len(board[0]))
	}
	queue := [][2]int{{x, y}}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		x, y = node[0], node[1]
		var count byte
		for i, v := range dirX {
			tx, ty := x+v, y+dirY[i]
			if tx < lx && tx >= 0 && ty < ly && ty >= 0 && board[ty][tx] == 'M' {
				count++
			}
		}
		if count > 0 {
			board[y][x] = count + '0'
		} else {
			board[y][x] = 'B'
			for i, v := range dirX {
				tx, ty := x+v, y+dirY[i]
				if tx < lx && tx >= 0 && ty < ly && ty >= 0 && board[ty][tx] == 'E' && !vis[ty][tx] {

						queue = append(queue, [2]int{tx, ty})
						vis[ty][tx] = true
				}
			}
		}
	}
	return board
}

//leetcode submit region end(Prohibit modification and deletion)
