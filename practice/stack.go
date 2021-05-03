package main

import "fmt"

type Stack []int

func (stack *Stack) Push(ele int) {
	*stack = append([]int{ele}, *stack...)
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Pop() int {
	old := *stack
	popped := old[0]
	if stack.IsEmpty() {
		panic("stack is empty")
	}
	*stack = old[1:]
	return popped
}

func main() {
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	for !stack.IsEmpty() {
		data := stack.Pop()
		fmt.Println(data)
	}
}
