# Django ORM

<br>
<br>
<br>
<br>
<br>

## Example Models

```python
from django.db import models

class MyUser(models.Model):
    name = models.CharField(max_length=64)

class MyProduct(models.Model):
    name = models.CharField(max_length=64)
    price = models.PositiveIntegerField()

class MyCart(models.Model):
    user = models.ForeignKey(to=MyUser, on_delete=models.CASCADE)
    products = models.ManyToManyField(to=MyProduct, through="MyCartProduct")

class MyCartProduct(models.Model):
    cart = models.ForeignKey(to=MyCart, on_delete=models.CASCADE)
    product = models.ForeignKey(to=MyProduct, on_delete=models.CASCADE)
    quantity = models.PositiveIntegerField(default=1)
```

<br>
<br>
<br>
<br>
<br>

## Model fields

<br>
<br>

### On Delete
- `on_delete=` parameter specifies what to do if the user is deleted.
- `models.CASCADE` means we should delete all the records referencing to that user.
- There are other options as well. Like:
    - `models.DO_NOTHING` means Django will do nothing and it’s up to the database to handle the deletion.
    - `models.SET_NULL` means set the foreign key to NULL.
    - `models.SET_DEFAULT` means set the foreign key to its default value.
    - `models.SET(...)` means set the foreign key to a given value or callable.
    - `models.PROTECT` means prevent deletion of the referenced object by raising a ProtectedError.

<br>
<br>

### Many to Many field
- In `ManyToManyField`, `through=` parameter is used to specify the schema for the intermediate table.
- This intermediate table should have columns referencing to both the models.

<br>

Setting ManyToManyField
- For `MyCart`, we first need to create the cart record and then assign products to that cart in the next statement.
- For assigning products we use `cart.products.set(some_products)` method.

<br>
<br>
<br>
<br>
<br>

## Create APIs

<br>
<br>

### Create objects

```python
user = MyUser.objects.create(name="Alice")
```

<br>
<br>

### Bulk create objects

```python
products = MyProduct.objects.bulk_create([
    MyProduct(name="Banana", price=20),
    MyProduct(name="Pineapple", price=40),
])
```

- It is not atomic by default. It may partially succeed depending on the database and error type.
- To make it atomic, wrap it with `transaction.atomic()`.

<br>
<br>
<br>
<br>

## Get APIs

<br>
<br>

### Get all objects

```python
all_users = MyUser.objects.all()
```

<br>
<br>

### Get a single object

```python
user = MyUser.objects.get(id=2)
```

- If matching record is not found, raises `MyUser.DoesNotExist` exception.
- If multiple records are found, raises `MyUser.MultipleObjectsReturned` exception.

<br>
<br>

### Filter objects

```python
users = MyUser.objects.filter(name="Alice")  # single filter
products = MyProduct.objects.filter(name="Apple", price=10)  # multiple filters
```

- Returns `QuerySet` of objects.
- If no records match the given condition(s), returns empty `QuerySet`.

<br>
<br>

### Exclude objects

```python
# get products having price greater than 0
# exclude products with id = 2 from above products
products = MyProduct.objects.filter(price__gt=0).exclude(id=2)
```

<br>
<br>

### Ordering objects

```python
# get all users order by id desc
users = MyUser.objects.all().order_by("-id")
```

- If `-` sign is used, order by `DESC`.
- If no sign is used, order by `ASC`. 

<br>
<br>

### Limiting objects

```python
# get first 3 objects
users = MyUser.objects.all()[:3]
```

- Django's ORM translates the slicing into an SQL LIMIT clause automatically.

<br>
<br>

### Counting objects

```python
count = MyProduct.objects.filter(price__gt=10).count()
```

- Django's ORM uses `COUNT(*)` query behind the scenes to get the count.

<br>
<br>

### Check if object exists

```python
exists = MyProduct.objects.filter(name="Banana").exists()
```

<br>
<br>

### Get first object

```python
first = MyProduct.objects.filter(name="Banana").first()
```

- If one or more objects match the filters, the first object is returned.
- If no object match the filters, `None` is returned.

<br>
<br>

### Get last object

```python
last = MyProduct.objects.filter(name="Banana").last()
```

- If one or more objects match the filters, the last object is returned.
- If no object match the filters, `None` is returned.

<br>
<br>

### Get specific columns

```python
# using .values()
users = MyUser.objects.filter(name="Alice").values("id", "name")
```

- Returns a queryset of dicts containing only required columns.
- Example: `<QuerySet [{'id': 1, 'name': 'Alice'}]>`.

<br>

```python
# using .values_list()
users = MyUser.objects.filter(name="Alice").values_list("id", "name")
```

- Returns a queryset of tuples containing only required columns.
- It omits the column names and preserves the column order.
- Example: `<QuerySet [(1, 'Alice')]>`.

<br>
<br>
<br>
<br>
<br>

## Field Lookups

- Field lookups are suffixes appended to field names using double underscores `(__)` to perform specific types of queries.
- They allow you to do more than just equality checks, such as comparisons, pattern matching, null checks, etc.

