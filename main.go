package main

import (
	"data-structure/stack"
)

func main() {
	stack := stack.NewStack[string]()
	stack.Push("C#")
	stack.Push("Golang")
	stack.Push("TypeScript")

	stack.Print()
}