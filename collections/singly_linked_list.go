package collections

import (
	"GoDataStructure/collections/node"
	"errors"
	"fmt"
)

var (
	errOutOfRangeIndex error = errors.New("out of range index")
	errEmptyList error = errors.New("linked list is empty")
)

type singlyLinkedList[Type node.Number | string] struct {
	head   *node.SingleNode[Type]
	length int
}

// NewSinglyLinkedList: create new instance of singlyLinkedList[Type]
func NewSinglyLinkedList[Type node.Number | string]() *singlyLinkedList[Type] {
	linkedList := &singlyLinkedList[Type]{
		head: nil,
		length: 0,
	}
	return linkedList
}

// Add: adds a new SingleNode[Type] to the singlyLinkedList[Type], default insertion - 0(1)
func (linkedList *singlyLinkedList[Type]) Add(data Type) {
	node := node.NewSingleNode(data)

	if (linkedList.length == 0) {
		linkedList.head = node
	} else {
		hookAtLastIndex := linkedList.searchNode(linkedList.length)
		hookAtLastIndex.Next = node
	}
	linkedList.length++
}

// InsertAt: adds a new SingleNode[Type] to the singlyLinkedList[Type] in the index if it is valid, otherwise return error - O(n)
func (linkedList *singlyLinkedList[Type]) InsertAt(index int, data Type) error {
	if (!linkedList.isValidIndexInsert(index)) {
		return errOutOfRangeIndex
	}

	if (index == linkedList.length) {
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
	linkedList.length++
	return nil
}

// Delete: removes the first element in the list, returning the data to the SingleNode[Type], default remotion - O(1)
func (linkedList *singlyLinkedList[Type]) Delete() (Type, error) {
	var data Type

	if (linkedList.length == 0) {
		return data, errEmptyList
	}

	hook := linkedList.head
	linkedList.head = hook.Next

	data = hook.Data
	hook.Next = nil

	linkedList.length--
	return data, nil
}

// RemoveAt: removes the element of the singlyLinkedList[Type] in the valid index returning it, otherwise return error - O(n)
func (linkedList *singlyLinkedList[Type]) RemoveAt(index int) (Type, error) {
	var data Type

	if (!linkedList.isValidIndexRemove(index)) {
		return data, errOutOfRangeIndex
	}

	if (index == 0) {
		return linkedList.Delete()
	}
	hook := linkedList.searchNode(index)
	hookAtIndex := hook.Next
	
	data = hookAtIndex.Data
	hook.Next = hookAtIndex.Next
	hookAtIndex.Next = nil

	linkedList.length--
	return data, nil
}

// Print: traverses through the singlyLinkedList[Type], printing the data to the existing SingleNode[Type] - O(n)
func (linkedList *singlyLinkedList[Type]) Print() {
	hook := linkedList.head
	
	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data)
		hook = hook.Next
	}
	fmt.Printf("Length: %v\n", linkedList.length)
}

// isValidIndexInsert: validates the index based in the list to InsertAt method
func (linkedList *singlyLinkedList[Type]) isValidIndexInsert(index int) bool {
	return index >= 0 && index <= linkedList.length
}

// isValidIndexRemove: validates the index based in the list to RemoveAt method
func (linkedList *singlyLinkedList[Type]) isValidIndexRemove(index int) bool {
	return index >= 0 && index < linkedList.length
}

// searchNode: searches for SingleNode[Type] in the valid index and returns the reference of the node before it
func (linkedList *singlyLinkedList[Type]) searchNode(index int) (*node.SingleNode[Type]) {
	hook := linkedList.head

	for {
		if (index - 1 == 0) {
			return hook
		}
		hook = hook.Next
		index--
	}
}