<br>
<br>

### exact
- Purpose: Exact match (case-sensitive for strings)
- Django: `filter(name__exact="Alice")`
- SQL: `WHERE name = 'Alice'`

<br>
<br>

### iexact
- Purpose: Exact match, case-insensitive (for strings)
- Django: `filter(name__iexact="alice")`
- SQL: `WHERE LOWER(name) = LOWER('alice')`

<br>
<br>

### contains
- Purpose: Substring match (case-sensitive)
- Django: `filter(name__contains="lic")`
- SQL: `WHERE name LIKE '%lic%'`

<br>
<br>

### icontains
- Purpose: Substring match, case-insensitive
- Django: `filter(name__icontains="lic")`
- SQL: `WHERE LOWER(name) LIKE LOWER('%lic%')`

<br>
<br>

### startswith, istartswith
- Purpose: Starts with (case-sensitive / case-insensitive)
- Django: `filter(name__startswith="Al")`
- SQL: `WHERE name LIKE 'Al%'`
- Django: `filter(name__istartswith="al")`
- SQL:` WHERE LOWER(name) LIKE LOWER('al%')`

<br>
<br>

### endswith, iendswith
- Purpose: Ends with (case-sensitive / case-insensitive)
- Django: `filter(name__endswith="ce")`
- SQL: `WHERE name LIKE '%ce'`
- Django: `filter(name__iendswith="CE")`
- SQL: `WHERE LOWER(name) LIKE LOWER('%ce')`

<br>
<br>

### in
- Purpose: Value is in a list
- Django: `filter(name__in=["Alice", "Bob"])`
- SQL: `WHERE name IN ('Alice', 'Bob')`

<br>
<br>

### gt, gte
- Purpose: Greater than, Greater than or equal to
- Django: `filter(age__gt=20)`
- SQL: `WHERE age > 20`
- Django: `filter(age__gte=20)`
- SQL: `WHERE age >= 20`

<br>
<br>

### lt, lte
- Purpose: Less than, Less than or equal to
- Django: `filter(age__lt=30)`
- SQL: `WHERE age < 30`
- Django: `filter(age__lte=30)`
- SQL: `WHERE age <= 30`

<br>
<br>

### range
- Purpose: Value is within a range
- Django: `filter(age__range=(18, 25))`
- SQL: `WHERE age BETWEEN 18 AND 25`
- Here 18 & 25 are both inclusive.
- SQL: `WHERE age >= 18 AND age <= 25`

<br>
<br>

### isnull
- Purpose: Field is (or is not) NULL
- Django: `filter(email__isnull=True)`
- SQL: `WHERE email IS NULL`
- Django: `filter(email__isnull=False)`
- SQL: `WHERE email IS NOT NULL`

<br>
<br>

### year, month, day
- Purpose: Extract parts of a DateField or DateTimeField
- Django: `filter(created_at__year=2023)`
- SQL: `WHERE EXTRACT(YEAR FROM created_at) = 2023`

<br>
<br>

### weekday
- Purpose: Check day of the week in a DateField or DateTimeField
- Django: `filter(created_at__week_day=6)`
- SQL: `WHERE EXTRACT(DOW FROM created_at) = 5`
- Note: Django uses 1 = Sunday, 7 = Saturday
- SQL (especially PostgreSQL) uses 0 = Sunday, 6 = Saturday

<br>
<br>

### regex, iregex
- Purpose: Regex pattern match (case-sensitive / case-insensitive)
- Django: `filter(name__regex=r'^[A-Z]')`
- SQL: `WHERE name ~ '^[A-Z]'`
- Django: `filter(name__iregex=r'^[a-z]')`
- SQL: `WHERE name ~* '^[a-z]'`

<br>
<br>
<br>
<br>
<br>

## Delete APIs

<br>
<br>

### Delete object

```python
deleted = MyProduct.objects.get(name="Banana").delete()
```

<br>

### Delete multiple objects

```python
deleted = MyProduct.objects.filter(id__gt=3).delete()
```

- `pre_delete` and `post_delete` signal will get triggered in this case.

<br>
<br>
<br>
<br>
<br>

## Update APIs

<br>
<br>

### Update single object

```python
product = MyProduct.objects.get(id=5)
product.price = 100
product.save()
```

<br>
<br>

### Update multiple objects

```python
products = MyProduct.objects.filter(id__gt=3)
products.update(price=22)
```

- `.update()` does not call `.save()` on each instance.
- That means, signals like `pre_save` or `post_save` won't trigger.

<br>
<br>
<br>
<br>
<br>

## Q objects
- Q objects are used to build complex queries with `OR`, `AND` and `NOT` conditions.
- Import Q object like this: `from django.db.models import Q`.
- We can use `&` (AND), `|` (OR) and `~` (NOT) operators.


<br>
<br>

### OR query

```python
q = Q(id__lt=2) | Q(price=22)
products = MyProduct.objects.filter(q)
```

