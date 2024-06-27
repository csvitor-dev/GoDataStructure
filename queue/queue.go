package queue

import (
	"data-structure/node"
	"errors"
	"fmt"
)

type queue[Type node.Number | string] struct {
	head   *node.Node[Type]
	Length int
}

// Newqueue: create new instance of queue[Type]
func NewQueue[Type node.Number | string]() *queue[Type] {
	queue := &queue[Type]{
		head: nil,
		Length: 0,
	}
	return queue
}

// Enqueue: adds a new Node[Type] to the queue[Type]
func (queue *queue[Type]) Enqueue(data Type) {
	node := node.NewNode(data)
	
	if queue.Length != 0 {
		node.Next = queue.head
	}
	queue.head = node
	queue.Length++
}

// Dequeue: removes the last element added, returning the data in the Node[Type]
func (queue *queue[Type]) Dequeue() (Type, error) {
	var data Type

	if (queue.Length == 0) {
		return data, errors.New("empty list error")
	}
	data = queue.head.Data
	queue.head = queue.head.Next
	queue.Length--

	return data, nil
}

// Print: scroll through the queue[Type], printing the data to the existing Node[Type]
func (queue *queue[Type]) Print() {
	hook := queue.head

	for i := range queue.Length {
		fmt.Printf("%d: %v, ", i, hook.Data)
		hook = hook.Next
	}
	fmt.Printf("Length: %v\n", queue.Length)
}