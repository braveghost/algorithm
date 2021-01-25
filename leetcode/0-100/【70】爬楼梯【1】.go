package __100

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 动态规划
// 👍 1408 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func ClimbStairs_Iter(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2 := 1, 2
	var f int
	for x := 3; x <= n; x++ {
		f1, f2, f = f2, f1+f2, f1+f2
	}
	return f
}

//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region begin(Prohibit modification and deletion)
var climbStairsMap = map[int]int{}

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var v1, v2 int
	if v, ok := climbStairsMap[n-1]; ok {
		v1 = v
	} else {
		v1 = ClimbStairs(n - 1)
		climbStairsMap[n-1] = v1
	}
	if v, ok := climbStairsMap[n-2]; ok {
		v2 = v
	} else {
		v2 = ClimbStairs(n - 2)
		climbStairsMap[n-2] = v2
	}
	return v1 + v2
}

//leetcode submit region end(Prohibit modification and deletion)

func ClimbStairs_Iter2(n int) int {
	if n <= 2 {
		return n
	}
	var (
		f = 0
		f1 = 1
		f2 = 2
	)
	for i := 3; i <= n; i++ {
		f1, f2, f = f2, f1+f2, f1+f2
	}
	return f
}
