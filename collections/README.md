# Talking About Familiar Concepts

> Now, we'll talk about essential concepts for both this projectand for **Go** itself!

For those who have studied a bit of the **C language**, they will understand the ideas more easily. But I'll try to explore each of these concepts sufficiently.

## Going Deeper

So, finally let's talk about the `struct` type and **_Go Way_** paradigm!

Basically, the definition of structs requires this format:

```go
package main

type NewStruct struct {
    // some definitions
}
```

Well, but to be more specific. Let's create an example that will be useful for the rest of this topic:

```go
package main

type Car struct {
    Brand string
    Model string
    year int
    motorState bool
}
```

> I know, I know... Some properties have the first letter capitalized, while others don't. But it's not just an aesthetic detail, I'll explain soon.

The above `struct` declaration is a classic example. Basically, the `struct` **Car** is a _blueprint_ for variable of this type.

For those who come from the **Clang**, this concept is already understood. But for those accustomed to the _OO paradigm_, understand that the `struct` "equates" to the `class` and the variable as the object, with this variable **_"being an instance of the `struct`"_**.

Note:

```go
// ...

func main() {
    car := Car{
        Brand: "Pagani",
        Model: "Zonda R",
        year: 2007,
        motorState: false,
    }

    println(car.Brand, car.Model, car.year, car.motorState)
    // will show `Pagani Zonda R 2007 false`
}
```

> This syntax reminds me a lot of the **C#** feature ([Object and Collection Initializers](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/classes-and-structs/object-and-collection-initializers)) which is literally the same way!

Sure, there's already quite a bit you can do. But there's something even more interesting that **Golang** provides: _methods_!

Methods in **Go** are functions specified for a `struct` type, which follow the following notation (exactly like the functions that have already been seen):

```go
func (car Car) run() {
    println(car.Brand, car.Model, "is running...")
}
```

So, the function (actually, _method of the Car `struct`_) will only be visible to instances of Car!

That's why this code:

```go
// ...
func main() {
    run()
}
```

It will throw a compile-time error:

```bash
# command-line-arguments
...: undefined: run
```

...

[//]: # "Falar sobre métodos, ponteiros e interfaces (Go Way), para aí sim avançar para o próximo tópico!"

## Singly Linked List Implementation

...

[//]: # "Falar sobre `error`, módulos, exportações e importações no Go e implementar a singly linked list de forma detalhada"
