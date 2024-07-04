package node

type SingleNode[Type Number | string] struct {
	Data Type
	Next *SingleNode[Type]
}

// NewSingleNode: create new instance of SingleNode[Type]
func NewSingleNode[Type Number | string](data Type) *SingleNode[Type] {
	node := &SingleNode[Type]{
		Data: data,
		Next: nil,
	}
	return node
}