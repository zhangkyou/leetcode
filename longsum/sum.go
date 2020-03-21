package longsum

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(a *ListNode, b *ListNode) *ListNode {
	c := ListNode{0, nil}

	p := a
	q := b
	r := &c

	var last = 0
	for{
		var tmp = 0
		if p == nil && q == nil && last == 0 {
			break
		}
		if p != nil {
			tmp += p.Val
		}

		if q != nil {
			tmp += q.Val
		}

		n := ListNode{0, nil}
		tmp += last
		if tmp >= 10 {
			last = 1
		} else {
			last = 0
		}
		n.Val = tmp%10
		r.Next = &n

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
		r = r.Next
	}

	return c.Next
}

func addTwoNumbersOptimiz(a *ListNode, b *ListNode) *ListNode {
	var last = 0
	p := a
	q := b
	for ; p.Next != nil; p = p.Next {
		tmp := p.Val + last
		if q != nil {
			tmp += q.Val
			q = q.Next
		}
		if tmp >= 10 {
			last = 1
		} else {
			last = 0
		}
		p.Val = tmp%10
	}

	var tmp = 0
	if q != nil {
		tmp = p.Val + q.Val + last
		p.Val = tmp%10
		if q.Next == nil {
			if tmp >= 10 {
				p.Next = &ListNode{1, nil}
			}
		} else {
			if tmp >= 10 {
				q.Next.Val += 1
			}
			p.Next = q.Next
		}
	} else {
		tmp = p.Val + last
		p.Val = tmp%10
		if tmp >= 10 {
			p.Next = &ListNode{1, nil}
		}
	}

	return a
}

func sum2(a *ListNode, b *ListNode) *ListNode {
	d := ListNode{0, nil}
	c := ListNode{0, &d}

	p := a
	q := b
	r := &c

	for{
		var tmp = 0
		if p == nil && q == nil {
			break
		}
		if p != nil {
			tmp += p.Val
		}

		if q != nil {
			tmp += q.Val
		}

		if tmp >= 10 {
			r.Val += 1
			r.Next.Val += tmp%10
		} else {
			r.Next.Val += tmp
		}
		n := ListNode{0, nil}
		r.Next.Next = &n

		p = p.Next
		q = q.Next
		r = r.Next
	}

	if c.Val == 0 {
		return c.Next
	} else {
		return &c
	}

}
