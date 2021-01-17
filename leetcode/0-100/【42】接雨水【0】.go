package __100

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 0 <= n <= 3 * 104
// 0 <= height[i] <= 105
//
// Related Topics 栈 数组 双指针 动态规划
// 👍 1943 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 每个元素的储存量等于两边最大高度的较小值减去当前高度的值
func Trap_Violence(height []int) int {
	var res int
	l := len(height)
	for i:= 1;i< l-1;i++{
		var left, right int
		for j:= i;j>=0;j--{
			if  height[j] > left{
				left = height[j]
			}
		}
		for j:= i;j<l;j++{
			if  height[j] > right{
				right = height[j]
			}
		}
		if left > right{
			res += right - height[i]
		}else {
			res += left - height[i]

		}
	}
	return res
	//
	//public int trap(int[] height) {
	//	int ans = 0;
	//	int size = height.length;
	//	for (int i = 1; i < size - 1; i++) {
	//		int max_left = 0, max_right = 0;
	//		for (int j = i; j >= 0; j--) { //Search the left part for max bar size
	//			max_left = Math.max(max_left, height[j]);
	//		}
	//		for (int j = i; j < size; j++) { //Search the right part for max bar size
	//			max_right = Math.max(max_right, height[j]);
	//		}
	//		ans += Math.min(max_left, max_right) - height[i];
	//	}
	//	return ans;
	//}



	//int trap(vector<int>& height)
	//{
	//	int ans = 0;
	//	int size = height.size();
	//	for (int i = 1; i < size - 1; i++) {
	//	int max_left = 0, max_right = 0;
	//	for (int j = i; j >= 0; j--) { //Search the left part for max bar size
	//		max_left = max(max_left, height[j]);
	//	}
	//	for (int j = i; j < size; j++) { //Search the right part for max bar size
	//		max_right = max(max_right, height[j]);
	//	}
	//	ans += min(max_left, max_right) - height[i];
	//}
	//	return ans;
	//}
	//
	//作者：LeetCode
	//链接：https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/
	//来源：力扣（LeetCode）
	//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}
//leetcode submit region end(Prohibit modification and deletion)
