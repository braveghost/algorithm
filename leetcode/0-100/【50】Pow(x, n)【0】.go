package __100

//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
//
//
//
// 示例 1：
//
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
//
//
// 示例 2：
//
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//
//
// 示例 3：
//
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2-2 = 1/22 = 1/4 = 0.25
//
//
//
//
// 提示：
//
//
// -100.0 < x < 100.0
// -231 <= n <= 231-1
// -104 <= xn <= 104
//
// Related Topics 数学 二分查找
// 👍 571 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func MyPow(x float64, n int) float64 {
	var f func(nn int) float64
	f = func(nn int) float64 {
		if nn == 0 {
			return 1
		}
		res := f(nn / 2)
		res = res * res
		if nn&1 != 0 {
			res *= x
		}
		return res
	}
	if n >= 0 {
		return f(n)
	}
	return 1 / f(-n)
}
func myPow_2(x float64, n int) float64 {
	var f func(nn int) float64
	f = func(nn int) float64 {
		switch nn {
		case 0: return 1
		case 1: return x
		case -1: return 1 / x
		}
		res1 := f(nn / 2)
		res2 := f(nn % 2)
		return res1 * res1 * res2
	}
	return f(n)
}

func myPow(x float64, n int) float64 {
	var f func(nn int) float64
	f = func(nn int) float64 {
		res := 1.0
		y := x
		for nn > 0 {

			if nn&1 == 1 {
				res *= y
			}
			y *= y
			nn >>= 1
		}
		return res

	}
	if n >= 0 {
		return f(n)
	}
	return 1.0 / f(-n)
}

//leetcode submit region end(Prohibit modification and deletion)
