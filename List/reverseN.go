package List

var after *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		after = head.Next
		return head
	}

	last := reverseN(head.Next, n - 1)
	head.Next.Next = head
	head.Next = after

	return last
}

func reverseMN(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}

	head.Next = reverseMN(head.Next, m - 1, n - 1)
	return head
}