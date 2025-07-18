## Index
- [SOLID Principles](#solid-principles)
- [Design Patterns](#design-patterns)
- [Creational Patterns](#creational-patterns)
- [Structural Patterns](#structural-patterns)
- [Behavioural Patterns](#behavioural-patterns)
- [C - Singleton](#c---singleton)
- [C - Factory](#c---factory)

<br>
<br>
<br>

### SOLID Principles

Single Responsibility Principle (SRP)
- A class should have only one job or responsibility.
- It should have only one reason to change.
- This keeps the code modular, easy to maintain & test.

<br>

Open/Closed Principle (OCP)
- Modules, classes & functions should be open for extension but closed for modification.
- New behaviour should be added by extending existing code, not changing it.
- This keeps the code stable and flexible.

<br>

Liskov Substitution Principle (LSP)
- Objects of parent class should be replaceable with objects of child class.
- Child classes must follow the behaviour (methods) expected from the parent class.
- This prevents unexpected bugs due to broken inheritance.

<br>

Interface Segragation Principle (ISP)
- A child class should not be forced to implement the methods it does not need.
- Prefer many small, specific interfaces over a large general-purpose one.
- This reduces the impact of changes and improves code clarity.

<br>

Dependency Inversion Principle (DIP)
- Higher level classes (business logic) should not depend on lower level classes (UI).
- Lower level classes should depend on higher level classes through abstract classes.
- This promotes decoupling & code reusability.

**Bad**
```python
class MySQLDatabase:
    def connect(self): ...

class UserService:
    def __init__(self):
        self.db = MySQLDatabase()
```
UserService depends on concrete MySQLDatabase implementation.

**Good**
```python
class DatabaseAbstract:
    def connect(self): ...

class MySQLDatabase(DatabaseAbstract):
    def connect(self): ...

class UserService:
    def __init__(self, db: DatabaseAbstract):
        self.db = db
```
UserService depends on a DatabaseAbstract, not a specific DB.



<br>
<br>
<br>

### Design Patterns

- Design patterns are proven solutions to common software development problems.
- They are not code snippets, but rather reusable templates.
- They provide a standard way of solving common problems.
- These design patterns are divided into 3 broad categories.
    - Creational patterns
    - Structural patterns
    - Behavioural patterns

<br>
<br>
<br>

### Creational Patterns

- These patterns deal with object creation mechanism.
- Some common creational design patterns are:
    - Singleton
    - Factory
    - Abstract Factory
    - Builder

<br>
<br>
<br>

### Structural Patterns

- These patterns deal with the relation/structure between objects and classes.
- They ensure the structures are flexible and efficient.
- Some common structural design patterns are:
    - a

<br>
<br>
<br>

### Behavioural Patterns

- These patterns deal with how objects interact and communicate with each other.
- Some common behavioural design patterns are:
    - a

<br>
<br>
<br>

### C - Singleton

**Overview**
- It ensures that a class has only 1 instance.
- Whenever we need to interact with the class, we use this same instance.
- It is useful when we need a single shared resource.
- Example: configuration manager, logger, database connection, etc.

<br>

```python
class MyClass:
    _instance = None
    
    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
        return cls._instance

a = MyClass()
b = MyClass()

print(a is b)  # True
```

<br>

**Explanation**
- `_instance` class variable holds the class instance.
- Whenever a new object of class is to be created, `__new__()` method is called. Inside this method, we check if class instance is already created.
- If the instance is not created, we create a new class instance, assign it to `_instance` variable and return it.
- If the instance is already created, we simply return that instance.
- `super()` is used to call the `__new__()` method of the parent class.

<br>
<br>
<br>

### C - Factory

**Overview**
- It defines an interface for creating objects.
- The type of object that will be created depends on the input.
- It decouples object creation from its usage.
- It maked code more flexible and extensible.
- Example
    - **ButtonFactory**: used to create **PrimaryButton**, **LinkButton**, **DisabledButton**, etc.
    - **DocumentFactory**: used to create **PDF**, **Word**, **Excel**, etc.
    - **DatabaseConnectionFactory**: **MySQL**, **PostgreSQL**, **SQLite** connection, etc

<br>

```python
from abc import ABC, abstractmethod

class Document(ABC):
    @abstractmethod
    def generate(self):
        pass

class PDFDocument(Document):
    def generate(self):
        print("generating pdf")

class ExcelDocument(Document):
    def generate(self):
        print("generating excel")

class DocumentFactory:
    @staticmethod
    def create_document(doc_type: str) -> Document:
        if doc_type == "pdf":
            return PDFDocument()
        elif doc_type == "excel":
            return ExcelDocument()
        else:
            raise ValueError("Unsupported document type")

doc_1 = DocumentFactory.create_document("pdf")
doc_1.generate() # Output: generating pdf

doc_2 = DocumentFactory.create_document("excel")
doc_2.generate() # Output: generating excel
```

<br>

**Explanation**
- We have an abstract class `Document` with an abstract method `generate()`.
- The `PDFDocument` and `ExcelDocument` classes implement this method to generate specific document types.
- The `DocumentFactory` class has a static method `create_document()` that returns the appropriate document instance (`PDFDocument` or `ExcelDocument`) based on the input string.
- This allows for flexible object creation without exposing the details of how they are created.

<br>
<br>
<br>

### 
