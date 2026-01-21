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

#### Bitmap index

- A Bitmap index uses bit arrays to represent the presence of values.
- Each distinct value in a column has a bitmap indicating which rows contain that value.

<br>

Example

```
| RowID | Name  | Gender | Department |
| ----- | ----- | ------ | ---------- |
|     1 | Alice | F      | HR         |
|     2 | Bob   | M      | IT         |
|     3 | Carol | F      | IT         |
|     4 | David | M      | HR         |
|     5 | Emma  | F      | MKT        |


---------------------------------------------

Bitmap index on column 'Gender'

   1 2 3 4 5    <- row number
F: 1 0 1 0 1    <- rows containing 'F' are marked as 1
M: 0 1 0 1 0    <- rows containing 'M' are marked as 1


---------------------------------------------

Bitmap index on column 'Department'

     1 2 3 4 5
IT:  0 1 1 0 0
HR:  1 0 0 1 0
MKT: 0 0 0 0 1
```

- To find rows `WHERE Gender='F' AND Department='HR'`, we simply take `AND` of the bitmaps.

```
     1 2 3 4 5
F:   1 0 1 0 1
HR:  1 0 0 1 0  AND

     1 0 0 0 0   -> row 1 is the required row
```

<br>

Pros

- It is very efficient for low-cardinality columns.
- It performs well for filtering using `AND / OR` conditions.
- It consumes relatively little storage for low-cardinality data.

<br>

Cons

- Not suitable for ordinal columns.
- It performs poorly for high-cardinality columns.

<br>

Usecases

- Columns like gender, status, or boolean flags.

<br>
<br>

#### Inverted index

- An inverted index maps values or terms to the list of rows that contain them.
- It is commonly used in full-text search systems.
- Uses algorithms like TF-IDF (Term Frequency-Inverse Document Frequency) to keep track of only important words.

<br>

```
| Doc ID | Text                                  |
| ------ | ------------------------------------- |
|     D1 | "database systems are fast"           |
|     D2 | "database index improves performance" |
|     D3 | "fast search uses inverted index"     |


---------------------------------------------

Inverted index on column 'Text'

| Term        | Posting List (Doc IDs) |
| ----------- | ---------------------- |
| database    | D1, D2                 |
| systems     | D1                     |
| fast        | D1, D3                 |
| index       | D2, D3                 |
| improves    | D2                     |
| ...         | ...                    |


---------------------------------------------

Search: "database" -> { D1, D2 }

Search: "fast index"
    - "fast" -> { D1, D3 }
    - "index" -> { D2, D3 }
    - "fast index" -> { D1, D3 } AND { D2, D3 } -> { D3 }
```

<br>

Pros

- It is highly efficient for text search and keyword matching.
- It supports searching for words, phrases, and text patterns.

<br>

Cons

- It requires more storage space.
- It is slower to update compared to traditional indexes.

<br>

Usecases

- Full text search in documents or articles.

<br>
<br>
