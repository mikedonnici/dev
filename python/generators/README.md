# Generators

A Python **generator** is a piece of specialized code able to produce a series of values, and to control the _iteration_ process. Often referred to as _iterators_, although there can be a subtle difference.

The `range()` function is a _generator_, which is an _iterator_.

```python
for i in range(5):
    print(i)
```

A _generator_ is a function that returns a set of values by (generally) being _implicitly_ invoked more than once.

In the above example `range()` is invoked 6 times - 4 to return values (0-4) and a final time to signal the end of the process.

## Iterator Protocol

The _iterator protocol_ defines the way an object should behave to conform to the rules imposed by the context of the `for` and `in` statements. An object conforming to the iterator protocol is called an _iterator_.

An _iterator_ must provide two methods:

`__iter__()` - returns the object itself, invoked once to start iteration

`__next__()` - returns the next value or `StopIteration` exception when done

#### Iterators and Iterables

_(this bit seems a tad clunky)_

- an _iterable_ is an object that can return an _iterator_
- an _iterator_ is an object that keeps state and produces the next value when you call `next()` on it

A list is _iterable_:

```python
l = ['a', 'b', 'c']
for i in l :
    print(i)
```

To create an _iterator_ from a list, use `iter()`:

```python
l = ['a', 'b', 'c']
il = iter(l)
print(next(il))
print(next(il))
print(next(il))
```

A range object is _iterable_:

```python
for n in range(5) :
    print(n)
```

...and can be used to create an _iterator_:

```python
itera = iter(range(3))
print(next(itera))
print(next(itera))
print(next(itera))
```

`enumerate()` returns an `enumerate` object that produces a sequence of tuples, each of which an index-value pair.

`list()` creates a list from the `enumerate` object.

```python
eo = enumerate(['a', 'b', 'c'])
print(list(eo))
[(0, 'a'), (1, 'b'), (2, 'c')]
```

## `yield` statement

Implementing the iterator protocol (interface) requires maintaining state
between each of the `__inter__()` invocations. This can be cumbersome.

The `yield` statement is a much easier way to create an iterator.

```python
def f(n):
    for i in range(n):
        yield i
```

In the above example `yield` is used instead of `return`. Obviously a `return` would provide only the first value after which the state of `i` would be lost.

The `yield` statement effectively turns the function into a generator. When invoked it returns the objects identifier and can be used like this:

```python
for i in f(5):
    print(i)
```

#### List comprehensions as generators

List comprehension, returns a list. That is, all the values are returned at once.

For example:

```python
lst = [x for x in range(5)]
for i in lst:
    print(i)
```

Similar list-comprehension syntax can be used to create a generator, with parenthesis in place of square brackets:

```python
genr = (x for x in range(5))
for i in genr:
    print(i)
```
