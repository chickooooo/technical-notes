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
- Useful for range-based logic.

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

## Array

- In Go, an array is a **fixed size**, **ordered** collection of elements of the **same type**.
- Arrays are mutable, i.e. we can update an element within an array.
- They are fixed size. We cannot add a new element or remove an existing element from an array.

<br>

### Key points

- The size of an array is part of its type. That means `[3]int` and `[4]int` are two completely different types.
- Arrays are **value types** and not reference types like Python lists.
- When passed to functions, arrays are copied if not passed by reference.

<br>
<br>

### Declaring and initialising array

```go
var arr [5]int // array of 5 integers
```

<br>

#### Default values

```go
var a [3]int // [0 0 0]
```

- If not initialised while declaration, array will hold **zero values**.

<br>

#### With values

```go
b := [3]int{1, 2, 3}
```

<br>

#### Let Go infer the length

```go
c := [...]int{10, 20, 30} // length = 3
```

<br>
<br>

### Accessing and modifying elements

```go
b[0] = 100
fmt.Println(b[0]) // Output: 100
```

- Out-of-range access will cause a **compile-time error** or a **runtime error**.

<br>
<br>

### Looping over array

```go
// using standard for loop
for i := 0; i < len(b); i++ {
    fmt.Println(b[i])
}
```

<br>

```go
// using range for loop
for index, value := range b {
    fmt.Println(index, value)
}
```

<br>
<br>

### Copying arrays

- Arrays are **value types**. Assigning a variable holding an array to another variable will create a new copy of that array.

```go
a := [3]int{1, 2, 3}
b := a
b[0] = 100

fmt.Println(a) // [1 2 3]
fmt.Println(b) // [100 2 3]
```

<br>
<br>

### Array length

```go
arr := [4]string{"a", "b", "c", "d"}
fmt.Println(len(arr)) // 4
```

<br>
<br>
<br>
<br>
<br>

## Slice

- In Go, a slice is a **variable sized**, flexible view of an array.
- Under the hood, a slice contains these 3 things:
    - A pointer to an underlying array
    - A length
    - A capacity
- Slices are of **reference type** in nature.

<br>
<br>

### Creating a slice

<br>

#### Empty slice

```go
var mySlice = []int{}
fmt.Println(mySlice) // []

mySlice = append(mySlice, 1, 2, 3)
fmt.Println(mySlice) // [1 2 3]
```

<br>

#### Direct initialization

```go
mySlice := []int{4, 5, 6}
```

<br>

#### From an array

- We can create a slice from an array by **slicing** that array.
- This newly created slice is just a view of the original array. That means, any changes made to the slice will be reflected in the original array as well.

```go
myArray := [5]int{1, 2, 3, 4, 5}
mySlice := myArray[2:5]

fmt.Println(myArray) // [1 2 3 4 5]
fmt.Println(mySlice) // [3 4 5]

mySlice[0] = 100
fmt.Println(mySlice) // [100 4 5]
fmt.Println(myArray) // [1 2 100 4 5]
```

<br>

#### Using make

- A slice created using `make()` is filled with zero values.
- This slice will be of `length` size.

```go
mySlice := make([]int, 3) // length 3, capacity 3
fmt.Println(mySlice, len(mySlice), cap(mySlice)) // [0 0 0] 3 3

mySlice2 := make([]int, 2, 5) // length 2, capacity 5
fmt.Println(mySlice2, len(mySlice2), cap(mySlice2)) // [0 0] 2 5
```

<br>
<br>

### Indexing and slicing

- We can access a single element or a slice of elements using the index.
- Slicing a slice or an array creates just a new view. Any changes made to this slice will be reflected in the original as well.

```go
mySlice := []int{10, 20, 30}
mySlice2 := mySlice[1:]

fmt.Println(mySlice[2]) // 30

fmt.Println(mySlice)  // [10 20 30]
fmt.Println(mySlice2) // [20 30]

mySlice2[0] = 100
fmt.Println(mySlice2) // [100 30]
fmt.Println(mySlice)  // [10 100 30]
```

<br>
<br>

### Slice is a reference type

- If two slices share the same underlying array, changes to one will affect the other.

```go
a := []int{1, 2, 3}
b := a
b[0] = 99
fmt.Println(a) // [99 2 3] → changed!
```

<br>
<br>

### Copying slices

- Use built-in `copy()` function to make a new copy of the slice.
- The destination slice should be of same length as the source slice to copy all elements.

```go
mySlice := []int{10, 20, 30}
mySlice2 := make([]int, len(mySlice))

fmt.Println(mySlice2) // [0 0 0]

copy(mySlice2, mySlice)
fmt.Println(mySlice2) // [10 20 30]
```

<br>
<br>

### Adding and removing elements

- Use `append()` method to add one or more elements to the slice.
- Use **Ellipsis operator** `mySlice...` to copy elements of one slice into another.

```go
mySlice := []int{10}

mySlice = append(mySlice, 20, 30)
fmt.Println(mySlice) // [10 20 30]

mySlice2 := []int{40, 50}

mySlice = append(mySlice, mySlice2...)
fmt.Println(mySlice) // [10 20 30 40 50]
```

