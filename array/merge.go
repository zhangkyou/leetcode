package array

type ListNode struct {
	Val int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2

	r := &ListNode{0, nil}
	z := r

	for {
		if p == nil || q == nil {
			break
		}
		z.Next = &ListNode{0, nil}
		z = z.Next

		if p.Val <= q.Val {
			z.Val = p.Val
			p = p.Next
		} else {
			z.Val = q.Val
			q = q.Next
		}
	}

	if p != nil {
		for {
			if p == nil {
				break
			}
			z.Next = &ListNode{0, nil}
			z = z.Next

			z.Val = p.Val
			p = p.Next
		}
	}

	if q != nil {
		for {
			if q == nil {
				break
			}
			z.Next = &ListNode{0, nil}
			z = z.Next

			z.Val = q.Val
			q = q.Next
		}
	}

	return r.Next
}

func MergeTwoListsOptimiz(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = MergeTwoListsOptimiz(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoListsOptimiz(l1, l2.Next)
		return l2
	}
}

