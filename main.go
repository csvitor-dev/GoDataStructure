package main

import (
	"data-structure/queue"
)

func main() {
	queue := queue.NewQueue[string]()
	queue.Enqueue("C#")
	queue.Enqueue("Golang")
	queue.Enqueue("TypeScript")
	queue.Dequeue()

	queue.Print()
}