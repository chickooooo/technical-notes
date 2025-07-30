## Index
- [ACID Properties](#acid-properties)
- [BASE Properties](#base-properties)
- [ACID vs BASE properties](#acid-vs-base-properties)
- [SQL Datatypes](#sql-datatypes)
- [SQL Keys](#sql-keys)
- [SQL Relationships](#sql-relationships)
- [SQL Constraints](#sql-constraints)
- [Data Normalization](#data-normalization)
- [Various Normal Forms](#various-normal-forms)
- [Data denormalization](#data-denormalization)
- [Referential Integrity](#referential-integrity)

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

### SQL Relationships

**One-to-One**

- One record in Table A is related to one and only one record in Table B.
- Used for splitting large tables for modularity.
- Example: Each user has one unique passport.

<br>

**One-to-Many**

- One record in Table A can relate to many records in Table B.
- Example: One customer can place many orders.

<br>

**Many-to-Many**

- Multiple records in Table A can relate to multiple records in Table B.
- An intermediate table is required to map this relationship.
- Example: Students can enroll in many classes, and each class has many students.

<br>
<br>
<br>

### SQL Constraints

**NOT NULL**

- Ensures a column cannot have NULL values.
- Example: `name VARCHAR(50) NOT NULL;`

<br>

**UNIQUE**

- Ensures all values in a column (or group of columns) are distinct.
- Allows one NULL (in most RDBMS).
- Example: `email VARCHAR(100) UNIQUE;`

<br>

**PRIMARY KEY**

- A combination of NOT NULL + UNIQUE.
- Only one primary key per table.
- Example: `id INT PRIMARY KEY;`

<br>

**FOREIGN KEY**

- Enforces **referential integrity** between tables.
- Links a column to the primary key of another table.
- Example: `customer_id INT REFERENCES customers(id);`

<br>

**CHECK**

- Ensures values in a column meet a specific condition.
- Example: `age INT CHECK (age >= 18);`

<br>

**DEFAULT**

- Sets a default value if no value is provided during insertion.
- Example: `status VARCHAR(20) DEFAULT 'active';`

<br>
<br>
<br>

### Data Normalization

- It is the process of reducing data redundancy and improve data integrity.
- It involves dividing the entities into two or more tables and defining relationship between these tables.

**Advantages**

- Eliminates Data Redundancy - No duplicate data across tables.
- Improves Data Integrity - Consistent and accurate data.
- Eliminates Update Anomalies - Only 1 record needs to be updated.

**Disadvantages**

- Complex Queries - Need to use joins to query data together.
- Using heavy joins induces performance overhead.
- Not ideal for ready-heavy system.

<br>
<br>
<br>

### Various Normal Forms

**First Normal Form (1NF)**

- Each record in a table should be uniquely identifiable.
- Each value in a column should be indivisible.

<br>

Violates 1NF 

| employee_id | name  | address              |
| ----------- | ----- | -------------------- |
| 1           | Alex  | Baker street, London |
| 2           | Brian | Townhall, Wales      |

<br>

Follows 1NF

| employee_id | name  | street_name  | city   |
| ----------- | ----- | ------------ | ------ |
| 1           | Alex  | Baker street | London |
| 2           | Brian | Townhall     | Wales  |

<br>
<br>

**Second Normal Form (2NF)**

- Table must be in 1NF.
- Applicable only in case of composite primary keys.
- Each non-key column in the table should dependent on the whole primary key.
- In other words, there should be no Partial Dependencies in the table.

Partial Dependency

- Table's primary key should be made up of two or more columns.
- If any non-primary key column depends only on a part of primary key, then we have partial dependency.

<br>

Violates 2NF 

| student_id | subject_id | marks | teacher |
| ---------- | ---------- | ----- | ------- |
| 1          | 101        | 60    | Ana     |
| 2          | 101        | 80    | Jasmine |

- Primary key is (student_id, subject_id).
- `teacher` column depends only on `subject_id` and not `student_id`.
- `teacher` column should be move to `subject` table.

<br>
<br>

**Third Normal Form (3NF)**

- Table must be in 2NF.
- It should not have Transitive Dependency.

Transitive Dependency.

- When a non-primary key column depends on another non-primary key column.

<br>

Violates 3NF 

| student_id | subject_id | marks | exam_type | total_marks |
| ---------- | ---------- | ----- | --------- | ----------- |
| 1          | 101        | 60    | Written   | 80          | 
| 2          | 101        | 18    | Oral      | 20          |

- Primary key is (student_id, subject_id).
- `exam_type` column depends on complete primary key.
- `total_marks` column depends only on `exam_type`.
- We should create a separate table for exam entity (exam_type, total_marks, exam_duration, etc.)

<br>
<br>

**Boyce-Codd Normal Form (BCNF)**
**Fourth Normal Form (4NF)**
**Fifth Normal Form (5NF)**

<br>
<br>
<br>

### Data denormalization

- We intentionally introduce redundancy into our database.
- It is done to reduce join queries and increase read performance.

**Advantages**

- Faster Read Performance
- Simpler Queries
- Reduced JOIN Overhead

**Disadvantages**

- Increased Data Redundancy
- Risk of Inconsistencies 
- More Complex Updates
- Higher Storage Usage

<br>
<br>
<br>

### Referential Integrity

- Referential integrity means that a foreign key in one table must either:
    - Match a primary key in another table
    - OR be null
- It ensures that relationships between records are valid and there are no orphan records.
- Orphan record means a reference pointing to a non-existing row.
- Referential integrity is enforced through `FOREIGN KEY` constraint.

<br>
<br>
<br>

### 

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
