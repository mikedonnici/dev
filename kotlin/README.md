# Kotlin

## Fundamentals

Functions can be defined at the top (package) level, unlike Java.

```kotlin
package foo

fun bar() {
    println("foo.bar()")
}
```

`if` is an expression:

```kotlin
fun main(args: Array<String>) {
    val name = if (args.isNotEmpty()) args[0] else "you"
    println("Hello, $name")
}
```

String templates allow expressions to be used inside string literals. A simple val/var can use `$`, a more complex expression `${}`:

```kotlin
fun main() {
    val iSay = "yum"
    print("I say $iSay, you say ${youSay()}")
}

fun youSay(): String {
    return "pie"
}
```

Variables declares using one of two keywords:

- `val` (value) - _immutable_ (readonly/assign once)
- `var` (variable) - _mutable_

> Use `val` wherever possible - makes code closer to _functional_ style.

Kotlin is statically typed and types can be declared explicity, or inferred by assignment:

```kotlin
val dad: String = "Mike"
val mum = "Christie"
```

> Should only omit type specification if the type is clear from the context. Otherwise, specify the type so the code is easier to read and reason about.

Note that an object that is pointed to by a `val` can still be modifed via the methods on the object. The `val` simply holds an immutable _reference_ to the object.

```kotlin
val name = "Mike"                   // cannot be re-assigned
val people = mutableListOf(name)    // people cannot be re-assigned...
people.add("Christie")              // ...but can modify the list object
```

Note that there are also _read-only_ list types which cannot be modified - see below.

## Functions

### `Unit` return type

A function that has no explicit return type returns `Unit`:

```kotlin
fun foo(): Unit {
    println("side effect")
}
```

Same as:

```kotlin
fun foo() {
    println("side effect")
}
```

### Expression body syntax

If a function returns a single expression, eg:

```kotlin
fun max(a: Int, b: Int): Int {
    return if (a > b) a else b
}
```

...can use _expression body_ syntax:

```kotlin
fun max(a: Int, b: Int) = if (a > b) a else b
```

### Context

Functions can be top-level, members of a class, or nested (local):

```kotlin
fun top() = 1

class A {
    fun member() = 2
}

fun outer() {
    // this is called a local function
    fun local() = 3
}
```

Note that top-level functions can be called from Java as a static function named after the file containing the function. For example, to call `foo()` from file `MyClass.kt` would be `MyClassKt.foo()`.

### Named arguments

Make code easier to reason about:

```kotlin
fun main() {
    println(listOf("a", "b", "c").joinToString(separator="", prefix="(", postfix=")"))
}
// (abc)
```

### Default arguments

Defaults values can be provided for function parameters:

```kotlin
fun main() {
    println(listOf("a", "b", "c").joinToString(postfix="."))
}

// a, b, c.
// default separator is ", ", default prefix is empty
```

To specify default values can use positional or named arguments:

```kotlin
fun main() {
    // can call in various ways:
    printDivider()         // accept defaults
    printDivider('-')     // override first arg (positional)
    printDivider('-', 20)     // override both defaults (positional)
    printDivider(num = 30, char = '+') // override both by name (any order)
}

fun printDivider(char: Char = '*', num: Int = 10) {
    repeat(num) {
        print(char)
    }
    println()
}
```

## Conditionals- `if` and `when`

`if` is an expression in Kotlin so there is no ternary operator.

```kotlin
val max = if (a > b) a else b
```

`when` is similar to a switch:

```kotlin
fun main() {
    val a = 1
    when (a) {
        1 -> println("a is 1")
        2 -> println("a is 2")
        else -> println("a is neither 1 nor 2")
    }
}
```

Can also check multiple values:

```kotlin
fun main() {
    val a = 2
    when (a) {
        1, 2 -> println("a is 1 or 2")
        else -> println("a is neither 1 nor 2")
    }
}
```

The argument to `when` is checked for equality with the branch conditions, so any expression can be used:

```kotlin
fun main() {
    when (setOf(1,2,3)) {
        setOf(1,2,3) -> println("Set contains 1, 2 and 3")
        else -> throw Exception("Set incomplete")
    }
}
```

`when` can be used for checking type:

```kotlin
// Note the type : Any is required for x, otherwise will get
// an Incompatible Type error in the branches
fun main() {
    val x: Any = "abc"
    when (x) {
        is Int -> println("x is an int value $x")
        is String -> print("x is a string with length ${x.length}")
    }
}
```

In the example below, `pet` will be _smart cast_ to the correct type in the matching branch. So if `pet` is a `Dog`, then`pet` is cast to `Dog` and the `pet.woof()` method is available:

