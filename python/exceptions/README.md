# Exceptions

In python a try-except block can contain contan one or more `except` branches, as well as and `else` branc and a `finally` branch.

For example:

```python
def reciprocal(n):
    try:
        n = 1 / n
    except ZeroDivisionError:
        print("Division failed")
        n = None
    else:
        print("Everything went fine")
    finally:
        print("It's time to say goodbye")
        return n

print(reciprocal(2))
print(reciprocal(0))
```

`else` will run if no `except` condition was met.

`finally` will run everytime.

#### Exceptions are classes

An `Exception` object can be examined through the use of an extended version
of the `except` statement:

```python
try:
    i = int("Hello!")
except Exception as e:
    print(e)
    print(e.__str__())
```

Any `Exception` class can be extended to create customised exceptions.

See: https://edube.org/learn/programming-essentials-in-python-part-2/exceptions-once-again
