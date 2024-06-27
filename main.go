package main

import (
	"fmt"
	"data-structure/queue"
)

func main() {
	queue := queue.NewQueue[string]()
	queue.Enqueue("C#")
	queue.Enqueue("Golang")
	queue.Enqueue("TypeScript")
	value, err := queue.Dequeue()
	queue.Print()

	fmt.Printf("%v %v\n", value, err)
}