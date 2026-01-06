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

## 

<br>
<br>
<br>
