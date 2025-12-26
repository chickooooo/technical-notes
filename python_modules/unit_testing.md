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

## 

<br>
<br>
<br>
<br>
<br>
