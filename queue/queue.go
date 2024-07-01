package queue

import (
	"data-structure/node"
	"errors"
	"fmt"
)

var (
	errEmptyQueue error = errors.New("queue is empty")
)

type queue[Type node.Number | string] struct {
	head   *node.SingleNode[Type]
	Length int
}

// NewQueue: create new instance of queue[Type]
func NewQueue[Type node.Number | string]() *queue[Type] {
	queue := &queue[Type]{
		head: nil,
		Length: 0,
	}
	return queue
}

// Enqueue: adds a new SingleNode[Type] to the queue[Type] - O(1)
func (queue *queue[Type]) Enqueue(data Type) {
	node := node.NewSingleNode(data)
	
	if queue.Length != 0 {
		node.Next = queue.head
	}
	queue.head = node
	queue.Length++
}

// Dequeue: removes the last element added, returning the data in the SingleNode[Type] - O(1)
func (queue *queue[Type]) Dequeue() (Type, error) {
	var data Type

	if (queue.Length == 0) {
		return data, errEmptyQueue
	}
	data = queue.head.Data
	queue.head = queue.head.Next
	queue.Length--

	return data, nil
}

// Contains: checks if queue[Type] contains the data in any SingleNode[Type],
// returning true if yes, false otherwise - O(n)
func (queue *queue[Type]) Contains(data Type) (bool, error) {
	if (queue.Length == 0) {
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

	for {
		if (hook == nil) {
			break
		}
		fmt.Printf("%v, ", hook.Data)
		hook = hook.Next
	}
	fmt.Printf("Length: %v\n", queue.Length)
}