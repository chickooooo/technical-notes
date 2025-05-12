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
select COUNT(*) AS total_admissions
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

14 - 

```sql
```

<br>
<br>

15 - 

```sql
```
