package LRU

import (
	"fmt"
	"sync"
)

type ListNode struct {
	Key int
	Val int
	Pre *ListNode
	Next *ListNode
}

var head = &ListNode{0, 0, nil, nil}
var tail = &ListNode{0, 0, nil, nil}

func ListConstructor() {
	head.Next = tail
	tail.Pre = head
}

func addNode(node *ListNode) {
	node.Next = head.Next
	head.Next.Pre = node
	node.Pre = head
	head.Next = node
}

func removeNode(node *ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

type LRUCache struct {
	cap int
	m map[int]*ListNode
}

func Constructor(capacity int) LRUCache {
	ListConstructor()
	return LRUCache{capacity, make(map[int]*ListNode, capacity)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		removeNode(node)
		addNode(node)
		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) error {
	sp := sync.Pool{New: func() interface{} {
		return &ListNode{key, value, nil, nil}
	}}

	if node, ok := this.m[key]; ok {
		node.Val = value
		removeNode(node)
		addNode(node)
		this.m[key] = node
	} else {
		cur := sp.Get().(*ListNode)
		if len(this.m) < this.cap {
			addNode(cur)
			this.m[key] = cur
		} else {
			//cur := &ListNode{key, value, nil, nil}
			delete(this.m, tail.Pre.Key)
			removeNode(tail.Pre)
			addNode(cur)
			this.m[key] = cur
		}
		sp.Put(cur)
	}

	return nil
}

func Print() {
	p := head
	for {
		if p == nil {
			break
		}

		fmt.Println(p.Key, p.Val)
		p = p.Next
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */