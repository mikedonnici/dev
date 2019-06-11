# Object-Oriented Programming

## Class Variables vs Instance Variables

In the example below, `cVar` is a _class_ variable and `iVar` is an _instance_ variable.

Class variables exists, and can be altered, in any object (instance) of the class. If this value is altered in _any_ instance then its value is altered in all instances.

Instance variables exist only within the object (instance) itself. Changes to its value only exist within the instance that the change was made.

```python
class Foo():

    ## Class var
    cVar = 0

    def __init__(self, val):
        # Instance var
        self.iVar = val
```

**Note**: Python differes from some OO languages in that instance vars are
created in the object itself. Hence it is possible for two objects of the same
class to have different instance vars. Coversely, you cannot assume that a
particular instance var exists in all objects of a class.

In the example below, the constructor will create either `a` or `b` depending on `val`:

```python
class Foo():
    def __init__(self, val):
        if val % 2 == 0:
            self.a = 1
        else:
            self.b = 1
```

The `hasattr` function can be used to check if an object/class contains a
specified property.

First arg is class/object name, second is string name of the property:

```python
cVarExists = hasattr(Foo, 'cVar')
iVarExists = hasattr(objName, 'iVar')
```

## Encapsulation

Vars (attributes) and methods can be hidden by prefixing their names with 2 or more underscores.

In the example below trying to access `Foo.__cVar` or `objName.__iVar` with throw an `AttributeError`.

```python
class Foo():

    __cVar = "hidden"

    def __init__(self, val)
        self.__iVar = val
```

Methods can be hidden in the same way:

```python
class Foo():

    def __hiddenMethod(self):
        pass
```

## Introspection and Reflection

**Introspection** - the ability of a program to examine the type or properties
of an object at runtime.

**Reflection** - the ability of a program to manipulate the values, properties
and / or functions of an object at runtime.

Each python class and objects has a set of attributes that can be used to
examine capability.

#### `__dict__`

`objName.__dict__` contains a `dict` of _instance_ vars with keys named in
the _mangled_ form: `_ClassName__varName`.

`ClassName.__dict__` will show all of the class attributes and methods.

Both instance and class attributes are directly accessible using these mangled names,
so encapsulation in Python is weak:

```python
f = Foo()
# hidden not hidden :)
print(f._Foo__iVar) # instance var
print(f._Foo__cVar) # class var
```

#### `__name__`

`ClassName.__name__` will return the string name of the class. **This attribute
only exists in classes**.

To find the class that an object belongs to use `type()` and `__name__`:

```python
class Foo:
    pass
obj = Foo()
print(type(obj).__name__)
```

#### `__module__`

Returns a string that is the name of the module that contains
the class - it works on both classes and objects:

```python
class Foo:
    pass
print(Foo.__module__)
obj = Foo()
print(obj.__module__)
```

#### `__bases__`

A tuple that contains classes (not class names) that are direct
_superclasses_ of the class in question. **This attribute
only exists in classes**.

## Inheritance

#### Class and subclass example:

```python
class Stack:
    """Implements a simple stack"""

    def __init__(self):
        """Initialises the stack with an empty list"""
        self.__items = []

    def push(self, item):
        """Pushes an item onto the stack"""
        self.__items.append(item)

    def pop(self):
        """Pops an item off the stack"""
        r = self.__items[-1]
        del self.__items[-1]
        return r


class AddingStack(Stack):

    def __init__(self):
        """Sub class constructor needs to explicitly
        call __init__() of super class"""
        Stack.__init__(self)
        self.__sum = 0

    def push(self, val):
        """Sub class push() overrides super class function
        (polymorphism) but needs to call push() on the super
        class to access __items. Note: self must be passed explicitly."""
        Stack.push(self, val)
        self.__sum += val

    def getSum(self):
        return self.__sum


s1 = AddingStack()
s1.push(1)
s1.push(2)
s1.push(3)
s1.push(4)
print(s1.getSum())
```

#### `issubclass(ClassOne, ClassTwo)`

Returns `True` if `ClassOne` is a subclass of `ClassTwo`.

**Note:** A class is considered subclass of itself.

#### `isinstance(objName, ClassName)`

Returns `True` if `objName` is an instance of `ClassName` or any of its superclasses.

#### `is` operator

Returns `True` if two vars point to the same object:

```python
objOne is objTwo
```

#### `super()`

The super and subclass example above shows how the subclass can call the superclass constructor with `SuperClassName.__init__(self, someVar)`. Using this method `self` must be passed explicitly.

The same thing can be achieved using the `super()` function, however there is no need to pass `self`:

```python
class Foo():
    def _init_(self, someVar):
        self.someVar = someVar

class SubFoo(Foo):
    def __init__(self, someVar):
        super().__init__(someVar)
```

Using `super()` does not require knowledge of the superclass name and provides access to all of the resources in the superclass.

When accessing entities from an object, python will:

1. Look at the object first, then
1. Look at the class hierarchy, from bottom to top
1. Raise an `<AttributeError>` exception

#### Multiple Inheritance

A class can inherit from more than one superclass:

```python
class Foo:
    pass

class Bar:
    pass

class FooBar(Foo, Bar):
    pass
```

#### Overriding

When a method or attribute is defined in a subclass with the same name as that in a superclass, the subclass property will _override_ that of the superclass.

In the case of _multiple inheritance_ the properties of the superclasses are scanned from _left_ to _right_ in the subclass declaration and the first match is the one that will be accessed.

For example:

```python
class Left:
    var = "L"
    varLeft = "LL"
    def fun(self):
        return "Left"

class Right:
    var = "R"
    varRight = "RR"
    def fun(self):
        return "Right"

class Sub(Left, Right):
    pass

obj = Sub()
print(obj.var, obj.varLeft, obj.varRight, obj.fun())
```