```kotlin
open class Pet

class Dog: Pet() {
    fun woof() {
        println("woof, woof")
    }
}

class Cat: Pet() {
    fun meow() {
        println("meeeoooowww")
    }
}

fun main() {
    val pet: Pet = Dog() // again, type must be specified here
    when (pet) {
        is Dog -> pet.woof()
        is Cat -> pet.meow()
    }
}
```

A new variable can be introduced inside the `when` parenthesis (since Kotlin 1.3):

```kotlin
// ...classes omitted...

fun main() {
    when (val pet = myPet()) {
        is Dog -> pet.woof()
        is Cat -> pet.meow()
    }
}

fun myPet(): Pet {
    return Dog()
}
```

This also makes it possible to use `when` in a function expression body. So refactoring the example above:

```kotlin
open class Pet

class Dog: Pet() {
    fun woof() =  "woof, woof"
}

class Cat: Pet() {
    fun meow() = "meeeoooowww"
}

fun main() {
    println(petNoise())
}

fun petNoise(): String =
    when (val pet = myPet()) {
        is Dog -> pet.woof()
        is Cat -> pet.meow()
        else -> throw Exception("Unknow pet")
    }

fun myPet(): Pet {
    return Dog()
}
```

`when` can be used without an argument, using any boolean expression on the left of each branch:

```kotlin
fun main() {
    when {
       1 + 1 == 2 -> println("1 and 1 is 2")
       1 + 1 == 3 -> println("not going to happen")
    }
}
```

## Loops - `for`, `while`, `do-while`

`while` and `do-while` are pretty standard:

```kotlin
while (condition) {
    /*...*/
}

do {
    /*...*/
} while (condition)
```

`for` loop iteration using the `in` keyword:

```kotlin
fun main() {
    val list = listOf("a", "b", "c")
    for (l in list) {
        println(l)
    }
    // generally, the element type is omitted, but can be explicit
    for (l: String in list) {
        println(l)
    }
}
```

Iterate over a list with indexes, storing values in a `Pair` (like a tuple):

```kotlin
fun main() {
    val list = listOf("a", "b", "c")
    for ((i, v) in list.withIndex()) {
        println("list[$i] = $v")
    }
}
```

Iterate over a map:

```kotlin
fun main() {
    val map = mapOf("a" to "A",
                    "b" to "B",
                    "c" to "C")
    for ((k, v) in map) {
        println("$k: $v")
    }
}
```

Iterate over a ranges:

```kotlin
fun main() {

    // includes upper bound, ie 0 to 9 inclusive
    for (i in 0..6) {
        println(i)
    }
    // 0123456

    // excludes upper bound
    for (i in 0 until 6) {
        println(i)
    }
    // 012345

    // more comlplex ranges
    for (i in 9 downTo 1 step 2) {
        print(i)
    }
    // 97531

    // strings - c will be Char type
    for (c in "abc") {
        print(c + 1)
    }
    // bcd

}
```

## Using `in` for checks and ranges

`in` has two use cases in Kotlin - iterating over a range, and checking if a value _occurs_ in a range.

```kotlin
// iteration
for (ch in 'a'..'z') {
    print(ch)
}

// within a range
if ('c' in 'a'..'z') {
    print(true)
}
```

`in` ranges can be used in when conditions:

```kotlin
fun main() {
    printIsAlphaNum('5')
}

fun printIsAlphaNum(ch: Char) {
    when (ch) {
        in 'a'..'z', in 'A'..'Z' -> println("Char is a letter")
        in '0'..'9' -> println("Char is a number")
        else -> println("Char is non-alphanumeric")
    }
}
// Char is a number
```

Ranges can be created from any comparible(?) type and can also be stored in a variable.

Ranges of strings are compared _lexographically_, so:

`"ball" in "a".."k" // true`

`"zoo" in "a".."k" // false`

`in` a collection

```kotlin
if (item in list) { ... }

// same as
if (list.contains(item)) { ... }
```

## Safe Casts

In the case where you want to call a method on a certain _type_ need to first check that a value is of that type.

For example:

```kotlin
fun main() {
    val a: Any = "string"
    if (a is String) {
        val b = a as String         // from abstract to specific type, note `as String` not required
        println(b.toUpperCase())    // call method on type
    } 
}
```

In the above example `toUpperCase()` method is not available on the `Any` type so we need to cast `a` to a `String` type first. This is only possible, of course, if the value of `a` is actually a String.

Kotlin has smart casting, so the example above can be achieved thus:

```kotlin
fun main() {
    val a: Any = "string"
    if (a is String) { // a is smart case to a string here
        println(a.toUpperCase())
    }
}
```

Note that `a` is an abstract type `Any` and an exception is thrown if the cast
is not possible. For example, if `a` was an `Int`:

```kotlin
fun main() {
    val a: Int = 10
    if (a is String) { // Incompatible types: String and Int
        println(a.toUpperCase())
    }
}
```

Instead, _safe cast_ `as?` can be used:

