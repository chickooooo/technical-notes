# REST APIs

<br>
<br>
<br>
<br>
<br>

## Index

- [Principles of REST](#principles-of-rest)
- [HTTP methods](#http-methods)
- [HTTP status codes](#http-status-codes)

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

### HTTP status codes

- Every HTTP request receives a response with an **HTTP status code**.
- These codes indicate whether the request was successful, failed, or requires additional action.
- Status codes are grouped into five categories, based on the first digit:

<br>

| Code Range | Category      | Meaning                                          |
| ---------- | ------------- | ------------------------------------------------ |
| **1xx**    | Informational | Request received, processing continues           |
| **2xx**    | Success       | Request was successfully processed               |
| **3xx**    | Redirection   | Further action needed (e.g., follow another URL) |
| **4xx**    | Client Error  | The request was incorrect or unauthorized        |
| **5xx**    | Server Error  | Server failed to process a valid request         |

<br>
<br>

#### 1xx: Informational Response

- These are informational responses.
- This means the server is communicating progress, not completing the request yet.

<br>

**100 Continue**
- The server has received the request headers and the client should continue sending the request body.
- Used for large uploads.
- Acts as a safety to check if the server would accept or reject th request.

<br>

**101 Switching Protocol**
- The server agrees to switch protocols as requested by the client via the `Upgrade` header.
- Example: Switching from HTTP/1.1 to WebSocket.

<br>

**102 Processing**
- Standard response for WebDAV.
- The server has accepted the request but is still processing it; it prevents timeouts for long operations.
- Example: Long-running WebDAV requests, such as copying large directory trees.

<br>

**103 Early Hints**
- The server sends preliminary headers before the final response.
- Commonly used for performance hints like sharing CSS & JS links.
- Sample response:

```
103 Early Hints
Link: </style.css>; rel=preload; as=style
```

<br>
<br>

#### 2xx: Success Response

- Indicates that the request was successfully processed.

**200 OK**

- The request succeeded and the server is returning the requested data in the response body.
- Common for:
    - Successful GET requests (returning a webpage, JSON, image, etc.)
    - Successful DELETE (with optional body)
    - Successful PUT/PATCH (when returning the updated resource)

<br>

**201 Created**

- A resource has been successfully created as a result of the request.
- Common for:
    - POST requests that create a new object (user, order, file, etc.)

<br>

**202 Accepted**

- The request has been successfully received for processing but is not yet complete.
- This is often used for long-running operations or background tasks.
- Where the server acknowledges the request and the client can then poll for the final result or be notified later.

<br>

**204 No Content**

- Request succeeded, but the response body is intentionally empty.
- Common for:
    - DELETE requests.
    - PUT/PATCH when no return body is returned.
    - Actions that succeed but have nothing to show

<br>

**206 Partial Content**

- Server is returning part of a resource due to a `Range` header.
- Used for:
    - Resumable downloads
    - Streaming media (video/audio players)
    - Large file fetches

<br>
<br>

#### 3xx: Redirection

- 3xx HTTP status codes represent redirection.
- They tell the client that it must take additional action (usually making a new request to a different URL) to complete the original request.

**301 Moved Permanently**

- The resource has been moved to a new permanent URL, and all future requests should use the new URL.
- Client behavior:
    - Browsers update bookmarks.
    - Search engines change indexed URL.
- Example: Redirecting from old website to new website.

<br>

**302 Found**

- Previously was 'Moved Temporarily'.
- The resource is temporarily at a different URL.
- Client behavior:
    - Use the new URL once.
    - Keep using original URL afterward.
- Example: After submitting a form, redirect to a 'success' page.
- **Note**: Commonly used historically for temporary redirects, although `303` & `307` are more precise today.

<br>

**303 See Other**

- The client should retrieve the resource using `GET` at another URL.
- Very common after POST requests to avoid resubmission if user reloads.
- Example flow: `POST` /checkout → redirect → `GET` /order/complete.

<br>

**304 Not Modified**

- The resource has not changed since the client’s last request, so the client should use its cached copy.
- Used for **Optimizing performance** and **Reducing bandwidth**.

<br>

**307 Temporary Redirect**

- The resource is temporarily at another URL.
- The HTTP method must NOT change when making request to the new URL.
- Important difference from `302`:
    - POST stays POST
    - GET stays GET

<br>

**308 Permanent Redirect**

- The resource has permanently moved.
- The method must remain the same.
- Important difference from `301`:
    - Method cannot change (POST stays POST)
- Used after **API version changes** or **URL restructuring**.

<br>
<br>
<br>
