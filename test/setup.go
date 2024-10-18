package test

import (
	"fmt"
	"testing"
)

type Input struct {
	Values []int
	Index int
}

type Output struct {
	Value int
	Err bool
}

// ErrorMessage: defines a generic message for failed tests
func ErrorMessage(t *testing.T, details string, expected any, received any) {
	message := fmt.Sprintf("%v | expected: %v, received: %v", details, expected, received)
	t.Error(message)
}
