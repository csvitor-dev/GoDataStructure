package linkedlist

import (
	"data-structure/node"
	"errors"
	"fmt"
)

var (
	errOutOfRangeIndex error = errors.New("out of range index")
	errEmptyList error = errors.New("linked list is empty")
)

type singleLinkedList[Type node.Number | string] struct {
	head   *node.SingleNode[Type]
	tail   *node.SingleNode[Type]
	Length int
}

// NewSingleLinkedList: ...
func NewSingleLinkedList[Type node.Number | string]() *singleLinkedList[Type] {
	linkedList := &singleLinkedList[Type]{
		head: nil,
		tail: nil,
		Length: 0,
	}
	return linkedList
}

// Add: insert default
func (linkedList *singleLinkedList[Type]) Add(data Type) {
	node := node.NewSingleNode(data)

	if (linkedList.Length == 0) {
		linkedList.head = node
		linkedList.tail = node
	} else {
		linkedList.tail.Next = node
		linkedList.tail = node
	}
	linkedList.Length++
}

// InsertAt: ...
func (linkedList *singleLinkedList[Type]) InsertAt(index int, data Type) error {
	if (linkedList.isValidIndex(index)) {
		return errOutOfRangeIndex
	}

	if (index == linkedList.Length) {
		linkedList.Add(data)
		return nil
	}
	node := node.NewSingleNode(data)

	if (index == 0) {
		node.Next = linkedList.head
		linkedList.head = node
	} else {
		hookAtIndex := linkedList.searchNode(index)
		node.Next = hookAtIndex.Next
		hookAtIndex.Next = node
	}
	linkedList.Length++
	return nil
}

// Delete: remove default
func (linkedList *singleLinkedList[Type]) Delete() (Type, error) {
	var data Type

	if (linkedList.Length == 0) {
		return data, errEmptyList
	}

	hook := linkedList.head
	linkedList.head = hook.Next

	data = hook.Data
	hook.Next = nil

	linkedList.Length--
	return data, nil
}

// Print: ...
func (linkedList *singleLinkedList[Type]) Print() {
	hook := linkedList.head
	
	for {
		if (hook == nil) {
			break
		}
		fmt.Printf("%v, ", hook.Data)
		hook = hook.Next
	}
	fmt.Printf("Length: %v\n", linkedList.Length)
}

// isValidIndex: ...
func (linkedList *singleLinkedList[Type]) isValidIndex(index int) bool {
	return index < 0 || index > linkedList.Length
}

// searchNode: ...
func (linkedList *singleLinkedList[Type]) searchNode(index int) *node.SingleNode[Type] {
	hook := linkedList.head

	for {
		if (index - 1 == 0) {
			return hook
		}
		hook = hook.Next
		index--
	}
}