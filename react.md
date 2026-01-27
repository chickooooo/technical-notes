## Index

- [What is React?](#what-is-react)
- [Why use React?](#why-use-react)
- [React vs Angular](#react-vs-angular)
- [What is JSX?](#what-is-jsx)
- [React.createElement()](#reactcreateelement)
- [useState Hook](#usestate-hook)
- [useReducer Hook](#usereducer-hook)
- [useState vs useReducer Hook](#usestate-vs-usereducer-hook)
- [Context API](#context-api)
- [useEffect Hook](#useeffect-hook)
- [Props and prop drilling](#props-and-prop-drilling)
- [useRef Hook](#useref-hook)
- [Controlled & Uncontrolled components](#controlled--uncontrolled-components)

<br>
<br>
<br>

### What is React?

- React is a JavaScript library for building user interfaces (UIs), especially single-page applications (SPAs).
- React follows a component-based architecture.
- It uses Virtual DOM for efficient UI rendering.
- It follows declarative programming instead of imperative programing. It focuses on **what you want to achieve**, instead of **how to achieve it**.

<br>
<br>
<br>

### Why use React?

- Declarative Nature: Focuses on **what you want to achieve**, instead of **how to achieve it**.
- Performance: Uses Virtual DOM for faster updates & rendering.
- State Management: Have various tools to ease state management process.
- Maintainability: Component-based structure make code maintainable and reusable.
- Large Ecosystem: Has tools like React Router, Redux, etc. and good community support.
- Scalability: Has features like concurrent rendering & suspense for strong performance on large scale apps.

<br>
<br>
<br>

### React vs Angular

- React is a library, while angular is a full-fledged framework.
- React has a cleaner code syntax as compared to angular.
- React is easier to learn as compared to angular.
- React used Virtual DOM for rerendering while angular uses Real DOM + with change detection

---

- Angular comes with batteries inculdes while you need to install separate libraries in React. For example: State Management.
- Angular is more opinionated while you get flexibility with react.
- The difference is same as FastAPI vs Django.

<br>
<br>
<br>

### What is JSX?

- JSX stand for JavaScript XML. It is a syntax extension for JavaScript used in React.
- Allows writing HTML-like code inside JavaScript.
- Behind the scenes, JSX elements are converted to `React.createElement()` calls. Using Babel.
- Supports embedding JavaScript expressions using `{}`.
- JSX expressions must have one parent element.

<br>

```jsx
// jsc element
const element = <h1>Hello, {user.name}!</h1>;

// converted to this
const element = React.createElement('h1', null, `Hello, ${user.name}!`);
```

<br>
<br>
<br>

### React.createElement()

- It is used to create a React element.
- It serves as an alternative to writing JSX.
- To create a react element, you need `element type`, `props` & `children`.

```js
const element = createElement(type, props, ...children);
```

<br>

- Behind the scenes, JSX elements are converted to `React.createElement()` calls. Using Babel.

```jsx
// JSX code
<h1 className="greeting">Hello <i>{name}</i>. Welcome!</h1>

// createElement code
createElement(
    'h1',
    { className: 'greeting' },
    'Hello ',
    createElement(
        'i',
        null,
        name,
    ),
    '. Welcome!'
)
```

<br>
<br>
<br>

### useState Hook

- This hook lets you add state to a functional component.
- Whenever this state is changed, our component gets re-rendered.
- State updates are asynchronous in nature (scheduled).

<br>

```js
const [state, setState] = useState(initialValue);
```

- `state`: Current state value.
- `setState`: Function to update the state.
- `initialValue`: Initial value of the state (can be any type).

<br>

Example:

```js
import { useState } from 'react';

function Counter() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>{count}</p>
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </div>
  );
}
```

<br>

- Note: Always use function form to get the latest state

```js
setCount(prev => prev + 1);
```

<br>
<br>
<br>

### useReducer Hook

- This hook is used for state management, often preferred for **complex state logic**.
- Similar to `useState`, but uses a reducer function to handle updates.

<br>

```js
const [state, dispatch] = useReducer(reducer, initialState);
```

- `state`: Current state value
- `reducer`: Function (state, action) => newState
- `dispatch`: Function to send actions
- `initialState`: Initial state value

<br>

Example:

```js
function reducer(state, action) {
  switch (action.type) {
    case 'increment':
      return { count: state.count + 1 };
    case 'decrement':
      return { count: state.count - 1 };
    default:
      return state;
  }
}

function Counter() {
  const [state, dispatch] = useReducer(reducer, { count: 0 });

  return (
    <div>
      <p>{state.count}</p>
      <button onClick={() => dispatch({ type: 'increment' })}>+</button>
      <button onClick={() => dispatch({ type: 'decrement' })}>-</button>
    </div>
  );
}
```

<br>

When to use:

- Complex state transitions
- Multiple sub-values in state
- When next state depends on previous state

<br>
<br>
<br>

### useState vs useReducer Hook

| Feature            | `useState`             | `useReducer`                         |
| ------------------ | ---------------------- | ------------------------------------ |
| Simplicity         | Simple state logic     | Complex state logic                  |
| State Shape        | Usually a single value | Often an object with multiple values |
| Update Logic       | Inline with `setState` | Via a reducer function               |
| Action Dispatching | Direct update          | Dispatch action                      |

<br>

- Start with `useState`. If the state logic becomes too complex or hard to manage, switch to `useReducer`.

<br>
<br>
<br>

### Context API

- Used to share state down the tree without prop drilling.
- Ideal for **themes**, **user auth**, **language settings**, etc.
- Contains 3 main parts:
  - `React.createContext()`: Creates a context object.
  - `<Context.Provider>`: Supplies the context value.
  - `React.useContext(Context)`: Accesses context value in a component.

<br>

**Create Context**

```js
const ThemeContext = React.createContext();
```

---

**Create Provider**

```js
const ThemeProvider = ({ children }) => {
  const [theme, setTheme] = React.useState("light");

  return (
    <ThemeContext.Provider value={{ theme, setTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};
```

---

**Consumer Context**

```js
const ThemedButton = () => {
  const { theme, setTheme } = React.useContext(ThemeContext);

  return (
    <button
      onClick={() => setTheme(theme === "light" ? "dark" : "light")}
    >
      Current Theme: {theme}
    </button>
  );
};
```

---

**Wrap with Provider**

```js
const App = () => (
  <ThemeProvider>
    <ThemedButton />
  </ThemeProvider>
);
```

<br>

**Key Points**

- Context should be avoided for high-frequency updates (can cause re-renders).
- Combine with `useReducer` for complex state logic (Redux).
- Prefer context for global but stable data.

<br>
<br>
<br>

### useEffect Hook

- It allows functional components to perform side effects.
- A **side effect** is an operation that does something outside the component's render process. Example:
  - Fetching data from an API
  - Interacting with browser APIs like localStorage
  - Setting up subscriptions or timers
- useEffect provides a safe and controlled place to run side effects **after rendering**.

<br>

Syntax

```jsx
useEffect(() => {
  // Side effect code here

  return () => {
    // Cleanup code (optional)
  };
}, [dependencies]);
```

- The first argument is a function containing the side effect.
- The second argument is a dependency array that controls when the effect runs.
- The cleanup function runs before the effect re-runs or when the component unmounts.

<br>

useEffect Running

- If dependency array is not provided, useEffect will run on every render. (Avoid this).
- If empty array is provided (`[]`), useEffect will run only once on component mount.
- If array is not empty, useEffect will run every time any of the dependency changes.

<br>

Example

```jsx
import { useEffect, useState } from "react";

function Users() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    // fetch and update users once
    fetch("https://api.example.com/users")
      .then(response => response.json())
      .then(data => setUsers(data));
  }, []);

  return (
    <ul>
      {users.map(user => (
        <li key={user.id}>{user.name}</li>
      ))}
    </ul>
  );
}
```

<br>

Cleanup

- Cleanup function helps prevent memory leaks.

```jsx
useEffect(() => {
  // start timer
  const timer = setInterval(() => {
    console.log("Running...");
  }, 1000);

  return () => {
    // cleanup timer
    clearInterval(timer);
  };
}, []);
```

<br>
<br>
<br>
<br>

### Props and prop drilling

#### What are props?

- Props is short form for properties.
- They allow to pass data from parent component to child component in React.
- The parent component can send values, functions or objects to the child component.

<br>

Key Points

- Props are read-only in nature, they cannot be modified.
- They make components reusable, modular and maintainable.

<br>

```jsx
function Button(props) {
  return <button>{props.label}</button>;
}

function App() {
  return (
    <>
      <Button label="Login" />
      <Button label="Signup" />
    </>
  );
}
```

<br>
<br>

#### What is prop drilling?

- Prop drilling is a situation where props are passed through multiple levels of components, even though some intermediate components do not need the data.
- This makes the code hard to read and debug.
- Intermediate components have to receive and pass on the props, which adds unnecessary overhead.

<br>

```jsx
function App() {
  return <Parent user="John" />;  // data origin
}

function Parent(props) {
  return <Child user={props.user} />;
}

function Child(props) {
  return <GrandChild user={props.user} />;
}

function GrandChild(props) {
  return <h1>Hello, {props.user}</h1>;  // data usage
}
```

<br>
<br>

#### How to avoid prop drilling?

- Use Context API to share data globally without passing props manually at every level.
- Use state management libraries like Redux & Zustand, to manage global state for large applications.
- Restructure component tree to keep the concerned components as close as possible.

<br>
<br>
<br>
<br>

### useRef Hook

- `useRef` is a React Hook that lets us create a **mutable reference** that persists across renders.
- Whenever the value of this reference changes, the UI **is not** re-rendered.
- `useRef` is commonly used to:
  - Persist data between renders that doesn't affect the UI.
  - Access or manipulate a DOM element directly.
  - Store and update a value without triggering re-renders

<br>

```jsx
import { useRef } from "react";

function InputFocus() {
  const inputRef = useRef(null);

  const focusInput = () => {
    inputRef.current.focus();
  };

  return (
    <>
      <input ref={inputRef} />
      <button onClick={focusInput}>Focus</button>
    </>
  );
}
```

<br>

#### Key points

- `useRef` works with mutable values while `useState` works with immutable values.
- Common usecases: store timers & intervals, access DOM elements, etc.

<br>
<br>
<br>
<br>

### Controlled & Uncontrolled components

- In React, controlled and uncontrolled components describe two different ways of handling form data.
- The key difference is who owns and controls the state of the form element: React or DOM.

<br>
<br>

#### Controlled component

- A controlled component is a form element whose value is fully controlled by React.
- `useState` is commonly used to manage the elements state.

<br>

```jsx
import { useState } from "react";

function ControlledInput() {
  const [name, setName] = useState("");

  const handleChange = (event) => {
    setName(event.target.value);
  };

  return (
    <div>
      <input
        type="text"
        value={name}
        onChange={handleChange}
      />
      <p>Your name is: {name}</p>
    </div>
  );
}
```

- Typing in the input, calls `handleChange` which updates the state via `setName`.
- The value of the input is always equal to `name`.
- Ideal when we want to validate input in real time.

<br>
<br>

#### Uncontrolled component

- An uncontrolled component is a form element that manages its own state internally in the DOM.
- React does not track the input value on every change.
- `useRef` is commonly used to access the element's state.

<br>

```jsx
import { useRef } from "react";

function UncontrolledInput() {
  const nameRef = useRef(null);

  const handleSubmit = (event) => {
    event.preventDefault();
    alert(`Your name is: ${nameRef.current.value}`);
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" ref={nameRef} />
      <button type="submit">Submit</button>
    </form>
  );
}
```

- The input manages its own value inside the DOM.
- React only reads the value when the form is submitted.
- Ideal for simple forms and performance.

<br>

- React suggests using Controlled components over uncontrolled components.

<br>
<br>
<br>
<br>

### 

<br>
<br>
<br>
<br>
