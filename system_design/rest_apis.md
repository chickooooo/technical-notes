# REST APIs

<br>
<br>
<br>
<br>
<br>

## Index

- [Principles of REST](#principles-of-rest)

<br>
<br>
<br>
<br>
<br>

## Core concepts

### Principles of REST

- REST stands for Representational State Tranfser.
- It is an architectural style for building networked applications.
- RESTful APIs follow a set of guiding principles (constraints), that make them scalable, simple, and reliable.

#### Client-Server separation

- The client (frontend) and server (backend) are independent.
    - The client only needs to know the API interface.
    - The server does not need to know how the client uses the data.
- **Benefit**: Scalability, flexibility, and independent evolution of client and server.

<br>

#### Statelessness

- Each request from a client to a server must contain all necessary information.
    - No session stored on the server.
    - Authentication must be included with every request (e.g., tokens).
- **Benefit**: Simpler, more scalable, easier to cache.

<br>

#### Resource-Based Architecture

- Everything is treated as a resource identified by a URI.
    - Use nouns, not verbs, in endpoints (`/books`, not `/getBooks`).
    - Actions are represented by HTTP methods.
- **Benefit**: Clear, consistent structure.

<br>

#### Use of Standard HTTP Methods

- REST uses HTTP semantics properly:
    - GET: retrieve a resource
    - POST: create a new resource
    - PUT: replace a resource
    - PATCH: update part of a resource
    - DELETE: delete a resource
- **Benefit**: Predictable behavior across APIs.

<br>

#### Layered System

- A REST system can have multiple layers (load balancers, proxies, gateways).
- Clients shouldn’t know whether they are communicating with the end server or an intermediary.
- **Benefit**: Security, scalability, modularity.

<br>

#### Cacheability

- Responses must be explicitly marked as cacheable or not.
- Use HTTP caching headers (e.g., `Cache-Control`, `ETag`).
- **Benefit**: Improves performance and reduces server load.

<br>
<br>
<br>
