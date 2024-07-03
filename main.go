package main

import (
	linkedlist "data-structure/linked_list"
)

func main() {
	list := linkedlist.NewSingleLinkedList[string]()
	list.Add("C#")
	list.Add("Golang")
	list.Add("TypeScript")

	list.InsertAt(1, "PHP")

	list.Delete()

	list.Print()
}