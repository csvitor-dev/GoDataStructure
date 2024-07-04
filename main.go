package main

import (
	"data-structure/collections"
)

func main() {
	list := collections.NewSingleLinkedList[string]()
	list.Add("C#")
	list.Add("Golang")
	list.Add("TypeScript")
	list.Add("PHP")

	list.Print()
	list.RemoveAt(3)

	list.Print()
}