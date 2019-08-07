# kotlin notes

## Collections

### Maps

Maps can be created thus:

```kotlin
fun main() {

    val nameToAge = mapOf(Pair("Mike", 48), Pair("Christie", 44))
    val nameToAge2 = mapOf(
        "Mike" to 48,
        "Christie" to 44
    )

    println(nameToAge == nameToAge2)
    println(nameToAge.keys)
    println(nameToAge.values)
    println(nameToAge.entries)

    val countryToPopulation = mutableMapOf(
        "Germany" to 80_000_000,
        "USA" to 300_000_000
    )
    countryToPopulation.put("Australia", 23_000_000)
    countryToPopulation.putIfAbsent("USA", 320_000_000)

    println(countryToPopulation)
    println(countryToPopulation.contains("Australia"))
    println(countryToPopulation.containsKey("France")) // same as above
    println(countryToPopulation.containsValue(23_000_000))
    println(countryToPopulation.get("Germany"))
    println(countryToPopulation.getOrDefault("France", 0))

    nameToAge.entries.forEach {
        val (name, age) = it
        println("$name is $age years old")
    }

}
```

Note that `to` is an _infix_ function, ie `"Mike" to 48` is the same as `"Mike".to(48)`.


## Named loops

This works for `break` and `continue`.

```kotlin
fun main() {
    outer@ for (i in 1..10) {
        for (j in 1..10) {
            if (j > 4) {
                break@outer
            }
            println("$i - $j")
        }
    }
}
```

## Object-oriented programming

### Class basics

Simplest class:

```kotlin
class Person
```

Class with no constructor:

```kotlin
class Person {
    var name: String = ""
    var age: Int = 0

    fun speak() {
        println("$name says hi!")
    }
}
```

Class with constructor (java style):

```kotlin
class Person(name: String) {

    // boiler plate to initialise object of this class
    val name: String

    init {
        this.name = name
    }

    fun speak() {
        println("$name says hi!")
    }
}
```

Class with named parameters:

```kotlin
class Person(val name: String) {

    // No boilerplate here!

    fun speak() {
        println("$name says hi!")
    }
}
```

Class with named params and default values:

```kotlin
class Person(val name: String = "John Doe") {

    fun speak() {
        println("$name says hi!")
    }
}
```

It is important to note that in Kotlin declaring a class _property_ differs from declaring a field in Java, and some other languages.

When a _property_ is created a _backing_ field is also created but access to the value is provided via automatically generated _getter_ and _setter_ methods.

A `var` will have a _getter_ and a _setter_, a `val` will only have a _getter_.

```kotlin
class Animal {
    var age: Int = 0
}

fun main() {
    val a1 = Animal()
    a1.age = 3          // calls setter
    println(a1.age)     // calls getter
}

// 3
```

Accessing `a1.age` looks like direct field access but internally it is using the _getter_ or _setter_ method.

The `get()` and `set()` methods can be overriden on a class:

```kotlin
class Animal {
    var age: Int = 0
        get() = field + 1
        set(value) {
            if (value > 0) {
                field = value
            }
        }
}

fun main() {
    val a1 = Animal()
    a1.age = -3
    println(a1.age)
}

// 1
```

### Open classes

Open classes are the first level of class that allows inheritence.

An `open` class can be instantiated as an object directly, or have its
properties and methods overriden by child classes.

```kotlin
open class Person(open val name: String) {

    open fun speak() {
        println("$name says hi!")
    }
}

class Doctor(override val name: String) : Person(name) {

    override fun speak() {
        println("Hi, my name is Dr $name")
    }
}

fun main() {

    val pr = Person(name = "Barry")
    pr.speak()
    val dr = Doctor(name = "Mike")
    dr.speak()
}

```

### Abstract classes

Abstract classes are classes that enable inheritence but cannot be directly
instantiated.

An `abstract` class must be overriden by a concrete child class and it is
implicitly an `open` class.

Methods in an `abstract` class may also be `open` or `abstract`,

```kotlin
abstract class Person(open val name: String) {

    open fun speakOpenly() {
        println("$name says hi!")
    }

    // This *must* be overriden in child classes
    abstract fun speakAbstractly()
}

class Patient(override val name: String) : Person(name) {

    override fun speakAbstractly() {
        println("I'm $name, with legs and arms")
    }
}

class Doctor(override val name: String) : Person(name) {

    override fun speakOpenly() {
        println("Hi, my name is Dr $name")
    }

    override fun speakAbstractly() {
        println("$name...medicine...helper")
    }
}

fun main() {

    // val pr = Person("Barry") // * compilation error *
    val pt = Patient("Sicko")
    pt.speakOpenly()
    pt.speakAbstractly()
    val dr = Doctor("Mike")
    dr.speakOpenly()
    dr.speakAbstractly()
}
```

