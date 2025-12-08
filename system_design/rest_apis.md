# REST APIs

<br>
<br>
<br>
<br>
<br>

## Index

- [Principles of REST](#principles-of-rest)
- [HTTP methods](#http-methods)

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

### HTTP methods

#### `GET` method

- Fetches data from the server.
- Characteristics:
    - Safe: Does not change anything on the server.
    - Idempotent: Same request produces the same result.

```
GET /users/10
```

<br>

#### `POST` method

- Sends data to the server to create something new.
- Characteristics:
    - Non-idempotent: Sending the same POST twice creates two resources.

```
POST /users
```

<br>

#### `PUT` method

- Updates an existing resource entirely.
- Characteristics:
    - Idempotent: Sending same PUT again results in same final state.

```
PUT /users/10
```

<br>

#### `PATCH` method

- Updates an existing resource partially.
- Characteristics:
    - Not necessarily idempotent, but often used that way.

```
PATCH /users/10
```

<br>

#### `DELETE` method

- Removes a resource from the server.
- Characteristics:
    - Idempotent: Deleting the same resource twice results in same outcome.

```
DELETE /users/10
```

<br>

#### `HEAD` method

What is it?
- `HEAD` is an HTTP method that is identical to `GET`.
- The only difference is that the server only returns the **response headers** and omits the body.

Usecases:
- Check if a resource exists (e.g. verify a URL).
- Check if a cached version is still valid before getting new data.
- Check resource metadata such as:
    - Content-Type
    - Content-Length (size)
    - Last-Modified
- **Health checks** or monitoring endpoints without retrieving content.

```
HEAD /users/10
```
```
# Sample response

200 OK
Content-Type: application/json
Content-Length: 85
Last-Modified: Tue, 11 Feb 2025 10:20:00 GMT
```

<br>

#### `OPTIONS` method

What is it?
- `OPTIONS` asks a server to describe the communication options for a resource or the server itself.
- It is used to determine what actions a client is allowed to perform before actually sending a request that might be disallowed.

Usecases:
- CORS (Cross-Origin Resource Sharing)
- Browsers send a preflight request using OPTIONS to ask:
    - Which clients are allowed to access the resource?
    - Which methods are allowed? (`Access-Control-Allow-Methods`)
    - Which headers can be sent?
- A client can check what HTTP methods a resource supports (e.g., GET, POST, PUT, DELETE).
- Useful when API testing/debugging to see available methods and limits on an endpoint.

<br>
<br>
<br>
