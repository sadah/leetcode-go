/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

package code

// @lc code=start
type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{[]int{}}
}

//lint:ignore ST1006 //
func (this *MyStack) Push(x int) {
	this.queue = append([]int{x}, this.queue...)
}

//lint:ignore ST1006 //
func (this *MyStack) Pop() int {
	q := this.queue[0]
	this.queue = this.queue[1:]
	return q
}

//lint:ignore ST1006 //
func (this *MyStack) Top() int {
	return this.queue[0]
}

//lint:ignore ST1006 //
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
