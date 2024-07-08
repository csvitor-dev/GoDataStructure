package collections

import (
	"data-structure/collections/node"
	"errors"
	"fmt"
)

var (
	errEmptyQueue error = errors.New("queue is empty")
)

// FIFO - First In, First Out
type queue[Type node.Number | string] struct {
	head   *node.SingleNode[Type]
	tail   *node.SingleNode[Type]
	length int
}

// NewQueue: create new instance of queue[Type]
func NewQueue[Type node.Number | string]() *queue[Type] {
	queue := &queue[Type]{
		head: nil,
		tail: nil,
		length: 0,
	}
	return queue
}

// Enqueue: adds a new SingleNode[Type] to the queue[Type] - O(1)
func (queue *queue[Type]) Enqueue(data Type) {
	node := node.NewSingleNode(data)
	
	if (queue.length == 0) {
		queue.head = node
	} else {
		queue.tail.Next = node
	}
	queue.tail = node
	queue.length++
}

// Dequeue: removes the first element added, returning the data in the SingleNode[Type] - O(1)
func (queue *queue[Type]) Dequeue() (Type, error) {
	var data Type

	if (queue.length == 0) {
		return data, errEmptyQueue
	}
	data = queue.head.Data

	if (queue.length == 1) {
		queue.tail = queue.tail.Next
	}
	queue.head = queue.head.Next
	queue.length--

	return data, nil
}

// Contains: checks if queue[Type] contains the data in any SingleNode[Type],
// returning true if yes, false otherwise - O(n)
func (queue *queue[Type]) Contains(data Type) (bool, error) {
	if (queue.length == 0) {
		return false, errEmptyQueue
	}
	hook := queue.head

	for {
		if (hook == nil) {
			return false, nil
		}
		if (hook.Data == data) {
			return true, nil
		}
		hook = hook.Next
	}
}

// Print: traverses through the queue[Type], printing the data to the existing SingleNode[Type] - O(n)
func (queue *queue[Type]) Print() {
	hook := queue.head

	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data)
		hook = hook.Next
	}
	fmt.Printf("Length: %v\n", queue.length)
}