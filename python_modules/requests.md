# Requests Module

<br>
<br>
<br>
<br>
<br>

## Index

<br>
<br>
<br>
<br>
<br>

## Request

### GET request

- Use `.get()` method to make a `GET` HTTP request.

```py
import requests

BASE_URL = "http://localhost:8080"

response = requests.get(BASE_URL + "/hello-json")

print(response.status_code)  # 200
print(response.json())  # {'message': 'Hello, world'}
```

<br>
<br>

### Query parameters

- Use `params` parameter to pass query parameters to the API.

```py
params = {"name": "john"}
response = requests.get(BASE_URL + "/hello-json", params=params)

print(response.status_code)  # 200
print(response.json())  # {'message': 'Hello, john'}
```

<br>
<br>

### Path parameters

- Path parameters should be added manually to the API endpoint.

```py
name = "john"
response = requests.get(BASE_URL + f"/hello/{name}")

print(response.status_code)  # 200
print(response.json())  # {'message': 'Hello, john'}
```

<br>
<br>

### POST request

- Use `.post()` method to make a `POST` HTTP request.

```py
response = requests.post(BASE_URL + "/orders")

print(response.status_code)  # 200
print(response.json())  # {'message': 'success'}
```

<br>
<br>

### POST request body

#### Plain text body

- Use `data` parameter to pass request body as plain text.

```py
body = "12345"
response = requests.post(BASE_URL + "/orders", data=body)

print(response.status_code)  # 200
print(response.json())  # {'message': '12345'}
```

<br>

#### JSON body

- Use `json` parameter to pass request body as a `JSON` object.
- It automatically adds `Content-Type: application/json` header and encodes the body.

```py
body = {"id": 101}
response = requests.post(BASE_URL + "/orders", json=body)

print(response.status_code)  # 200
print(response.json())  # {'message': {'id': 101}}
```

<br>

#### Form data

- Use `data` parameter to pass **Form data** in request body.
- It automatically adds `Content-Type: application/x-www-form-urlencoded` header and encodes the dict as form data.

```py
body = {
    "username": "admin",
    "password": "admin",
}
response = requests.post(BASE_URL + "/login", data=body)

print(response.status_code)  # 200
print(response.json())  # {'username': 'admin', 'password': 'admin'}
```

<br>
<br>

### 

<br>
<br>

### 

<br>
<br>
<br>
<br>
<br>

## Request Advanced 

<br>
<br>
<br>
<br>
<br>

## Response

<br>
<br>
<br>
<br>
<br>
