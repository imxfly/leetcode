package p232

import "container/list"

// URL: https://leetcode-cn.com/problems/implement-queue-using-stacks/

// MyQueue my queue
type MyQueue struct {
	l *list.List
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{list.New()}
}

// Push Push element x to the back of queue.
func (myQueue *MyQueue) Push(x int) {
	myQueue.l.PushBack(x)
}

// Pop Removes the element from in front of queue and returns that element.
func (myQueue *MyQueue) Pop() int {
	var top int
	l := list.New()
	for e := myQueue.l.Back(); e != nil; e = e.Prev() {
		if e.Prev() == nil {
			top, _ = e.Value.(int)
		} else {
			l.PushBack(e.Value)
		}
	}

	myQueue.l = myQueue.l.Init()
	for e := l.Back(); e != nil; e = e.Prev() {
		myQueue.l.PushBack(e.Value)
	}

	return top
}

// Peek Get the front element.
func (myQueue *MyQueue) Peek() int {
	var top int
	for e := myQueue.l.Back(); e != nil; e = e.Prev() {
		if e.Prev() == nil {
			top, _ = e.Value.(int)
		}
	}

	return top
}

// Empty Returns whether the queue is empty.
func (myQueue *MyQueue) Empty() bool {
	return myQueue.l.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
