# Context Managers

- A context manager is a function that is set up in a certain way
- Allows a block of code to be run within a _context_
- Cleans up once the context block is complete
- `with` statement is a _compound_ statement containing code to be run in the context
- ie, the `with` statement signals entry into a _context_
- Context manager can take arguments like any normal function

```python
# start context
with <context-manager>(<args>):
   # statement 1
   # statement 2
# end context   
```

- Context manager can return a var that can be used inside the context

```python
with <context-manager>(<args>) as <var-name>:
   # code inside context
```

- the `open` function is a context manager 

```python
with open('file.csv') as fp:
   length = len(fp.read())
print(f"File is {length} chars long")
```

- Two ways to create a context manager:
   - Function-based
   - Class-based
   
## Creating a context manager function

1. Define the function
2. Add set up code (optional)
3. Use the `yield` keyword
4. Add teardown code (optional)
5. Add the `@contextlib.contextmanager` decorator

```python
@contextlib.contextmanager
def my_context():
  print("hello")
  yield 42 
  print("goodbye")


with my_context() as num:
    print(f"num is {num}")
# hello
# 42
# goodbye
```

- `yield` signifies a return value but allows the rest of the the function to be run subsequently
- the `yield` value is assigned to `var` using `with my_context() as num`
- `yield` is used in generators, and a context manager is _technically_ a generator that yields a single value

- Context managers often contain setup and teardown code, for example:

```python
@contextlib.contextmanager
def database(url):
    # set up
    db = postgres.connect(url)
    # return connection
    yield db
    # teardown
    db.disconnect() 

with database(DB_URL) as db:
    rows = db.execute("SELECT * from table")
``` 

- a context manager can yield `None` to return control back to caller, and then run code after the context is finished:

```python
@contextlib.contextmanager
def in_dir(path):
    old_dir = os.getcwd()  # save initial dir
    os.chdir(path)         # change dir
    yield                  # yield None, return control
    os.chdir(old_dir)      # return to original dir
```

## Nested contexts

- Nested with statements are legal
- For example, copying a file too large for memory could uses a nested context:

```python
with open('file1.csv') as in_file:
   # open dest file inside source file context
   with open('file2.csv') as out_file:
      # copy the file one line at a time
      for line in in_file:
         out_file.write(line)
```

## Handling errors

- when n exception is raise may not be able to run rest of context function code 
- use `try` -> [`except`] -> `finally`

```python
import contextlib

@contextlib.contextmanager
def foo(key):
   data = {'a': 'A', 'b': 'B'}    
   try:
      print(data[key])
      yield
   except KeyError:
      print(f"No data for key = {key}")
      yield  
   finally:
      print("Running important code after yield OR error")

with foo('a'):
    pass
# A
# Running important code after an error

with foo('c'):
    pass
# No data for key = c
# Running important code after yield OR error
```

## Context manager patterns

- Code that follows these patterns are candidates for implementing context managers:
   - OPEN - CLOSE
   - LOCK - RELEASE
   - CHANGE - RESET
   - ENTER - EXIT
   - START - STOP
   - SETUP - TEARDOWN
   - CONNECT - DISCONNECT
   



 







    
