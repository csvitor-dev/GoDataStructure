package types

type DoubleNode[Type T] struct {
	data     Type
	next     *DoubleNode[Type]
	previous *DoubleNode[Type]
}

// NewDoubleNode: create new instance of DoubleNode[Type]
func NewDoubleNode[Type T](data Type) *DoubleNode[Type] {
	node := &DoubleNode[Type]{
		data:     data,
		next:     nil,
		previous: nil,
	}
	return node
}

func (d *DoubleNode[Type]) Data() Type {
	return d.data
}

func (d *DoubleNode[Type]) Next() *DoubleNode[Type] {
	return d.next
}
func (d *DoubleNode[Type]) AddReferenceOnNext(newNode *DoubleNode[Type]) error {
	if d == nil {
		return errNilReference
	}
	d.next = newNode
	return nil
}

func (d *DoubleNode[Type]) Previous() *DoubleNode[Type] {
	return d.previous
}
func (d *DoubleNode[Type]) AddReferenceOnPrevious(newNode *DoubleNode[Type]) error {
	if d == nil {
		return errNilReference
	}
	d.previous = newNode
	return nil
}
