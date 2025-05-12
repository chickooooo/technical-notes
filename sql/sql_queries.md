## SQL Queries

<br>
<br>

Update the patients table for the allergies column. If the patient's allergies is null then replace it with 'NKA'

```sql
UPDATE patients
SET allergies='NKA'
WHERE allergies IS NULL;
```

<br>
<br>

1 - Show first name, last name, and gender of patients whose gender is 'M'

```sql
SELECT
	first_name,
    last_name,
    gender
FROM patients
WHERE gender='M';
```

<br>
<br>

2 - Show first name and last name of patients who does not have allergies. (null)

```sql
SELECT
	first_name,
    last_name
FROM patients
WHERE allergies IS NULL;
```

<br>
<br>

3 - Show first name of patients that start with the letter 'C'

```sql
SELECT first_name
FROM patients
WHERE first_name LIKE "C%";
```

<br>
<br>

4 - Show first name and last name of patients that weight within the range of 100 to 120 (inclusive)

```sql
SELECT
	first_name,
    last_name
FROM patients
WHERE
	weight BETWEEN 100 AND 120;
```

<br>
<br>

5 - Update the patients table for the allergies column. If the patient's allergies is null then replace it with 'NKA'

```sql
UPDATE patients
SET allergies='NKA'
WHERE allergies IS NULL;
```

<br>
<br>

6 - Show first name and last name concatinated into one column to show their full name.

```sql
SELECT
	CONCAT(first_name, ' ',last_name) AS full_name
FROM patients; 
```

<br>
<br>

7 - Show first name, last name, and the full province name of each patient.

```sql
SELECT
	p.first_name,
    p.last_name,
    pn.province_name
FROM patients as p
LEFT JOIN province_names AS pn
ON p.province_id = pn.province_id;
```

```
If we don't specific JOIN type, by default INNER JOIN is used.
```

<br>
<br>

8 - Show how many patients have a birth_date with 2010 as the birth year.

```sql
SELECT COUNT(*)
FROM patients
WHERE YEAR(birth_date) = 2010;
```

<br>
<br>

9 - Show the first_name, last_name, and height of the patient with the greatest height.

```sql
SELECT
	first_name,
    last_name,
    height
FROM patients
ORDER BY height DESC
LIMIT 1;
```

```sql
-- using subquery
SELECT
	first_name,
    last_name,
    height
FROM patients
WHERE height = (SELECT MAX(height) FROM patients);
```

<br>
<br>

10 - Show all columns for patients who have one of the following patient_ids: 1,45,534,879,1000

```sql
SELECT *
FROM patients
WHERE patient_id IN (1,45,534,879,1000);
```

<br>
<br>

11 - Show the total number of admissions

```sql
SELECT COUNT(*) AS total_admissions
FROM admissions;
```

<br>
<br>

12 - Show all the columns from admissions where the patient was admitted and discharged on the same day.

```sql
SELECT *
FROM admissions
WHERE admission_date = discharge_date;
```

<br>
<br>

13 - Show the patient id and the total number of admissions for patient_id 579.

```sql
SELECT
	patient_id,
    COUNT(*) AS total_admissions
FROM admissions
WHERE patient_id=579
GROUP BY patient_id;  -- this line is optional
```

<br>
<br>

14 - Based on the cities that our patients live in, show unique cities that are in province_id 'NS'.

```sql
SELECT DISTINCT city AS unique_cities
FROM patients
WHERE province_id='NS';
```

<br>
<br>

15 - Write a query to find the first_name, last name and birth date of patients who has height greater than 160 and weight greater than 70

```sql
SELECT
	first_name,
    last_name,
    birth_date
FROM patients
WHERE
	height > 160
    AND weight > 70;
```

<br>
<br>

16 - Write a query to find list of patients first_name, last_name, and allergies where allergies are not null and are from the city of 'Hamilton'

```sql
SELECT
	first_name,
    last_name,
    allergies
FROM patients
WHERE
	allergies is NOT NULL
    AND city='Hamilton';
```

<br>
<br>

17 - Show unique birth years from patients and order them by ascending.

```sql
SELECT
	DISTINCT YEAR(birth_date) AS birth_year
FROM patients
ORDER BY birth_year ASC;
```

<br>
<br>

18 - Show unique first names from the patients table which only occurs once in the list.
<br>
For example, if two or more people are named 'John' in the first_name column then don't include their name in the output list. If only 1 person is named 'Leo' then include them in the output.

```sql
SELECT first_name
FROM patients
GROUP BY first_name
HAVING COUNT(*) = 1;
```

```sql
SELECT first_name
FROM (
	SELECT
  		first_name,
  		COUNT(*) AS first_name_count
  	FROM patients
  	GROUP BY first_name
)
WHERE first_name_count = 1;
```

<br>
<br>

19 - Show patient_id and first_name from patients where their first_name start and ends with 's' and is at least 6 characters long.

```sql
SELECT
	patient_id,
    first_name
FROM patients
WHERE
	first_name LIKE 's%s'
    AND LEN(first_name) >= 6;
```

<br>
<br>

20 - Show patient_id, first_name, last_name from patients whos diagnosis is 'Dementia'. Primary diagnosis is stored in the admissions table.

```sql
SELECT
	p.patient_id,
    p.first_name,
    p.last_name
FROM patients as p
JOIN admissions as a
ON p.patient_id = a.patient_id
WHERE a.diagnosis = 'Dementia';
```

```sql
SELECT
  patient_id,
  first_name,
  last_name
FROM patients
WHERE patient_id IN (
    SELECT patient_id
    FROM admissions
    WHERE diagnosis = 'Dementia'
);
```