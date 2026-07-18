type MinListNode struct {
	Val  int
	Min  int
	Back *MinListNode
}

type MinStack struct {
	Head *MinListNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {

	min := val

	if this.Head != nil && this.Head.Min < val {
		min = this.Head.Min
	}

	this.Head = &MinListNode{
		Val:  val,
		Min:  min,
		Back: this.Head,
	}
}

func (this *MinStack) Pop() {
	if this.Head != nil {
		this.Head = this.Head.Back
	}
}

func (this *MinStack) Top() int {
	return this.Head.Val
}

func (this *MinStack) GetMin() int {
	return this.Head.Min
}
