# Unit Testing

<br>
<br>
<br>
<br>
<br>

## Index

- [Fundamentals](#fundamentals)
    - [What is unit testing?](#what-is-unit-testing)
    - [Why unit testing is important?](#why-unit-testing-is-important)
    - [What makes a good unit test?](#what-makes-a-good-unit-test)
    - [Arrange Act Assert (AAA) pattern](#arrange-act-assert-aaa-pattern)
- [Pytest Package](#pytest-package)
    - [What is pytest?](#what-is-pytest)
    - [Why use pytest?](#why-use-pytest)
    - [Test cases discovery rules](#test-cases-discovery-rules)
- [Writing Test cases](#writing-test-cases)
    - [A simple test case](#a-simple-test-case)
    - [Testing exceptions](#testing-exceptions)

<br>
<br>
<br>
<br>
<br>

## Fundamentals

### What is unit testing?

- Unit testing is the process of testing individual units of code.
- A unit usually refers to a function or method.
- It verifies that the unit behaves correctly to different inputs.
- We write unit test cases to test all possible scenarios.

<br>
<br>

### Why unit testing is important?

- It helps catch bugs early on.
- It ensures all edge cases are covered.
- It makes refactoring the code safer.
- It serves as a living documentation.

<br>
<br>

### What makes a good unit test?

- Isolated
    - Tests only one unit of functionality.
    - Self contained and can run independently.
    - Makes use of mocks when needed.

- Fast and Deterministic
    - Runs quickly.
    - Always produces the same result given the same input.

- Readable and Clear
    - Easy to understand what is being tested and why.

- Tests Behavior, Not Implementation
    - Focuses on *what* the code does, not *how* it does it.

<br>
<br>

### Arrange Act Assert (AAA) pattern

- The **AAA pattern** is a widely used structure for writing clean and readable unit tests.
- **Arrange**: Setup the test data, objects and inputs.
- **Act**: Call the function or method being tested.
- **Assert**: Verify the result matches the expected outcome.

<br>

```py
def test_addition():
    # Arrange
    a = 2
    b = 3
    expected = 5

    # Act
    result = add(a, b)

    # Assert
    assert result == expected
```

<br>
<br>
<br>
<br>
<br>

## Pytest Package

### What is pytest?

- Pytest is a modern Python testing framework.
- It is simple, readable and has a rich set of features.
- It supports unit, integration, and functional testing.
- Works with plain `assert` statements.

<br>
<br>

### Why use pytest?

- Simple syntax, uses `assert` to verify the result.
- Automatic test cases discovery.
- Provides detailed failure output.
- Rich set of features: fixtures, mocking, parameterization, etc.
- Large plugin ecosystem: code coverage, HTML report, etc.
- Can run tests in parallel (with plugins).
- Easy integration with CI/CD pipelines.

<br>
<br>

### Test cases discovery rules

- File naming:
    - `test_*.py`
    - `*_test.py`
- Function naming:
    - `def test_*():`
- Class naming:
    - `class Test*():`
    - Class should not have `__init__` method.
- Class method naming:
    - `def test_*():`

<br>
<br>
<br>
<br>
<br>

## Writing Test cases

### A simple test case

- We use `assert` statement to verify that the result matches the expected output.
- A single test case can have multiple `assert` statements.

<br>

```py
# src/my_functions.py

def add(a: int, b: int) -> int:
    return a + b
```

```py
# tests/test_my_functions.py

def test_add():
    # Arrange
    a = 3
    b = 5
    expected = 8

    # Act
    result = add(a, b)

    # Assert
    assert result == expected
```

<br>
<br>

### Testing exceptions

- Use `pytest.raises()` method to test functions that throw exceptions.
- We can also verify the error message using the `match` parameter.

<br>

```py
# src/my_functions.py

def divide(a: int, b: int) -> float:
    if b == 0:
        raise ZeroDivisionError("'b' cannot be zero")

    return a / b
```

```py
# tests/test_my_functions.py

import pytest

def test_divide():
    a = 10
    b = 4
    expected = 2.5

    result = divide(a, b)

    assert result == expected


def test_divide_by_zero():
    # Arrange
    a = 10
    b = 0

    # Act and Assert that exception is raised
    with pytest.raises(ZeroDivisionError, match="'b' cannot be zero"):
        divide(a, b)
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
