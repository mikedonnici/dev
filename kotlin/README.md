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