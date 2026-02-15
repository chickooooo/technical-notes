# JavaScript Core

<br>
<br>
<br>
<br>
<br>

## Index

- [Datatypes](#datatypes)
- [Operators](#operators)
- [Conditional statements](#conditional-statements)
- [Declaring variables](#declaring-variables)
- [Hoisting](#hoisting)
- [Temporal Dead Zone](#temporal-dead-zone)
- [Scope](#scope)
- [Currying](#currying)
- [Async-Await](#async-await)

<br>
<br>
<br>
<br>
<br>

## Datatypes

- JavaScript has 8 fundamental data types.
- These are divided into primitive and non-primitive (reference) types.

<br>

### Number

- The `number` type represents both integer and floating-point values.
- JS does not have separate types for integers and decimals; both are number.

```js
let age = 25;
let price = 99.99;
```

<br>

### BigInt

- `BigInt` is used to represent integers larger than the safe limit for regular numbers
(> 2⁵³ − 1 or < −(2⁵³ − 1)).
- A BigInt is created by appending n to the end of a number.
- BigInt allows safe calculations with very large integers.

```js
let bigNumber = 12345678901234567890n;
```

<br>

### String

- A string represents textual data. It can be written using single quotes, double quotes, or backticks.
- Strings are immutable (they cannot be changed directly).

```js
let name = "Alice";
let city = 'London';
```

<br>

### Boolean

- A boolean represents logical values: `true` or `false`.
- Booleans are commonly used in conditions and comparisons.

```js
let isLoggedIn = true;
let hasPermission = false;
```

<br>

### Null

- `null` represents the intentional absence of a value. It is assigned explicitly.
- It means 'no value' or 'empty'.

```js
let selectedUser = null;
```

<br>

### Undefined

- `undefined` means a variable has been declared but not assigned a value.
- It indicates that a value does not exist yet.

```js
let score;
console.log(score); // undefined
```

<br>

### Symbol

- A `symbol` is a unique and immutable value, often used as object property keys.
- Each symbol is unique, even if created with the same description.

```js
let id = Symbol("id");
```

<br>

### Object

- An `object` is a collection of related data stored as key-value pairs.
- **Note**: Arrays also fall under the `object` type.

```js
let person = {
  name: "John",
  age: 30
};

let arr = [1, 2, 3];
console.log(typeof arr); // "object"
```

<br>
<br>

### Primitive datatypes

- First 7 datatypes are primitive in nature.
- They are immutable and are stored by value.
- They generally hold a single value.

```js
let a = 10;
let b = a;
b = 20;

console.log(a); // 10 (original value unchanged)
```

<br>

### Non-primitive datatypes

- Also known as **reference types**.
- `object` is the only non-primitive datatype.
- They are mutable and are stored by reference.
- They can hold multiple values.

<br>
<br>
<br>
<br>
<br>

## Operators

### Arithmetic operators

- Used to perform basic mathematical operations.

<br>

| Operator | Description         | Example | Result |
| -------- | ------------------- | ------- | ------ |
| `+`      | Addition            | `5 + 2` | `7`    |
| `-`      | Subtraction         | `5 - 2` | `3`    |
| `*`      | Multiplication      | `5 * 2` | `10`   |
| `/`      | Division            | `5 / 2` | `2.5`  |
| `%`      | Modulus (remainder) | `5 % 2` | `1`    |

<br>
<br>

### Assignment operators

- Used to update variable values.

<br>

| Operator | Meaning             | Example  | Equivalent To |
| -------- | ------------------- | -------- | ------------- |
| `+=`     | Add and assign      | `x += 5` | `x = x + 5`   |
| `-=`     | Subtract and assign | `x -= 5` | `x = x - 5`   |
| `*=`     | Multiply and assign | `x *= 5` | `x = x * 5`   |
| `/=`     | Divide and assign   | `x /= 5` | `x = x / 5`   |

<br>
<br>

### Increment and decrement

- Used to increase or decrease a value by 1.

<br>

| Operator | Description    |
| -------- | -------------- |
| `++`     | Increment by 1 |
| `--`     | Decrement by 1 |

<br>

```js
let count = 5;

count++; // 6
count--; // 5
```

```js
let x = 5;

console.log(++x); // 6 (increments first, then prints)
console.log(x++); // 6 (prints first, then increments)
console.log(x);   // 7
```

<br>
<br>

### `typeof` operator

- Used to check the data type of a value.

```js
typeof 42;          // "number"
typeof "Hello";     // "string"
typeof true;        // "boolean"
typeof undefined;   // "undefined"
typeof null;        // "object" (known JavaScript quirk)
```

<br>
<br>
<br>
<br>
<br>

## Conditional statements

### `if`, `else if`, `else`

- Used to execute code based on conditions.

```js
let age = 18;

if (age > 18) {
  console.log("Adult");
} else if (age === 18) {
  console.log("Just became adult");
} else {
  console.log("Minor");
}
```

<br>
<br>

### Comparison operators

- Used to compare values.
- Always prefer `===` instead of `==` to avoid type coercion issues.

| Operator | Meaning                     |
| -------- | --------------------------- |
| `>`      | Greater than                |
| `>=`     | Greater than or equal       |
| `<`      | Less than                   |
| `<=`     | Less than or equal          |
| `===`    | Strict equal (value + type) |
| `!==`    | Strict not equal            |

<br>
<br>

### Logical operators

- Used to combine conditions.

| Operator | Meaning                     |
| -------- | --------------------------- |
| `&&`     | AND (both must be true)     |
| `ll`     | OR (at least one true)      |
| `!`      | NOT (reverse boolean value) |


<br>
<br>

### Truthy and falsy values

- In JS, values are automatically converted to `true` or `false` in conditions.
- Falsy values: `false`, `0`, `""`, `null`, `undefined`, `NaN`.
- All other values are truthy. 

<br>
<br>

### Short circuit evaluation

- OR (`||`) always returns the first truthy value.
- If both are falsy, returns the second value.

```js
let username = "";
let defaultName = username || "Stranger";

console.log(defaultName); // "Stranger"
```

<br>

- AND (`&&`) always returns the first falsy value.
- If both are truthy, returns the second value.
- Second expression is evaluated only if first is truthy.

```js
let isLoggedIn = true;
isLoggedIn && console.log("Welcome!");

// Welcome!
```

<br>
<br>

### Ternary operator

- Short form of `if-else`.

```js
let age = 20;

let message = age >= 18 ? "Adult" : "Minor";
console.log(message);  // Adult
```

<br>
<br>

### Switch statement

- Used when checking multiple possible values of the same variable.
- `break` is important to prevent fall-through.

```js
let day = 2;

switch (day) {
  case 1:
    console.log("Monday");
    break;
  case 2:
    console.log("Tuesday");
    break;
  default:
    console.log("Invalid day");
}
```

<br>
<br>
<br>
<br>
<br>

## Declaring variables

- In JS variables can be declared using 3 keywords: `var`, `let` and `const`.
- `var` is generally used in legacy JS code.
- `let` and `const` are generally used in modern JS code.

<br>
<br>

### `var`

- With `var`, we can declare a variable and initialise it later.
- This declared variable will be `undefined` until it is initialised.
- We can also do declaration and initialization at the same time.

```js
var x;  // undefined
x = 10;

var y = 10;
```

<br>

Key points

- Variables declared with `var` are **function-scoped**. They can be accessed anywhere inside the function they are declared. Accessing outside the function will cause `ReferenceError`.
- Note that these variables are not block-scoped. Meaning, they can be accessed outside `if` and `for` blocks.
- Due to hoisting, variable declared using `var` can be accessed before their declaration line.

```js
// function scoped

if (true) {
  var a = 5;
}

console.log(a); // 5
```

```js
// hoisting

console.log(x);  // undefined
var x = 10;
```

<br>
<br>

### `let`

- To fix the issues with `var`, JS introduced `let` keyword.
- Just like `var`, we can declare a variable and initialise it later OR do both on the same line.

```js
let x;  // undefined
x = 10;

let y = 10;
```

<br>

Key points

- Variable declared using `let` are **block-scoped**. This means they exist only inside the nearest `{}` block.
- `let` variables are also hoisted, but they cannot be accessed before the line they are declared on, due to the Temporal Dead Zone.
- These 2 features make the variables behaviour much more predictable.

```js
// block scoped

if (true) {
  let a = 5;
}

console.log(a); // ReferenceError
```

<br>
<br>

### `const`

- Variables declared using `const` must be initialised on the same line.
- These variables are also **block-scoped** just like `let`.
- A variable declared with `const` cannot be reassigned to another value.

```js
const x = 1;

x = 10;  // TypeError
```

Key points

- Note that `const` variable are not immutable. If the value they are pointing to is mutable, it can change.
- Just like `let`, `const` variables are also hoisted, but they cannot be accessed before their declaration.

```js
// mutability

const x = { a: 1 };

x.a = 10;
```

<br>
<br>

### Usage model

- Use `const` by default.
- Use `let` only when reassignment is required.
- Avoid `var` entirely in modern code.

<br>
<br>
<br>
<br>
<br>

## Hoisting

- Before JS executes the code, it performs a preprocessing step. During this step, JS scans the code and allocates memory to variables and functions. This behaviour is called hoisting.
- This means that JS makes the variables available before the line they appears on.
- Variables are always hoisted at the top of their scope. 

<br>
<br>

### Impact on `var`

- Because of hoisting, variable declared with `var` can be accessed before the line they are declared on.

```js
console.log(x);  // undefined
var x = 10;
```

- The above code does not throw `ReferenceError` because memory is already allocated to `x` and it currently holds nothing (`undefined`).

<br>
<br>

### Impact on `let` and `const`

- Variable declared with `let` and `const` are also hoisted but cannot be accessed before declaration due to the Temporal Dead Zone.

<br>
<br>
<br>
<br>
<br>

## Temporal Dead Zone

- Temporal Dead Zone is the period between the start of a scope and the point where a `let` or `const` variable is declared.
- During this period trying to access the variable results in a `ReferenceError`.
- This is done to avoid access to hoisted variables that are not yet initialised in the code.

<br>
<br>
<br>
<br>
<br>

## Scope

- In JavaScript, scope defines where a variable or a function is accessible.
- There are 3 main types of scope:
    - Global
    - Function
    - Block

<br>
<br>

### Lexical scope

- JavaScript uses lexical scoping, which means scope is determined at write-time, not run-time.
- In simple words, A variables's scope is defined by where it is written in the source code and not by where it is called from.

```js
function outer() {
  const x = 10;

  function inner() {
    console.log(x);  // 10
  }

  inner();
}

outer();
```

- Here, `inner` can access `x` because `x` exists in the lexical environment where `inner` is defined.

<br>
<br>

### Global scope

- Anything declared outside of a function or a block lives in the global scope.
- Global variables declared using `var` behaves differently than `let` or `const`.
- `var` creates a property on the global object (`window` in browsers, `global` in Node.js).
- `let` and `const` do not attach themselves to the global object.

```js
var x = 10;
let y = 20;

console.log(window.x); // 10
console.log(window.y); // undefined
```

<br>
<br>

### Function scope

- Variables declared inside a function are scoped to that function. They cannot be accessed outside the function.
- Variables declared using `var` are function-scoped.

```js
function test() {
  var a = 10;
  let b = 20;
  const c = 30;
}

console.log(a, b, c); // ReferenceError
```

<br>
<br>

### Block scope

- Variables declared inside a block (`{}`) are scoped to that block. They cannot be accessed outside the block.
- Variables declared using `let` and `const` are block-scoped.

```js
if (true) {
  var a = 1;
  let b = 2;
}

console.log(a); // 1
console.log(b); // ReferenceError
```

<br>
<br>

### Scope chain & Variable resolution

- When JavaScript tries to resolve a variable, it follows the **scope chain**.
- Scope chain defines a chain of scopes. JS tries to find the variable in the first scope, if not present, it moves on to the next scope in the chain.
- Here's how the lookup process looks like:
    - Check the current scope.
    - If not found, move to the parent scope.
    - Continue until the global scope.
    - Throw a `ReferenceError` if not found.

```js
const a = 10;

function foo() {
  const b = 20;

  function bar() {
    const c = 30;
    console.log(a, b, c);  // 10 20 30
  }
  bar();
}
foo();
```

<br>
<br>

### Scope vs Execution context

- Scope is static and is determined at parse time.
- Execution context is dynamic and created at runtime when code executes.

```js
function foo() {
  console.log(x);  // ReferenceError: x is not defined
}

function bar() {
  const x = 10;
  foo();
}

bar();
```

- When `bar()` is called, even though `x` is defined, trying to access `x` in `foo()` will throw `ReferenceError` because it is not present in the scope.
- JS will search for `x` in `foo()` -> `global scope`. Not in `bar()`.

<br>
<br>

### Scope and loops

```js
for (var i = 0; i < 3; i++) {
  setTimeout(() => console.log(i), 1000);
}

// 3
// 3
// 3
```

- `var` is function-scoped, All callbacks share the same `i`.

```js
for (let i = 0; i < 3; i++) {
  setTimeout(() => console.log(i), 1000);
}

// 0
// 1
// 2
```

- Using `let`, each iteration gets its own block-scoped `i`.

<br>
<br>

### Relationship between Scope, Hoisting & TDZ

- Scope defines where a variable exists.
- Hoisting defines when it becomes available.
- TDZ defines when it can be accessed safely.

<br>
<br>

### Key points

- JavaScript uses lexical scoping (write-time), not dynamic scoping (run-time).
- `var` is function-scoped, while `let` and `const` are block-scoped.
- Global `var` attaches to the global object; `let` and `const` do not.
- Variable resolution is done using scope chain.

<br>
<br>
<br>
<br>
<br>

## Currying

- Currying is a functional programming technique where a function that takes multiple arguments is transformed into a sequence of functions, each taking a single argument.
- This allows better code reuse and composable functions.
- Currying is highly used in functional libraries like Redux.

<br>

```js
// normal function
function multiplyNormal(multiplier, num) {
  return multiplier * num;
}

// curried function
function multiplyCurried(multiplier) {
  return function (num) {
    return multiplier * num;
  };
}

// code reuse
const multiplyBy10 = multiplyCurried(10);
const multiplyBy50 = multiplyCurried(50);

// usecase
multiplyBy10(3); // 30
multiplyBy50(3); // 150
multiplyCurried(100)(3); // 300
```

<br>
<br>
<br>
<br>
<br>

## Async-Await

- Async-await is a syntax in JavaScript that allows us to write asynchronous code in a synchronous looking manner.
- It is built on top of Promises and makes asynchronous code easier to read, write, and maintain.

<br>
<br>

### `async`

- The `async` keyword is used before a function.
- An `async` function always returns a Promise.
- If we return a value, JavaScript automatically wraps it in a resolved Promise.

<br>

```js
async function getData() {
  return "Hello";
}

getData().then(result => console.log(result));

// Output:
// Hello
```

<br>
<br>

### `await`

- The `await` keyword can be used inside an async function only.
- It pauses the execution of the function until the Promise is resolved or rejected.
- It returns the resolved value of the Promise.

<br>

```js
function fetchData() {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve("Data received");
    }, 2000);
  });
}

async function showData() {
  const result = await fetchData();
  console.log(result);
}

showData(); // Data received
```

<br>
<br>

### Error handling

- We handle errors in async-await using `try-catch` blocks.
- If an error is raised, rest of the `try` block code is skipped and the error is caught in the `catch` block.

```js
async function getUser() {
  try {
    const response = await fetch("https://wrong-url.com");
    const data = await response.json();
    console.log(data);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

getUser(); // Error occurred: TypeError: fetch failed
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
