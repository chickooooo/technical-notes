# SQL Revision Notes

## Table Schema: `employees`

| Column Name | Data Type | Description           |
|-------------|-----------|-----------------------|
| id          | INT       | Primary key           |
| name        | VARCHAR   | Employee name         |
| department  | VARCHAR   | Department name       |
| salary      | INT       | Monthly salary (USD)  |
| hire_date   | DATE      | Date of joining       |

<br>

## Table Schema: `departments`

| Column Name | Data Type | Description           |
|-------------|-----------|-----------------------|
| dept_name   | VARCHAR   | Department name (PK)  |
| manager     | VARCHAR   | Manager's name        |

<br>
<br>
<br>

### Q1: Names of employees in 'Sales' department

```sql
SELECT name
FROM employees
WHERE department = 'Sales';
```

<br>
<br>

### Q2: Top 3 highest-paid employees (name and salary)

```sql
SELECT name, salary
FROM employees
ORDER BY salary DESC
LIMIT 3;
```

<br>
<br>

### Q3: Average salary in 'Engineering' department

```sql
SELECT AVG(salary)
FROM employees
WHERE department = 'Engineering';
```

<br>
<br>

### Q4: Total salary by department

```sql
SELECT department, SUM(salary) AS total_salary
FROM employees
GROUP BY department;
```

<br>
<br>

### Q5: Departments with total salary > 100,000

```sql
SELECT department, SUM(salary) AS total_salary
FROM employees
GROUP BY department
HAVING total_salary > 100000;
```

<br>
<br>

### Q6: List employees with their department manager

```sql
SELECT e.name, e.department, d.manager
FROM employees AS e
INNER JOIN departments AS d
ON e.department = d.dept_name;
```

<br>
<br>

### Q7: Employees earning above average salary

```sql
SELECT name
FROM employees
WHERE salary > (SELECT AVG(salary) FROM employees);
```
