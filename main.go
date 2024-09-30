package main

import (
	c "GoDataStructure/collections"
)

func main() {
	c1 := c.NewDoublyLinkedList[string]()
	c2 := c.NewSinglyLinkedList[string]()
	c3 := c.NewQueue[string]()
	c4 := c.NewStack[string]()

	c1.Add("C#")
	c2.Add("TypeScript")
	c3.Enqueue("Golang")
	c4.Push("An Ordinary Software Engineering")

	c.Print(c1)
	c.Print(c2)
	c.Print(c3)
	c.Print(c4)
}