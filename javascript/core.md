# JavaScript Core

<br>
<br>
<br>
<br>
<br>

## Index

- [Declaring variables](#declaring-variables)
- [Hoisting](#hoisting)
- [Temporal Dead Zone](#temporal-dead-zone)

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

## Temporal Dead Zone

- Temporal Dead Zone is the period between the start of a scope and the point where a `let` or `const` variable is declared.
- During this period trying to access the variable results in a `ReferenceError`.
- This is done to avoid access to hoisted variables that are not yet initialised in the code.

<br>
<br>
<br>
<br>

## 

<br>
<br>
<br>
<br>
