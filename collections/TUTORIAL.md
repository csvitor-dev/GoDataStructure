# Implementation Tutorial

> This tutorial will serve to detail the implementation of one of the data structures (in this case, the **Singly Linked List**), which is the foundation for many other data structures!

The main reason for this implementation (besides revisiting the same learnings I had in **Clang**) is to show how _pointer_ manipulation works!
So, don't forget to review [**this topic**](../README.md).

## Previous Details

Before starting to implement (I promise it will be the last topic before implementation), I will talk about two concepts that are present in every implementation: the `error` type (in addition to the `errors` package) and _generic_ types.

### The `error` Type

In **Go**, there are no exceptions and no `try-catch` statement blocks. So, how do you deal with _error handling_?
In a simpler and more modular way: just one `if` statement and one "instance" of `error`.

`error` type is defined as an _interface_:

```go
type error interface {
    Error() string
}
```

> This an important detail: it allows you to create custom errors!

This is the most common statement block in **Go**:

```go
err := doSomeThing() // returns an error
if err != nil {
    // handling the error
}
```

Or using an `if-init` statement block:

```go
if err := doSomeThing(); err != nil {
    // handling the error
}
```

As `error` is a _reference type_, its **zero value** ¹ (or _default_ value) is `nil`.

¹ **_Zero value_** is the value that a variable receives when declared, but not initialized.
**Primitive types** receive **0** (if they are numeric), **false** (if they are boolean), or an **empty string**.
Whereas **reference types** (errors, functions, pointers, and interfaces) receive `nil`.
As a result, structures that hold primitive or reference fields will receive their respective zero velues.

I understood what the `error` type is, but how can I make it more concise when handling it?
Similar to other programming languages that support the `Exception` type (like in **C#**), the `error` type can be specified with a _message_. That would be precisely the implementation of the `error` interface.

There are two ways to define custom error messages: with the `Errorf()` function (from the `fmt` package) or with the `New()` function (from the [`errors`](https://pkg.go.dev/errors) ² package) that work similarly.

² The `errors` package is much broader, focusing on error handling.

```go
package main

import (
    "fmt"
    "errors"
)

// with `fmt` pkg
func AnError() error {
    return fmt.Errorf("Generate an error")
}

// or with `errors` pkg
func OtherError() error {
    return errors.New("Generate another error")
}

func main() {
    err1, err2 := AnError(), OtherError()
    fmt.Println(err1, err2)
    // it will show `Generate an error Generate another error`
}
```

> It will be a very useful feature when we implement methods that need to manipulate our list, but it isn't a valid operation. Keep that in mind!

### Generic types

...

## Implementation (for real)

...

[//]: # "Falar sobre generics no Go e implementar a singly linked list de forma detalhada"
