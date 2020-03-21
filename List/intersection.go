package List


func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]int)

	p := headA
	for {
		if p == nil {
			break
		}

		m[p] = p.Val
		p = p.Next
	}

	q := headB
	for {
		if q == nil {
			break
		}

		if _, ok := m[q]; ok {
			return q
		}
		q = q.Next
	}

	return nil
}

/**
optimize a + b == b + a
 */
func GetIntersectionNodeOptimize(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p := headA
	q := headB

	for {
		if p == q {
			break
		}

		if p == nil {
			p = headB
		}
		p = p.Next

		if q == nil {
			q = headA
		}
		q = q.Next
	}

	return p
}

