package List

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	q := head.Next
	p.Next = nil
	if q.Next == nil {
		q.Next = p
		return q
	}

	z := q.Next

	for {
		if z == nil {
			break
		}

		q.Next = p
		p = q
		q = z
		z = z.Next
	}

	q.Next = p

	return q
}

func reverseListRecurve(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := reverseListRecurve(head.Next)
	head.Next.Next = head
	head.Next = nil

	return cur
}

func ReversePart(head *ListNode, m, n int) *ListNode {
	dummy := &ListNode{0, head}

	p := dummy
	pCount := 0
	q := dummy
	qCount := 0

	for {
		if qCount == n + 1 || q == nil {
			break
		}

		if pCount < m - 1 {
			p = p.Next
			pCount += 1
		}

		if qCount < n + 1 {
			q = q.Next
			qCount += 1
		}
	}

	cur := recurveByPart(p.Next, p, q)
	p.Next = cur

	return dummy.Next
}

func recurveByPart(head *ListNode, before *ListNode, after *ListNode) *ListNode {
	if head == after || head.Next == after {
		return head
	}

	cur := recurveByPart(head.Next, before, after)
	head.Next.Next = head
	head.Next = after

	return cur
}

