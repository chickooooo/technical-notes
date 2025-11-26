# SQL

<br>
<br>
<br>
<br>
<br>

## Index

- [Subquery](#subquery)
- [Common table expression (CTE)](#common-table-expression-cte)
- [Temp table](#temp-table)

<br>
<br>
<br>
<br>
<br>

### Subquery

- A subquery is a SELECT statement inside another SELECT statement.
- Used for filtering, comparison or generating derived results.
- Can be a scalar, row, column or a table.

```sql
-- Get all products having price greater than average price
SELECT name, price
FROM products
WHERE price > (SELECT AVG(price) from products);
```

<br>
<br>
<br>

### Common table expression (CTE)

- A CTE is a temporary named result set.
- It is created using the `WITH` clause.
- Can be referenced multiple times within the same query.
- Can create multiple CTEs within the same query.
- Exists only for the duration of that statement; not stored.
- Improves readability and modularity of complex queries.

```sql
WITH max_table AS (
  SELECT category, MAX(price)
  FROM products
  GROUP BY category
)

SELECT *
FROM products AS p
LEFT JOIN max_table AS mt
  ON p.category = mt.category;
```

<br>
<br>
<br>

### Temp table

- A physical table stored in temp database.
- Can be modified (INSERT/UPDATE/DELETE).
- Used to store intermediate results that need indexing or reuse across multiple steps.
- Better for large datasets, repeated use, or when performance tuning needs indexing.
- Lives for the session or connection.

<br>
<br>
<br>

### 
