package collections

import (
	"data-structure/collections/node"
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

// NewSingleLinkedList: create new instance of singleLinkedList[Type]
func NewSingleLinkedList[Type node.Number | string]() *singleLinkedList[Type] {
	linkedList := &singleLinkedList[Type]{
		head: nil,
		tail: nil,
		Length: 0,
	}
	return linkedList
}

// Add: adds a new SingleNode[Type] to the singleLinkedList[Type], default insertion - 0(1)
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

// InsertAt: adds a new SingleNode[Type] to the singleLinkedList[Type] in the index if it is valid, otherwise return error - O(n)
func (linkedList *singleLinkedList[Type]) InsertAt(index int, data Type) error {
	if (linkedList.isValidIndexInsert(index)) {
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

// Delete: removes the first element in the list, returning the data to the SingleNode[Type], default remotion - O(1)
func (linkedList *singleLinkedList[Type]) Delete() (Type, error) {
	var data Type

	if (linkedList.Length == 0) {
		return data, errEmptyList
	}

	hook := linkedList.head
	linkedList.head = hook.Next

	if (linkedList.Length == 1) {
		linkedList.tail = hook.Next
	}

	data = hook.Data
	hook.Next = nil

	linkedList.Length--
	return data, nil
}

// RemoveAt: removes the element of the singleLinkedList[Type] in the valid index returning it, otherwise return error - O(n)
func (linkedList *singleLinkedList[Type]) RemoveAt(index int) (Type, error) {
	var data Type

	if (linkedList.isValidIndexRemove(index)) {
		return data, errOutOfRangeIndex
	}

	if (index == 0) {
		return linkedList.Delete()
	}

	hookAtIndex := linkedList.searchNode(index)

	if (hookAtIndex.Next == linkedList.tail) {
		linkedList.tail = hookAtIndex
	}
	data = hookAtIndex.Next.Data

	hook := hookAtIndex.Next
	hookAtIndex.Next = hook.Next
	hook.Next = nil

	linkedList.Length--
	return data, nil
}

// Print: traverses through the singleLinkedList[Type], printing the data to the existing SingleNode[Type] - O(n)
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

// isValidIndexInsert: validates the index based in the list to Insert method
func (linkedList *singleLinkedList[Type]) isValidIndexInsert(index int) bool {
	return index < 0 || index > linkedList.Length
}

func (linkedList *singleLinkedList[Type]) isValidIndexRemove(index int) bool {
	return index < 0 || index >= linkedList.Length
}

// searchNode: searches for SingleNode[Type] in the valid index and returns the reference of the node before it
func (linkedList *singleLinkedList[Type]) searchNode(index int) (*node.SingleNode[Type]) {
	hook := linkedList.head

	for {
		if (index - 1 == 0) {
			return hook
		}
		hook = hook.Next
		index--
	}
}