package collections

import "fmt"

type Generic interface {
	print()
	Length() int
}

func Print(collection Generic) {
	collection.print()
	fmt.Printf("\nLength: %v\n", collection.Length())
}