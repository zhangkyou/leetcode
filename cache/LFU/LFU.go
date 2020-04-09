package LFU

import (
	"fmt"
)

type ListNode struct {
	Key int
	Val int
	Freq int
	Pre *ListNode
	Next *ListNode
}

func remove(node *ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Next = nil
	node.Pre = nil
}

type DoubleLink struct {
	Head *ListNode
	Tail *ListNode
}

func DLinkConstructor() *DoubleLink {
	head := &ListNode{-1,-1,0,nil,nil}
	tail := &ListNode{-1,-1,0, nil, nil}
	head.Next = tail
	tail.Pre = head
	return &DoubleLink{Head:head, Tail:tail}
}

func (d *DoubleLink) addNode2Head(node *ListNode) {
	node.Next = d.Head.Next
	node.Pre = d.Head
	d.Head.Next.Pre = node
	d.Head.Next = node
}

type LFUCache struct {
	KeyMap map[int]*ListNode
	FreqMap map[int]*DoubleLink
	Cap int
	MinFrequency int
}

func Constructor(capacity int) LFUCache {
	keyMap := make(map[int]*ListNode)
	freqMap := make(map[int]*DoubleLink)
	return LFUCache{KeyMap:keyMap, FreqMap:freqMap, Cap:capacity, MinFrequency:-1}
}

func (this *LFUCache) Get(key int) int {
	if this.Cap == 0 {
		return -1
	}

	if node, ok := this.KeyMap[key];ok {
		this.updateMinFrequency(node)
		remove(node)
		node.Freq++

		this.addFreqMap(node)
		return node.Val
	} else {
		return -1
	}
}

func (this *LFUCache) updateMinFrequency(node *ListNode) {
	if this.MinFrequency == -1 { //init min frequency
		this.MinFrequency = node.Freq
		return
	}

	if this.MinFrequency < node.Freq {
		return
	} else if this.MinFrequency == node.Freq {
		if node.Pre.Key == -1 && node.Next.Key == -1 { //only this node
			this.MinFrequency = node.Freq + 1
			return
		}
	} else {
		this.MinFrequency = node.Freq
	}
	return
}

func (this *LFUCache) addFreqMap(node *ListNode) {
	//add to freq map
	if dlink, ok := this.FreqMap[node.Freq]; ok {
		dlink.addNode2Head(node)
	} else {
		dlink = DLinkConstructor()
		dlink.addNode2Head(node)
		this.FreqMap[node.Freq] = dlink
	}
}

func (this *LFUCache) Put(key int, value int)  {
	if this.Cap == 0 {
		return
	}

	if node, ok := this.KeyMap[key]; ok {
		this.updateMinFrequency(node)
		remove(node)
		node.Freq++
		node.Val = value
		this.addFreqMap(node)
		return
	}

	if len(this.KeyMap) == this.Cap {
		var needDeleteNode *ListNode
		needDeleteNode = this.FreqMap[this.MinFrequency].Tail.Pre
		//this.updateMinFrequency(needDeleteNode)
		remove(needDeleteNode)
		delete(this.KeyMap, needDeleteNode.Key)
	}

	node := &ListNode{key, value, 1, nil, nil}
	this.KeyMap[key] = node
	if dlink, ok := this.FreqMap[node.Freq]; ok {
		dlink.addNode2Head(node)
	} else {
		dlink = DLinkConstructor()
		dlink.addNode2Head(node)
		this.FreqMap[node.Freq] = dlink
	}
	this.MinFrequency = node.Freq
}

func (this *LFUCache) ShowKeyMap() {
	for i, node := range this.KeyMap {
		fmt.Println(i, node.Key, node.Val, node.Freq)
	}
}

func (this *LFUCache) ShowFreqMap() {
	fmt.Printf("current min frequency is : %d\n", this.MinFrequency)
	for i, dlink := range this.FreqMap {
		fmt.Printf("freq: %d\n", i)
		if dlink != nil {
			cur := dlink.Head
			for {
				if cur == nil {
					break
				}
				fmt.Println(cur.Key, cur.Val, cur.Freq)
				cur = cur.Next
			}
		}
	}
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */