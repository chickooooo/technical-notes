# Datetime Module

<br>
<br>
<br>
<br>
<br>

## Index

- [`datetime.time`](#datetimetime)
- [`datetime.timedelta`](#datetimetimedelta)

<br>
<br>
<br>
<br>
<br>

## `datetime.time`

- Represents time of day: hour, minute, second, microsecond
- Has support for timezone.

<br>

### Creation

```py
datetime.time(hour=0, minute=0, second=0, microsecond=0, tzinfo=None)
```

<br>

### Comparison

- Can compare two time objects (`<`, `>`, `==`).
- Both times must have the same `tzinfo` (or both `None`).

<br>

### Arithmetic

- Cannot add two `time` objects.
- Cannot add `time` + `timedelta`.
- **Reason**: `time` has no concept of a day boundary.

<br>

### Current time

- No `now()` method for `time`.
- To get current time use `datetime.datetime.now().time()`.

<br>

### Key points

- `time` is not suitable for arithmetic.
- Use `datetime.datetime` or `datetime.timedelta` for calculations.

<br>

```py
from datetime import time, timedelta


t1 = time()
t2 = time(hour=2, minute=1)
td1 = timedelta(minutes=3)

print(t1)  # 00:00:00
print(t2)  # 02:01:00
print(t1 > t2)  # False
```

<br>
<br>
<br>
<br>
<br>

## `datetime.timedelta`

- Represents a **duration** or **time difference**.

<br>

### Creation

```py
datetime.timedelta(
    weeks: float = 0,
    days: float = 0,
    hours: float = 0,
    minutes: float = 0,
    seconds: float = 0,
    microseconds: float = 0,
    milliseconds: float = 0,
)
```

<br>

### Comparison

- Can compare two `timedelta` objects.
- Comparisons are based on **total duration**.

<br>

### Arithmetic

- Can add or subtract `timedelta` with:
    - Another `datetime.timedelta`
    - `datetime.datetime`
    - `datetime.date`
- Cannot be added to `time`.
- Supports negative durations.

<br>

### Normalization

- Automatically normalizes values.
- Example: 90 seconds is converted to 1 minute 30 seconds.

<br>

### Useful methods

- `total_seconds()`: returns duration in total seconds as a `float`.

<br>

### Key points

- No months or years attribute.
- Always use `total_seconds()` for precise calculations.

<br>

```py
from datetime import timedelta


td1 = timedelta()
td2 = timedelta(weeks=2, days=1.5, hours=5)

print(td1)  # 0:00:00
print(td2)  # 15 days, 17:00:00
```

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
