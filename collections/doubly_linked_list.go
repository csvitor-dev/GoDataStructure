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
func (dll *doublyLinkedList[Type]) Add(data Type) {
	node := t.NewDoubleNode(data)

	if (dll.length == 0) {
		dll.head = node
	} else {
		dll.tail.AddReferenceOnNext(node)
		node.AddReferenceOnPrevious(dll.tail)
	}
	dll.tail = node
	dll.length++
}

// InsertAt: adds a new DoubleNode[Type] to the doublyLinkedList[Type] in the index if it is valid, otherwise return error - O(n)
func (dll *doublyLinkedList[Type]) InsertAt(index int, data Type) error {
	if (!dll.isValidIndexInsert(index)) {
		return errOutOfRangeIndex
	}

	if (index == dll.length) {
		dll.Add(data)
		return nil
	}
	node := t.NewDoubleNode(data)

	if (index == 0) {
		node.AddReferenceOnNext(dll.head)
		dll.head.AddReferenceOnPrevious(node)
		dll.head = node
	} else {
		hookAtIndex := dll.searchNode(index)
		node.AddReferenceOnNext(hookAtIndex.Next())
		node.AddReferenceOnPrevious(hookAtIndex)
		hookAtIndex.Next().AddReferenceOnPrevious(node)
		hookAtIndex.AddReferenceOnNext(node)
	}
	dll.length++
	return nil
}

// Delete: removes the first element in the list, returning the data to the DoubleNode[Type], default remotion - O(1)
func (dll *doublyLinkedList[Type]) Delete() (Type, error) {
	var data Type

	if (dll.length == 0) {
		return data, errEmptyList
	}

	hook := dll.head
	dll.head = hook.Next()

	if (dll.head != nil) {
		dll.head.AddReferenceOnPrevious(nil)
	}

	if (dll.length == 1) {
		dll.tail = hook.Previous()
	}

	data = hook.Data()
	hook.AddReferenceOnNext(nil)

	dll.length--
	return data, nil
}

// RemoveAt: removes the element of the doublyLinkedList[Type] in the valid index returning it, otherwise return error - O(n)
func (dll *doublyLinkedList[Type]) RemoveAt(index int) (Type, error) {
	var data Type

	if (!dll.isValidIndexRemove(index)) {
		return data, errOutOfRangeIndex
	}

	if (index == 0) {
		return dll.Delete()
	}

	hookAtIndex := dll.searchNode(index)

	if (hookAtIndex.Next() == dll.tail) {
		dll.tail = hookAtIndex
	}
	data = hookAtIndex.Next().Data()

	hook := hookAtIndex.Next()
	hookAtIndex.AddReferenceOnNext(hook.Next())

	if (hook.Next() != nil) {
		hook.Next().AddReferenceOnPrevious(hookAtIndex)
	}
	hook.AddReferenceOnNext(nil)
	hook.AddReferenceOnPrevious(nil)

	dll.length--
	return data, nil
}

// print: traverses through the doublyLinkedList[Type], printing the data to the existing DoubleNode[Type];
// option parameter adllows define the order of print: true is ascending and false is descending - O(n)
func (dll *doublyLinkedList[Type]) print() {
	if (dll.length == 0) {
		return
	}
	fmt.Println("Print Ascendent...")
	dll.printAscendent()

	fmt.Println("\nPrint Descendent...")
	dll.printDescendent()
}

// printAscendent: ...
func (dll *doublyLinkedList[Type]) printAscendent() {
	hook := dll.head
	
	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data())
		hook = hook.Next()
	}
}

// printDescendent: ...
func (dll *doublyLinkedList[Type]) printDescendent() {
	hook := dll.tail
	
	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data())
		hook = hook.Previous()
	}
}

func (dll *doublyLinkedList[Type]) Length() int {
	return dll.length
}

// isValidIndexInsert: validates the index based in the list to InsertAt method
func (dll *doublyLinkedList[Type]) isValidIndexInsert(index int) bool {
	return index >= 0 && index <= dll.length
}

// isValidIndexRemove: validates the index based in the list to RemoveAt method
func (dll *doublyLinkedList[Type]) isValidIndexRemove(index int) bool {
	return index >= 0 && index < dll.length
}

// searchNode: searches for DoubleNode[Type] in the valid index and returns the reference of the node before it
func (dll *doublyLinkedList[Type]) searchNode(index int) (*t.DoubleNode[Type]) {
	hook := dll.head

	for {
		if (index - 1 == 0) {
			return hook
		}
		hook = hook.Next()
		index--
	}
}