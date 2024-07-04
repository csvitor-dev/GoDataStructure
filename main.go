package main

import (
	"data-structure/collections"
)

func main() {
	list := collections.NewDoublyLinkedList[string]()
	list.Add("C#")
	list.Add("Golang")
	list.Add("TypeScript")
	list.Add("PHP")
	list.Print(true)

	list.InsertAt(1, ".NET")
	list.Print(true)

	list.Delete()
	list.RemoveAt(2)

	list.Print(true)
}