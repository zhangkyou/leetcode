package List

import "fmt"

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head

	count := 1
	for {
		if p == nil {
			break
		}
		count += 1
		p = p.Next
	}
	fmt.Println(count)

	middlePos := count/2
	p = head
	fmt.Println(middlePos, p.Val)
	for {
		if middlePos == 0 {
			break
		}
		p = p.Next
		middlePos -= 1
	}
	return p
}
