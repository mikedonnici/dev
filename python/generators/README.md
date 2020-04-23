# Generators

A Python **generator** is a piece of specialized code able to produce a series of values, and to control the _iteration_ 
process. Often referred to as _iterators_, although there can be a subtle difference.

The `range()` function is a _generator_, which is an _iterator_.

```python
for i in range(5):
    print(i)
```

A _generator_ is a function that returns a set of values by (generally) being _implicitly_ invoked more than once.

In the above example `range()` is invoked 6 times - 4 to return values (0-4) and a final time to signal the end of the 
process.




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
