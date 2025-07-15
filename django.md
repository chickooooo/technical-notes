## Index
- [Django](#django)
- [Django vs Flask](#django-vs-flask)
- [Project vs App](#project-vs-app)
- [Django Architecture](#django-architecture)
- [Django Features](#django-features)
- [Django Project Structure](#django-project-structure)
- [Common Commands](#common-commands)
- [Admin Interface](#admin-interface)
- [Django Urls](#django-urls)
- [Django Views](#django-views)
- [Django Models](#django-models)
- [ORM .values() vs .values_list()](#orm-values-vs-values_list)
- [Makemigrations vs Migrate](#makemigrations-vs-migrate)
- [Rollback Migration](#rollback-migration)
- [Sessions](#sessions)
- [CSRF Prevention](#csrf-prevention)
- [Middlewares](#middlewares)
- [Request-Response Lifecycle](#request-response-lifecycle)
- [Signals](#signals)
- [Middleware Short-Circuiting](#middleware-short-circuiting)
- [Permissions & Groups](#permissions--groups)
- [Caching](#caching)

<br>
<br>
<br>

### Django

- Django is a fullstack web development framework.
- It is based on Model-View-Template (MTV) architecture.
- It has built-in support for common functionalities like:
    - Admin panel
    - ORM
    - Middlewares	
    - Testing, etc.

<br>
<br>
<br>

### Django vs Flask

- Django is a fullstack web framework with built-in features like ORM, Templating, Admin panel, security, etc.
- Flask is a micro framework suitable for quickly creating APIs.
- Flask depends on third party plugins to provide the functionality Django provides.
- Django is suitable for large scale applications while Flask is preferred quick small applications.

<br>
<br>
<br>

### Project vs App

- A project is our complete web development project
- An app is a part of our project
- One project can have multiple app
- An app within itself is sufficient to perform any task
    - e.g. Authentication, Payment, etc.

<br>
<br>
<br>

### Django Architecture

- Djang is based on Model-View-Template architecture
    - Model: responsible for data management
    - View:
        - sits between models & templates
        - responsible for fetching data from models & rendering respective templates
    - Template: responsible for rendering HTML on the browser

<br>
<br>
<br>

### Django Features

- ORM
- Middlewares
- Templating
- Caching
- Testing
- Regex based URL routing
- Admin interface
- Security

<br>
<br>
<br>

### Django Project Structure

- urls.py: contains url patterns
- views.py: contains views
- settings.py: contains projectwide settings
- models.py: contains database models
- templates/: contains html templates
- admin.py: used for registering models for admin interface
- app.py: contains configuration for that app
- wsgi.py: synchronous web server gateway interface
- asgi.py: asynchronous server gateway interface
- manage.py: responsible for running management commands

<br>
<br>
<br>

### Common Commands

- Create django project: `django-admin startproject <PROJECT_NAME>`
- Create django app: `python manage.py startapp <APP_NAME>`
- Start development server: `python manage.py runserver ?<PORT> ?<HOST:PORT>`
- Create migrations file: `python manage.py makemigrations ?<APP_NAME>`
- Migrate changes to database: `python manage.py migrate`
- Create super user: `python manage.py createsuperuser`
- A django project can have many super users

<br>
<br>
<br>

### Admin Interface

- Admin interface is used to perform administrative tasks
- We can create normal/staff users and manage their roles & permissions
- We can alaso add/delete data from database tables

<br>
<br>
<br>

### Django Urls

- Url paths are used for matching incoming requests
- django also supports regex for URL matching
- when an incoming request is matched with an url, that request is then forwaded to the view attached with that url path.

<br>
<br>
<br>

### Django Views

- Django views are responsible for processing incoming requests.
- It has access to the request data.
- It executes the business logic present in that view and then send a response back to the user.
- The response can be a HttpResponse, JSONResponse or even a FileResponse.

<br>
<br>
<br>

### Django Models

- Django models map an entity to it's database schema.
- Models are also responsible for querying & updating data in the database.
- Often django ORM is used to handle database operations.

<br>
<br>
<br>

### ORM .values() vs .values_list()

- `Model.objects.filter().values(“name”, “age”)`
    - Returns a QuerySet of dictionaries where each dict maps field names to values.
    - Example: [{'name': 'Alice', 'age': 25}, {'name': 'Bob', 'age': 30}]

<br>

- `Model.objects.filter().values_list(“name”, “age”)`
    - Returns a QuerySet of tuples with field values only (no field names).
    - Example: [('Alice', 25), ('Bob', 30)]
    - Use flat=True when querying a single field to get a flat list: ['Alice', 'Bob']

<br>
<br>
<br>

### Makemigrations vs Migrate

- `makemigrations` command creates a migration file.
- This file contains the changes that needs to be made to database.
- E.g. create a table, alter a field, remove a field, set default, etc.
- These changes are detected from changes made to the models.

<br>

- `migrate` command is used to apply those changes to the database.
- It may apply changes from more than 1 migration file.
- It also keeps track of which migrations have already been applied in "django_migrations" table.

<br>
<br>
<br>

### Rollback Migration

- We can also rollback / revert migrations in django.
- If we have our latest migration file as "0003_auto_20250311_1234.py",
- We can rollback to any previous migration using this command:
    - `python manage.py migrate <APP_NAME> 0002`

<br>
<br>
<br>

### Sessions

- Session allows to store and access user data across multiple requests. e.g. login status, preferences, etc.
- They are dictionary like objects generally stored in session database, cache or server itself.
- Whenever a session is created on the server, the session id is sent to the client in the request cookie.
- In the subsequent request, this ID is used to fetch session data.

- In django, sessions are managed by `SessionMiddleware`. It is responsible for sending session ID in the cookie and making the session data available through the request.
- By default sessions are stored in the database in `django_session` table.

<br>
<br>
<br>

### CSRF Prevention

- CSRF stands for Cross site request forgery
- In this type of attack, the attacker tricks the (logged in) user into performing unwanted actions like funds transfer, information leak, etc. without them knowing.
- It uses scripts / links causing the victim’s browser send request to the wep app using victim’s credentials.

- Django prevents CSRF attacks using CSRF token.
- For each form, django adds an unique hidden `csrf_token` to that form. When this form is submitted, the token is also sent with the request and verified by the `CSRFMiddleware`.

<br>
<br>
<br>

### Middlewares

- Middlewares operate between request-response cycle.
- They can execute some code before the request enters the view and after the response is generated by the view.
- If you have multiple middlewares, they are executed from top to bottom.

```
Incoming Request:
User → MID_A → MID_B → View

Outgoing Response:
View → MID_B → MID_A → User
```

<br>

Practical usecases:
- Managing authenticated routes.
- Logging request data & response time for analytics.
- Rate limiting & geo based restriction.

<br>

```python
# app/middlewares.py
class LoggerMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        logger.info("Before the view...")

        response = self.get_response(request)

        logger.info("After the view...")

        return response

# settings.py
MIDDLEWARE = [
    ...
    "app.middlewares.LoggerMiddleware",
]
```

<br>
<br>
<br>

### Request-Response Lifecycle

- The user clicks on a button to fetch products.
- The browser makes a TCP connection with our webserver (NGINX) and makes a GET request to /products endpoint.
- As per our web server configuration, this request is forwarded to a gunicorn worker.
- This worker converts the HTTP request to a python object (HTTPRequest). This HTTPRequest object is forwarded to the Django WSGI handler.
- Inside Django, this request (HTTPRequest) passes through our middlewares in top-bottom order.
- Then Django's URL router matches the request path and forwards the request to the corresponding Django view.
- Inside the Django view, the request is processed and response (HTTPResponse) is generated.
- This response then passes through the middlewares in reverse order.
- Then Django's WSGI handler sends the response back to gunicorn worker.
- The worker then converts the response python object (HTTPResponse) to HTTP response. This HTTP response is then send back to the webserver (NGINX).
- The webserver sends the response for the GET request and the TCP connection is closed.
- The browser receives the response and displays the list of products.

<br>

```
User -> Browser -> Web server (NGINX) -> Gunicorn worker -> Django WSGI handler -> Middlewares (A -> B -> C) -> URL router -> View -> Response
User <- Browser <- Web server (NGINX) <- Gunicorn worker <- Django WSGI handler <- Middlewares (A <- B <- C) <- URL router <- Response
```

<br>
<br>
<br>

### Signals

- Signals in Django follow the "Observer" pattern.
- Whenever an action is performed (new User record created in DB), Django will trigger a function (send welcome email) that is observing this action.
- Signals decouple the action from it's reaction.
- Django signals are synchronous in nature. That means if a signal takes 5 seconds to execute, that request will wait for 5 seconds before sending back response.
- Some built-in signals:
    - pre_save
    - post_save
    - pre_delete
    - post_delete
    - request_started
    - request_finished
    - got_request_exception: when exception is raised

<br>

```python
# app/signals.py
from django.db.models.signals import post_save
from django.dispatch import receiver, Signal
from app.models import Product

@receiver(post_save, sender=Product)
def log_product_added(sender, instance, created, **kwargs):
    # add reaction here
    pass

# app/app.py
class TutorialsConfig(AppConfig):
    name = "tutorials"

    def ready(self):
        import app.signals  # Register signals on app startup
```

<br>

- We can also create custom signals and trigger them manually.

```python
# app/signals.py
my_custom_signal = Signal()

@receiver(my_custom_signal)
def my_custom_receiver(sender, **kwargs):
    # add reaction here
    pass

# app/views.py
my_custom_signal.send(sender=123)
```

<br>
<br>
<br>

### Middleware Short-Circuiting

- It refers to stopping a request inside a middleware and sending a response directly from that middleware.
- In this case, the request does not pass through the rest of the middlewares in the stack and it doesn't even reach the view.
- We can short-ciruit a middleware by returning `HttpResponse` directly from that middleware.
- Examples: auth checks, rate limiting, maintenance, etc.

<br>

```python
from django.http import HttpResponse

class AuthRequiredMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        # Short-circuit: block unauthenticated users
        if not request.user.is_authenticated:
            return HttpResponse("Authentication required", status=401)

        # Otherwise continue to view
        response = self.get_response(request)
        return response
```

Request Flow
- A request comes in.
- `AuthRequiredMiddleware` checks if the user is authenticated.
- If not authenticated, it returns a response immediately.
- If authenticated, the request continues to the next middleware / view.

<br>
<br>
<br>

### Permissions & Groups

Permissions
- For every model, django automatically creates 4 permissions
    - view_<MODELNAME>
    - add_<MODELNAME>
    - change_<MODELNAME>
    - delete_<MODELNAME>
- These permissions grant the ability to view, add, modify & delete instance of that model.
- We can also create custom permissions like below:

```python
class MyModel(models.Model):
    ...
    class Meta:
        permissions = [
            ("can_publish", "Can publish content"),
        ]
```

<br>

Groups
- A group is a way to assign permissions to multiple users at once.
- Multiple permissions can be assigned to a group and multiple users can be part of that group.

```python
from django.contrib.auth.models import Group, Permission

# Create group
editors = Group.objects.create(name='Editors')

# Add permission to group
perm = Permission.objects.get(codename='change_article')
editors.permissions.add(perm)

# Assign user to group
user.groups.add(editors)
```

<br>

- When checking permissions, django checks:
    - Permissions directly assigned to the user.
    - Permissions assigned via users group.

```python
user.has_perm('app_label.codename')       # Boolean
user.get_user_permissions()               # Direct permissions
user.get_group_permissions()              # Permissions via groups
user.get_all_permissions()                # Combined
```

<br>

- **Note:** superuser bypasses permission check.
- Just like we check permissions for model operations, we can also check permissions for view access.

<br>
<br>
<br>

### Caching

Caching views

```python
from django.views.decorators.cache import cache_page

@cache_page(60 * 15)  # Cache for 15 minutes
def my_view(request):
    ...
```

<br>

Manual Caching

```python
from django.core.cache import cache

# Set value
cache.set('my_key', 'value', timeout=300)

# Get value
value = cache.get('my_key')

# Delete cache
cache.delete('my_key')
```

<br>

Common Cache Backends
- LocMemCache: in-memory per-process (default; good for dev)
- Memcached: production-ready, fast
- RedisCache: highly performant, popular for large apps
- FileBasedCache: stores cache in files
- DatabaseCache: stores cache in DB (slower)

<br>

- Make sure to invalidate cache using `cache.delete('my_key')` methods.
- Django also supports cache versioning: `cache.set('my_key', 'value', version=2)`.
- A common practise is to invalidate cache in `post_save` or `post_delete` signals.

<br>
<br>
<br>

### Core

- Rate limiting
- Django Testing
- Django + Celery Integration
- Django types of views & viewsets
- Gunicorn + Nginx
- WSGI/ASGI


### ORM

- prefetch_related
- Complex queries (Q, F, Subquery, OuterRef, annotate, aggregate)
- Transactions (atomic)
- Raw SQL when necessary
