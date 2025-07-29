## Index
- [ACID Properties](#acid-properties)
- [BASE Properties](#base-properties)
- [ACID vs BASE properties](#acid-vs-base-properties)
- [SQL Datatypes](#sql-datatypes)
- [SQL Keys](#sql-keys)

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

### SQL Datatypes

**Numeric**

| Data Type                             | Description                                                                         |
| ------------------------------------- | ----------------------------------------------------------------------------------- |
| `INT` / `INTEGER`                     | Whole numbers. Usually 4 bytes. Range: -2^31 to 2^31 - 1                            |
| `SMALLINT`                            | Smaller range integer (2 bytes).                                                    |
| `BIGINT`                              | Large integer values (8 bytes).                                                     |
| `DECIMAL(p,s)` / `NUMERIC(p,s)`       | Fixed-point numbers with precision (p) and scale (s). Suitable for monetary values. |
| `FLOAT` / `REAL` / `DOUBLE PRECISION` | Approximate floating-point numbers. Use when precision isn’t strict.                |

<br>

**String**

| Data Type    | Description                                                                                             |
| ------------ | ------------------------------------------------------------------------------------------------------- |
| `CHAR(n)`    | Fixed-length character string (e.g., CHAR(10) always takes 10 bytes).                                   |
| `VARCHAR(n)` | Variable-length string, up to `n` characters. Saves space if content varies in length.                  |
| `TEXT`       | Stores long strings. No predefined limit. Slightly slower for indexing/searching compared to `VARCHAR`. |

<br>

**Date & Time**

| Data Type   | Description                                         |
| ----------- | --------------------------------------------------- |
| `DATE`      | Stores only date (YYYY-MM-DD).                      |
| `TIME`      | Stores only time (HH\:MM\:SS).                      |
| `TIMESTAMP` | Stores date and time (YYYY-MM-DD HH\:MM\:SS).       |
| `INTERVAL`  | Represents a time duration (e.g., 5 days, 2 hours). |

<br>

**Boolean**

| Data Type | Description                        |
| --------- | ---------------------------------- |
| `BOOLEAN` | Stores `TRUE`, `FALSE`, or `NULL`. |

<br>

**Binary**

| Data Type      | Description                                                                   |
| -------------- | ----------------------------------------------------------------------------- |
| `BINARY(n)`    | Fixed-length binary data.                                                     |
| `VARBINARY(n)` | Variable-length binary data.                                                  |
| `BLOB`         | Binary Large Object. Used to store large binary content like files or images. |

<br>

**JSON & XML**

| Data Type | Description                                                        |
| --------- | ------------------------------------------------------------------ |
| `JSON`    | Stores JSON data as text.                                          |
| `JSONB`   | Stores JSON in binary format. Allows indexing and faster querying. |
| `XML`     | Stores XML-formatted data.                                         |

<br>
<br>
<br>

### SQL Keys

- In relational databases, keys are used to uniquely identify rows.
- They also establish relationship between tables.

<br>

**Primary Key**

- Uniquely identifies each record in a table.
- Must contain unique values.
- Cannot have `NULL` values.
- Only one primary key per table.

<br>

**Candidate Key**

- A minimal set of attributes that can uniquely identify a row.
- There can be multiple candidate keys in a table.
- One of them is chosen as the primary key.
- Example: In a `Students` table: both `roll_number` and `email` could be candidate keys.

<br>

**Alternate Key**

- Candidate keys that are not chosen as the primary key.
- Essentially, it's a leftover candidate key.

<br>

**Composite Key**

- A key formed by combining two or more columns to uniquely identify a row.
- Used when a single column is not sufficient to identify uniqueness.

**Foreign Key**

- A column or set of columns that creates a link between two tables.
- Refers to the primary key of another table.

**Super Key**

- Any set of attributes that uniquely identifies a row in a table.
- May contain extra attributes beyond what's necessary.
- Every candidate key is a super key, but not all super keys are candidate keys.

**Unique Key**

- Ensures all values in a column (or group of columns) are unique.
- Unlike primary key, it can contain a single NULL value.

<br>
<br>
<br>


- Type of relationships (one-to-one, one-to-many, many-to-many, etc.)
- Type of constraints (unique, not null, etc.)

- What is data normalization? Why is it important in DBMS?
- Explain various normal forms.
- What is denormalization? How does it differ from normalization?
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
