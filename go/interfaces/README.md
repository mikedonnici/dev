# Interfaces in Go

An interface is a type that defines a set of method signatures and are used to
achieve a form of _polymorphism_ in Go.

Interfaces do not _implement_ their methods but are mapped to _concrete types_
that do. If a concrete type implements methods with the same signatures as an
interface's methods, then that concrete type is said to _satisfy the interface_.

A concrete type specifies data as well as methods and may contain additional
methods unrelated to a particular interface. A concrete type may also satfisfy
more that one interface.

In this example, the concrete type `Square` satisfies the `Shape` interface
because it implements the `Area()` and `Perimeter()` functions specified in
the interface definition.

```go
type Shape interface {
    Area()     float64
    Permimeter() float64
}

type Square struct {
    side float64
}

func (s Square) Area() float64 {
    return s.side * s.side
}

func (s Square) Perimeter() float64 {
    return s.side * 4
}
```

## Interface Values

Once an interface has been defined, values of that interface type can be
created and passed around, just like any other type.

```go
var s Shape // <nil>
```

The interface value has two components: a **Dynamic Type** and a
**Dynamic Value**. These are assigned based on the underlying concrete
type associated with the interface value.

```go
var sh Shape
sq := Square{side : 4}
fmt.Println(sq.Area()) // works, of course
sh = sq                // dynamic type and value assignment
fmt.Println(sh.Area()) // also works
```

<https://play.golang.org/p/r9rBxxpVxkz>

In the above example the interface value `sh` is dynamically assigned
the _type_ `Square` and the _value_ `sq`. Hence the `Area()` method can be called
on `sh` - and interface value of type `Shape`.

## Interface Type Assertions

Interfaces are useful for concealing concrete type difference,for example using
`Shape` inplace of concrete types `Square`, `Triangle` etc.

However, there are occasions when the underlying concrete type of an interface
needs to be exposed or _disambiguated_.

This can be done with type assertions using the _comma, ok_ idiom often used for
maps.

```go
func DrawShape(s Shape) {
    sq, ok := s.(Square)
    if ok {
        DrawSquare(sq)
    }
    tri, ok := s.(Triangle)
    if ok {
        DrawTriangle(tri)
    }
}
```

Could also us a `type switch`:

```go
func DrawShape(s Shape) {
    switch sh := s.(type) {
        case Square:
            DrawSquare(sh)
        case Triangle:
            DrawTriangle(sh)
    }
}
```

In this example `sh` is assigned the concrete value of `s`.
