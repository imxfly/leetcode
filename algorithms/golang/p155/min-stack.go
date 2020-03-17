package p155

// URL: https://leetcode-cn.com/problems/min-stack/

// MinStack stack
type MinStack struct {
	min  int   // 最小元素
	data []int // 栈存储空间
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{}
}

// Push push element to stack
func (stack *MinStack) Push(x int) {
	if len(stack.data) == 0 {
		stack.min = x
	} else if stack.min > x {
		stack.min = x
	}
	stack.data = append(stack.data, x)
}

// Pop pop element from stack
func (stack *MinStack) Pop() {
	len := len(stack.data)
	if len == 0 {
		return
	}

	// 只有当出栈的元素是当前最小元素时，去重新选择最小元素
	if stack.min == stack.data[len-1] {
		min := stack.data[0]
		for i := 0; i < len-1; i++ {
			if min > stack.data[i] {
				min = stack.data[i]
			}
		}
		stack.min = min
	}

	stack.data = stack.data[:len-1]
}

// Top get the top element of stack
func (stack *MinStack) Top() int {
	return stack.data[len(stack.data)-1]
}

// GetMin get the minimum value of stack
func (stack *MinStack) GetMin() int {
	return stack.min
}
