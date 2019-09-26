package main

/*
请用栈实现一个队列，支持如下四种操作：

push(x) – 将元素x插到队尾；
pop() – 将队首的元素弹出，并返回该元素；
peek() – 返回队首元素；
empty() – 返回队列是否为空；
注意：

你只能使用栈的标准操作：push to top，peek/pop from top, size 和 is empty；
如果你选择的编程语言没有栈的标准库，你可以使用list或者deque等模拟栈的操作；
输入数据保证合法，例如，在队列为空时，不会进行pop或者peek等操作；
样例
MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
*/

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/* Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{stack1: make([]int, 0), stack2: make([]int, 2)}
}

/* Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
	return
}

/* Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for i := len(this.stack1) - 1; i > 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	a := this.stack1[0]
	this.stack1 = this.stack1[1:]
	return a
}

/* Get the front element. */
func (this *MyQueue) Peek() int {
	for i := len(this.stack1) - 1; i > 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	a := this.stack1[0]

	return a
}

/* Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}

func main() {

}
