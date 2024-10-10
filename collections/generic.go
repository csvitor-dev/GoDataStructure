package collections

import (
	t "GoDataStructure/collections/types"
	"fmt"
)

type Generic[Type t.T] interface {
	find(int)   (Type, error)
	print()
	Length() int
}

func Print[Type t.T](collection Generic[Type]) {
	collection.print()
	fmt.Printf("\nLength: %v\n", collection.Length())
}

func FindOnIndex[Type t.T](collection Generic[Type], index int) (Type, error) {
	return collection.find(index)
}