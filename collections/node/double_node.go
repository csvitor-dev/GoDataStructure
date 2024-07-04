package node

type DoubleNode[Type Number | string] struct {
	Data     Type
	Next     *DoubleNode[Type]
	Previous *DoubleNode[Type]
}

// NewDoubleNode: create new instance of DoubleNode[Type]
func NewDoubleNode[Type Number | string](data Type) *DoubleNode[Type] {
	node := &DoubleNode[Type]{
		Data:     data,
		Next:     nil,
		Previous: nil,
	}
	return node
}