package __100

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
// Related Topics 递归 链表
// 👍 778 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	start := head.Next
	var last, next *ListNode
	for head != nil && head.Next != nil {
		head.Next.Next, head.Next, next = head, head.Next.Next, head.Next
		if last != nil {
			last.Next = next
		}
		last, head = head, head.Next
	}
	return start
}


//leetcode submit region end(Prohibit modification and deletion)

func SwapPairs_Recursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	start := head.Next
	head.Next.Next, head.Next = head, SwapPairs_Recursive(head.Next.Next)
	return start
}


//leetcode submit region end(Prohibit modification and deletion)
