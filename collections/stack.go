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
func (s *stack[Type]) Push(data Type) {
	node := t.NewSingleNode(data)

	node.AddReferenceOnNext(s.head)
	s.head = node
	s.length++
}

// Pop: removes the last element added in the stack, returning the data in the SingleNode[Type] - O(1)
func (s *stack[Type]) Pop() (Type, error) {
	var data Type
	
	if (s.length == 0) {
		return data, errEmptyStack
	}
	data = s.head.Data()
	s.head = s.head.Next()
	s.length--

	return data, nil
}

// print: traverses through the stack[Type], printing the data to the existing SingleNode[Type] - O(n)
func (s *stack[Type]) print() {
	hook := s.head

	for  (hook != nil) {
		fmt.Printf("%v,\n", hook.Data())
		hook = hook.Next()
	}
}

func (s *stack[Type]) find(index int) (Type, error) {
	hook := s.head
	var v Type

	if index < 0 || index >= s.length {
		return v, errOutOfRangeIndex
	}

	for (index > 0) {
		hook = hook.Next()
		index--
	}
	v = hook.Data()
	
	return v, nil
}

func (s *stack[Type]) Length() int {
	return s.length
}
