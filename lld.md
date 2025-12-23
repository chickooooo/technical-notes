## Index
- [SOLID Principles](#solid-principles)
- [Design Patterns](#design-patterns)
- [Creational Patterns](#creational-patterns)
- [Structural Patterns](#structural-patterns)
- [Behavioural Patterns](#behavioural-patterns)
- [C - Singleton](#c---singleton)
- [C - Factory](#c---factory)
- [C - Abstract Factory](#c---abstract-factory)
- [C - Builder](#c---builder)
- [C - Prototype](#c---prototype)
- [S - Adapter](#s---adapter)
- [S - Decorator](#s---decorator)
- [S - Composite](#s---composite)
- [S - Proxy](#s---proxy)
- [B - Observer](#b---observer)
- [B - Strategy](#b---strategy)
- [B - Command](#b---command)
- [B - Template Method](#b---template-method)
- [B - State](#b---state)
- [B - Chain of Responsibility](#b---chain-of-responsibility)

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
    - Prototype

<br>
<br>
<br>

### Structural Patterns

- These patterns deal with the relation/structure between objects and classes.
- They ensure the structures are flexible and efficient.
- Some common structural design patterns are:
    - Adapter
    - Decorator
    - Composite
    - Proxy

<br>
<br>
<br>

### Behavioural Patterns

- These patterns deal with how objects interact and communicate with each other.
- Some common behavioural design patterns are:
    - Observer
    - Strategy
    - Command
    - Template Method
    - State
    - Chain of Responsibility

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

### C - Abstract Factory

**Overview**
- It defines an interface for creating families of related objects.
- Factory pattern focuses on creating individual objects (e.g. button), while abstract factory focuses on creating a set of related objects (e.g. button & input & popup).
- It is like a factory for factories.

<br>

```python
from abc import ABC, abstractmethod

class Button(ABC):
    """Abstract Product"""

    @abstractmethod
    def click(self):
        raise NotImplementedError()

class MacButton(Button):
    """Concrete Product"""

    def click(self):
        print("Mac button clicked")

class WindowsButton(Button):
    """Concrete Product"""

    def click(self):
        print("Windows button clicked")

class Input(ABC):
    """Abstract Product"""

    @abstractmethod
    def validate(self):
        raise NotImplementedError()

class MacInput(Input):
    """Concrete Product"""

    def validate(self):
        print("Mac input validated")

class WindowsInput(Input):
    """Concrete Product"""

    def validate(self):
        print("Windows input validated")

class AbstractGUIFactory(ABC):
    """Abstract Factory"""

    @abstractmethod
    def create_button(self) -> Button:
        raise NotImplementedError()

    @abstractmethod
    def create_input(self) -> Input:
        raise NotImplementedError()

class MacGUIFactory(AbstractGUIFactory):
    """Concrete Factory"""

    def create_button(self) -> Button:
        return MacButton()

    def create_input(self) -> Input:
        return MacInput()

class WindowsGUIFactory(AbstractGUIFactory):
    """Concrete Factory"""

    def create_button(self) -> Button:
        return WindowsButton()

    def create_input(self) -> Input:
        return WindowsInput()

def process_form(factory: AbstractGUIFactory):
    input = factory.create_input()
    input.validate()
    button = factory.create_button()
    button.click()

factory = WindowsGUIFactory()
# factory = MacGUIFactory()  # easily replacable

process_form(factory)
# Output:
# Mac input validated
# Mac button clicked
```

<br>

**Explanation**
- `Button`, `Input`: These are abstract products that a factory will create.
- `WindowsButton`, `MacButton`, etc.: These are the actual implementations of the products for each platform.
- `AbstractGUIFactory`: This is an abstract factory or factory for factories. It has methods for creating the abstract products.
- `WindowsGUIFactory`, `MacGUIFactory`: These are concrete factories.
- `process_form`: This function is used to render a form using given factory.

<br>

**Advantages**
- Separation of concerns: Each factory handles its own plotform.
- Extensibility: New platform like Linux can easily be added with modifying existing code.
- Interchangeability: One factory can easily be switch with another factory.

<br>
<br>
<br>

### C - Builder

**Overview**
- It is used to construct complex objects step by step.
- We can produce a variety of objects using the same construction process.
- It extracts the object construction code out of its own class and move it to separate objects called **builders**.

**When to use?**
- When object construction has a complex logic.
- When object constructor has many optional parameters.
- When you want to create different representation of the same object.

**Key Components**
- Product: The complex object under construction.
- Builder: The object builder class.
- Director: Constructs the object using builder.

<br>

```python
# Product
class Computer:
    def __init__(self):
        self.cpu = None
        self.ram = None
        self.storage = None
        self.gpu = None

    def __str__(self):
        return f"CPU: {self.cpu}, RAM: {self.ram}, Storage: {self.storage}, GPU: {self.gpu}"

# Builder
class ComputerBuilder:
    def __init__(self):
        self.computer = Computer()

    def add_cpu(self, cpu):
        self.computer.cpu = cpu
        return self

    def add_ram(self, ram):
        self.computer.ram = ram
        return self

    def add_storage(self, storage):
        self.computer.storage = storage
        return self

    def add_gpu(self, gpu):
        self.computer.gpu = gpu
        return self

    def build(self):
        return self.computer

# Director
class Director:
    def __init__(self, builder: ComputerBuilder):
        self.builder = builder

    def build_gaming_pc(self):
        return (
            self.builder.add_cpu("Intel i9")
            .add_ram("32GB")
            .add_storage("1TB SSD")
            .add_gpu("NVIDIA RTX 4090")
            .build()
        )

    def build_office_pc(self):
        return (
            self.builder.add_cpu("Intel i5")
            .add_ram("16GB")
            .add_storage("512GB SSD")
            .build()
        )

# Usage
builder = ComputerBuilder()
director = Director(builder)

gaming_pc = director.build_gaming_pc()
print(gaming_pc)  # CPU: Intel i9, RAM: 32GB, Storage: 1TB SSD, GPU: NVIDIA RTX 4090

office_pc = director.build_office_pc()
print(office_pc)  # CPU: Intel i5, RAM: 16GB, Storage: 512GB SSD, GPU: NVIDIA RTX 4090
```

<br>

**Explanation**
- `Computer` is the product we want to build.
- `ComputerBuilder` breaks down the building process into individual parts.
- `Director` makes use of the builder to build a variety of computers.

<br>
<br>
<br>

### C - Prototype

**Overview**
- It uses a prototype object to initialise new objects.
- Used when the cost of creating new object is expensive.
- Instead of creating a new object using the constructor, we use a copy of an existing object (i.e., a prototype).

**When to use**
- When object creation is costly or time-consuming.
- When creating an object involves a complex configuration.
- When you need a copy of an object at runtime.

<br>

```py
import copy

class Car:
    def __init__(self, brand, model, features):
        self.brand = brand
        self.model = model
        self.features = features  # list of features

    def __str__(self):
        return f"{self.brand} {self.model} with features {self.features}"

    def clone(self):
        # Return a deep copy of the object
        return copy.deepcopy(self)

# Create a prototype car
prototype_car = Car(
    brand="Tesla",
    model="Model S",
    features=["Autopilot", "Electric", "Panoramic Roof"],
)

# Clone the prototype to create a new car
car1 = prototype_car.clone()
car1.model = "Model X"
car1.features.append("ABS")

print(prototype_car)  # Original remains unchanged
print(car1)  # Modified clone

# Output:
# Tesla Model S with features ['Autopilot', 'Electric', 'Panoramic Roof']
# Tesla Model X with features ['Autopilot', 'Electric', 'Panoramic Roof', 'ABS']
```

<br>

**Explanation**
- `Car` is our product.
- `prototype_car` is the prototype object. New objects will be cloned from this object.
- `car1` is the new object created via cloning `prototype_car`.
- After creating `car1`, we can modify it's properties without affecting the prototype.

<br>
<br>
<br>

### S - Adapter

**Overview**
- It allows objects with incompatible interfaces to work together.
- It acts as a bridge between two incompatible interfaces.
- Just like a universal socket adapter acts as a bridge between US electricity plug and an Indian cord.

**When to use**
- You want to reuse an existing class but its interface doesn’t match the one you need.
- You need to integrate a third-party class with your system. (e.g. custom frappe client).

**Components**
- **Target Interface:** The interface your client code expects.
- **Adaptee:** The existing class with an incompatible interface.
- **Adapter:** A class that implements the target interface and translates the calls to the adaptee.

<br>

```py
# Target Interface
class PaymentProcessor:
    def pay(self, amount_in_usd):
        raise NotImplementedError

# Adaptee (Legacy system with incompatible interface)
class LegacyPay:
    def make_payment_in_rupees(self, amount_in_rupees):
        print(f"Paid ₹{amount_in_rupees} using LegacyPay.")

# Concrete implementation that fits the interface
class StripeAPI(PaymentProcessor):
    def pay(self, amount_in_usd):
        print(f"Paid ${amount_in_usd} using Stripe.")

# Adapter to make LegacyPay compatible with PaymentProcessor
class LegacyPayAdapter(PaymentProcessor):
    def __init__(self, legacy_pay):
        self.legacy_pay = legacy_pay

    def pay(self, amount_in_usd):
        # Convert USD to INR (just assume $1 = ₹80 for simplicity)
        amount_in_rupees = amount_in_usd * 80
        self.legacy_pay.make_payment_in_rupees(amount_in_rupees)

# Client code
def complete_purchase(payment_processor: PaymentProcessor, amount):
    payment_processor.pay(amount)

# Usage
print("Using Stripe:")
stripe = StripeAPI()
complete_purchase(stripe, 100)

print("\nUsing LegacyPay via Adapter:")
legacy = LegacyPay()
adapter = LegacyPayAdapter(legacy)
complete_purchase(adapter, 100)
```

<br>

**Explanation**
- `PaymentProcessor` is the target interface. The client code expects this interface.
- `LegacyPay` is the adaptee. It processes the payment in rupees.
- `StripeAPI` is the concrete implementation of the target interface. Which is already compatible with the clients code.
- `LegacyPayAdapter` is the **Adapter**. It makes `LegacyPay` compatible with `PaymentProcessor`.
- `LegacyPayAdapter` is used with the clients code and it makes calls to `LegacyPay` internally.

<br>
<br>
<br>

### S - Decorator

**Overview**
- It allows to add new behaviour to objects, without modifying their existing code.
- It wraps around exisiting code to add more functionality.
- It is used in place of creating all the permutations of subclasses.

**When to use**
- You want to add responsibilities to objects dynamically.
- You want to follow the **Open/Closed Principle**: classes should be open for extension, but closed for modification.
- Subclassing would lead to an explosion of classes to cover all combinations of behaviors.

**Key components**
- Component: The original object interface.
- ConcreteComponent: The actual object being wrapped.
- Decorator: Has a reference to a `Component` and implements the same interface.
- ConcreteDecorator: Extends the behavior of the Component.

<br>

```py
from abc import ABC, abstractmethod

# Component
class Handler(ABC):
    @abstractmethod
    def handle(self, request):
        pass

# Concrete Component
class BasicHandler(Handler):
    def handle(self, request):
        return f"Handling request: {request}"

# Base Decorator
class HandlerDecorator(Handler):
    def __init__(self, handler):
        self._handler = handler

    def handle(self, request):
        return self._handler.handle(request)

# Concrete Decorator 1: Logging
class LoggingDecorator(HandlerDecorator):
    def handle(self, request):
        print(f"[LOG] Received request: {request}")
        response = self._handler.handle(request)
        print(f"[LOG] Response: {response}")
        return response

# Concrete Decorator 2: Authentication
class AuthDecorator(HandlerDecorator):
    def handle(self, request):
        user = request.get("user")
        if not user or not user.get("is_authenticated"):
            return "401 Unauthorized"
        return self._handler.handle(request)

# --- Client Code ---

# A simulated request object
request = {
    "user": {"name": "Alice", "is_authenticated": True},
    "payload": "GET /home"
}

# Compose decorators
handler = BasicHandler()
handler = AuthDecorator(handler)
handler = LoggingDecorator(handler)

# Execute request
print(handler.handle(request))


# Output:
# [LOG] Received request: {'user': {'name': 'Alice', 'is_authenticated': True}, 'payload': 'GET /home'}
# [LOG] Response: Handling request: {'user': {'name': 'Alice', 'is_authenticated': True}, 'payload': 'GET /home'}
# Handling request: {'user': {'name': 'Alice', 'is_authenticated': True}, 'payload': 'GET /home'}

# Unauthenticated output:
# "401 Unauthorized"
```

<br>

**Explanation**
- `Handler` class is the component.
- `BasicHandler` is the concrete component.
- `HandlerDecorator` is the base decorator which implements the component.
- `LoggingDecorator` & `AuthDecorator` are the concrete decorators that implement the base decorator.










<br>
<br>
<br>

### S - Composite

**Overview**
- It treats individual objects and compositions of objects (i.e groups of objects) uniformly.
- It is useful when dealing with tree structures like **file systems**, **GUI elements**, etc.
- Here, both the leaf node (file, button) and composite node (folder, div) should be treated similarly.

**Components**
- Component: Interface for all objects in the composition.
- Leaf: Leaf object in the composition.
- Composite: Composite object. Has a list of children.

<br>

```py
from abc import ABC, abstractmethod

# Component
class FileSystemComponent(ABC):
    @abstractmethod
    def show(self, indent=0):
        pass

# Leaf
class File(FileSystemComponent):
    def __init__(self, name):
        self.name = name

    def show(self, indent=0):
        print(' ' * indent + f"File: {self.name}")

# Composite
class Directory(FileSystemComponent):
    def __init__(self, name):
        self.name = name
        self.children = []

    def add(self, component: FileSystemComponent):
        self.children.append(component)

    def remove(self, component: FileSystemComponent):
        self.children.remove(component)

    def show(self, indent=0):
        print(' ' * indent + f"Directory: {self.name}")
        for child in self.children:
            child.show(indent + 2)

# Usage
if __name__ == "__main__":
    file1 = File("file1.txt")
    file2 = File("file2.txt")
    file3 = File("file3.txt")

    dir1 = Directory("Documents")
    dir1.add(file1)
    dir1.add(file2)

    dir2 = Directory("Downloads")
    dir2.add(file3)
    dir2.add(dir1)  # Nested directory

    dir2.show()
```

<br>

**Explanation**
- `FileSystemComponent` is the common interface.
- `File` is the concrete implementation for leaf node.
- `Directory` is the concrete implementation for composite node.

<br>
<br>
<br>

### S - Proxy

**Overview**
- It provides a **placeholder** or **surrogate** for another object to control access to it.
- A proxy can be used for:
    - Access control
    - Lazy initialization
    - Logging or auditing
    - Caching

**Components**
- Subject: Common interface for the Real Object and Proxy.
- Real Subject: The actual object that performs the real operations.
- Proxy: Controls access to the Real Subject.

<br>

```py
from abc import ABC, abstractmethod
import time

# Subject
class Image(ABC):
    @abstractmethod
    def display(self):
        pass

# Real Subject
class RealImage(Image):
    def __init__(self, filename):
        self.filename = filename
        self.load_from_disk()

    def load_from_disk(self):
        print(f"Loading image from disk: {self.filename}")
        time.sleep(1)  # Simulate delay

    def display(self):
        print(f"Displaying image: {self.filename}")

# Proxy
class ProxyImage(Image):
    def __init__(self, filename):
        self.filename = filename
        self.real_image = None

    def display(self):
        if self.real_image is None:
            print("Creating RealImage object...")
            self.real_image = RealImage(self.filename)
        else:
            print("RealImage already created.")
        self.real_image.display()

# Usage
if __name__ == "__main__":
    print("Creating proxy...")
    image = ProxyImage("photo.jpg")
    
    print("\nFirst display call (loads image):")
    image.display()
    
    print("\nSecond display call (uses cached image):")
    image.display()
```

<br>

**Explanation**
- `Image` is our subject interface.
- `RealImage` is the real subject.
- `ProxyImage` is the proxy that controls the access to `RealImage`.

<br>
<br>
<br>

### B - Observer

**Overview**
- It allows **Subject** to notify a list of **Observers** when an event occures in subject.
- Subject & Observers are coupled loosly.
- Promotes **Open/Close** principle, easy to add new observers without modifying the Subject.

**Components**
- Subject: interface for subject.
- Observer: interface for observer.
- ConcreteSubject: maintains a list of observers and notifies them of any changes.
- ConcreteObserver: object that should be notified.

<br>

```py
from abc import ABC, abstractmethod

# Observer Interface
class Observer(ABC):
    @abstractmethod
    def update(self, message: str):
        pass

# Concrete Observer
class EmailSubscriber(Observer):
    def __init__(self, name):
        self.name = name

    def update(self, message: str):
        print(f"{self.name} received email notification: {message}")

# Subject Interface
class Subject(ABC):
    @abstractmethod
    def attach(self, observer: Observer):
        pass

    @abstractmethod
    def detach(self, observer: Observer):
        pass

    @abstractmethod
    def notify(self, message: str):
        pass

# Concrete Subject
class YouTubeChannel(Subject):
    def __init__(self):
        self.subscribers = []

    def attach(self, observer: Observer):
        self.subscribers.append(observer)

    def detach(self, observer: Observer):
        self.subscribers.remove(observer)

    def notify(self, message: str):
        for subscriber in self.subscribers:
            subscriber.update(message)

    # Business logic that changes state
    def upload_video(self, video_title: str):
        print(f"Channel: Uploaded a new video: {video_title}")
        self.notify(f"New video uploaded: {video_title}")

# === Usage ===
if __name__ == "__main__":
    # Create subject
    channel = YouTubeChannel()

    # Create observers
    alice = EmailSubscriber("Alice")
    bob = EmailSubscriber("Bob")

    # Attach observers to subject
    channel.attach(alice)
    channel.attach(bob)

    # Change state
    channel.upload_video("Observer Pattern in Python")
```

<br>

**Explanation**
- `Observer` is the observer interface and `EmailSubscriber` is the concrete implementation.
- `Subject` is the subject interface and `YouTubeChannel` is the concrete implementation.
- `alice` and `bob` are notified when a new video is uploaded.

<br>
<br>
<br>

### B - Strategy

**Overview**
- It defines a family of algorithms and makes them interchangable at runtime.
- Use when you want to performs a task in different ways (strategies).
- Think of **Google Maps**, user can choose different routes based on different **Strategies**.
    - Shortest distance
    - Fastest time
    - Avoid tolls
- Promotes **Open/Closed** Principle (you can add new strategies without changing existing code).

**Components**
- Strategy: interface for the algorithm.
- Concrete Strategies: different implementations of the Strategy.
- Context: makes use of a Strategy.

<br>

```py
from abc import ABC, abstractmethod

# Strategy Interface
class PaymentStrategy(ABC):
    @abstractmethod
    def pay(self, amount: float):
        pass

# Concrete Strategies
class CreditCardPayment(PaymentStrategy):
    def __init__(self, card_number):
        self.card_number = card_number

    def pay(self, amount: float):
        print(f"Paid ${amount:.2f} using Credit Card ending with {self.card_number[-4:]}.")

class PayPalPayment(PaymentStrategy):
    def __init__(self, email):
        self.email = email

    def pay(self, amount: float):
        print(f"Paid ${amount:.2f} using PayPal account {self.email}.")

# Context
class ShoppingCart:
    def __init__(self):
        self.items = []
        self.payment_strategy: PaymentStrategy = None

    def add_item(self, item, price):
        self.items.append((item, price))

    def set_payment_strategy(self, strategy: PaymentStrategy):
        self.payment_strategy = strategy

    def checkout(self):
        if not self.payment_strategy:
            raise Exception("Payment strategy not set!")

        total = sum(price for _, price in self.items)
        self.payment_strategy.pay(total)

# === Usage ===
if __name__ == "__main__":
    cart = ShoppingCart()
    cart.add_item("Book", 12.99)
    cart.add_item("Pen", 1.99)

    # Use Credit Card
    cart.set_payment_strategy(CreditCardPayment("1234-5678-9876-5432"))
    cart.checkout()

    # Switch to PayPal
    cart.set_payment_strategy(PayPalPayment("user@example.com"))
    cart.checkout()
```

<br>
<br>
<br>

### B - Command

**Overview**
- It allows to encapsulate actions as objects.
- It decoupling client (one who invokes the operation) from the receiver (one who performs the operation).
- Often used in **undo/redo** and **transactional** systems.
- Follows **Open/Close principle**: new commands can be added without modification.

**Components**
- Command: Interface that declares `execute()`.
- Concrete Command: Implements the Command.
- Receiver: Knows how to perform the operation.
- Invoker: Asks the command to carry out a request.
- Client: Creates commands and sets their receivers.

<br>

```py
# Command Interface
class Command:
    def execute(self):
        pass

    def undo(self):
        pass

# Receiver
class Light:
    def turn_on(self):
        print("Light is ON")

    def turn_off(self):
        print("Light is OFF")

# Concrete Commands
class LightOnCommand(Command):
    def __init__(self, light):
        self.light = light

    def execute(self):
        self.light.turn_on()

    def undo(self):
        self.light.turn_off()

class LightOffCommand(Command):
    def __init__(self, light):
        self.light = light

    def execute(self):
        self.light.turn_off()

    def undo(self):
        self.light.turn_on()

# Invoker
class RemoteControl:
    def __init__(self):
        self.command_history = []

    def submit(self, command):
        command.execute()
        self.command_history.append(command)

    def undo_last(self):
        if self.command_history:
            last_command = self.command_history.pop()
            last_command.undo()

# Client Code
if __name__ == "__main__":
    light = Light()
    light_on = LightOnCommand(light)
    light_off = LightOffCommand(light)

    remote = RemoteControl()

    remote.submit(light_on)   # Light is ON
    remote.submit(light_off)  # Light is OFF
    remote.undo_last()        # Light is ON
```

<br>

**When to use**
- When you want to encapsulate operations as objects.
- Perfect for features like: queuing, logging, undo/redo.

<br>
<br>
<br>

### B - Template Method

**Overview**
- It defines skeleton of an algorithm in a method called **template method**.
- It allows subclasses to redefine certain steps of the algorithm **without changing its skeleton**.
- Promotes code reuse and avoid duplication.

**Components**
- Component: defines the template method (non-overridable).
- Concrete Component: implements parts of the algorithm.
- Client Code: calls the template method on concrete classes.

<br>

```py
from abc import ABC, abstractmethod

# Component
class DataProcessor(ABC):
    # Template method
    def process(self):
        self.load_data()
        self.process_data()
        self.save_data()

    @abstractmethod
    def load_data(self):
        pass

    @abstractmethod
    def process_data(self):
        pass

    @abstractmethod
    def save_data(self):
        pass

# Concrete Component 1
class CSVDataProcessor(DataProcessor):
    def load_data(self):
        print("Loading data from CSV file")

    def process_data(self):
        print("Processing CSV data")

    def save_data(self):
        print("Saving processed data to CSV")

# Concrete Component 2
class JSONDataProcessor(DataProcessor):
    def load_data(self):
        print("Loading data from JSON file")

    def process_data(self):
        print("Processing JSON data")

    def save_data(self):
        print("Saving processed data to JSON")

# Usage
csv_processor = CSVDataProcessor()
csv_processor.process()

print("---")

json_processor = JSONDataProcessor()
json_processor.process()
```

<br>

**When to use**
- When you have a generic algorithm where parts of it can vary, but the overall structure should remain the same.

<br>
<br>
<br>

### B - State

**Overview**
- It allows an object to change its behaviour when it's internal state changes.
- It appears as if object has changed its class.

**Components**
- Component: interface for all concrete states.
- Concrete Component: implements Component for various states.
- Context: holds reference to state object and delegates task to it.

<br>

```py
from abc import ABC, abstractmethod

# Component
class CaseState(ABC):
    @abstractmethod
    def write(self, text: str) -> str:
        raise NotImplementedError

# Concrete Component 1
class NormalCaseState(CaseState):
    def write(self, text: str) -> str:
        return text

# Concrete Component 2
class UpperCaseState(CaseState):
    def write(self, text: str) -> str:
        return text.upper()

# Concrete Component 3
class LowerCaseState(CaseState):
    def write(self, text: str) -> str:
        return text.lower()

# Context
class TextEditor:
    def __init__(self):
        self._case_state: CaseState = NormalCaseState()

    def set_state(self, case_state: CaseState):
        self._case_state = case_state

    def add_text(self, text: str):
        return self._case_state.write(text)

# ----- Client code

editor = TextEditor()
print(editor.add_text("Hello World!"))  # Hello World!

editor.set_state(UpperCaseState())
print(editor.add_text("Hello World!"))  # HELLO WORLD!

editor.set_state(LowerCaseState())
print(editor.add_text("Hello World!"))  # hello world!
```

<br>

**When to use**
- When an object must change its behaviour on its internal state change.
- When you want to avoid large `if-else` or `switch` statements.

<br>
<br>
<br>

### B - Chain of Responsibility

**Overview**
- It lets you pass requests along a chain of handlers.
- Each handler decides:
    - Should it process the request.
    - Or pass it to the next handler in the chain.
- A good example is **Django Middlewares**.

**Components**
- Component: defines one method for request handling and one for passing the request.
- Concrete Component: implements the request handling logic.
- Client: creates and links handlers into a chain.

<br>

```py
from abc import ABC, abstractmethod

# --- Handler Interface ---
class SupportHandler(ABC):
    def __init__(self):
        self.next_handler = None

    def set_next(self, handler: 'SupportHandler') -> 'SupportHandler':
        self.next_handler = handler
        return handler  # Allows chaining

    @abstractmethod
    def handle(self, issue: str, level: int):
        pass

# --- Concrete Handlers ---
class LevelOneSupport(SupportHandler):
    def handle(self, issue: str, level: int):
        if level == 1:
            print(f"Level 1 handled issue: {issue}")
        elif self.next_handler:
            self.next_handler.handle(issue, level)
        else:
            print("Issue unhandled.")

class LevelTwoSupport(SupportHandler):
    def handle(self, issue: str, level: int):
        if level == 2:
            print(f"Level 2 handled issue: {issue}")
        elif self.next_handler:
            self.next_handler.handle(issue, level)
        else:
            print("Issue unhandled.")

class ManagerSupport(SupportHandler):
    def handle(self, issue: str, level: int):
        if level >= 3:
            print(f"Manager handled issue: {issue}")
        elif self.next_handler:
            self.next_handler.handle(issue, level)
        else:
            print("Issue unhandled.")

# --- Client Code ---
if __name__ == "__main__":
    # Setup chain: Level1 -> Level2 -> Manager
    level1 = LevelOneSupport()
    level2 = LevelTwoSupport()
    manager = ManagerSupport()

    level1.set_next(level2).set_next(manager)

    # Test different issues
    level1.handle("Forgot password", level=1)
    level1.handle("System crash", level=2)
    level1.handle("Legal escalation", level=3)
    level1.handle("Unknown issue", level=4)
```

<br>

**When to use**
- When you have multiple handlers that can handle a request.
- When you want to decouple sender from receiver.
- When you want the request to be processed by one or many handlers.

<br>
<br>
<br>

### 
