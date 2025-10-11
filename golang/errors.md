# `errors` Package

<br>
<br>
<br>
<br>
<br>

## Create a new error

- Use `errors.New()` to create a new error.
- When a function uses `errors.New()`, it creates a new error value each time the function is called.
- In Go, error is an **interface type** and interfaces are value types.

<br>

- Note: Error message should start with a lowercase letter and contain no punctuations.

```go
import "errors"

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

<br>
<br>
<br>

## Wrap an error

- Use `fmt.Errorf()` to wrap an error with extra context.
- Always use `%w` verb to preserve the original error for later unwrapping.
- We can wrap an error any number of times.

```go
baseErr := errors.New("file not found")
wrappedErr := fmt.Errorf("failed to open config: %w", baseErr)
```

<br>
<br>
<br>

## Check for specific error

- Use `errors.Is()` to check for a specific error.
- It filters through wrapping layers to match the base error.

```go
var ErrPermission = errors.New("permission denied")

func checkPermission(user string) error {
    if user != "admin" {
        return fmt.Errorf("access denied: %w", ErrPermission)
    }
    return nil
}

func main() {
    err := checkPermission("guest")
    if errors.Is(err, ErrPermission) {
        fmt.Println("Caught permission error!")
    }
}
```

<br>
<br>
<br>

## Unwrap an error

- Use `errors.Unwrap()` to unwrap an error and get the base error.
- If no base error is present, it returns `nil`.

```go
err := fmt.Errorf("operation failed: %w", errors.New("timeout"))

unwrapped := errors.Unwrap(err)
fmt.Println(unwrapped) // "timeout"
```

<br>
<br>
<br>

## Type assertion on error

- Use `errors.As()` to extract the custom error from the error tree.
- `errors.As()` performs type assertion through wrapped layers and uses the first matched type.
- Once the error is extracted, we can access that custom error's fields.

```go
type APIError struct {
	Code    int
	Message string
}

func (a *APIError) Error() string {
	return fmt.Sprintf("API Error Code: %d, Message: %s", a.Code, a.Message)
}

func CallAPI() error {
	return &APIError{
		Code:    401,
		Message: "Unauthorized",
	}
}

func main() {
	err := CallAPI()

	if err != nil {
		fmt.Println(err) // API Error Code: 401, Message: Unauthorized
	}

	var apiErr *APIError
	if errors.As(err, &apiErr) {
		// If type matched, we can access the error fields
		fmt.Println(apiErr.Code, apiErr.Message) // 401 Unauthorized
	}
}
```

<br>
<br>
<br>

##
