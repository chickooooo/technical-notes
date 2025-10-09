# `fmt` Package

<br>
<br>
<br>
<br>
<br>

## Printing output

### `Print`

- Prints to the standard output **without a newline** ending.
- Spaces are added between operands when neither is a string.

```go
a := []int{10, 20, 30}

// [10 20 30]
fmt.Print(a)

// [10 20 30] [1 2]
fmt.Print(a, []int{1, 2})

// Hello
fmt.Print("Hello")

// HelloWorld
fmt.Print("Hello", "World")
```

<br>
<br>

### `Println`

- Prints to the standard output **with a newline** ending.
- Spaces are always added between operands.

```go
a := []int{10, 20, 30}

// [10 20 30]
fmt.Println(a)

// [10 20 30] [1 2]
fmt.Println(a, []int{1, 2})

// Hello
fmt.Println("Hello")

// Hello World
fmt.Println("Hello", "World")
```

<br>
<br>

### `Printf`

- Formats the string and prints to the standard output.
- Newline (`\n`) needs to be manually appended at the end, if needed.

```go
name, age, active := "Alice", 20, true

// Name: Alice, Age: 20, Status: true
fmt.Printf("Name: %s, Age: %d, Status: %t\n", name, age, active)


data := struct{ Name string }{"Bob"}

// Data: {Bob}
fmt.Printf("Data: %v\n", data)

// Data: {Name:Bob}
fmt.Printf("Data: %+v\n", data)

// Data: struct { Name string }{Name:"Bob"}
fmt.Printf("Data: %#v\n", data)
```

<br>
<br>

### Format specifiers

| Verb   | Description                              | Example Input      | Output Example          |
| ------ | ---------------------------------------- | ------------------ | ----------------------- |
| `%s`   | String                                   | `"Alice"`          | `Alice`                 |
| `%q`   | Quoted string                            | `"Alice"`          | `"Alice"`               |
| `%d`   | Decimal integer                          | `42`               | `42`                    |
| `%b`   | Binary                                   | `42`               | `101010`                |
| `%o`   | Octal                                    | `42`               | `52`                    |
| `%x`   | Hexadecimal (lowercase)                  | `42`               | `2a`                    |
| `%X`   | Hexadecimal (uppercase)                  | `42`               | `2A`                    |
| `%f`   | Decimal floating-point                   | `3.14159`          | `3.141590`              |
| `%.2f` | Floating-point with 2 decimal places     | `3.14159`          | `3.14`                  |
| `%t`   | Boolean                                  | `true`             | `true`                  |
| `%v`   | Default format (any value)               | `123`, `"hi"`      | `123`, `hi`             |
| `%+v`  | Struct with field names                  | `Person{Name:"A"}` | `{Name:A}`              |
| `%#v`  | Go syntax representation                 | `Person{Name:"A"}` | `main.Person{Name:"A"}` |
| `%T`   | Type of the value                        | `"hello"`          | `string`                |
| `%%`   | A literal percent sign                   | –                  | `%`                     |

<br>
<br>
<br>

## Formatting strings without printing

### `Sprintf`

- Formats according to the **Format specifiers** and return the resulting string.

```go
name, age, active := "Alice", 20, true

statement := fmt.Sprintf("Name: %s, Age: %d, Status: %t", name, age, active)

// Name: Alice, Age: 20, Status: true
fmt.Println(statement)
```

<br>
<br>
<br>

## Reading input

### `Scanln`

- Reads input from standard input.
- Keeps scanning until a new line.

```go
var age int

fmt.Scanln(&age)

// Your age is: 10
fmt.Println("Your age is:", age)
```

<br>
<br>
<br>
