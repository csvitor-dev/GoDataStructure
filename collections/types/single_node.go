package types

import "errors"

var (
	errNilReference error = errors.New("the reference is null")
)

type SingleNode[Type T] struct {
	data Type
	next *SingleNode[Type]
}

// NewSingleNode: create new instance of SingleNode[Type]
func NewSingleNode[Type T](data Type) *SingleNode[Type] {
	node := &SingleNode[Type]{
		data: data,
		next: nil,
	}
	return node
}

func (s *SingleNode[Type]) Data() Type {
	return s.data
}

func (s *SingleNode[Type]) Next() *SingleNode[Type] {
	return s.next
}

func (s *SingleNode[Type]) AddReferenceOnNext(newNode *SingleNode[Type]) error {
	if s == nil {
		return errNilReference
	}
	s.next = newNode
	return nil
}