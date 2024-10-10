package collections

import (
	t "GoDataStructure/collections/types"
	"errors"
	"fmt"
)

var (
	errOutOfRangeIndex error = errors.New("out of range index")
	errEmptyList error = errors.New("linked list is empty")
)

type singlyLinkedList[Type t.T] struct {
	head   *t.SingleNode[Type]
	length int
}

// NewSinglyLinkedList: create new instance of singlyLinkedList[Type]
func NewSinglyLinkedList[Type t.T]() *singlyLinkedList[Type] {
	linkedList := &singlyLinkedList[Type]{
		head: nil,
		length: 0,
	}
	return linkedList
}

// Add: adds a new SingleNode[Type] to the singlyLinkedList[Type], default insertion - 0(1)
func (sll *singlyLinkedList[Type]) Add(data Type) {
	node := t.NewSingleNode(data)

	if (sll.length == 0) {
		sll.head = node
	} else {
		hookAtLastIndex := sll.searchNode(sll.length)
		hookAtLastIndex.AddReferenceOnNext(node)
	}
	sll.length++
}

// InsertAt: adds a new SingleNode[Type] to the singlyLinkedList[Type] in the index if it is valid, otherwise return error - O(n)
func (sll *singlyLinkedList[Type]) InsertAt(index int, data Type) error {
	if (!sll.isValidIndexInsert(index)) {
		return errOutOfRangeIndex
	}

	if (index == sll.length) {
		sll.Add(data)
		return nil
	}
	node := t.NewSingleNode(data)

	if (index == 0) {
		node.AddReferenceOnNext(sll.head)
		sll.head = node
	} else {
		hookAtIndex := sll.searchNode(index)
		node.AddReferenceOnNext(hookAtIndex.Next())
		hookAtIndex.AddReferenceOnNext(node)
	}
	sll.length++
	return nil
}

// Delete: removes the first element in the list, returning the data to the SingleNode[Type], default remotion - O(1)
func (sll *singlyLinkedList[Type]) Delete() (Type, error) {
	var data Type

	if (sll.length == 0) {
		return data, errEmptyList
	}

	hook := sll.head
	sll.head = hook.Next()

	data = hook.Data()
	hook.AddReferenceOnNext(nil)

	sll.length--
	return data, nil
}

// RemoveAt: removes the element of the singlyLinkedList[Type] in the valid index returning it, otherwise return error - O(n)
func (sll *singlyLinkedList[Type]) RemoveAt(index int) (Type, error) {
	var data Type

	if (!sll.isValidIndexRemove(index)) {
		return data, errOutOfRangeIndex
	}

	if (index == 0) {
		return sll.Delete()
	}
	hook := sll.searchNode(index)
	hookAtIndex := hook.Next()
	
	data = hookAtIndex.Data()
	hook.AddReferenceOnNext(hookAtIndex.Next())
	hookAtIndex.AddReferenceOnNext(nil)

	sll.length--
	return data, nil
}

// print: traverses through the singlyLinkedList[Type], printing the data to the existing SingleNode[Type] - O(n)
func (sll *singlyLinkedList[Type]) print() {
	hook := sll.head
	
	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data())
		hook = hook.Next()
	}
}

func (sll *singlyLinkedList[Type]) find(index int) (Type, error) {
	hook := sll.head
	var v Type

	if index < 0 || index >= sll.length {
		return v, errOutOfRangeIndex
	}

	for (index > 0) {
		hook = hook.Next()
		index--
	}
	v = hook.Data()
	
	return v, nil
}

func (sll *singlyLinkedList[Type]) Length() int {
	return sll.length
}

// isValidIndexInsert: validates the index based in the list to InsertAt method
func (sll *singlyLinkedList[Type]) isValidIndexInsert(index int) bool {
	return index >= 0 && index <= sll.length
}

// isValidIndexRemove: validates the index based in the list to RemoveAt method
func (sll *singlyLinkedList[Type]) isValidIndexRemove(index int) bool {
	return index >= 0 && index < sll.length
}

// searchNode: searches for SingleNode[Type] in the valid index and returns the reference of the node before it
func (sll *singlyLinkedList[Type]) searchNode(index int) (*t.SingleNode[Type]) {
	hook := sll.head

	for (index - 1 > 0) {
		hook = hook.Next()
		index--
	}
	return hook
}