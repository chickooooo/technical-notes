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
- [Window functions](#window-functions)

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

### Window functions

- A window function performs a calculation across a set of rows that are related to the current row, without collapsing rows like **GROUP BY** does.
- Think of them as: "Aggregations without losing individual rows."

<br>
<br>

#### Basic syntax

```sql
function_name(column) OVER (
  PARTITION BY ...
  ORDER BY ...
)
```

- `function_name`: SUM, AVG, COUNT, ROW_NUMBER, RANK, LAG, etc.
- `PARTITION BY`: Like GROUP BY, divides rows into groups (optional).
- `ORDER BY`: Defines order of rows inside each partition (optional, but common).
- `OVER`: Tells SQL that this is a window function.

<br>
<br>

#### `ROW_NUMBER()`

- Assigns a unique number to each row within a partition.

```sql
SELECT
  name,
  department,
  salary,
  ROW_NUMBER() OVER (PARTITION BY department ORDER BY salary) AS row_num
FROM employees;
```

| name  | department | salary | row_num |
| ----- | ---------- | ------ | ------- |
| Alice | HR         | 60000  | 1       |
| Bob   | HR         | 80000  | 2       |
| Carol | IT         | 70000  | 1       |
| Erin  | IT         | 75000  | 2       |
| Dan   | IT         | 80000  | 3       |

<br>
<br>

#### `RANK()` and `DENSE_RANK()` 

- `RANK()`: Allows gaps while numbering (1,2,2,4).
- `DENSE_RANK()`: Has no gaps while numbering (1,2,2,3).

```sql
SELECT
  name,
  department,
  salary,
  RANK() OVER (PARTITION BY department ORDER BY salary) AS rnk,
  DENSE_RANK() OVER (PARTITION BY department ORDER BY salary) AS dense_rnk
FROM employees;
```

| name  | department | salary | rnk | dense_rnk |
| ----- | ---------- | ------ | --- | --------- |
| Alice | HR         | 50000  | 1   | 1         |
| Bob   | HR         | 60000  | 2   | 2         |
| Carol | IT         | 70000  | 1   | 1         |
| Erin  | IT         | 70000  | 1   | 1         |
| Dan   | IT         | 80000  | 3   | 2         |

<br>
<br>

#### `SUM()`

- Calculate running total of a column.

```sql
SELECT
  name,
  salary,
  SUM(salary) OVER (ORDER BY salary) AS running_total
FROM employees;
```

| name  | salary | running_total |
| ----- | ------ | ------------- |
| Bob   | 50000  | 50000         |
| Alice | 60000  | 110000        |
| Carol | 70000  | 180000        |
| Erin  | 75000  | 255000        |
| Dan   | 80000  | 335000        |

<br>
<br>

#### `AVG()` and `COUNT()`

- `AVG()`: Calculate average of a column per partition.
- `COUNT()`: Calculate number of rows per partition.

```sql
SELECT
  name,
  department,
  salary,
  AVG(salary) OVER (PARTITION BY department) AS avg_salary,
  COUNT(*) OVER (PARTITION BY department) AS rows_count
FROM employees
ORDER BY salary;
```

| name  | department | salary | avg_salary | rows_count |
| ----- | ---------- | ------ | ---------- | ---------- |
| Bob   | HR         | 50000  | 55000      | 2          |
| Alice | HR         | 60000  | 55000      | 2          |
| Carol | IT         | 70000  | 75000      | 3          |
| Erin  | IT         | 75000  | 75000      | 3          |
| Dan   | IT         | 80000  | 75000      | 3          |

<br>
<br>

#### `LAG()` and `LEAD()`

- `LAG()`: Used to access value of a column from the previous row.
- `LEAD()`: Used to access value of a column from the next row.

```sql
SELECT
  name, 
  department,
  salary,
  LAG(salary) OVER (
    PARTITION BY department
    ORDER BY salary
  ) AS prev_salary
FROM employees;
```

| name  | department | salary | prev_salary |
| ----- | ---------- | ------ | ----------  |
| Bob   | HR         | 50000  | NULL        |
| Alice | HR         | 60000  | 50000       |
| Carol | IT         | 70000  | NULL        |
| Erin  | IT         | 75000  | 70000       |
| Dan   | IT         | 80000  | 75000       |

<br>
<br>

#### Window frames

- Window frames define which rows participate in the window calculation.

<br>
<br>
<br>

### 

<br>
<br>
<br>
