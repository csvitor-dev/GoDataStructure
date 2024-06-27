package node

type Node[Type Number | string] struct {
	Data Type
	Next *Node[Type]
}

// NewNode: create new instance of Node[Type]
func NewNode[Type Number | string](data Type) *Node[Type] {
	node := &Node[Type]{
		Data: data,
		Next: nil,
	}
	return node
}