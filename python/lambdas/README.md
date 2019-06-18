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
