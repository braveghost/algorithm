package __100

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//
//
// 示例：
//
// 给你这个链表：1->2->3->4->5
//
// 当 k = 2 时，应当返回: 2->1->4->3->5
//
// 当 k = 3 时，应当返回: 3->2->1->4->5
//
//
//
// 说明：
//
//
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
// Related Topics 链表
// 👍 837 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	var pre, sub, lastSub *ListNode // 除了第一段每段的头指针   每一段的尾指针
	cur, tmp, end := head, head, head
	for cur != nil {
		lastSub = sub
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil && i < k-1 {
				if lastSub != nil {
					lastSub.Next = cur
				}
				return tmp
			}
		}
		sub = cur // 每段头指针 翻转就成了尾指针
		for i := 0; i < k; i++ {
			cur, cur.Next, pre = cur.Next, pre, cur
		}
		if tmp == head {
			tmp = pre // 入口为第一段的现头指针
		}
		if lastSub != nil {
			lastSub.Next = pre // 每段的尾指针 指向下一段的头指针
		}
		pre = nil
	}
	return tmp
	//
	//var pre *ListNode // next
	//cur := head
	//for cur != nil {
	//	cur, cur.Next, pre = cur.Next, pre, cur
	//}
	//return pre
}

//leetcode submit region end(Prohibit modification and deletion)

//ReverseKGroup2 官方题解优化版
func ReverseKGroup2(head *ListNode, k int) *ListNode {
	tmp := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := tmp
	for head != nil{
		tail := pre // 从开始结点(非子链表的头结点)寻找尾结点， k=2  tail=tail.Next  tail=tail.Next == head.Next
		for i :=0;i<k;i++{
			tail = tail.Next
			if tail == nil{
				return tmp.Next
			}
		}
		head, tail = reverse(head, tail)
		//// 当前链表的入口节点从头结点变为尾结点，因为翻转了，head是翻转前的尾
		//pre.Next = head
		//// 当前节点的尾节点成为下一段子链表的父节点
		//pre = tail
		//// 下一段链表的头结点
		//head = tail.Next
		pre.Next, pre, head = head, tail, tail.Next
	}
	return tmp.Next
}

/**
 * @Description:
 * @param head: 当前链表的头结点
 * @param tail: 当前链表的尾结点
 * @return *ListNode: 当前链表的头结点(翻转之前的尾结点)
 * @return *ListNode: 当前链表的尾结点(翻转之前的头结点)
 */
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next // 新子链表的开头, 当前链表 head 的头变尾之后指向这个开头   k=2 pre = head.Next.Next
	cur := head
	for pre!= tail{  // k=2 tail=head.Next
		cur, cur.Next, pre = cur.Next, pre, cur   // k=2 pre=head   pre=head.Next
	}
	return tail, head
}



//ReverseKGroup_Stack 堆栈
func ReverseKGroup_Stack(head *ListNode, k int) *ListNode {
	//tmp := &ListNode{
	//	Val:  -1,
	//	Next: head,
	//}
	//var stack []*ListNode
    // pre := tmp
	// todo  k=2
	// 头结点
	// 入栈2个  1=pre.Next   2=pre.Next.Next
	// 记录接下来下面子链表的头节点 x 、、、
	// 依次出栈并将next指向下一个出栈的元素、最后将
	// pre.Next = 2   pre = pre.Next   pre=1  pre.Next = x
	//
	// 入栈2个  1=pre.Next   2=pre.Next.Next
	// 记录接下来下面子链表的头节点 x 、、、
	// 依次出栈并将next指向下一个出栈的元素、最后将
	// pre.Next = 2   pre = pre.Next   pre=1   pre.Next = x
	//
	// 中间记得判断剩余子链表长度不够的情况

	//pre := tmp
	//for pre.Next != nil{
	//	for i:=0;i<k;i++{
	//		t := pre.Next
	//		if t == nil{
	//			return pre
	//		}
	//		stack = append(stack, t)
	//	}
	//	x:= k
	//	for x >0{
	//		pre.Next = stack[x-1]
	//		pre = pre.Next
	//		x --
	//	}
	//
	//}
	return head
}

// todo 递归  从最后一部分子链表处理、翻转后返回结果链表、然后接上前一个链表的翻转结果