# Golang

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

## 

