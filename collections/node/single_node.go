package node

type SingleNode[Type T] struct {
	Data Type
	Next *SingleNode[Type]
}

// NewSingleNode: create new instance of SingleNode[Type]
func NewSingleNode[Type T](data Type) *SingleNode[Type] {
	node := &SingleNode[Type]{
		Data: data,
		Next: nil,
	}
	return node
}