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

// Using interface as type
// using panic in case of type mismatch
/*
package main

import (
	"fmt"
	"reflect"
)

type Stack []interface{}

func (stack *Stack) Push(ele interface{}) bool {
	if len(*stack) != 0 {
		lastInserted := (*stack)[len(*stack)-1]

		if reflect.TypeOf(lastInserted) != reflect.TypeOf(ele) {
			panic("Type mismatch")
		}
	}

	*stack = append([]interface{}{ele}, *stack...)

	return true
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Pop() interface{} {
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
	stack.Push("abc")

	for !stack.IsEmpty() {
		data := stack.Pop()
		fmt.Println(data)
	}
}

*/
