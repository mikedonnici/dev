# kotlin notes

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
}

// Painting red
```
