package p225

import (
	"container/list"
)

// MyStack my stack
type MyStack struct {
	q *list.List
}

// Constructor Initialize your data structure here.
func Constructor() MyStack {
	stack := MyStack{list.New()}

	return stack
}

// Push element x onto stack.
func (myStack *MyStack) Push(x int) {
	myStack.q.PushBack(x)
}

// Pop Removes the element on top of the stack and returns that element.
func (myStack *MyStack) Pop() int {
	var top int
	q2 := list.New()
	for e := myStack.q.Front(); e != nil; e = e.Next() {
		if e.Next() == nil {
			top, _ = e.Value.(int)
		} else {
			q2.PushBack(e.Value)
		}
	}

	myStack.q = q2

	return top
}

// Top Get the top element.
func (myStack *MyStack) Top() int {
	var top int
	for e := myStack.q.Front(); e != nil; e = e.Next() {
		if e.Next() == nil {
			top, _ = e.Value.(int)
		}
	}

	return top
}

// Empty Returns whether the stack is empty.
func (myStack *MyStack) Empty() bool {
	return myStack.q.Len() == 0
}
