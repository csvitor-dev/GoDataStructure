package main

import (
	"GoDataStructure/collections"
	"fmt"
)

func main() {

	doublyLinkedList := collections.NewDoublyLinkedList[string]()
	doublyLinkedList.InsertAt(0, "C#")
	doublyLinkedList.InsertAt(0, ".NET")
	doublyLinkedList.InsertAt(2, "Golang")
	doublyLinkedList.InsertAt(2, "TypeScript")
	
	value, err := doublyLinkedList.RemoveAt(2)
	doublyLinkedList.Print(true)
	fmt.Printf("%v, %v\n", value, err)
}