<br>
<br>

### AND query

```python
q = Q(id__lte=4) & Q(price=22)
products = MyProduct.objects.filter(q)
```

<br>
<br>

### NOT query

```python
q = Q(price=22) & ~Q(name="Banana")
products = MyProduct.objects.filter(q)
```

<br>
<br>

### Complex query

```python
q = (Q(id__lte=4) & Q(price=22)) | Q(price=30)
products = MyProduct.objects.filter(q)
```

<br>
<br>
<br>
<br>
<br>

## Related Objects

<br>
<br>

### Select Related

- Works with `ForeignKey()` and `OneToOneField()` relationships.
- Performs a join and includes the related object in the original query.
- Solves the N+1 query problem.

```python
# user data is also fetched when fetching cart data
cart = MyCart.objects.select_related("user").get(id=1)
```

```python
# fetching multiple foreign records
cart_product = MyCartProduct.objects.select_related("cart", "product").get(id=1)
```

<br>
<br>

### Prefetch Related

- Works with `ManyToManyField()` and reverse `ForeignKey()` relationships.
- Executes separate queries and combines them in Python.
- Also solves the N+1 query problem, but via multiple queries instead of SQL joins.

```python
# Prefetch products when fetching carts (M2M relationship)
carts = MyCart.objects.prefetch_related("products")
```

```python
# Prefetch MyCartProduct items for a cart (reverse FK relationship)
cart = MyCart.objects.prefetch_related("mycartproduct_set").get(id=1)
```
- `mycartproduct_set` is the default related name for the reverse FK from `MyCartProduct` to `MyCart`.

<br>
<br>

### Select Related vs Prefetch Related

- `select_related()` works by creating an SQL join and including the fields of the related object in the SELECT statement.
- That means, `select_related()` gets the related objects in the same database query.
- `prefetch_related()` works by executing separate queries for each related model and then combines the results in Python.
- It is used when a SQL JOIN isn't practical or possible, such as with **ManyToMany** or **reverse ForeignKey** relationships.

<br>
<br>

### N+1 Query Problem

- Happens when your code makes 1 query to get N objects, and then N additional queries to fetch related data for each object.
- Total queries: 1 (main) + N (users) = N+1 queries
- To solve this problem, use:
    - `select_related()` for **ForeignKey** and **OneToOne** relationship.
    - `prefetch_related()` for **ManyToMany** and **Reverse ForeignKey** relationships.

<br>
<br>
<br>

### Querying using related object

```python
# get all carts belonging to user named 'Alice'
carts = MyCart.objects.filter(user__name='Alice')
```

- Django performs SQL join behind the scenes when querying using related objects.

<br>
<br>
<br>
<br>
<br>

## Aggregation
- The `.aggregate()` method is used to perform an aggregate calculation over a **Queryset** and return a dictionary containing the computed values.
- It runs aggregate queries directly on the database.
- It returns a single dictionary with the result(s), not a QuerySet.

<br>

Example:

```python
from django.db.models import Avg

# using default alias
res = MyProduct.objects.all().aggregate(Avg("price"))
print(res)  # `{'price__avg': 21.0}`

# using custom alias
res = MyProduct.objects.all().aggregate(price=Avg("price"))
print(res)  # `{'price': 21.0}`
```

<br>
<br>

### Aggregate methods

<br>
<br>

| Django Aggregate Method | Description                                               | SQL Equivalent                |
| ----------------------- | --------------------------------------------------------- | ----------------------------- |
| `Count('field')`        | Counts the number of rows (or non-null values in a field) | `COUNT(column)` or `COUNT(*)` |
| `Sum('field')`          | Calculates the sum of values in a column                  | `SUM(column)`                 |
| `Avg('field')`          | Calculates the average value of a column                  | `AVG(column)`                 |
| `Min('field')`          | Finds the minimum value in a column                       | `MIN(column)`                 |
| `Max('field')`          | Finds the maximum value in a column                       | `MAX(column)`                 |


<br>
<br>
<br>
<br>
<br>

## Annotation
- Annotation allows to add **Calculated Fields** to each item in a queryset.
- It makes use of `.annotate()` method.
- These fields are usually aggregations or expressions computed per-row or per-group.
- We can access these fields just like any other model attribute.

<br>

Example:

```python
from django.db.models import Count

# get count of products in each cart
res = MyCart.objects.all().annotate(products_count=Count("products"))
for item in res:
    print(item.id, item.products_count)  # 1 2
```

<br>

`.annotate()` vs `.aggregate()`

| `.annotate()`                       | `.aggregate()`                     |
| ----------------------------------- | ---------------------------------- |
| Adds calculated fields **per row**  | Returns a **single summary** value |
| Returns a QuerySet                  | Returns a dictionary               |
| Used for grouping/adding extra info | Used for overall statistics        |

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

- what is a queryset?
- lazy initialization of queryset
- transactions in django

- Django field types: `OneToOne`, `ForeignKey`, etc.
