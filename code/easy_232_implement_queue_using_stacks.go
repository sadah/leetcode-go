/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

package code

// @lc code=start
type MyQueue struct {
	queue []int
}

func Constructor2() MyQueue {
	return MyQueue{[]int{}}
}

//lint:ignore ST1006 //
func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

//lint:ignore ST1006 //
func (this *MyQueue) Pop() int {
	ret := this.queue[0]
	this.queue = this.queue[1:]
	return ret
}

//lint:ignore ST1006 //
func (this *MyQueue) Peek() int {
	return this.queue[0]
}

//lint:ignore ST1006 //
func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