### Interfaces

An `interface` is the highest level of abstraction for classes and describes a
contract for a set of properties/methods that must be present for a class to
satisfy the `interface`.

Properties and methods defined in an interface are implicitly `abstract`.

```kotlin
interface Staff {

    val employeeId: String

    fun available(): Boolean
}

abstract class Person(open val name: String) {

    fun identify() {
        println("My name is $name")
    }
}

class Nurse(override val name: String, override val employeeId: String) : Person(name), Staff {

    override fun available(): Boolean {
        return true
    }
}

fun main() {

    val n1 = Nurse("Cratchet", "N-abc123")
    n1.identify()
    println(n1.available())
}
```

Interface values as function parameters:

```kotlin
interface Staff {
    val employeeId: String

    fun fullName(): String
}

abstract class Person(open val name: String)

class Nurse(override val name: String, override val employeeId: String) : Person(name), Staff {

    override fun fullName() = "Nurse $name"
}

class Doctor(override val name: String, override val employeeId: String) : Person(name), Staff {

    override fun fullName() = "Dr $name"
}

fun identify(staff: Staff) {
    println("My name is ${staff.fullName()}")
}

fun main() {

    val staff1 = Nurse("Cratchet", "N-abc123")
    val staff2 = Doctor("Feelgood", "D-gdft")
    identify(staff1)
    identify(staff2)
}
```

### Preventing `override` using `final`

Successive subclasses can override an inherited method:

```kotlin
abstract class Teller {
    abstract fun tellStory()
}

open class Parent : Teller() {
    override fun tellStory() {
        println("Parent tells story")
    }
}

open class Child : Parent() {
    override fun tellStory() {
        println("Child tells story")
    }
}

class Toy : Child() {
    override fun tellStory() {
        println("Toy tells story")
    }
}

fun main() {
    val storyTeller = Toy()
    storyTeller.tellStory()
}

// Toy tells story
```

Use the `final` keyword to specify a limit to subclass override:

```kotlin
abstract class Teller {
    abstract fun tellStory()
}

open class Parent : Teller() {
    override fun tellStory() {
        println("Parent tells story")
    }
}

open class Child : Parent() {
    final override fun tellStory() { // * final *
        println("Child tells story")
    }
}

class Toy : Child() {
    override fun tellStory() { // * error *
        println("Toy tells story")
    }
}

fun main() {
    val storyTeller = Toy()
    storyTeller.tellStory()
}

// Error: 'tellStory' in 'Child' is final and cannot be overridden
```

Note that `final` is implicit when declaring a property or method _unless_ it is `open` or `abstract`.

However, when using `override` the `final` keyword must be used explicitely for the desired effect. Otherwise, the property/method can be _overridden_ in child classes.

In summary: _*overridden*_ properties/methods are `open` by default, and this can be stopped by using the `final` keyword.

### Inheriting multiple properties/methods with the same name

When a class inherits multiple implementations of a method it _must_
override that method to resolve the ambiguity:

```kotlin
interface Teller {
    // In this example the interface has an implementation for tellStory()
    fun tellStory() {
        println("Anyone can tell a story!")
    }
}

abstract class StoryTeller {
    open fun tellStory() {
        println("StoryTeller tells the story")
    }
}

open class Parent : StoryTeller(), Teller {
    // No tellStory() here!
}

fun main() {
    val st = Parent()
    st.tellStory()
}

// Error:(14, 6) Kotlin: Class 'Parent' must override public open fun tellStory():
// Unit defined in StoryTeller because it inherits many implementations of it
```

The inheriting class _must_ impement the `tellStory()` method to resolve the ambiguity. However, one of the parent implemetations can be specified:

```kotlin
interface Teller {
    fun tellStory() {
        println("Anyone can tell a story!")
    }
}

abstract class StoryTeller {
    open fun tellStory() {
        println("StoryTeller tells the story")
    }
}

open class Parent : StoryTeller(), Teller {
    // Resolves the ambiguity
    override fun tellStory() {         // must be implemented
        // super.tellStory()           // still ambiguous
        // super<Teller>.tellStory()   // works
        super<StoryTeller>.tellStory() // works
    }
}

fun main() {
    val st = Parent()
    st.tellStory()
}

// StoryTeller tells the story
```

