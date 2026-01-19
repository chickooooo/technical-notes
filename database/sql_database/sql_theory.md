# SQL Theory

<br>
<br>
<br>
<br>
<br>

## Index

- [What is a SQL database?](#what-is-a-sql-database)
- [SQL datatypes](#sql-datatypes)
- [SQL relationships](#sql-relationships)
- [Referential integrity](#referential-integrity)
- [SQL keys](#sql-keys)
- [Database views](#database-views)
- [Non-materialised view](#non-materialised-view)
- [Materialised view](#materialised-view)
- [Data normalization](#data-normalization)

<br>
<br>
<br>
<br>
<br>

## What is a SQL database?

- A SQL database is a structured data storage system.
- It stores data in tables. Each row represents a record and each column represents an attribute of that record.
- Tables can be linked with each another using relations.
- It uses SQL query language for querying and managing the data.
- Example database: SQLite, MySQL, PostgreSQL, Oracle, etc.

<br>
<br>
<br>
<br>

## SQL datatypes

- Each SQL table column can have one of the following data types:
    - Numeric
    - String
    - Date & Time
    - Boolean
    - Binary
    - JSON / XML

<br>
<br>

### Numeric

- `INTEGER`: Whole numbers usually 4 bytes. Range: -2^31 to 2^31 - 1.
- `SMALLINT`: Smaller range integers with 2 bytes.
- `BIGINT`: Large integer values with 8 bytes. 
- `DECIMAL(p,s)`: Fixed-point numbers with precision (p) and scale (s). Suitable for monetary values.
- `FLOAT`: Floating-point numbers. Use when strict precision isn't needed. 

<br>
<br>

### String

- `CHAR(n)`: Fixed-length character string. `CHAR(10)` always takes 10 bytes.  
- `VARCHAR(n)`: Variable-length string, up to `n` characters. Saves space if content varies in length.
- `TEXT`: Stores long strings. No predefined limit. Slightly slower for indexing/searching compared to `VARCHAR`.

<br>
<br>

### Date & Time

- `DATE`: Stores only the date (YYYY-MM-DD).
- `TIME`: Stores only the time (HH:MM:SS).
- `TIMESTAMP`: Stores both date and time (YYYY-MM-DD HH:MM:SS).
- `INTERVAL`: Represents a time duration (e.g., 5 days, 2 hours).

<br>
<br>

### Boolean

- `BOOLEAN`: Stores `true`, `false`, or `NULL`.

<br>
<br>

### Binary

- `BINARY(n)`: Fixed-length binary data.
- `VARBINARY(n)`: Variable-length binary data.
- `BLOB`: Binary Large Object. Used to store large binary content like files or images.

<br>
<br>

### JSON / XML

- `JSON`: Stores JSON data as text.
- `JSONB`: Stores JSON in binary format. Allows indexing and faster querying.
- `XML`: Stores XML-formatted data.

<br>
<br>
<br>
<br>

## SQL relationships

- A relationship in SQL defines how records of one table are related to records of another table.
- There are 3 main types of SQL relations:
    - one-to-one relation
    - one-to-many relation
    - many-to-many relation

<br>
<br>

### One-to-one relation

- One record in Table A is related to one and only one record in Table B.
- Used for splitting large tables for modularity.
- Example: Each user has one unique passport.

<br>
<br>

### One-to-many relation

- One record in Table A can relate to many records in Table B.
- Example: One customer can place many orders.

<br>
<br>

### Many-to-many relation

- Multiple records in Table A can relate to multiple records in Table B.
- An intermediate table is required to map this relationship.
- Example: Students can enroll in many classes, and each class has many students.

<br>
<br>
<br>
<br>

## Referential integrity

- Referential integrity means that a foreign key in one table must either:
    - Match a primary key in another table
    - OR be `NULL`
- It ensures that relationships between records are valid and there are no orphan records.
- Orphan record means a reference pointing to a non-existing row.
- Referential integrity is enforced through `FOREIGN KEY` constraint.

<br>
<br>
<br>
<br>

## SQL keys

- Keys in SQL are constraints used to uniquely identify rows in a table.
- They are also used to establish and enforce relationships between tables.
- Keys help maintain data integrity and improve query performance through indexing.

<br>
<br>

### Unique key

- Ensures all values in a column are unique.
- `NULL` values are allowed (usually one).
- Used to prevent duplicate data.
- Example: `email` in users table.

<br>
<br>

### Primary key

- Uniquely identifies each record in a table.
- Cannot be `NULL`.
- Only one primary key per table.
- Example: `student_id` in students table.

<br>
<br>

### Candidate key

- A column (or a minimal set of columns) that can become a primary key.
- Uniquely identifies records in a table.
- A table can have multiple candidate keys.
- Example: `passport_number`, `email` in users table.

<br>
<br>

### Alternate key

- A candidate key not chosen as the primary key.
- Example: If `student_id` is primary, `email` becomes an alternate key.

<br>
<br>

### Composite (Compound) key

- A key made using two or more columns together.
- Used when a single column is not enough to ensure uniqueness.
- Example: `(order_id, product_id)`

<br>
<br>

### Super key

- Any combination of columns that uniquely identifies a row.
- May include extra columns.
- All candidate keys are super keys, but not all super keys are candidate keys.

<br>
<br>

### Natural key

- A key derived from real-world data.
- Has some business meaning.
- Example: `admission_id`, `email`, etc.
- Can change over time (less stable).

<br>
<br>

### Surrogate key

- An artificially created key.
- Has no business meaning.
- Usually auto-incremented.
- Example: `id`, `UUID`, etc.
- More stable and commonly used.

<br>
<br>

### Foreign key

- A column that references the primary key of another table.
- Maintains **referential integrity**.
- Can be `NULL`, depending on requirements.
- Example: `student_id` in enrollments table referencing `id` in students table.

<br>
<br>
<br>
<br>

## Database views

- A database view is just a saved named SQL query.
- There are 2 types of database views:
    - Non-materialised view (does not store data)
    - Materialised view (stores data)

<br>
<br>
<br>
<br>

## Non-materialised view

- A non-materialised view is just a saved named SQL query.
- It is does not stores data separately.
- Each time it is queried, it executes the underlying `SELECT` statement.

<br>

```sql
CREATE VIEW view_name AS
    SELECT column1, column2
    FROM table_name
    WHERE condition;
```

<br>
<br>

### Advantages

- When queried, always responds with up-to-date data.
- No additional storage required.
- Lightweight and simple.

<br>
<br>

### Disadvantages

- Complex joins are performed everytime it is queried.
- No performance gain over normal queries.
- Cannot be indexed.

<br>
<br>

### Usecases

- Restricting access to sensitive data. Only required data is made visible through the view.
- Simplifying complex quries and providing a single source of usage with optimised implementation.

<br>
<br>
<br>
<br>

## Materialised view

- A materialized view stores the query result physically on the disk.
- It basically caches the query result.

---

- This cached result is not refreshed automatically when the underlying data changes.
- A refresh technique needs to be specified.
- Provides eventual consistency.
- Can support writes and updates for simple queries without joins.

<br>

```sql
CREATE MATERIALIZED VIEW view_name AS
    SELECT column1, column2
    FROM table_name
    WHERE condition;
```

<br>
<br>

### Advantages

- Improved performance: No complex joins are executed in realtime. Stored data is used.
- Columns of a materialised view can be indexed for faster queries.

<br>
<br>

### Disadvantages

- Extra space is used to store cached results.
- Data is not up-to-date until next refresh cycle.
- Refreshing data adds more complexity.
- Not ideal for write heavy workloads.

<br>
<br>

### Refresh strategies

- Full refresh:
    - Recomputes the entire materialized view from scratch.
    - Expensive for large datasets.
- Incremental refresh:
    - Updates only changed data since the last refresh.
    - Much faster than full refresh.
    - More complex and requires change tracking.
- Scheduled refresh:
    - Refresh runs at fixed intervals (cron-like).
    - Data is stale between runs.
    - Ideal for reporting dashboard views.
- On-demand refresh:
    - Refresh occurs only when explicitly triggered.
    - Ideal for pipeline driven operations.

<br>
<br>

### Usecases

- Analytics & reporting dashbaords.
- For read heavy systems with less frequent writes.
- When eventual consistency is acceptable.

<br>
<br>
<br>
<br>

## Data normalization

- It is the process of reducing data redundancy and improve data integrity.
- It involves dividing the entities into two or more tables and defining relationship between these tables.

<br>

Advantages

- Eliminates Data Redundancy: No duplicate data across tables.
- Improves Data Integrity: Consistent and accurate data.
- Eliminates Update Anomalies: Only 1 record needs to be updated.

<br>

Disadvantages

- Complex Queries: Need to use joins to query data together.
- Low Performance: Using heavy joins induces performance overhead.
- Not ideal for ready-heavy system.

<br>
<br>

### Various normal forms

First normal form (1NF)

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

---

<br>

Second normal form (2NF)

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

---

<br>

Third normal form (3NF)

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

---

<br>

- Boyce-Codd Normal Form (BCNF)
- Fourth normal form (4NF)
- Fifth normal form (5NF)

<br>
<br>
<br>
<br>

## 

<br>
<br>
<br>
<br>
