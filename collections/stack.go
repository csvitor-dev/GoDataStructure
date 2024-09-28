package collections

import (
	t "GoDataStructure/collections/types"
	"errors"
	"fmt"
)

var (
	errEmptyStack error = errors.New("stack is empty")
)

// LIFO - Last In, First Out
type stack[Type t.T] struct {
	head   *t.SingleNode[Type]
	length int
}

// NewStack: create new instance of stack[Type]
func NewStack[Type t.T]() *stack[Type] {
	stack := &stack[Type]{
		head: nil,
		length: 0,
	}
	return stack
}

// Push: adds a new SingleNode[Type] at the top of the stack[Type] - O(1)
func (stack *stack[Type]) Push(data Type) {
	node := t.NewSingleNode(data)

	node.AddReferenceOnNext(stack.head)
	stack.head = node
	stack.length++
}

// Pop: removes the last element added in the stack, returning the data in the SingleNode[Type] - O(1)
func (stack *stack[Type]) Pop() (Type, error) {
	var data Type
	
	if (stack.length == 0) {
		return data, errEmptyStack
	}
	data = stack.head.Data()
	stack.head = stack.head.Next()
	stack.length--

	return data, nil
}

// Print: traverses through the stack[Type], printing the data to the existing SingleNode[Type] - O(n)
func (stack *stack[Type]) Print() {
	hook := stack.head

	for  (hook != nil) {
		fmt.Printf("%v,\n", hook.Data())
		hook = hook.Next()
	}
	fmt.Printf("Length: %v\n", stack.length)
}