<br>

- To pop the last element, we can using slicing to slice off the last element.
- Similarly, we can use slicing to slice off any number of elements from any place.

```go
mySlice := []int{10, 20, 30, 40, 50}

mySlice = mySlice[:len(mySlice)-1]
fmt.Println(mySlice) // [10 20 30 40]
```

<br>
<br>

### Nil vs empty slice

- A slice created without initialization has no underlying array, hence it is `nil`.
- A slice initialized with an empty array is **not nil**.

```go
var mySlice []int      // nil slice (no underlying array)
var mySlice2 = []int{} // empty slice (0-length, but allocated)

fmt.Println(mySlice == nil)  // true
fmt.Println(mySlice2 == nil) // false
```

<br>
<br>
<br>
<br>
<br>

## Map

- In Go, a `map` is a data structure that holds **key-value** pairs. Similar to Python `dict`.
- Maps are of **reference type** in nature.

<br>
<br>

### Creating a map

#### Declaration only

```go
var hashmap map[string]int

fmt.Println(hashmap == nil) // true
```

- When a map is declared without initialization, its value is `nil`.
- Go doesn't allocate memory to maps that are not initialized.

<br>

#### Using make

```go
hashmap := make(map[string]int)

fmt.Println(hashmap == nil) // false
fmt.Println(hashmap)        // map[]
```

- A map created using `make` is initialised at the time of declaration.
- Such map is empty and contains no key-value pairs.

<br>

#### Direct initialization

```go
hashmap := map[string]int{}

fmt.Println(hashmap == nil) // false
fmt.Println(hashmap)        // map[]

hashmap2 := map[string]int{
    "alice": 1,
    "bob":   2,
}

fmt.Println(hashmap2) // map[alice:1 bob:2]
```

- Initialization can be done with one or more key-value pairs, or no key-value pairs.

<br>
<br>

### Adding and updating key-value pair

```go
hashmap := map[string]int{}

hashmap["Alice"] = 10
hashmap["Bob"] = 20

fmt.Println(len(hashmap)) // 2
fmt.Println(hashmap)      // map[Alice:10 Bob:20]
```

- Syntax for adding & updating a key-value pair is same as Python.
- If the key already exists, it is updated, otherwise a new key-value pair is created.
- Also, `len()` function can be used to get the length of the map.

<br>
<br>

### Accessing data

#### Direct access

```go
hashmap := map[string]int{
    "Alice": 10,
}

fmt.Println(hashmap["Alice"]) // 10
fmt.Println(hashmap["Bob"])   // 0
```

- Syntax for accessing a value for the given key is same as Python.
- If the key is not present in the map, Go **will not** throw an error, instead it will return a zero value.

<br>

#### Checking if exists

```go
hashmap := map[string]int{"Alice": 10}
fmt.Println(hashmap["Alice"]) // 10

value, exists := hashmap["Bob"]

fmt.Println(value, exists) // 0, false
```

- We can use `value, exists := myMap[key]` syntax to check if the key exists in the map.

<br>
<br>

### Deleting key-value pair

```go
hashmap := map[string]int{
    "Alice": 10,
    "Bob":   20,
}

delete(hashmap, "Alice")
fmt.Println(hashmap) // map[Bob:20]
```

- Use built-in `delete()` method to delete a key from the map.
- If the map is `nil` or if the key does not exist, delete does nothing.

<br>
<br>

### Iterating a map

- We can use `for range` loop to iterate over a map.

#### Iterating over keys and values

```go
hashmap := map[string]int{
    "Alice": 10,
    "Bob":   20,
}

for key, value := range hashmap {
    fmt.Println(key, value)
}
```

<br>

#### Iterating over keys only

```go
hashmap := map[string]int{
    "Alice": 10,
    "Bob":   20,
}

for key := range hashmap {
    fmt.Println(key)
}
```

<br>

#### Iterating over values only

```go
hashmap := map[string]int{
    "Alice": 10,
    "Bob":   20,
}

for _, value := range hashmap {
    fmt.Println(value)
}
```

<br>
<br>
<br>
<br>
<br>

## Pointer

- A pointer is a variable that stores the **memory address** of another variable.
- We use `&x` syntax to get the memory address of a variable `x`.
- And `*p` syntax to get the value, present at the memory address, where pointer `p` is pointing to.

```go
x := 10 // integer variable
p := &x // 'p' is a pointer to 'x' (using &)

fmt.Println("Value of x:", x)                // 10
fmt.Println("Value of p (address of x):", p) // 0xc0000100e0
fmt.Println("Value pointed to by p:", *p)    // 10 (dereferencing)
```

<br>
<br>

### Pass by reference

- When passing variables to a function, Go uses **pass by value**, by default.
- We can make use of pointers to **pass by reference**.

```go
func modify(x *int) {
	*x = 100
}

x := 10
modify(&x)
fmt.Println(x) // 100
```

<br>
<br>

### Declaring a pointer

#### Zero value of a pointer

- If we declare a pointer without initialization, its value is `nil`.

