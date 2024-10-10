package collections

import (
	t "GoDataStructure/collections/types"
	"errors"
	"fmt"
)

var (
	errEmptyQueue error = errors.New("queue is empty")
)

// FIFO - First In, First Out
type queue[Type t.T] struct {
	head   *t.SingleNode[Type]
	tail   *t.SingleNode[Type]
	length int
}

// NewQueue: create new instance of queue[Type]
func NewQueue[Type t.T]() *queue[Type] {
	queue := &queue[Type]{
		head: nil,
		tail: nil,
		length: 0,
	}
	return queue
}

// Enqueue: adds a new SingleNode[Type] to the queue[Type] - O(1)
func (q *queue[Type]) Enqueue(data Type) {
	node := t.NewSingleNode(data)
	
	if (q.length == 0) {
		q.head = node
	} else {
		q.tail.AddReferenceOnNext(node)
	}
	q.tail = node
	q.length++
}

// Dequeue: removes the first element added, returning the data in the SingleNode[Type] - O(1)
func (q *queue[Type]) Dequeue() (Type, error) {
	var data Type

	if (q.length == 0) {
		return data, errEmptyQueue
	}
	data = q.head.Data()

	if (q.length == 1) {
		q.tail = q.tail.Next()
	}
	q.head = q.head.Next()
	q.length--

	return data, nil
}

// Contains: checks if queue[Type] contains the data in any SingleNode[Type],
// returning true if yes, false otherwise - O(n)
func (q *queue[Type]) Contains(data Type) (bool, error) {
	if (q.length == 0) {
		return false, errEmptyQueue
	}
	hook := q.head

	for {
		if (hook == nil) {
			return false, nil
		}
		if (hook.Data() == data) {
			return true, nil
		}
		hook = hook.Next()
	}
}

// print: traverses through the queue[Type], printing the data to the existing SingleNode[Type] - O(n)
func (q *queue[Type]) print() {
	hook := q.head

	for (hook != nil) {
		fmt.Printf("%v, ", hook.Data())
		hook = hook.Next()
	}
}

func (q *queue[Type]) find(index int) (Type, error) {
	hook := q.head
	var v Type

	if index < 0 || index >= q.length {
		return v, errOutOfRangeIndex
	}

	for (index > 0) {
		hook = hook.Next()
		index--
	}
	v = hook.Data()
	
	return v, nil
}

func (q *queue[Type]) Length() int {
	return q.length
}