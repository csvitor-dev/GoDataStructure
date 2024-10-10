package test

import "fmt"

// message: defines a generic message for failed tests
func message(details string, expected any, received any) string {
	return fmt.Sprintf("%v | expected: %v, received: %v", details, expected, received)
}