```go
var p *string
fmt.Println(p == nil) // true
```

<br>

#### Direct initialization

```go
x := 10
p := &x
```

<br>
<br>
<br>
<br>
<br>

## Structure

- A Structure is defined using the `struct` keyword.
- It is a **Composite data type** that groups together zero or more fields with varying types.
- Think of it as a Python `class` without methods, just data.

<br>
<br>

### Creating a struct

```go
type Person struct {
	Name string
	Age  int
}
```

<br>
<br>

### Using struct

#### Declaration only

- If we create a struct without assigning values, all its fields get **zero values**.

```go
var p1 Person
fmt.Println(p1) // { 0}
```

<br>

#### Direct initialization

```go
alice := Person{
    Name: "Alice",
    Age:  30,
}

fmt.Println(alice) // {Alice 30}
```

<br>
<br>

### Accessing and modifying fields

```go
alice := Person{
    Name: "Alice",
    Age:  30,
}

alice.Age = 40

fmt.Println(alice.Name) // Alice
fmt.Println(alice)      // {Alice 40}
```

<br>
<br>

### Pointer to struct

- We can use pointers to `struct`, to avoid copying large data structures.
- When accessing / updating fields, we can use `pointer.Name` syntax. No need for `(*pointer).Name`.

```go
alice := Person{"Alice", 30}

pointer := &alice
pointer.Name = "Bob"

fmt.Println(alice) // {Bob 30}
```

<br>
<br>

### Nested struct

- One structure can contain another structure within itself.

```go
type Address struct {
    Street string
    Zip    string
}

type Employee struct {
    Name    string
    Address Address
}


e := Employee{
    Name: "Eve",
    Address: Address{
        Street: "123 Main St",
        Zip:    "12345",
    },
}
```

<br>
<br>
<br>
<br>
<br>

## Method

- In Go, methods are functions that are associated with a **type**, usually a `struct`.
- A **Go Method** is equivalent of a method of a Python class.

<br>
<br>

### Syntax

```go
func (receiver Type) methodName(params) returnType {
    // method body
}
```

- `receiver` is like `self` in Python, but explicitly named.
- Receiver and it's type is placed before the method name.
- The receiver associates the method with a type.

<br>
<br>

### Creating a method

#### Value receiver

```go
type Person struct {
	Name string
	Age  int
}

func (p Person) greet() {
	fmt.Printf("My name is %s and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
	alice := Person{"Alice", 25}
	alice.greet() // calling a method
}
```

- The method `greet()` is bound to type `Person`.
- The receiver `p Person` is passed by value (copied).

<br>

#### Pointer receiver

- We can also pass a **pointer to the receiver** in a method.
- This allows us to modify the receiver inside the method.

```go
func (p *Person) incrementAge() {
	p.Age++
}

func main() {
	alice := Person{"Alice", 25}
	alice.incrementAge()
	alice.greet() // My name is Alice and I'm 26 years old.
}
```

<br>
<br>

### Method for non-struct type

- Go allows defining methods for any **user-defined types**, not only `struct`.

```go
type Celsius float64

func (c Celsius) toFahrenheit() float64 {
    return float64(c)*1.8 + 32
}
```

<br>
<br>
<br>
<br>
<br>

## Interface

- An **Interface** in Go, defines a set of method signatures.
- Any type that implements all the methods of an interface is said to satisfy that interface.
- Go has **implicit** interface implementation behaviour. Unlike other languages like Python or Java where we have to explicitly state that a class implements an interface.
- Go uses **duck typing**: "If it walks like a duck and quacks like a duck, then its a duck."

<br>
<br>

### Creating an interface

```go
type Speaker interface {
	Speak() string
}
```

- Any type that has a `Speak() string` method, automatically implements the `Speaker` interface.

<br>
<br>

### Implementing an interface

```go
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}
```

- Both `Dog` & `Cat` structures now implement `Speaker` interface as they have a `Speak()` method.

<br>
<br>

### Using an interface

```go
// Function that accepts type Speaker
func makeNoise(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	var speaker Speaker

    // Speaker can be a dog
	speaker = Dog{}
	makeNoise(speaker) // Woof!

    // Speaker can be a cat
	speaker = Cat{}
	makeNoise(speaker) // Meow!
}
```

- An interface is a **type**, that means we can use it as a paramter, return value, field type, etc.

<br>
<br>

### Empty interface

- A special type `interface{}` can hold any type. Because every type has zero or more methods.

```go
var anything interface{}

anything = 5
anything = "hello"
anything = true
```

<br>
<br>

### Interface composition

- We can **compose** interfaces by using other interfaces into an interface.

```go
type Speaker interface{ Speak() string }

type Feelings interface{ Happy() string }

// Composition
type Animal interface {
	Speaker
	Feelings
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

func (d Dog) Happy() string { return "Wag tail!" }

func main() {
	var animal Animal
	animal = Dog{}

	fmt.Println(animal.Speak()) // Woof!
	fmt.Println(animal.Happy()) // Wag tail!
}
```

<br>
<br>
<br>
<br>
<br>

## 

<br>
<br>
<br>
<br>
<br>
