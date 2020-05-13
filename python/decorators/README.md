# Decorators in Python

Decorators provide a way to modify the behavior of functions.


## Function fundamentals

### Functions are objects

- function are just a type of object, they can be assigned to vars and passed as args
- nested functions can be defined inside other functions, also called _inner_, _helper_ or _child_ functions
- functions can also be returned

### Function scope

- Python interpreter checks scope in the following order: 
   - _local_ (function) scope
   - _non-local (parent function) scope - nested functions only
   - _global_ scope
   - _builtin_ scope
- By default, only read access is given to variables defined outside of current scope

```python
x = 1
def foo():
   x = 2
   print(x)
foo()
print(x)
# 2
# 1
```

- `global` keyword can give access to a global var - not good practice

```python
x = 1
def foo():
   global x
   x = 2
   print(x)
foo()
print(x)
# 2
# 2
```

- `nonlocal` keyword does the same thing for nested functions

```python
def foo():
   x = 1 
   def bar():     
      nonlocal x
      x = 2
   print(x)
   bar()
   print(x)
 
foo()
# 1
# 2
```

### Closures

- In Python, a closure is a _tuple_ of variables that are no longer in scope, but that a function needs in order to run

```python
def foo():
   x = 5
   def p():
      print(x)
   return p

f = foo()
f()
# 5
```

- In the example above, `foo()` returns `p()` which does not have `x` in its scope
- However, `p()` requires the value of `x` to run
- Python attaches required, non-local to the `p` object as a _tuple_ in the `.__closure__` attribute

```python
print(type(f.__closure__))
# <class 'tuple'>
print(len(f.__closure__))
# 1
print(p.__closure__[0].cell_contents)
# 5
``` 

- Values persists in the closure even when the source var is deleted or overwritten 

## Decorators

- a decorator is a wrapper around a function that can modify inputs, behaviour and outputs
- denoted by `@decorator_name` immediately above the definition of the function being decorated
- a decorator is a function that takes a function as an argument, and returns a modified version of that function

```python
def double_args(func):
   # same signature as func  
   def wrapper(a, b):
      # returns value of func with doubled args  
      return func(a*2, b*2)
   # returns the decorated func 
   return wrapper
      
def multiply(a, b):
   return a * b

multiply(2, 2)
# 4

multiply = double_args(multiply)
multiply(2, 2)
# 16
```

- The above works because of function scope and closure fundamentals discussed above
- Decorator syntax provides a convenient way to do this:

```python
def double_args(func):
   def wrapper(a, b):
      return func(a*2, b*2)
   return wrapper

@double_args
def multiply(a, b):
   return a * b

multiply(2, 2)
```

- So, equivalent syntax for decorating a function

```python
func = decorator(func)
# same as...
@decorator
def func():
   pass
```

## Useful examples

### Timing a function

```python
import time

def timer(func):
    """Time the execution of a function.
    
    Args:
       func (function): The function being decorated
    
    Returns:
       function: The decorated function 
    """
    def wrapper(*args, **kwargs):
        start = time.time()
        # store result of the func (if any)
        result = func(*args, **kwargs)
        end = time.time()    
        print(f"{func.__name__} time: {end-start}")
        return result
    return wrapper    

@timer
def wait_for(seconds):
    time.sleep(seconds)
    print("finished")

wait_for(2)
# finished
# wait_for time: 2.000854969024658
```


### Memoize cache

```python
import time
import json

def memoize(func):
    """Cache function calls to improve performance.
    
    Args:
       func (function): The function being decorated
       
    Returns:
       function: The decorated function
    """
    cache = {} # called once when function is defined?
    def wrapper(*args, **kwargs):
        # create a string key
        key = json.dumps((args, kwargs))
        if (key) not in cache:
            cache[key] = func(*args, **kwargs)
        return cache[key]     
    return wrapper

@memoize
def add(a, b):
    print("add() function sleeping...")
    time.sleep(3)
    return a + b
```

```python
print(add(1,2))
# add() function sleeping...
# 3
print(add(1,2))
# 3
```

- Use decorators when you want to add some common code to multiple functions instead of repeating code (DRY)

## Decorators amd metadata

- One issue with decorators is that they obscure the decorated function's metadata
- eg, `.__doc__` or `.__name__` will show values from the nested func in the decorator
- `functools.wraps` is a decorator for the nested function that fixes this problem

```python
from functools import wraps

def decorator(func):
   # wraps is a decorator that takes an argument
   @wraps(func)
   def wrapper(*args, **kwargs):
      return func(*args, **kwargs)
   return wrapper
```

- `wraps` also provides access to the underlying, undecorated function using `.__wrapped__`
- a bit easier than accessing it via the closure

## Passing arguments to decorators

- Need to create a decorator _factory_
- ie a decorator that _returns_ a decorator rather than _is_ a decorator

```python
def repeat(n):
    def decorator(func):
        def wrapper(*args, **kwargs):
            for i in range(n):
                func(*args, **kwargs)
        return wrapper
    return decorator

@repeat(3)
def show(s):
    print(s)

show("hello")
# 3
# 3
# 3
```












 

