# Python Notes

(From DataCamp course)

### Iterators vs Iterables

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

...and can be used to create an _interator_:

```python
itera = iter(range(3))
print(next(itera))
print(next(itera))
print(next(itera))
```
