# Caching

<br>
<br>
<br>
<br>
<br>

## Index

- [What is caching?](#what-is-caching)
- [Types of cache](#types-of-cache)
- [Caching strategies](#caching-strategies)
- [Cache eviction policies](#cache-eviction-policies)

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

### Caching strategies

- Depending on consistency requirements and traffic pattern, following caching strategies can be used:

<br>

#### Cache aside

- The application first checks the cache for the requested data. If the data is found (cache hit), it is returned immediately.
- If the data is not found (cache miss), the application fetches the data from the database, stores it in the cache, and then returns it to the user.

---

This strategy is ideal when:

- Read operations are more frequent than write operations.
- Not all data needs to be cached.

<br>

#### Read through cache

- Here, the cache sits between the application and the database.
- The application always reads from the cache.
- If the data is not present, the cache itself retrieves the data from the database, stores it, and returns it to the application.

---

This strategy is ideal when:

- We want to outsource caching implementation.
- Consistent data access patterns are required.

<br>

#### Write through cache

- Here again, the cache sits between the application and the database.
- The application writes data to the cache and the cache immediately writes data to the database.
- Write is considered complete only after both cache & database writes succeed.

---

This strategy is ideal when:

- Strong consistency between cache and database is required.
- Read performance is critical.

<br>

#### Write behind cache

- The application writes data only to the cache.
- The cache then updates the database asynchronously after a short delay.

---

This strategy is ideal when:

- High write performance is required.
- Eventual database consistency is acceptable.

<br>

#### Write around cache

- Here, write operations go directly to the database, bypassing the cache.
- The cache is updated only when data is read again (using any of the read strategies).

---

This strategy is ideal when:

- Write operations are frequent but reads are infrequent.
- We want to avoid polluting the cache with rarely accessed data.

<br>
<br>
<br>
<br>

### Cache eviction policies

- When a cache becomes full, it must decide which existing item to remove to make space for a new one.
- This process is known as cache eviction. Following are some of the commonly used cache eviction policies:

<br>

#### Least recently used (LRU)

- LRU removes the item that has not been accessed for the longest time.
- It assumes that if something hasn't been used recently, it is less likely to be used soon.
- This policy is simple to implement and most commonly used.

<br>

#### Least frequently used (LFU)

- LFU removes the item that has been accessed the fewest number of times.
- It assumes that items accessed more frequently in the past will continue to be important.
- This policy is ideal when we have popular items in the system.

<br>

#### First in first out (FIFO)

- FIFO removes the oldest inserted item, regardless of usage.
- It does not consider recency or frequency.
- Ideal when cache freshness is more important than usage pattern.

<br>

#### Time to live (TTL)

- Items expire automatically after a fixed time.
- TTL is often combined with LRU or LFU.
- Ideal when cache freshness and usage pattern are equally important.
- This policy is commonly used for API response and DB queries.

<br>
<br>
<br>
<br>

### 

<br>
<br>
<br>
<br>