### Data classes

A `data class` provides convenience methods for when a classes primary purpose is as a data structure. They are not restricted to properties only and can be given additional methods. However, they cannot be _abstract_ and are therefore more useful for concrete classes that contain a lot of data.

```kotlin
data class Thing(val a: String, val b: String, val c: String)

fun main() {
        val t1 = Thing("AAA", "BBB", "CCC")
        println(t1)
}

// Thing(a=AAA, b=BBB, c=CCC)
```

By default, a `data class` has a nicer `toString()` method than a standard class.

Data classes are easier to test for equality:

```kotlin
class NormalThing(val a: String, val b: String, val c: String)

data class DataThing(val a: String, val b: String, val c: String)

fun main() {
    val nt1 = NormalThing("AAA", "BBB", "CCC")
    val nt2 = NormalThing("AAA", "BBB", "CCC")
    println(nt1.equals(nt2)) // false

    val dt1 = DataThing("AAA", "BBB", "CCC")
    val dt2 = DataThing("AAA", "BBB", "CCC")
    println(dt1.equals(dt2)) // true
}
```

Data classes can also be copied easily:

```kotlin
data class Thing(val a: String, val b: String, val c: String)

fun main() {
    val t1 = Thing("AAA", "BBB", "CCC")
    val t2 = t1.copy()
    println(t1)
    println(t2)
    println(t1.equals(t2)) // true
}
```

Properties on a copy are easy to modify:

```kotlin
data class Thing(val a: String, val b: String, val c: String)

fun main() {
    val t1 = Thing("AAA", "BBB", "CCC")
    val t2 = t1.copy(a = "A*A")
    println(t2)
}

// Thing(a=A*A, b=BBB, c=CCC)
```

Data classes can be easily decomposed into values:

```kotlin
data class Thing(val a: String, val b: String, val c: String)

fun main() {
    val t1 = Thing("AAA", "BBB", "CCC")
    val (a, b, c) = t1
    println(a)
    println(b)
    println(c)
}

// AAA
// BBB
// CCC
```

Data classes are also useful for creating _hash sets_. These sets cannot contain duplicate values so using `hashSetOf()` with instances of data classes will ensure there are no duplicates. In the example below, the `hashSetOf()` excludes the duplicate `data class Fauna` object but not the duplicate for `class Flora`.

```kotlin
data class Fauna(val a: String, val b: String, val c: String)

class Flora(val a: String, val b: String, val c: String)

fun main() {

    val a1 = Fauna("Cat", "Dog", "Bird")
    val a2 = Fauna("Cat", "Dog", "Bird") // same as t1
    val a3 = Fauna("Quoll", "Snake", "Wombat")
    val set1 = hashSetOf(a1, a2, a3)

    val p1 = Flora("Rose", "Jasmine", "Bamboo")
    val p2 = Flora("Rose", "Jasmine", "Bamboo") // same as p1
    val p3 = Flora("Dionaea", "Drosera", "Nepenthes")
    val set2 = hashSetOf(p1, p2, p3)

    println(set1)
    println(set2)
}

// [Fauna(a=Cat, b=Dog, c=Bird), Fauna(a=Quoll, b=Snake, c=Wombat)]
// [Flora@eed1f14, Flora@65ab7765, Flora@1b28cdfa]
```

### Objects / Singletons

A singleton is a design pattern where an object of a particular class is only instatiated once.

Some argue it is an anti-pattern as it effectively creates global data and also makes testing difficult.

Either way, in Kotlin this is done using an _object declaration_:

```kotlin
object DatabaseConnection {

    val dsn = "mysql@localhost"

    fun execute(sql: String) {
        println("Execute query: $sql")
    }

}

fun main() {

    val db = DatabaseConnection
    db.execute("SELECT * FROM row WHERE 1")
}
```

### Enum classes

Enum classes are used for storing a set of _possible_ values and help to ensure type safety:

```kotlin
enum class Colour {
    RED, GREEN, BLUE
}

class Paint(val clr: Colour) {

    fun brush() {
        when (clr) {
            Colour.RED -> println("Painting red")
            Colour.GREEN -> println("Painting green")
            Colour.BLUE -> println("Painting blue")
        }
    }
}

fun main() {
    val pRed = Paint(Colour.RED)
    pRed.brush()
    println(Colour.RED) // prints meaningful value
}

// Painting red
// RED
```

### Information hiding

