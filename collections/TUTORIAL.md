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

Finally, I want to talk about _generic types_ in **Golang**!
An interesting detail is that this feature was added in version _1.18_ (with some reluctance from part of the community).

Before _generics_, a "strategy" to support various types in a function/method **was to use empty interfaces**. Agreeing with some **Go** developers, here in Brazil, this would be called POG _Programação Orientada a Gambiarra_ (trying to solve something in a unrecommended way, with potential **_side effects_**).

Even though the use of the _empty interface_ feature has its place (like in the `println()` function or in `fmt.Println()` and similar -- but now with the `any` **constraint**) it's not advisable to use it in more complex cases.

Note the following case:

```go
func Sum(i ...interface{}) (interface{}, error) {
    sum := i[0]
    for j := 1; j < len(i); j++ {
        switch sum.(type) { // type assertion
        case int:
            if v, ok := i[j].(int); ok {
                sum = sum.(int) + v
            } else {
                return nil, fmt.Errorf("Error: incompatible types")
            }
        case float64:
            if v, ok := i[j].(float64); ok {
                sum = sum.(float64) + v
            } else {
                return nil, fmt.Errorf("Error: incompatible types")
            }
        case float32:
            if v, ok := i[j].(float32); ok {
                sum = sum.(float32) + v
            } else {
                return nil, fmt.Errorf("Error: incompatible types")
            }
        default:
            return nil, fmt.Errorf("Error: unsupported type")
        }
    }
    return sum, nil
}
```

> It's a generic sum function with variable parameters (**_variadic function_**). You can see that the code becomes very rigid and it will be necessary to modify it to add new types, _meaning it's open for implementation_.

But with generic types, we get a cleaner result.

```go
func Sum[T int | float64 | float32](t ...T) T {
    var sum T

    for _, v := range t {
        sum += v
    }
    return sum
}
```

> We'll use this feature a lot in ourt list nodes, since it can handle **strings**, **integers**, and **floats**, for example.

## Implementation (for real)

Now, let's get to work!

[//]: # "implementar a singly linked list de forma detalhada"
