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

On Delete
- `on_delete=` parameter specifies what to do if the user is deleted.
- `models.CASCADE` means we should delete all the records referencing to that user.
- There are other options as well. Like:
    - `models.DO_NOTHING` means Django will do nothing and it’s up to the database to handle the deletion.
    - `models.SET_NULL` means set the foreign key to NULL.
    - `models.SET_DEFAULT` means set the foreign key to its default value.
    - `models.SET(...)` means set the foreign key to a given value or callable.
    - `models.PROTECT` means prevent deletion of the referenced object by raising a ProtectedError.

<br>

Many to Many field
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

**Create objects**

```python
user = MyUser.objects.create(name="Alice")
```

<br>

**Get all objects**

```python
all_users = MyUser.objects.all()
```

<br>

**Get a single object**

```python
user = MyUser.objects.get(id=2)
```

- If matching record is not found, raises `MyUser.DoesNotExist` exception.
- If multiple records are found, raises `MyUser.MultipleObjectsReturned` exception.

<br>

**Filter objects**

```python
users = MyUser.objects.filter(name="Alice")  # single filter
products = MyProduct.objects.filter(name="Apple", price=10)  # multiple filters
```

- Returns `QuerySet` of objects.
- If no records match the given condition(s), returns empty `QuerySet`.

<br>

**Exclude objects**

```python
# get products having price > 0
# exclude products with id = 2 from above products
products = MyProduct.objects.filter(price__gt=0).exclude(id=2)
```

<br>

**Ordering objects**

```python
# get all users order by id desc
users = MyUser.objects.all().order_by("-id")
```

<br>

**Limiting objects**

```python
# get first 3 objects
users = MyUser.objects.all()[:3]
```
- Django's ORM translates the slicing into an SQL LIMIT clause automatically.

<br>

**Counting objects**

```python
count = MyProduct.objects.filter(price__gt=10).count()
```

<br>

**Check if object exists**

```python
exists = MyProduct.objects.filter(name="Banana").exists()
```

<br>

**Get first object**

```python
first = MyProduct.objects.filter(name="Banana").first()
```

<br>

**Get last object**

```python
last = MyProduct.objects.filter(name="Banana").last()
```

<br>

**Delete object**

```python
deleted = MyProduct.objects.get(name="Banana").delete()
```

<br>

**Bulk create objects**

```python
products = MyProduct.objects.bulk_create([
    MyProduct(name="Banana", price=20),
    MyProduct(name="Pineapple", price=40),
])
```

<br>

**Update single object**

```python
product = MyProduct.objects.get(id=5)
product.price = 100
product.save()
```

<br>

**Update multiple objects**

```python
products = MyProduct.objects.filter(id__gt=3)
products.update(price=22)
```

<br>

**Q objects: OR query**

```python
q = Q(id__lt=2) | Q(price=22)
products = MyProduct.objects.filter(q)
```

<br>

**Q objects: AND query**

```python
q = Q(id__lte=4) & Q(price=22)
products = MyProduct.objects.filter(q)
```

<br>

**Q objects: complex query**

```python
q = (Q(id__lte=4) & Q(price=22)) | Q(price=30)
products = MyProduct.objects.filter(q)
```

<br>
