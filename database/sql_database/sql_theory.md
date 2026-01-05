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

## 

<br>
<br>
<br>
<br>
