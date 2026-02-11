# Caching

<br>
<br>
<br>
<br>
<br>

## Index

- [What is caching?](#what-is-caching)
- [Types of cache](#types-of-cache)

<br>
<br>
<br>
<br>
<br>

### What is caching?

- Caching is a technique used to store **frequently accessed** or **computationally expensive** data in a fast storage layer.
- So that future requests for the same data can be served more quickly.
- Caching improves system performance, lowers infrastructure cost, and enhances user experience.

<br>

#### Important concepts

- Cache hit: When the requested data is found in the cache.
- Cache miss: When the requested data is not found in the cache.

<br>
<br>
<br>
<br>

### Types of cache

- Depending on where the data is stored, caches can be divided into following types:

<br>

#### External / shared cache

- Data is stored on an external server that is running Redis, Memcached, etc.
- This is also known as shared cache because multiple application servers can access this cache.
- Idea for distributed systems where multiple servers need a shared cache.
- If the load on a single cache server increases, multiple instances can be added creating a cache cluster.

<br>

#### Server cache

- Here, the data is stored on the application server itself.
- It is the fastest form of caching as it avoids making any external network calls.
- Ideal when we want low latency and want to avoid frequent external calls.
- Example: caching lookup table / configs, etc.

<br>

#### Client side cache

- Data is stored on the client's device in browser or mobile app storage.
- Web browsers usually cache images, CSS & JS files so they do not need to be downloaded again.
- This reduces server load and speeds up page loading time.

<br>

#### Content Delivery Network (CDN)

- A CDN is a geographically distributed network of computers that stores frequently used data.
- A user can access this data from its closest CDN, resulting in faster response time.
- Commonly used to cache static media files. But can also cache API responses, HTML files, etc.

<br>

#### Database cache

- A database can cache frequently queried data to reduce repeated reads and expensive queries.
- Some databases have built-in caching mechanisms, such as query caching or buffer pools.
- This improves database performance and avoid the need for writing our own caching logic.

<br>
<br>
<br>
<br>

### 

<br>
<br>
<br>
<br>
