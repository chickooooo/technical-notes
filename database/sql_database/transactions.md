# SQL Transactions

<br>
<br>
<br>
<br>
<br>

## Index

- [SQL Transaction](#sql-transaction)
- [ACID properties](#acid-properties)
- [Transaction rollback](#transaction-rollback)
- [Transaction savepoint](#transaction-savepoint)
- [Isolation levels](#isolation-levels)

<br>
<br>
<br>
<br>
<br>

## SQL Transaction

- A transaction is a sequence of one or more SQL operations (like INSERT, UPDATE, etc.).
- These operations are treated as one single unit of work.
- A SQL transaction follows ACID properties.

<br>

```sql
BEGIN TRANSACTION;   -- Start the transaction

UPDATE accounts SET balance = balance - 500 WHERE id = 1;
UPDATE accounts SET balance = balance + 500 WHERE id = 2;

COMMIT;              -- Save all changes
```

```sql
ROLLBACK;            -- Undo all changes
```

<br>
<br>
<br>

## ACID properties

- Atomicity:
    - All operations inside a transaction succeed or all fail.
    - There is no partial success (like 3 operations succeed and 2 failed).
    - If one operation fails, the entire transaction is rolled back.
- Consistency:
    - The database moves from one valid state to another.
    - Rules like constraints are always preserved.
- Isolation:
    - Multiple Transactions don't interfere with each other.
    - The final result is as if they were run one after another (serially).
- Durability:
    - Once committed, the changes are permanent.
    - The data is preserved even if the system crashes.

<br>
<br>
<br>

## Transaction rollback

- A transaction rollback is the process of undoing all database changes done since the start of the transaction.
- These are all uncommitted changes of that transaction.
- A transaction rollback can be triggered manually or automatically by the database.
- A rollback can be:
    - a full rollback (till the start of transaction).
    - OR partial rollback (till the last savepoint).

<br>
<br>

### Manual rollback

- The developer or the application explicitly decides to undo the transaction.
- This can happen due to:
    - A business rule fails (e.g., insufficient balance).
    - Validation error, etc.

```sql
BEGIN TRANSACTION;

UPDATE accounts SET balance = balance - 1000 WHERE id = 1;

-- Condition fails
ROLLBACK;
```

### Automatic rollback

- The database automatically rolls back when an error occurs.
- Common causes:
    - SQL error (syntax or runtime error)
    - Constraint violation (PRIMARY KEY, FOREIGN KEY, CHECK, NOT NULL)
    - Deadlock detected

```sql
BEGIN TRANSACTION;

INSERT INTO users (id, email) VALUES (1, 'a@test.com');
INSERT INTO users (id, email) VALUES (1, 'b@test.com'); -- Duplicate PK

-- Database automatically rolls back
```

<br>
<br>

### How rollback is done?

- When a transaction runs, the database keeps undo logs.
- On rollback, it uses these logs to reverse the changes.

<br>
<br>
<br>

## Transaction savepoint

- A savepoint is a named point inside a transaction that allows us to roll back part of the transaction, instead of rolling back the entire transaction.
- It works as a checkpoint within a transaction.

<br>

```sql
BEGIN TRANSACTION;

-- Operation 1
SAVEPOINT op_1_savepoint;

-- Something goes wrong here
ROLLBACK TO op_1_savepoint;

COMMIT;
```

<br>
<br>

### Key points

- Savepoints allow partial rollback of a transaction.
- They can be used only inside a transaction.
- Rolling back to a savepoint:
    - Preserves changes done before that savepoint.
    - Removes changes done after that savepoint.

<br>
<br>
<br>

## Isolation levels

- Isolation levels define how and when the changes made by one transaction become visible to other transactions.
- Depending upon the isolation level selected, it has impact on following things:
    - Concurrency
    - Data consistency
    - Database Performance

<br>
<br>

### Why isolation levels are needed?

- When transactions run concurrently, the following problems can happen.
- Isolation levels decide which of these problems are allowed.

#### Dirty Read

- Reading uncommitted data from another transaction.

Example

- Transaction A updates salary from `5000` to `8000` but does not commit yet.
- Transaction B reads salary as `8000`.
- Transaction A rolls back.
- Transaction B read wrong (dirty) data.

<br>

#### Non-Repeatable Read

- Reading the same row twice and getting different values.

Example

- Transaction A reads salary as `5000`.
- Transaction B updates salary to `7000` and commits.
- Transaction A reads salary again as `7000`.

<br>

#### Phantom Read

- Re-running a query returns extra or missing rows.

Example

- Transaction A runs: `SELECT * FROM employees WHERE dept = 'HR'`. Gets 5 rows.
- Transaction B inserts a new HR employee and commits.
- Transaction A runs the same query again. Gets 6 rows.
- A new **phantom row** appeared.

<br>
<br>

### Isolation levels (lowest to highest isolation)

#### READ UNCOMMITTED

- A transaction can read data that has not been committed yet by another transaction.
- No locking or minimal locking.

Example

- Transaction A updates salary but doesn't commit yet.
- Transaction B reads that updated salary resulting in a dirty read.

Features

- Allows dirty reads.
- Allows non-repeatable reads.
- Allows phantom reads.

---

- Low consistency, high performance. 
- Rarely used.

<br>

#### READ COMMITTED

- A transaction can read only committed data.

Example

- Transaction A reads salary = 5000.
- Transaction B updates salary to 6000 but does not commit yet.
- Transaction A reads again, sees 5000.

Features

- Does not allow dirty reads.
- Allows non-repeatable reads.
- Allows phantom reads.
- Good balance of consistency and performance.
- Default in PostgreSQL, Oracle, SQL Server.

<br>

#### REPEATABLE READ

- Once a row is read, it cannot be changed by other transactions until the current transaction finishes.
- Reading same row always returns same data.
- It uses row level read locks which blocks writing but allows other transactions to read that row.

Example

- Transaction A reads salary = 5000.
- Transaction B reads salary = 5000.
- Transaction B tries to update, but is blocked.
- Transaction A reads again and gets 5000.

Features

- Prevents dirty reads.
- Prevents non-repeatable reads.
- Allows phantom reads.
- Used when consistent reads are important.
- Default in MySQL, InnoDB.

<br>

#### SERIALIZABLE

- Transactions run as if they are executed one after another.
- Uses strict locking.

Example

- If one transaction is reading a range of rows, no other transaction can insert or update rows in that range.

Features

- Prevents dirty reads.
- Prevents non-repeatable reads.
- Prevents phantom reads.
- Maximum data consistency but slowest performance.
- Used in financial systems.

<br>
<br>

### Key points

- Isolation level controls visibility of data between transactions.
- Higher isolation means better consistency but lower performance.
- `READ COMMITTED` is most commonly used.
- `SERIALIZABLE` is safest but slowest.

<br>
<br>
<br>

## 

<br>
<br>
<br>
