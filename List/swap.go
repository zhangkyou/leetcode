package List

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{0, head}

	p := dummy
	q := p.Next

	for {
		if q == nil || q.Next == nil {
			break
		}
		p.Next = q.Next
		q.Next = p.Next.Next
		p.Next.Next = q
		p = q
		q = q.Next
	}

	return dummy.Next
}