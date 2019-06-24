# Lambdas

A _lambda_ function is an anonymous function (although it can be assigned to a var) that is used to simplify code in some circumstances.

The declaration of a lambda function is in this format: `lambda parameters : expression`

For example (these lambdas are assigned names):

```python
# no params
two = lambda : 2
# one param
sqr = lambda x : x ** 2
# two params
pwr = lambda x, y : x ** y

print(two())    # 2
print(sqr(2))   # 4
print(pwr(2,3)) # 8
```

Lambdas are used to make code more readable - not sure that is always the case, but definately more compact.

Lambdas can also be handy when used with the `map()` and `filter()` functions.

The `map()` functions takes a minimum of two args - the first is a function, and
the second is an _iterable_. It applies the function passed as the first
argument to all of the elements in the second argument, returning an iterator
that delivers all subsequent function results.

Remember, it returns an _iterator_ (one result at a time), not a _list_ (all the
results at once). You can use `list()` to convert the _iterator_ to a _list.
create a  

For example:

```python
def sqr(x):
    return x**2

lst = [1,2,3,4]
for r in map(sqr, lst):
    print(r)

print(list(map(sqr, lst)))
```

Using a lambda function:

```python
lst = [1,2,3,4]
for r in map(lambda x : x ** 2, lst):
    print(r)

print(list(map(lambda x : x ** 2, lst)))
```

So, lambdas are good when a function is only needed right there and then, and
nowhere else. Otherwise, I'd argue they are not really more readable.


