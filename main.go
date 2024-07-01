package main

import (
	"data-structure/queue"
	"fmt"
)

func main() {
	queue := queue.NewQueue[string]()
	queue.Enqueue("C#")
	queue.Enqueue("Golang")
	queue.Enqueue("TypeScript")

	queue.Print()
	fmt.Println(queue.Contains("F#"))
}