fun main() {
    val a = "string"
    println((a as? String)?.toUpperCase()) // "STRING"

    val b = 10
    println((b as? String)?.toUpperCase()) // null
}

Safe cast returns either the smart-cast value, is the value can be cast, or
`null` if not.

## Exceptions

In Kotlin, `throw` is an expression:

```kotlin
val percentage = 
    if (number in 0..100)
        number
    else 
        throw IllegalArgumentException(
            "A percentage must be between 0 and 100: $number")
```

`try` is also an expression:

```kotlin
fun main() { 
   println(strToNum("abc")) // null
   println(strToNum("123")) // 123
}

fun strToNum(str: String) = try {
        Integer.parseInt(str)
    } catch (e: NumberFormatException) {
        null
    }
```

## Collections

In Kotlin there is a distinction between _read-only_ and _mutable_ collections:

val mutableList = mutableListOf("Go", "Python") mutableList.add("Kotlin")

val readOnlyList = listOf("Go", "Python") mutableList.add("Kotlin") // computer says 'no'

````
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
````

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

An `open` class can be instantiated as an object directly, or have its properties and methods overriden by child classes.

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

Abstract classes are classes that enable inheritence but cannot be directly instantiated.

An `abstract` class must be overriden by a concrete child class and it is implicitly an `open` class.

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

An `interface` is the highest level of abstraction for classes and describes a contract for a set of properties/methods that must be present for a class to satisfy the `interface`.

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

In summary: **overridden** properties/methods are `open` by default, and this can be stopped by using the `final` keyword.

### Inheriting multiple properties/methods with the same name

When a class inherits multiple implementations of a method it _must_ override that method to resolve the ambiguity:

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

Importing the enum class removes the need to preface the constants with the enum class name:

```kotlin
package enumdemo

import enumdemo.Colour.*

enum class Colour {
    RED, GREEN, BLUE
}

class Paint(val clr: Colour) {

    fun brush() {
        when (clr) {
            RED -> println("Painting red")
            GREEN -> println("Painting green")
            BLUE -> println("Painting blue")
        }
    }
}

fun main() {
    val pRed = Paint(RED)
    pRed.brush()
}
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

When a classes or functions needs to accommodate more than one data type, _generic_ parameters can be used.

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

## Extension Functions

Extension functions extend the class. They are defined outside of the class but can be called as a regular member to the class.

The _type_ that the function extends is called a **receiver**.

In the example below, `String` is the receiver and can be referenced in the body of the function using the `this` keyword:

```kotlin
fun String.lastChar() = this.get(this.length - 1)
```

The `this` keyword can be omitted:

```kotlin
fun String.lastChar() = get(length-1)
```

An extension defined outside of a package must be imported:

```kotlin
import com.example.extensions.lastChar
val c: Char = "abc".lastChar()
```

...or:

```kotlin
import com.example.extensions.*
val c: Char = "abc".lastChar()
```

Extensions are an important part of Kotlin and the Kotlin Standard Library is the **Java Standard Library + extentions**.

Under the hood, extensions are _static_ functions. In this example `obj` is of type `Parent` so the _static_ function `Parent.foo()` is called rather than `Child.foo()`:

```kotlin
open class Parent
class Child: Parent()

fun Parent.foo() = "parent"
fun Child.foo() = "child"

fun main() { 
  val obj: Parent = Child() // type is Parent, so...
  println(obj.foo())         // ...static fun Parent.foo() called
}
// parent
```

### Extension vs Member

A member function always wins out over an extension with the same name:

```kotlin
fun main() {
    println("abc".get(1))
}

fun String.get(index: Int) = "*"
// b
```

Compiler will issue a warning: `Extension is shadowed by a member:...`

However, extensions can _overload_ members if the signature is different:

```kotlin
fun main() {
    println("abc".get(1)) // calls member
    println("abc".get(1, true)) // calls extension override
}

