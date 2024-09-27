package types

type DoubleNode[Type T] struct {
	Data     Type
	Next     *DoubleNode[Type]
	Previous *DoubleNode[Type]
}

// NewDoubleNode: create new instance of DoubleNode[Type]
func NewDoubleNode[Type T](data Type) *DoubleNode[Type] {
	node := &DoubleNode[Type]{
		Data:     data,
		Next:     nil,
		Previous: nil,
	}
	return node
}