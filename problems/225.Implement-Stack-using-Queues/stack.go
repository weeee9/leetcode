package leetcode

type MyStack struct {
	enqueue []int
	dequeue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.enqueue = append(this.enqueue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	length := len(this.enqueue)
	for i := 0; i < length-1; i++ {
		this.dequeue = append(this.dequeue, this.enqueue[0])
		this.enqueue = this.enqueue[1:]
	}
	top := this.enqueue[0]
	this.enqueue = this.dequeue
	this.dequeue = []int{}
	return top
}

/** Get the top element. */
func (this *MyStack) Top() int {
	top := this.Pop()
	this.enqueue = append(this.enqueue, top)
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.enqueue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
