## Index

- [What is React?](#what-is-react)
- [Why use React?](#why-use-react)
- [React vs Angular](#react-vs-angular)
- [What is JSX?](#what-is-jsx)
- [React.createElement()](#reactcreateelement)
- [useState Hook](#usestate-hook)
- [useReducer Hook](#usereducer-hook)
- [useState vs useReducer Hook](#usestate-vs-usereducer-hook)

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

### 

- What is React?
- Why use React?
- React vs Angular
- What is JSX?
- React.createElement()

- What are props in React?
- What is prop drilling and how can you avoid it?
- What are keys in React and why are they important?
- Class Component vs Functional Component
- controlled vs uncontrolled components
- What are higher-order components (HOC)?
- What are synthetic events in React?

- What are side effects?
- useEffect hook
- useLayoutEffect
- useEffect vs useLayoutEffect

- useRef hook
- usePortal hook
- useMemo hook
- useCallback hook
- useMemo vs useCallback

- What are hooks in React?
- How to create custom

- What is lifting state up in React?

- context API (deepdive)
- redux
- react redux toolkit
- What are middleware in Redux?
- Redux vs Context API

- What are React Router and its main components?
- Protected routes
- Role based access control (RBAC)

- virtual dom
- diffing algorithm
- reconcilliation
- react fibre
- What are render props in React?
- concurrent rendering
- suspense

- How do you optimize performance in a React app?

- What is Next.js and how is it different from React?
- What is server-side rendering (SSR) in React?
