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

## Functions

<br>
<br>

### Basic function

```go
// Greet the user
func greet() {
	fmt.Println("Hello Go!")
}
```

- `func` keyword is used to define a function.
- Documentation for a function is added above the function definition.
- Functions must be defined outside other functions. No nested functions alowed.

<br>
<br>

### Function with parameters and return value

```go
func add(x int, y int) int {
	return x + y
}
```

- We must specify the type for each parameter.
- Return type should be specified after the parameters and before `{}`.

<br>
<br>

### Multiple return values

```go
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return x / y, nil
}

result, err := divide(4, 2)
```

- When returning multiple values, the return types should be enclosed in parenthesis `()`.
- Multiple return values are commonly used for returning result + error.

<br>
<br>

### Named return values

- We can name the return variables in the function signature.
- These variables are already declared. We just have to reassign them.

```go
func subtract(x, y int) (res int) {
	res = x - y  // note `=` is used and not `:=`
	return  // empty return
}
```

<br>
<br>

### Variadic functions

- Just like `*args` in Python, use `...type` to accept a variable number of arguments.

```go
// Add all nums and return the total
func addAll(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return total
}

result := addAll(1, 2, 5)
```

<br>
<br>
<br>
<br>
<br>

## Conditional statements

<br>
<br>

### if-else statement

```go
if x > 10 {
    fmt.Println("Good")
} else if x > 5 {
    fmt.Println("OK")
} else {
    fmt.Println("Bad")
}
```

- There are no parentheses around the condition.
- Curly brackets are mandatory even for one-liners.

<br>
<br>

### Short statement in `if`

- Go allows a **short statement** after the `if` keyword and before the condition.
- This variable is scoped only within the if-else block.

```go
if x := 10; x > 0 {
    fmt.Println(x, "Positive")
} else {
    fmt.Println(x, "Negative")
}
```

<br>
<br>

### Note

- Go does not have a ternary operator `(cond ? a : b)`. Use standard `if` block instead.

<br>
<br>
<br>
<br>
<br>

## Logical operators

<br>
<br>

### Logical AND (`&&`)

```go
a := true
b := false

fmt.Println(a && b)  // false
```

- Both operands should be true.
- Short-circuits: If the first condition is false, second one is not evaluated.

<br>
<br>

### Logical OR (`||`)

```go
a := true
b := false

fmt.Println(a || b)  // true
```

- Any one operand should be true.
- Short-circuits: If the first condition is true, second one is not evaluated.

<br>
<br>

### Logical NOT (`!`)

```go
a := true

fmt.Println(!a)  // false
```

- Inverts the boolean value.

<br>
<br>

### Grouping conditions

- Use parentheses for clear grouping of logical operators.

```go
if (score >= 90 && score <= 100) || bonus {
    fmt.Println("Eligible for reward")
}
```

<br>
<br>

### Note

- Boolean expressions must be **explicit**.
- Unlike Python, Go does not allow implicit boolean coercion.

```go
x := 5

// Incorrect: non-boolean used as condition
if x {}

// Correct: boolean value used as condition
if x != 0 {}
```

<br>
<br>
<br>
<br>
<br>

## Switch statement

<br>
<br>

### Standard switch statement

```go
day := "Sunday"

switch day {
case "Monday":
    fmt.Println("Weekday")
case "Friday":
    fmt.Println("Almost Weekend")
default:
    fmt.Println("No Idea")
}
```

- Go evaluates the expression (`day`) and runs the first matching `case`.
- If no cases match, it runs the `default` block (if provided).
- No `break` needed. Go automatically breaks after a matching case.

<br>
<br>

### Multiple values matching

- We can even match multiple values in a single case.
- Add these values separated by a comma.

```go
day := "Tuesday"

switch day {
case "Monday", "Tuesday":  // multiple values matching
    fmt.Println("Weekday")
case "Friday":
    fmt.Println("Almost Weekend")
// optional default case
}
```

<br>
<br>

### Switch without expression

- If you omit the expression in a `switch` statement, it acts like an `if-else` statement.
- Conditions are matched from top to bottom.
- Only the first matching condition runs.

```go
marks := 85

switch {
case marks >= 90:
    fmt.Println("Grade A")
case marks >= 80:
    fmt.Println("Grade B")
case marks >= 70:
    fmt.Println("Grade C")
}
```

<br>
<br>

### `fallthrough` keyword

- Unlike C/C++, Go does not fall through the cases by default.
- We can force it using the `fallthrough` keyword.
- Use `fallthrough` carefully, it ignores the next case's condition.

```go
num := 1

switch num {
case 1:
    fmt.Println("One") // This is logged
    fallthrough
case 2:
    fmt.Println("Two") // This is logged
case 3:
    fmt.Println("Three")
}
```

<br>
<br>

### Short statement in `switch`

- Just like in `if` statement, we can declare and use a variable inside the `switch` statement.

```go
switch day := "Monday"; day {
case "Monday":
    fmt.Println("Weekday")
case "Friday":
    fmt.Println("Almost Weekend")
}
```

<br>
<br>
<br>
<br>
<br>

## 
