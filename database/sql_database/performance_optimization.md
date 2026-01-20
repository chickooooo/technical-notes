# Performance Optimization

<br>
<br>
<br>
<br>
<br>

## Index

- [Types of index](#types-of-index)

<br>
<br>
<br>
<br>
<br>

### Types of index

#### B-Tree index

- A B-Tree index stores data in a balanced tree structure.
- Here the keys are kept in a sorted order.
- Search, insert, and delete operations take logarithmic time.

<br>

Pros

- It supports **equality** and **range queries** efficiently.
- It works well with `ORDER BY` clause.
- It is the default index type in most relational databases.

<br>

Cons

- It is not efficient for searching inside complex data types like arrays or JSON.
- It provides limited benefit for very low-cardinality columns.

<br>

Usecases

- Searching by primary key or foreign key.
- Range queries with `WHERE` clause: `WHERE salary BETWEEN 50000 AND 100000`.
- Sorting queries using `ORDER BY`.

<br>
<br>

#### B-Tree vs B+ Tree

B-Tree

- Keys and data (records) are stored together in all nodes.
- When searching for a record, it can be found at internal or leaf node.
- It is less efficient for sequential data access than B+ tree.

<br>

B+ Tree

- A B+ Tree is a variation of B-Tree, optimized for databases and file systems.
- All internal nodes store only the keys.
- All data records are stored in the leaf nodes.
- Leaf nodes are linked together (like a linked list).
- More efficient for sequential data access.

<br>
<br>

#### Hash index

- A Hash index uses a hash function to map keys directly with the record.
- It is optimized for **exact match** lookups.

<br>

Pros

- It provides very fast performance for equality comparisons.
- The lookup time is close to constant time.

<br>

Cons

- It does not support range queries.
- In many databases, B-Tree indexes are generally used instead of hash indexes.

<br>

Usecases

- Exact search queries like `WHERE user_id = 12345`.

<br>
<br>
