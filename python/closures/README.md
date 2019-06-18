# Closures

Closure is a technique which allows the storing of values in spite of the fact that the context in which they have been created does not exist anymore.

Example:

```python
def outer(par):
	loc = par
	def inner():
		return loc
	return inner

var = 1
fun = outer(var)
print(fun())
```

In is example, `outer()` returns a _copy_ of `inner()` with all of the local variables intact - a kind of snapshot frozen in time, in fact, a **closure**.

When `fun()` is called, `outer()` no longer exists but the closure contains a copy of the environment and hence the value of `loc` is available.

In the above example the closure has no params, so is \_invoked\_ with no args. However, a closure can be created with params as well:

```python
def makeclosure(par):
	loc = par
	def power(p):
		return p ** loc
	return power

fsqr = makeclosure(2)
fcub = makeclosure(3)
for i in range(5):
	print(i, fsqr(i), fcub(i))
```
