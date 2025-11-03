# System Design Fundamentals

<br>
<br>
<br>
<br>
<br>

## Pagination

<br>
<br>
<br>

### Offset based pagination

- It uses `page number` and `page size` to find the required page.
- It skips records from previous `n-1` pages to get the `nth` page.

<br>

- For example, lets say we want page `3` and the page size is `10`.
- In this case, offset based pagination will skip the starting `20` records and return the next `10` records.
- Related SQL query:

```sql
SELECT * FROM products
ORDER BY created_at DESC
LIMIT 10 OFFSET 20;
```

<br>

Advantages:
- We can easily go to the `nth` page.
- Easy to implement and understand.

Disadvantages:
- Performance degrades with large offset. The database still has to scan the starting records that will be discarded later.
- If rows are inserted or deleted while paginating, some records may be skipped or can appear on multiple pages.

<br>
<br>
<br>

### Cursor based pagination

- It uses a `cursor` and `page size` to get the current page.
- A cursor is basically a **unique identifier** of the last record from the last page.
- To get the records of the current page, return `page size` records after the `cursor`.
- With every page, the `cursor` is send to the client. And the client requests the next page using the `cursor` and `page size`.

<br>

- For example, if the cursor is `2025-11-03T12:00:00Z`, then the records on the next page can be found using:

```sql
SELECT * FROM products
WHERE created_at < '2025-11-03T12:00:00Z'
ORDER BY created_at DESC
LIMIT 10;  -- page size
```

<br>

Advantages:
- Much faster for large datasets. It uses index based filtering (`WHERE` clause) instead of scanning and discarding rows.
- Consistent behaviour. Records are not skipped or duplicated even if rows are inserted or deleted.

Disadvantages:
- We cannot go directly to `nth` page. We can only get the next page or previous page.
- `cursor` needs to be encoded & decoded and is exposed to the client.
- More complex to implement compared to **Offset based pagination**.

<br>
<br>
<br>
<br>
<br>

## 

<br>
<br>
<br>
<br>
<br>
