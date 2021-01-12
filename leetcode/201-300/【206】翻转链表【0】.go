package _01_300

//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
// 👍 1430 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


type ListNode struct {
    Val int
    Next *ListNode
}
//func ReverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	var last *ListNode
//	for head.Next != nil{
//		// 记录 下一个 节点
//		next :=head.Next
//		// 将当前节点设置为last节点
//		head.Next = last
//		last = head
//		head = next
//
//	}
//	head.Next = last
//	return head
//}
//leetcode submit region end(Prohibit modification and deletion)
func ReverseList(head *ListNode) *ListNode {

	var pre *ListNode // next
	cur := head
	for cur != nil {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
	return pre
}
func ReverseList_Recursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	cur := ReverseList_Recursive(head.Next)

	head.Next.Next = head
	head.Next = nil

	return cur
}

