# Exploring Data Structure in [_Golang_](https://go.dev)

> Experimenting with **Go** features to create data structures (remember **Clang**?), such as generic types, methods and modules.

## Get Started

<p align="center">
<img width="300" src="./assets/go.png" alt="Go programming language logo"/>
</p>

The **Go** programming language was designed at **_Google_** in 2007 by _Robert Griesemer_ (he worked on V8 engine), _Rob Pike_ and _Ken Thompson_ (both worked on the UTF-8). In 2009, it became open-source.

**Golang** is static and strongly typed, compiled and with a mix of OO and Structured paradigms (called the **_Go Way_** by the **Go** community). Futhermore, it has native support for _concurrency_ and _parallelism_ (through **_go routines_**), and is also memory-safe with the presence of **Garbage Collection**. And Also with a robust standard library that offers several features, from network protocols ([`net/http`](https://pkg.go.dev/net/http) pkg, for example), JSON serialization ([`encoding/json`](https://pkg.go.dev/encoding/json) pkg) and even contexts ([`context`](https://pkg.go.dev/context) pkg).

Finally, **Go** is syntactically similar to **Clang** (in addition to inheriting many concepts from it, such as _pointers_), but much more orthogonal, bringing a relatively smaller learning curve. So, whoever enjoys learning a programming language that has low-level resources, **Go** can be a worthwhile learning experience!

> For more:
>
> - [Go (programming language)](<https://en.wikipedia.org/wiki/Go_(programming_language)>)
> - [Documentation - The Go Programming Language](https://go.dev/doc/)

**Golang**'s mascot is called the _Gopher_, just out of curiosity:

<p align="center">
    <img 
        width="265"
        src="./assets/gopher.png" 
        alt="Gopher, mascot of the Golang"
    />
</p>

## Elaboration

Well, after defining what the **Go** programming language is and what it offers, we can move on to the goals of this project, which are:

- learn the basic concepts and syntax of **Go**;
- build some data structures and their algorithms (insertions, removals, etc.), then we will learn additionally about DSA (Data Structures and Algorithms, one of the main foundations of _Computer Science_).

Okay, let's started!

### A Bit of the Basics

> Important! I won't dwell on downloading the _SDK_ or programming basics (such as data types, for example), assuming that you, the reader, already have this foundation. My intention in discussing syntax is because of some details that **Golang** allows and that I haven't seen in other programming languages I've studied.

So, let's start with declaring variables. To achieve this, there are two ways to define variables in **Go**: by declaration and assignment or by initialization. Declaration and assignment follows this format: `var identifier type = value`. While initialization allows the compiler to infer the type, with the syntax: `identifier := value`.

```go
// all Go source code files start with package declaration
package main

// declaration the 'something' variable
var something int
// assign the value to it
something = 10

// or, simplifying...
var something int = 10

// the declaration below is possible because the compiler also infers its type
var somethingElse = "allowed"

// initialization the 'other' variable
other := 10
```

> Notice that the type is after the identifier!
> About the packages, we'll talk about them later.

It's important to make it clear that the initialization statement needs to be within the scope of a function, otherwise it will throw a compile-time error:

```bash
ERROR!
# command-line-arguments
...: syntax error: non-declaration statement outside function body
```

**Go** offers another way of declaring variables, which is `const`, and it is **immutable**. Similar to declaring a regular variable, but using `const` instead of `var`: `const identifier type`.

```go
// with `const`, only the initialization statement is allowed, as it is technically expected
const something = 23 // the compiler will infer the type
const somethingElse string = "allowed" // the type was specified

const other string // the declaration generates an error
other = "error"
```

> As expected, declaration statements with `:=` are specific to **mutable variables**!

Futhermore, it's possible to declare multiple variables (or constants) using the syntax:

```go
var (
    // declare your variables
)

const (
    // declare your constants
)
```

> It's interesting to group only the declaration statements that are about the same context or the same type, avoiding all of them together. Because it would be difficult to understand and track your code, in my opinion.

Good, but the functions? Well, it's similar to declaring variables, just changing a few keywords.

The functions follow the following syntax:

```go
package main

func sum(a int, b int) int {
    return a + b
}

// ...

func main() {
    result := sum(3, 5) // returns 8
    println(result)
}
```

> As you can see, you define what it is (a function, with the keyword `func`), its formal parameters with the same notation as the variables and the type of return.
> If the function doesn't have a return, leave the return type empty, like the `main` function, which by the way, is the entry point for **Go** applications.

**Go** also allows multiple returns, just separate each type of return with commas inside parentheses, like a _tuple_.

```go
package main

func Operations(a, b int) (int, int, int) {
    return a + b, a - b, a * b
}

// ...

func main() {
    addition, subtraction, multiplication := Operations(2, 2)
    println(addition, subtraction, multiplication) // '4 0 4' will be displayed
}
```

> Note that the formal parameters of the `Operations` function have only one type definition, because **Go** allows parameters of the same type to be identified with a single type definition.

> The main difference between **Go** and **Clang** is that **Go** allows functions to return multiple values, which is different from **C#** (my main stack for development, which is part of the **.NET** environment), for example, where it is necessary to **wrap multiple values in a _tuple_ to return them**.
> I say this because **Go** will require values to be individually assigned to each variable as done above. Whereas in **C#**, there are two possible ways: deconstructing the _tuple_ or assigning the returned _tuple_ to a variable, both are allowed.

See the following examples:

```cs
using System;

public class CSharpExample
{
    public static void Main(string[] args)
    {
        var tuple = Operations(2, 2);
        var (addition, subtraction, multiplication) = Operations(2, 2);

        Console.WriteLine(tuple); // it will show `(4, 0, 4)`
        Console.WriteLine($"{addition} {subtraction} {multiplication}"); // it will show `4 0 4`
    }

    public static (int, int, int) Operations(int a, int b)
        => (a + b, a - b, a * b);
}
```

While in **Golang**, the assignment statement will generate compile-time error:

```go
package main

func Operations(a, b int) (int, int, int) {
    return a + b, a - b, a * b
}

// ...

func main() {
    tuple := Operations(2, 2) // invalid assignment
    println(tuple)
}
```

```bash
# command-line-arguments
...: assignment mismatch: 1 variable but Operations returns 3 values
```

Great, but if not all values will be used in the code, what to do?

Well, both **Go** and **C#** offer a feature that allows _discarding some return values_: the [_blank identifier_](https://www.geeksforgeeks.org/what-is-blank-identifierunderscore-in-golang/) in **Golang**, and the [_discard variable_](https://learn.microsoft.com/en-us/dotnet/csharp/fundamentals/functional/discards) in **C#**.

To illustrate, let's go back to the previous example. The `Operations` function returns three values, but let's imagine that only the addition value is useful. So, let's discard the other values:

In **C#**, we have the following code:

```cs
// same previous code

public static void Main(string[] args)
{
    var tuple = Operations(2, 2);
    var (addition, _, _) = Operations(2, 2);

    Console.WriteLine(tuple); // it will show `(4, 0, 4)`
    Console.WriteLine(addition); // now, it will only show `4` because other values do not exist in current context
}

// same previous code
```

In **Go**, it would be like this:

```go
// same previous code

func main() {
    addition, _, _ := Operations(2, 2)
    println(addition) // only '4' will be displayed because of the same reason
}
```

> I talk about it because **Go** will throw a compile-time error if there are variables (or _packages imports_) that have been declared but are not being used. Fortunately, the **Go** compiler is very strict about memory usage!

### About Packages

Phew, we finally got to _packages_!
Just like _namespaces_ in **C#** or even _packages_ in **Java**, **Golang** provides a high-level feature for organizing and segmenting **_modules_** into "logical directories", so that there can't be two packages defined in the same directory, only one.

Besides being used for organizing code into modules, **it is through packages that we can work with _exports_ and _imports_**, with the package name allowing us to access parts of the module.

> But, there's a subtle detail: you may have noticed that I declared functions with _the first letter capitalized_, right? But the functions `main` and `println` aren't declared like that, why?
>
> In **Go**, the **_visibility_** of any type of data (whether variables, constants, functions, etc.) is determined by this characteristic: if _the first letter of the identifier is lowercase_, then it will only be visible at the package level (**meaning it's not exported**), **otherwise it will be public**.

For example, in a folder with any name (but it's interesting that the folder name is the same as the package), let's create a module called `math`, so the files in this folder will have instruction `package math`.

First, let's create the first function of package: `sum`.

```go
// in `math/sum.go`
package math

func sum(a, b int) int {
    return a + b
}
```

Now, at the _application entry point_ (thus, the `main` function in the `main` package), let's use the first function from our first package:

```go
// in `main.go`
package main

// using the `math` package
import "your-module/math"

func main() {
    // all declarations made in the `math` package are exposed by it
    println(math.sum(2, 5))
}
```

But, as expected, this will generate a compile-time error:

```bash
# command-line-arguments
...: undefined: math.sum
```

Remember that only declarations in **which the identifier has the first letter capitalized will be exportable**!

Fixing:

```go
// in `math/sum.go`
package math

func Sum(a, b int) int {
    return a + b
}
```

Now it will work! And that's why every package you use (or create) will have its members with an uppercase initial letter.

The package I will use the most going forward is [`fmt`](https://pkg.go.dev/fmt), which is used string and terminal output, and will replace the `println` function in the next examples.

### About Modules

You must have noticed that in the previous example the instruction `import "your-module/..."` was used.
That's because your module is the module declared to aggregate the entire application, but how you declare a module? Through the **Golang _CLI_**!

You need to run the following command in your terminal to initialize your project module (this should be the first step of your project), and this way, the import instruction will work as expected!

```bash
go mod init your-module
```

The above command creates a `go.mod` file that serves the same purpose as the `package.json` in a **Node.js** application, for example.

## Project Tree

> ```bash
> GoDataStructure/
>   ├─── assets/
>   └─── collections/
>         └─── node/
> ```
>
> 15 files and 3 folders.

**To continue, [_click here_](./collections/README.md)!**
