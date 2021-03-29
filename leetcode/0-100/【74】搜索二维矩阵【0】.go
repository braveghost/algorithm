package __100

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
// 示例 1：
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
// 示例 2：
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
// 提示：
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -104 <= matrix[i][j], target <= 104
//
// Related Topics 数组 二分查找
// 👍 340 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix) - 1
    if m < 0 {
        return false
    }
    n := len(matrix[0]) - 1
    l, r := 0, m
    for l <= r {
        mid := (l + r) >> 1
        if matrix[mid][0] <= target && target <= matrix[mid][n] {
            l, r = 0, n
            for l <= r {
                mid2 := (l + r) >> 1
                if matrix[mid][mid2] == target {
                    return true
                } else if matrix[mid][mid2] > target {
                    r = mid2 - 1
                } else {
                    l = mid2 + 1
                }
            }
            return false
        } else if matrix[mid][0] > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return false

}

//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix_other(matrix [][]int, target int) bool {
    // 二维数组坐标模拟成一维、、降维
    return false

}

//leetcode submit region end(Prohibit modification and deletion)