It is best practice to control access to the properties and methods of a class by implementing the highest level of information hiding possible while still providing the required functionality via a well-defined public interface.

The highest level of infomration hiding is `private`. A property declared as `private` is only visible within the class:

```kotlin
class Animal {
    private var age: Int = 0
}

fun main() {
    val a1 = Animal()
    a1.age = 1
    println(a1.age)
}

// Error: Cannot access 'age': it is private in 'Animal'
```

The next level of information hiding is `protected`. A `protected` property is accessible within the class and sub classes:

```kotlin
open class Animal {

    protected var name: String = "Animal"
}

class Dog : Animal() {

    fun call() {
        println("Hey $name")
    }
}

fun main() {

    val f = Dog()
    f.call()
}

// "Hey Animal"
```

The next level of modifier is `internal` which limits visibility to the same _module_ - ie set of related packages.

```kotlin
class Animal {
    internal genus: String = ""
}
```

The final (and default) visibility modifier is `public` which allows access to a property or method from anywhere. Generally `public` methods are used to create a well-defined public interface for a class:

```kotlin
class Animal {

    public fun identify() {
        // ...
    }
}
```

Note that the _getters_ and _setters_ will have the same visibility as the properties.

## Packages

Applications are organised into packages and packages can be nested by creating subfolders in a top-level package directory.

Kotlin can also use Java packages, and vice-versa.

```kotlin
import java.util.Date

fun main() {
    println(today())
}

fun today(): Date {
    return Date()
}
```

Can also used _qualified package name_:

```kotlin
fun main() {
    println(today())
}

fun today(): java.util.Date {
    return java.util.Date()
}
```

Importing packages, or parts thereof, can be done ina few ways:

Import all classes from a package:

```kotlin
import java.util.*
import com.mikedonnici.somepkg.*
```

Import a specific class:

```kotlin
import java.util.Date
import com.mikedonnici.somepkg.Thingo
```

Import a specific top-level function:

```kotlin
import com.mikedonnici.somepkg.Thingo.foo

fun main() {
    foo()
}
```

Import a method from an _object_, ie a _singleton_, not a class:

```kotlin
// foo/Foo.kt
package foo

object FooFactory {
    fun bar() {
        println("Here's the bar!")
    }
}
```

...can be imported, thus:

```kotlin
import foo.FooFactory.bar

fun main() {
    bar()
}

// Here's the bar!
```

Can also import enum constants:

```kotlin
import pkg.Colours.BLUE

fun main() {
    val blue = BLUE
}
```

Packages that are publically available are conventionally named with a reverse domain name prefix:

```kotlin
import com.mikedonnici.proj.pkg.subpkg
```

## Generics

When a classes or functions needs to accomodate more than one data type, _generic_ parameters can be used.

Generic parameters are specifed with angled brackets `<>` and a parameter name that is often a single, upper-case letter, eg. `<E>`, `<T>`.

Generics with classes:

```kotlin
class Stack<E>(vararg val items: E) {

    var elements = items.toMutableList()

    fun push(item: E) {
        elements.add(item)
    }

    fun pop(): E? {
        if (elements.isEmpty()) {
            return null
        }
        return elements.removeAt(elements.size - 1)
    }
}

class Person(val name: String)

fun main() {

    val p1 = Person("Mike")
    val p2 = Person("Chris")
    val p3 = Person("Maia")
    val p4 = Person("Leo")

    val st = Stack(p1)
    st.push(p2)
    st.push(p3)
    st.push(p4)

    // ? is 'safe' call as can return null
    // see: https://kotlinlang.org/docs/reference/null-safety.html#safe-calls
    println(st.pop()?.name)
    println(st.pop()?.name)
    println(st.pop()?.name)
    println(st.pop()?.name)
    println(st.pop()?.name)
}

```

Note: **vararg** params are like _variadic_(?) params in Go - a comma-seperated list of args of a particular type.

Generics with functions:

```kotlin
class Stack<E>(vararg val items: E) {

    var elements = items.toMutableList()

    fun push(item: E) {
        elements.add(item)
    }

    fun pop(): E? {
        if (elements.isEmpty()) {
            return null
        }
        return elements.removeAt(elements.size - 1)
    }
}

fun <T> stackOf(vararg items: T): Stack<T> {

    // * is spread operator which 'spreads' the array of items
    // back out into a vararg for the Stack() constructor
    return Stack(*items)
}

fun main() {

    val s1 = stackOf("a", "b", "c")
    for (i in 0..2) {
        println(s1.pop())
    }
}

```
