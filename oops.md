## Index
- [4 Pillars of OOPS](#4-pillars-of-oops)
- [Access modifiers](#access-modifiers)
- [Getters & Setters](#getters--setters)
- [Method overriding vs Method overloading](#method-overriding-vs-method-overloading)
- [Class method vs Static method](#class-method-vs-static-method)
- [Abstract class vs Interface](#abstract-class-vs-interface)

<br>
<br>
<br>

### 4 Pillars of OOPS

- Abstraction
- Encapsulation
- Inheritance
- Polymorphism

Abstraction
- Hiding complex implementation details.
- Providing a cleaner API to interact with the object.
- It allows to define the structure of the class.
- It focuses on 'what object does' instead of 'how it does it'.

Encapsulation
- It bundles data (attributes) & methods that operate on this data into a single class.
- It also restricts access to object's internal state using access modifiers.
- Access modifiers control how data is accessed and modified.

Inheritance
- Inheritance allows child class to access data (attributes) & methods of parent class.
- This promotes code reuse and establish child-parent hierarchy.
- It enables child class to extend the functionality of the parent class.

Polymorphism
- Polymorphism means Many Forms.
- It allows objects of different classes to be treated as object of common superclass.
- It enables a method to behave in different ways, depending on the object that invoke it.

<br>
<br>
<br>

### Access modifiers

- Access modifiers are used to control the access of class attributes & methods.
- There are 3 types of access modifiers: public, private, protected.
- Public data members can be accessed within as well as outside the class.
- Private data members can be accessed only within the current class.
- Protected data members can be accessed within the current class & all the subclasses that inherit this class.

---

- In python, private data members starts with "__" (double underscore).
- Protected data members start with "_" (single underscore).
- Note: Python does not enforce strict access control like Java or C++.

<br>
<br>
<br>

### Getters & Setters

- A getter is a special method used to get value of private & protected attributes.
- A setter is used to set the value of private & protected attributes.
- In python, we use `@property` decorator to implement getters & setters.

```python
class Tree:
    _age = 10

    @property
    def age(self):
        return self._age

    @age.setter
    def age(self, value: int):
        self._age = value
```

<br>
<br>
<br>

### Method overriding vs Method overloading

- In method overriding, the method in the child class replaces the method in the parent class.
- The method signature (parameters & return type) remains the same.
- When we call the method using child's object, only the new method is accessible.

---

- Method overloading refers to defining multiple methods with the same name but different signatures.
- The method signature is changed in this case (different parameters or return type).
- When we call the method using class object, both new and old methods can be called separately (depending on the arguments passed).

---

- Python only supports method overriding.
- To implement functionality like method overloading, we can use `*args` & `**kwargs`.

<br>
<br>
<br>

### Class method vs Static method

- A class method is bound to the class rather than the instance.
- It can access class level attributes & methods.
- It is defined using @classmethod decorator.

---

- A static method is neither bound to the class nor the instance.
- It does not have acess to any attributes or methods.
- They are generally stateless and are used as utility functions.
- It is defined using @staticmethod decorator.

```python
class MyClass:
    _value = 10

    @classmethod
    def get_value(cls):
        return cls._value


class MathUtilities:
    @staticmethod
    def add(x, y):
        return x + y


MyClass.get_value()  # Outputs: 10
MathUtilities.add(5, 3)  # Outputs: 8
```

<br>
<br>
<br>

### Abstract class vs Interface

- Abstract class can have both concrete as well as abstract methods.
- Can have a constructor.
- Common behaviour (methods) can be implemented in abstract class, leaving specific implementation to child class.

---

- Interface can have only abstract methods.
- Cannot have a constructor.
- No implementation can be done in an interface. All methods must be implemented in the child class.

---

- Both provides a blueprint for the child class to implement.
- Python only supports abstract class using abc module.

```python
from abc import ABC, abstractmethod

class BaseClass(ABC):
    @abstractmethod
    def print_hello(self):
        pass

class ChildClass(BaseClass):
    def print_hello(self):
        print("hello")

child = ChildClass()
child.print_hello()  # Output: hello
```
