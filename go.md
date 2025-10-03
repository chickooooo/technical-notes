# Golang

<br>
<br>
<br>
<br>
<br>

## Declaring variables

- In golang, variables are statically typed. That means each variable has a fixed datatype.
- There are 3 ways of declaring variables in Go:

<br>
<br>

### 1. Explicit type

```go
// First declaration then initialization
var name string
name = "Alice"

// Declaration and initialization together
var age int = 20
```

<br>

**Zero values**

- Zero values are the default values assigned to variables when they are declared without an explicit initialization.
- Here are some common zero values by type:

| Type                                 | Zero Value                                                         |
| ------------------------------------ | ------------------------------------------------------------------ |
| `int`, `int32`, `int64`, `uint` etc. | `0`                                                                |
| `float32`, `float64`                 | `0.0`                                                              |
| `bool`                               | `false`                                                            |
| `string`                             | `""` (empty string)                                                |
| Pointers                             | `nil`                                                              |
| Slices                               | `nil`                                                              |
| Maps                                 | `nil`                                                              |
| Channels                             | `nil`                                                              |
| Interfaces                           | `nil`                                                              |
| Functions                            | `nil`                                                              |
| Arrays                               | All elements are set to the zero value of the array's element type |
| Structs                              | Each field is set to its respective zero value                     |

<br>
<br>

### 2. Inplicit type

- Go can automatically infer the type if we provide a value.

```go
var isActive = true  // type is bool
```

<br>
<br>

### 3. Short variable declaration

- This is a Go-specific feature and is only valid inside functions (not at package level).
- This is the most common way to declare and initialise variables inside a function.

```go
message := "Hello Go!"
```

<br>
<br>

### 4. Declaring multiple variables

- We can also declare multiple variables at once.
- While declaring, if 2 or more variables have the same datatype, we can omit all the previous types and add type only once at the end.

```go
var a, b int = 10, 20
firstName, lastName := "Alice", "Johnson"
```

<br>
<br>

### 5. Block declaration

- We can also declare multiple variables in a single block.
- These variables are also known as grouped variables.

```go
var (
    a         = 10
    b         = "Alice"
    c bool    = true
    d float64 = 3.1415
)
```

<br>
<br>

### 6. Declaring constants

- Constants are immutable values that are known at compile time and cannot be changed at runtime.

```go
const Pi float32 = 3.1415
```

<br>
<br>

### Note

- In Go, we cannot declare a variable and leave it unused. This is causes a **compile-time** error.

<br>
<br>
<br>
<br>
<br>

## Loops

- Go has only one looping keyword i.e. `for`.
- We can use this single keyword to behave like a `for` loop, `while` loop, `enumerate`, etc.

<br>
<br>

### Classic loop

- Standard `for` loop to iterate over a range of numbers.

```go
// Prints 0 1 2 3 4
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

<br>
<br>

### Range loop

- Range `for` loop to iterate over a range of numbers.

```go
// Prints 0 1 2 3 4
for i := range 5 {
    fmt.Println(i)
}
```

<br>
<br>

### While-style loop

- We can omit the init and post parts, and just use the condition.

```go
// Prints 1 2 3 4 5
i := 1
for i < 6 {
    fmt.Println(i)
    i++
}
```

<br>
<br>

### Infinite loop

- To create an infinite loop, omit everything.
- Use `break` keyword to exit the loop.

```go
import "math/rand"

var num int
for {
    num = rand.Intn(11)
    fmt.Println(num)
    if num == 10 {
        break
    }
}
```

<br>
<br>

### Looping over collections

- We can use `range` keyword to loop over collections liked array, string, slice, etc.
- The `range` keyword provides access to both the index and the value.
- If we don't require the index or the value, we can use `_` to ignore it.

```go
var nums = []int{1, 3, 5, 7}

// Prints 0 1  1 3  2 5  3 7
for index, num := range nums {
    fmt.Println(index, num)
}
```

<br>

- When looping over a string, `range` provides access to the index and character unicode.
- We can use `%c` get back the character value from that unicode.

```go
var s = "abcd"

// Prints a b c d
for _, unicode := range s {
    fmt.Printf("Char: %c\n", unicode)
}
```

<br>
<br>

### Breaking and Continuing

- Use `break` keyword to exit the loop.
- Use `continue` keyword to skip to the next iteration.

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue  // skip 3
    }
    if i == 5 {
        break  // exit loop when i == 5
    }
    fmt.Println(i)
}
```

<br>
<br>
<br>
<br>
<br>

## 