fun String.get(index: Int, uppercase: Boolean): Char {
    if (uppercase) {
        return get(index).toUpperCase()
    }
    return get(index)
}
// b
// B
```

## Functional Programming

### Lambda syntax

Lambdas are anonymous functions that can be used as an expression.

In Kotlin, lambdas are written in curly braces:

```kotlin
{ x: Int, y: Int -> x + y }
```

Arguments on the left: `x: Int, y: Int`, function body on the right: `x + y`.

A lambda can be passed to a function:

```kotlin
fun main() {
    val list = listOf(1,2,3,4,5,6,7,8,9)
	val newList = list.filter({ i: Int -> i > 5 })
    print(newList)
}
// [6, 7, 8, 9]
```

If the lambda is the last argument, it can be placed outside the function
parenthesis:

```kotlin
fun main() {
    val list = listOf(1,2,3,4,5,6,7,8,9)
	val newList = list.filter() { i: Int -> i > 5 }
    print(newList)
}
// [6, 7, 8, 9]
```

If it is the only argument, the function parenthesis can be omitted (an idea
borrowed rom Ruby):

```kotlin
fun main() {
    val list = listOf(1,2,3,4,5,6,7,8,9)
	val newList = list.filter { i: Int -> i > 5 }
    print(newList)
}
// [6, 7, 8, 9]
```

Type can also be ommitted if it is clear from the context:

```kotlin
fun main() {
    val list = listOf(1,2,3,4,5,6,7,8,9)
	val newList = list.filter { i -> i > 5 }
    print(newList)
}
// [6, 7, 8, 9]
```

If there is only one argument you can use the automatically created `it`:

```kotlin
fun main() {
    val list = listOf(1,2,3,4,5,6,7,8,9)
	val newList = list.filter { it > 5 }
    print(newList)
}
// [6, 7, 8, 9]
```

For a *multi-line* lambda the last expression is the result:

```kotlin
list.any {
    println("Processing $it")
    it > 0 // result expression
}
```

Lambda arguments such as Map entries or Pairs can be _destructured_ as well.

So instead of this:

```kotlin
fun main() {
    val map = mapOf("one" to 1, "two" to 2)
	val map2 = map.mapValues { entry -> "${entry.key} -> ${entry.value}" }
}
```

...can do:

```kotlin
fun main() {
    val map = mapOf("one" to 1, "two" to 2)
	val map2 = map.mapValues { (key, value) "$key -> $value" }
    println(map2)
}
```

Unused parameters can be discarded with an underscore:

```kotlin
fun main() {
    val map = mapOf("one" to 1, "two" to 2)
	val map2 = map.mapValues { (_, value) "$value" }
    println(map2)
}
```

### Collection extensions

The Kotlin libray contains many extension functions for working with
collections.

Many of these work are _functional_ and work with lambdas.

#### `filter`

Filters out the content of a list and keeps only the elements that satisfy the
given predicate:

```kotlin
fun main() {
    println(listOf(1,2,3,4,5,6,7,8,9).filter { it > 5 })
}
// [6, 7, 8, 9]
```

#### `map`

Transforms each element in a collection:

```kotlin
fun main() {
    println(listOf(1,2,3,4,5,6,7,8,9).map { it * it })
}
```

#### `any`, `all`, `none`

Return a boolean based on the predicate being satisfied by the appropriate
number of elements:

```kotlin
fun main() {
    val list = listOf(1,2,3,4,5,6,7,8,9)
    println(list.any { it > 10 }) // false
    println(list.all { it < 10})  // true
    println(list.none { it < 4 }) // false
}
```

#### `find`, `first`, `firstOrNull`

Returns first element that satisfies the predicate.

If no element satisfies, `find` and `firstOrNull` return `null`, `first` will throw an exception.

```kotlin
fun main() {
    val list = listOf(1,2,3,4,5,6,7,8,9)
    
    println(list.find { it > 5 })  // 6
    println(list.find { it > 10 }) // null
    
    
    println(list.first { it < 10}) 		// 1
    try {
    	println(list.first { it > 10})  // Exception    
    } 
    catch(e: Exception) {
        println(e.message) 
        // Collection contains no element matching the predicate. 
    }
    
    println(list.firstOrNull { it > 10 }) // null
}

```

#### `count`

Counts the number of elements that satisfy the given predicate:

```kotlin
fun main() {
    val list = listOf(1,2,3,4,5,6,7,8,9)
    println(list.count { it > 7 })
}
// 2
```

#### `partition`

Splits the original list into a pair of lists, where the first list contains
elements for which the predicate yielded true and the second list, false.

```kotlin
fun main() {
    val list = listOf(1,2,3,4,5,6,7,8,9)
    val lists = list.partition { it % 2 == 0 }
    println(lists.first)
    println(lists.second)
}
// [2, 4, 6, 8]
// [1, 3, 5, 7, 9]
```

#### `groupBy`

Partitions a collection into multiple sub-collections. In this example, the
collection is split into a map that is keyed by `genus`.

```kotlin
class Carnivore(val genus: String, val species: String)

fun main() {
    val plants = listOf(
    	Carnivore("Drosera", "spathulata"),
    	Carnivore("Drosera", "pygmaea"),
        Carnivore("Sarracenia", "flava"),
        Carnivore("Sarracenia", "leucophylla"),
    	Carnivore("Drosera", "binata"),
    	Carnivore("Drosera", "capensis"),
        Carnivore("Dionaea", "muscipula"),
        Carnivore("Drosphyllum", "lusitanicum")
    )
    val genera = plants.groupBy { it.genus }
    genera["Drosera"]?.forEach {
        println("${it.genus} ${it.species}")
    }
}
// Drosera spathulata
// Drosera binata
// Drosera pygmaea
// Drosera capensis
```
