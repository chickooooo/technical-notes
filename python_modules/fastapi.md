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
- [Why is FastAPI Considered Fast?](#why-is-fastapi-considered-fast)

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

## Why is FastAPI Considered Fast?

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

## 

<br>
<br>
<br>
<br>
<br>
