## Index
- [ACID Properties](#acid-properties)
- [BASE Properties](#base-properties)
- [ACID vs BASE properties](#acid-vs-base-properties)

<br>
<br>
<br>

### ACID Properties

- ACID properties ensures following transactional guarantee in a relational database.
- **Atomicity**: All operations in a transaction either complete fully or not at all. (All-or-nothing).
- **Consistency**: Transaction takes DB from one valid state to another, maintaining database integrity.
- **Isolation**: Concurrent transactions don’t interfere; result is as if run serially.
- **Durability**: Once committed, changes persist even after a crash.

<br>
<br>
<br>

### BASE Properties

- BASE properties are used in NoSQL databases.
- Databases that prioritize availability & scalability over strict consistency.
- **Basically Available**: System guarantees availability (even during partial failures).
- **Soft State**: System’s state may change over time, even without input (eventual consistency).
- **Eventual Consistency**: Updates will propagate eventually; covering all replicas over time.

<br>
<br>
<br>

### ACID vs BASE properties

| Feature      | ACID (SQL)                | BASE (NoSQL)                       |
| ------------ | ------------------------- | ---------------------------------- |
| Focus        | Consistency & integrity   | Availability & partition tolerance |
| Consistency  | Strong consistency        | Eventual consistency               |
| Transactions | Strict, reliable          | Looser, more flexible              |
| Use-case     | Banking, critical systems | Social media, AWS analytics dash.  |

<br>
<br>
<br>

- What is data normalization? Why is it important in DBMS?
- Explain various normal forms.
- What is denormalization? How does it differ from normalization?
- Type of datatypes (int, char, varchar, jsonb, etc.)
- Type of keys (primary, candidate, foreign, composite, etc.)
- Type of relationships (one-to-one, one-to-many, many-to-many, etc.)
- Type of constraints (unique, not null, etc.)
- Type of joins (inner, left, right, outer, full, cross, self, etc.)
- What are database views?
- Difference between materialise & non materailise views
- Difference between delete, truncate & drop 
- What is indexing, how does it works?
- How indexes are implemented internally?
- Type of indexes (primary, composite, etc.)
- What is a transaction? What are the properties of a transaction?
- What are triggers?
- What is a stored procedure?
- What is a stored function?
- Difference between UNION & UNION ALL
- Difference between COMMIT and ROLLBACK operations
- What is referential integrity in DBMS?
- What is database Sharding?
- What is database Partioning?
- Sharding vs Partioning 
- PostgreSQL vs MySQL vs SQLite

- SQL vs NoSQL
- What are document databases?
- Explain features of MongoDB & usecase
- What are key-value databases?
- Explain features of Redis & usecase
- What are wide column databases?
- Explain features of Cassandra & usecase
- What are graph databases?
- Explain features of neo4j & usecase
- Explain features of DynamoDB & usecase
- Why NoSQL databases are more scalable?*
