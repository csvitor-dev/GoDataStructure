package collections

import (
	"data-structure/collections/node"
	"fmt"
)

type doublyLinkedList[Type node.Number | string] struct {
	head   *node.DoubleNode[Type]
	tail   *node.DoubleNode[Type]
	length int
}

// NewDoublyLinkedList: create new instance of doublyLinkedList[Type]
func NewDoublyLinkedList[Type node.Number | string]() *doublyLinkedList[Type] {
	linkedList := &doublyLinkedList[Type]{
		head: nil,
		tail: nil,
		length: 0,
	}
	return linkedList
}

// Add: adds a new DoubleNode[Type] to the doublyLinkedList[Type], default insertion - 0(1)
func (linkedList *doublyLinkedList[Type]) Add(data Type) {
	node := node.NewDoubleNode(data)

	if (linkedList.length == 0) {
		linkedList.head = node
	} else {
		linkedList.tail.Next = node
		node.Previous = linkedList.tail
	}
	linkedList.tail = node
	linkedList.length++
}

// InsertAt: adds a new DoubleNode[Type] to the doublyLinkedList[Type] in the index if it is valid, otherwise return error - O(n)
func (linkedList *doublyLinkedList[Type]) InsertAt(index int, data Type) error {
	if (!linkedList.isValidIndexInsert(index)) {
		return errOutOfRangeIndex
	}

	if (index == linkedList.length) {
		linkedList.Add(data)
		return nil
	}
	node := node.NewDoubleNode(data)

	if (index == 0) {
		node.Next = linkedList.head
		linkedList.head.Previous = node
		linkedList.head = node
	} else {
		hookAtIndex := linkedList.searchNode(index)
		node.Next = hookAtIndex.Next
		node.Previous = hookAtIndex
		hookAtIndex.Next.Previous = node
		hookAtIndex.Next = node
	}
	linkedList.length++
	return nil
}

// Delete: removes the first element in the list, returning the data to the DoubleNode[Type], default remotion - O(1)
func (linkedList *doublyLinkedList[Type]) Delete() (Type, error) {
	var data Type

	if (linkedList.length == 0) {
		return data, errEmptyList
	}

	hook := linkedList.head
	linkedList.head = hook.Next

	if (linkedList.head != nil) {
		linkedList.head.Previous = nil
	}

	if (linkedList.length == 1) {
		linkedList.tail = hook.Previous
	}

	data = hook.Data
	hook.Next = nil

	linkedList.length--
	return data, nil
}

// RemoveAt: removes the element of the doublyLinkedList[Type] in the valid index returning it, otherwise return error - O(n)
func (linkedList *doublyLinkedList[Type]) RemoveAt(index int) (Type, error) {
	var data Type

	if (!linkedList.isValidIndexRemove(index)) {
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

	if (hook.Next != nil) {
		hook.Next.Previous = hookAtIndex
	}
	hook.Next = nil
	hook.Previous = nil

	linkedList.length--
	return data, nil
}

// Print: traverses through the doublyLinkedList[Type], printing the data to the existing DoubleNode[Type];
// option parameter allows define the order of print: true is ascending and false is descending - O(n)
func (linkedList *doublyLinkedList[Type]) Print(option bool) {
	if (linkedList.length == 0) {
		fmt.Printf("Length: %v\n", linkedList.length)
		return
	}
	
	if (option) {
		linkedList.printAscendent()
		return
	}
	linkedList.printDescendent()
}

// printAscendent: ...
func (linkedList *doublyLinkedList[Type]) printAscendent() {
	hook := linkedList.head
	
	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data)
		hook = hook.Next
	}
	fmt.Printf("Length: %v\n", linkedList.length)
}

// printDescendent: ...
func (linkedList *doublyLinkedList[Type]) printDescendent() {
	hook := linkedList.tail
	
	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data)
		hook = hook.Previous
	}
	fmt.Printf("Length: %v\n", linkedList.length)
}

// isValidIndexInsert: validates the index based in the list to InsertAt method
func (linkedList *doublyLinkedList[Type]) isValidIndexInsert(index int) bool {
	return index >= 0 && index <= linkedList.length
}

// isValidIndexRemove: validates the index based in the list to RemoveAt method
func (linkedList *doublyLinkedList[Type]) isValidIndexRemove(index int) bool {
	return index >= 0 && index < linkedList.length
}

// searchNode: searches for DoubleNode[Type] in the valid index and returns the reference of the node before it
func (linkedList *doublyLinkedList[Type]) searchNode(index int) (*node.DoubleNode[Type]) {
	hook := linkedList.head

	for {
		if (index - 1 == 0) {
			return hook
		}
		hook = hook.Next
		index--
	}
}