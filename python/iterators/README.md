# Iterables and Iterators

An _iterable_ is an object that can be looped over, such as a list, string or dictionary.

An _iterator_ is an object that can do the _work_ of _iterating_ over the _iterable_.

The _iterator protocol_ defines the way an object should behave to conform to the rules imposed by the context of the 
`for` and `in` statements. An object conforming to the iterator protocol is called an _iterator_.

An _iterator_ must provide two methods:

`__iter__()` - returns the object itself, invoked once to start iteration

`__next__()` - returns the next value or `StopIteration` exception when done


- _iterable_ is an object that has an `iter()` method and can return an _iterator_, eg a list, string, dictionary
- _iterator_ is an object that has a `next()` method, it keeps state and produces the _next_ value when `next()` is invoked

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
# [(0, 'a'), (1, 'b'), (2, 'c')]
```

`zip()` creates an iterator of tuples from two lists:

```python
l1 = ["a", "b", "c"]
l2 = ["A", "B", "C"]
z = zip(l1, l2)
print(list(z))
# [('a', 'A'), ('b', 'B'), ('c', 'C')]
```

Can iterate all at once using the splat operator, `*`. It unpacks all of the values of an iterator:

```python
wrd = "Data"
it = iter(wrd)
print(*it)
# D a t a
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