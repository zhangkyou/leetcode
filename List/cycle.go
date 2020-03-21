package List

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	p := head
	q := p.Next
	for {
		if p == nil || q == nil {
			break
		}

		if q.Next == nil {
			break
		}

		if p == q {
			return true
		}

		p = p.Next
		q = q.Next.Next
	}

	return false
}
