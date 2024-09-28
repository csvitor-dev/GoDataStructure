package collections

import (
	t "GoDataStructure/collections/types"
	"fmt"
)

type doublyLinkedList[Type t.T] struct {
	head   *t.DoubleNode[Type]
	tail   *t.DoubleNode[Type]
	length int
}

// NewDoublyLinkedList: create new instance of doublyLinkedList[Type]
func NewDoublyLinkedList[Type t.T]() *doublyLinkedList[Type] {
	linkedList := &doublyLinkedList[Type]{
		head: nil,
		tail: nil,
		length: 0,
	}
	return linkedList
}

// Add: adds a new DoubleNode[Type] to the doublyLinkedList[Type], default insertion - 0(1)
func (ll *doublyLinkedList[Type]) Add(data Type) {
	node := t.NewDoubleNode(data)

	if (ll.length == 0) {
		ll.head = node
	} else {
		ll.tail.AddReferenceOnNext(node)
		node.AddReferenceOnPrevious(ll.tail)
	}
	ll.tail = node
	ll.length++
}

// InsertAt: adds a new DoubleNode[Type] to the doublyLinkedList[Type] in the index if it is valid, otherwise return error - O(n)
func (ll *doublyLinkedList[Type]) InsertAt(index int, data Type) error {
	if (!ll.isValidIndexInsert(index)) {
		return errOutOfRangeIndex
	}

	if (index == ll.length) {
		ll.Add(data)
		return nil
	}
	node := t.NewDoubleNode(data)

	if (index == 0) {
		node.AddReferenceOnNext(ll.head)
		ll.head.AddReferenceOnPrevious(node)
		ll.head = node
	} else {
		hookAtIndex := ll.searchNode(index)
		node.AddReferenceOnNext(hookAtIndex.Next())
		node.AddReferenceOnPrevious(hookAtIndex)
		hookAtIndex.Next().AddReferenceOnPrevious(node)
		hookAtIndex.AddReferenceOnNext(node)
	}
	ll.length++
	return nil
}

// Delete: removes the first element in the list, returning the data to the DoubleNode[Type], default remotion - O(1)
func (ll *doublyLinkedList[Type]) Delete() (Type, error) {
	var data Type

	if (ll.length == 0) {
		return data, errEmptyList
	}

	hook := ll.head
	ll.head = hook.Next()

	if (ll.head != nil) {
		ll.head.AddReferenceOnPrevious(nil)
	}

	if (ll.length == 1) {
		ll.tail = hook.Previous()
	}

	data = hook.Data()
	hook.AddReferenceOnNext(nil)

	ll.length--
	return data, nil
}

// RemoveAt: removes the element of the doublyLinkedList[Type] in the valid index returning it, otherwise return error - O(n)
func (ll *doublyLinkedList[Type]) RemoveAt(index int) (Type, error) {
	var data Type

	if (!ll.isValidIndexRemove(index)) {
		return data, errOutOfRangeIndex
	}

	if (index == 0) {
		return ll.Delete()
	}

	hookAtIndex := ll.searchNode(index)

	if (hookAtIndex.Next() == ll.tail) {
		ll.tail = hookAtIndex
	}
	data = hookAtIndex.Next().Data()

	hook := hookAtIndex.Next()
	hookAtIndex.AddReferenceOnNext(hook.Next())

	if (hook.Next() != nil) {
		hook.Next().AddReferenceOnPrevious(hookAtIndex)
	}
	hook.AddReferenceOnNext(nil)
	hook.AddReferenceOnPrevious(nil)

	ll.length--
	return data, nil
}

// Print: traverses through the doublyLinkedList[Type], printing the data to the existing DoubleNode[Type];
// option parameter allows define the order of print: true is ascending and false is descending - O(n)
func (ll *doublyLinkedList[Type]) Print(option bool) {
	if (ll.length == 0) {
		fmt.Printf("Length: %v\n", ll.length)
		return
	}
	
	if (option) {
		ll.printAscendent()
		return
	}
	ll.printDescendent()
}

// printAscendent: ...
func (ll *doublyLinkedList[Type]) printAscendent() {
	hook := ll.head
	
	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data())
		hook = hook.Next()
	}
	fmt.Printf("Length: %v\n", ll.length)
}

// printDescendent: ...
func (ll *doublyLinkedList[Type]) printDescendent() {
	hook := ll.tail
	
	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data())
		hook = hook.Previous()
	}
	fmt.Printf("Length: %v\n", ll.length)
}

// isValidIndexInsert: validates the index based in the list to InsertAt method
func (ll *doublyLinkedList[Type]) isValidIndexInsert(index int) bool {
	return index >= 0 && index <= ll.length
}

// isValidIndexRemove: validates the index based in the list to RemoveAt method
func (ll *doublyLinkedList[Type]) isValidIndexRemove(index int) bool {
	return index >= 0 && index < ll.length
}

// searchNode: searches for DoubleNode[Type] in the valid index and returns the reference of the node before it
func (ll *doublyLinkedList[Type]) searchNode(index int) (*t.DoubleNode[Type]) {
	hook := ll.head

	for {
		if (index - 1 == 0) {
			return hook
		}
		hook = hook.Next()
		index--
	}
}