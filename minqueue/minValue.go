package minqueue

type MinQueue struct {
	q1 []int
	min []int
}

func Constructor() MinQueue {
	return MinQueue{
		make([]int,0),
		make([]int,0),
	}
}

func (this *MinQueue) Min_value() int {
	if len(this.min) == 0{
		return -1
	}
	return this.min[0]
}

func (this *MinQueue) Push_back(value int)  {
	this.q1 = append(this.q1,value)
	for len(this.min) != 0 && value < this.min[len(this.min)-1]{
		this.min = this.min[:len(this.min)-1]
	}
	this.min = append(this.min,value)
}

func (this *MinQueue) Pop_front() int {
	n := -1
	if len(this.q1) != 0{
		n = this.q1[0]
		this.q1 = this.q1[1:]
		if this.min[0] == n{
			this.min = this.min[1:]
		}
	}
	return n
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */