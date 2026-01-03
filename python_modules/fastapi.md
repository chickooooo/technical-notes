# FastAPI

<br>
<br>
<br>
<br>
<br>

## Index

- [What is FastAPI?](#what-is-fastapi)
- [Flask vs Django vs FastAPI](#flask-vs-django-vs-fastapi)
- [Advantages and limitations of FastAPI](#advantages-and-limitations-of-fastapi)
- [Why is FastAPI considered fast?](#why-is-fastapi-considered-fast)
- [Starlette in FastAPI](#starlette-in-fastapi)
- [Sync vs Async endpoints in FastAPI](#sync-vs-async-endpoints-in-fastapi)
- [Dependency injection](#dependency-injection)
- [Dependency override](#dependency-override)

<br>
<br>
<br>
<br>
<br>

## What is FastAPI?

- FastAPI is a modern, high-performance Python web framework used to build REST APIs.
- It is built on Starlette (web layer) and Pydantic (data validation).
- It is designed for speed, type safety, and developer productivity.
- FastAPI supports **Async APIs** natively.

<br>
<br>
<br>

## Flask vs Django vs FastAPI

Flask

- Flask is a lightweight web framework.
- It is simple and easy to learn for beginners.
- No built-in ORM, authentication, etc. (you add what you need).
- Good for:
    - Small projects
    - Prototypes

<br>

Django

- Django is a full-featured, batteries-included framework.
- It comes with built-in ORM, authentication, admin panel, template engine, etc.
- Django has its own way of doing things and a steeper learning curve than Flask.
- Good for:
    - Fullstack web applications
    - Large or complex applications
    - When we want batteries included

<br>

FastAPI

- FastAPI is a modern, high-performance framework.
- It is more fast, pythonic and has strong type hinting.
- It has Asynchronous (async/await) support by default.
- Good for:
    - APIs and microservices
    - High-performance & modern applications

<br>
<br>
<br>

## Advantages and limitations of FastAPI

Advantages:

- Very fast (comparable to Node.js & Go).
- Native async support.
- Fast data validation and serialization using Pydantic.
- Auto-generated API docs (OpenAPI & Swagger).
- Pythonic and easy to learn.

<br>

Limitations:

- Smaller ecosystem compared to Django.
- Not ideal for full-stack apps.
- Async requires understanding of async programming.

<br>
<br>
<br>

## Why is FastAPI considered fast?

- FastAPI is fast in both **Execution** and **Development**.

<br>

Execution:

- Built on top of Starlette (ASGI):
    - ASGI is designed to be lightweight and fast compared to WSGI.
    - It avoids some of the overhead of thread or process-based concurrency.
- Uses Pydantic for validation and serialization:
    - Pydantic is implemented in pure Python with C speedups under the hood.
    - This allows request parsing and validation to be very fast compared to traditional approaches.
- Minimal overhead:
    - FastAPI has minimal middleware overhead. You add what you need.
    - This enables faster request processing.

<br>

Development:

- Less boilerplate code.
- Automatic request parsing & validation using Pydantic.
- Auto API docs generation using OpenAPI and Swagger.

<br>
<br>
<br>
<br>
<br>

## Starlette in FastAPI

- Starlette is an asynchronous server gateway interface (ASGI) framework.
- FastAPI is built on top of Starlette.
- Starlette provides the core web framework, while FastAPI adds data validation, type safety, and other features.

<br>
<br>

### Key points:

- Starlette handles routing, middlewares, request-response, websockets, etc.
- FastAPI uses:
    - Starlette for web framework + async handling.
    - Pydantic for data validation & serialization.
    - It's own features like type safety, dependency injection, API docs, etc.

<br>
<br>
<br>
<br>
<br>

## Sync vs Async endpoints in FastAPI

Sync endpoints

- `sync` endpoints run normally and block the worker thread.
- Run in a threadpool managed by FastAPI.
- Suitable for blocking operations (sync database calls, file IO).
- Only one request is handled by a thread at a time.

<br>

Async endpoints

- `async` endpoints are non-blocking in nature.
- Run directly on the event loop.
- Ideal for non-blocking IO operations (async DB drivers, HTTP calls).
- Allows handling many requests concurrently.

<br>

Key points

- Use sync endpoints if the code relies on blocking libraries.
- Use async endpoints if the code supports async calls (databases, APIs, I/O).
- If we make blocking calls in an async endpoint, it will block the event loop.
- A FastAPI application can have both sync and async endpoints exist together.

<br>
<br>
<br>
<br>
<br>

## Dependency injection

- In FastAPI, Dependency Injection (DI) is a way to provide reusable logic or resources to an endpoint.
- These resources can be DB sessions, authenticated user, extracted params, etc.
- FastAPI uses the `Depends()` function to declare dependencies.

<br>

Key points

- A dependency can be a class, function or a generator.
- Dependencies are resolved automatically per request.
- FastAPI handles execution order and cleanup (using `yield` dependencies).
- DI promotes clean, reusable and testable code.

<br>

```py
from fastapi import Depends, FastAPI

app = FastAPI()

def get_db():
    db = "db_connection"
    try:
        yield db
    finally:
        print("Closing DB")

@app.get("/profile")
def profile(db=Depends(get_db)):
    ...
```

<br>
<br>
<br>
<br>
<br>

## Dependency override

- Dependency override allows replacing a dependency with another implementation.
- Commonly used when testing endpoints. Both unit and integration tests.
- It is used to mock dependencies like database, external services, etc.

<br>

```py
# test file

from fastapi import FastAPI

app = FastAPI()

app.dependency_overrides[get_db] = get_test_db
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
