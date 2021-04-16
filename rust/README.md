# Rust

- [The Rust Book](https://doc.rust-lang.org/stable/book)
- [Rust By Example](https://doc.rust-lang.org/stable/rust-by-example)  
- [Cargo Guide](https://doc.rust-lang.org/cargo/guide/)
- [Crate Register](https://crates.io/)

## Install and Update

```script
# Installation
$ $ curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh

# rustup toolchain manager
$ rustup --version

# rust compiler
$ rustc --version

# update
$ rustup update

# uninstall
$ rustup self uninstall
```

## Local Docs

```shell
$ rustup doc
```

## Build & Run

With the compiler:

```script
$ rustc main.rs
$ ./main
```

With cargo:

```script
# create a new package
$ cargo new pkgname
$ cd pkgname

# build
$ cargo build
# run
$ ./target/debug/pkgname

# build and run
$ cargo run

# compilation check only (no build)
$ cargo check
```

## Language Fundamentals

### Variables - Mutability, Constants, Shadowing

- Variables are immutable by default
- Use `mut` to set a var to _mutable_  

```rust
let x = 5;      // immutable
let mut y = 10; // mutable
```

- Constants are always immutable
- Convention is UPPERCASE_UNDERSCORES
- Can only set with a _constant expression_, not func call or computed value 

```rust
const MAX_NUM: u32 = 9;
```

- The `let` keyword allows for a var to be _shadowed_
- Effectively creating a new var with same name so differs from `mut`

```rust
let x = 1;
let x = 2; // shadows x above 
```

### Data Types

- Rust is _statically typed_ so types must be known at compile time

#### Scalar Types

- Represent a _single_ value
- Rust has 4 scalar types:
   - integers: `23`
   - floating-point numbers: `2.45`
   - booleans: `true`, `false`
   - characters: `'a'`, `'*'`
   
#### Compound Types

##### Tuples

- Tuples: grouping of typed values into a single type
- Fixed length once declared
- Each position has a type:

```rust
fn main() {
    // Type annotation is optional and can be inferred
    let tup: (i32, f64, u8) = (500, 6.4, 1);
}
```

- values can be accessed by _destructuring_:

```rust
fn main() {
    let tup = (500, 6.4, 1);
    let (x, y, z) = tup;
    println!("{}, {}, {}", x, y, z)
}
```

- ...or by dot notation, ie `var.idx`

```rust
fn main() {
    let tup = (500, 6.4, 1);
    println!("{}, {}, {}", tup.0, tup.1, tup.2)
}
```

##### Arrays

- Elements must be same type
- Like tuples, arrays have a fixed length
- Allocated on the stack, not the heap

```rust
fn main() {
  // type specification [type; count]:
  let a1: [u32; 5] = [1, 2, 3, 4, 5];
  
  // Without...
  let a2 = [1, 2, 3, 4, 5];
  
  // initialising with same value for each element
  let a3 = [3; 5]; // same as [3, 3, 3, 3, 3]
}
```

- An array is a single chuck of memory on the stack, so can be accessed with 
  indexing:
  
### Functions

- Snake case is convention for func and var names
- `{}` block is an expression so return is implicit:

```rust
fn main() {
  let first_num = 2;
  let second_num = 4;   
  let s = sum_numbers(2, 4);
  println!("Sum of {} and {} is {}", first_num, second_num, s);
}

fn sum_numbers(a: u32, b: u32) -> u32 {
  // return of {} block implicit if no ; after expression a + b
  // If the ; was there get a compilation error
  a + b
  // Could also do 
  // return a + b;
}
```

### Control Flow

#### `if`, `else if` and `else`

- The usual syntax

```rust
fn main() {
  let n = 4;
  if n > 5 {
    println!("n is greater than 5");
  } else if n == 4 {
    println!("n is equal to 4");
  } else {
    println!("n is less than 4");
  }
}
```

- `if` is an expression so can use like this:

```rust
fn main() {
  let condition = true;
  let number = if condition { 5 } else { 6 };
  println!("The value of number is: {}", number);
}
```

#### loop, while and for

- `loop`:

```rust
fn main() {
  loop {
    println!("this goes forever...or until control c")
  }
}
```

- can return an expression from `loop` at the `break`

```rust
fn main() {
  let mut count = 0;
  let result = loop {
    count += 1;
    println!("Count is now {}", count);
    if count >= 4 {
      println!("break and return value from loop");
      break count + 100;
    }
  };
  println!("result from loop is {}", result);
}
```

- `while`:

```rust
fn main() {
    let mut number = 3;

    while number != 0 {
        println!("{}!", number);

        number -= 1;
    }
    println!("LIFTOFF!!!");
}
```

- `for` loops are safer as can't run out of range:

```rust
fn main() {
  let a = [1, 2, 3];
  for n in a.iter() {
    println!("{}", n);
  }
}
```

- can loop over a range `(n..m)`:

```rust
fn main() {
  // range 1,2,3
  for n in 1..4 {
    println!("{}", n);
  }
  // reversed  
  for n in (1..4).rev() {
    println!("{}", n);
  }
}
```

## Ownership

ref: <https://doc.rust-lang.org/book/ch04-01-what-is-ownership.html>

- Some languages require explict memory allocation and release, eg C
- Others have garbage collection that takes care of memory management, eg Go
- Rust uses _ownership_ - a central feature that ensures memory safety without 
  the need for garbage collection

### Stack and Heap Memory

- Stack memory is _LIFO_ - Last In First Out
- Items are _pushed_ onto stack and _popped_ off stack
- Items stored on the stack must be of a known, fixed size
- The stack is faster
- Items of unknown or variable size must be stored on the heap which is less 
  organised and slower  
- Memory allocator finds a space that is big enough and returns a _pointer_ 
  (the address) to the memory location - this process is called _allocating on 
  the heap_ or just _allocating_.
- The _pointer_ itself is a memory location which is of a known and fixed size, 
  so the pointer can be stored on the stack.
- The data itself is retrieved by following the pointer to the heap location.
- When a function is called the arguments that get passed to the function, and 
  the functions local variables, get pushed to the stack. When the function is 
  over these values are popped off the stack.

### Ownership Rules

- Each value in Rust has a variable that is its _owner_
- Can only be _one_ owner at a time
- When the owner goes out of scope, the value will be dropped

### Scope, Move and Copy

- Example using strings...
- A string literal is immutable, known and fixed size so its value can be 
  hardcoded into the final executable - faster:
  
```rust
let s = "a string literal";
```

- Rust has another string type - `String` which is allocated onto the heap and 
  therefore can be used to store an amount of text that is unknown at compile 
  time
- Heap memory is requested from the memory allocator at run time - slower
- When variable is out of scope Rust automatically calls the `drop` function 
  which has the code to free the assigned heap memory:

```rust
{
    let s = String::from("a variable string"); // s in scope
} // s out of scope, so drop is called by Rust
```

- Semantics are similar to a slice in Go
- For example:

```rust
let s1 = String::from("Mike");
let s2 = s1;
println!("s1 = {}", s1);
```
- Here s1 and s2 both contain a `ptr`, `len` and `capacity`
- The `ptr` points to the same location in the heap, so both `s1` and `s2` are 
  pointing to the same location on the heap, and hence the same value.
- In other languages this would be referred to as a _shallow_ copy - ie, copying 
  the `ptr`, `len` and `capacity` but not the actual data.   
- However, when `s1` is out of scope Rust would call `drop` to release the heap 
  memory. This means that `s2` would have nothing to point to and would result 
  in what's called a _double-free_ error - That's why the above won't compile.
- Instead what Rust does is referred to as a _move_. This means it _invalidates_ 
  `s1` and effectively _moves_ the data to `s2`.
  
- Deep copies _can_ be made, eg:

```rust
let s1 = String::from("Mike");
let s2 = s1.clone();
println!("s1 = {}", s1);
```

- Scalar values that are stored on the stack do _not_ go out of scope like this
- Instead of being _moved_ thought invocation of the `drop` trait, they are 
  _copied_ through invocation of the `copy` trait:
  
```rust
let a = 5; 
let b = a;
println!("a and b are still both in scope, a = {}, b = {}", a, b);
```

- When args are passed to a function they function takes ownership and they go 
  out of scope:
  
```rust
// Here a and b stay in scope as they are stack values
let a = 5;
let b = 7;
let sum = sum_numbers(a, b);
println!("{}, {}", a, b); // <-- OK

// But here they are deactivated (or moved)
let a = String::from("dog");
let b = String::from("cat");
print_strings(a, b);
println!("{}, {}", a, b); // <-- compilation error
```

### References and Borrowing

ref: <https://doc.rust-lang.org/book/ch04-02-references-and-borrowing.html>

- As in Go, referencing operator is `&`, dereferencing operator is `*`
- Instead of passing ownership of a value to a function, can pass a _reference_:

```rust
fn main() {
  let s = String::from("abc");
  // &s is a reference to the value of s, but does not own it
  let len = str_len(&s);
  println!("string_len({}) = {}", s, len)
}

fn str_len(s: &String) -> usize {
  s.len()
} // here s goes out of scope but as it does not have ownership of what 
// it refers to nothing happens (maybe the stored memory address is removed??)
```

- Using _references_ as function parameters is called _borrowing_

```rust
fn change_str(s: &String) {
   s.push_str(", world"); // <- will error with:
   // ^ `s` is a `&` reference, so the data it refers to cannot be borrowed as mutable
}
```

- So, need to specify `mut` for the var and `&mut` for the reference:

```rust
fn main() {
  let mut s = String::from("Hello");
  change_str(&mut s);
}

fn change_str(s: &mut String) {
   s.push_str(", world");
}
```

- **Note**: can only have _one_ mutable reference to the same data in a 
  particular scope. This is to prevent data _races_ at compile time.
  
```rust
let mut s = String::from("hello");
let r1 = &mut s;
let r2 = &mut s; // < -- not allowed
```

- Can use curly brackets to create a new scope and allow for multiple, mutable
  references:

```rust
let mut s = String::from("hello");
{
    let r1 = &mut s;
} // r1 is now out of scope so new ref is allowed
let r2 = &mut s;
```

- Multiple immutable references are ok, but cannot mix immutable and mutable 
  references and this could effectively alter an immutable reference.
  
```rust
let s = String::from("hello");
let r1 = &s; // ok
let r2 = &s; // ok
let r3 = &mut s; // NOT ok
println!("{}, {}, and {}", r1, r2, r3);
```

- Scope of a reference starts where it is created and ends at its last use, so 
  this is ok:
  
```rust
fn main() {
    let mut s = String::from("hello");

    let r1 = &s; // no problem
    let r2 = &s; // no problem
    println!("{} and {}", r1, r2);
    // r1 and r2 are no longer used after this point

    let r3 = &mut s; // no problem
    println!("{}", r3);
}
```

- _Dangling references_ are pointers to data that has been de-allocated
- Rust prevents dangling references:

```rust
fn dangle() -> &String { // dangle returns a reference to a String
    // But the string itself is created inside this function...
    let s = String::from("hello");
    &s // return a reference to the String
} // but here s goes out of scope, and is dropped. Its memory goes away.
```

- To avoid this can just return the string:

```rust
fn no_dangle() -> String {
    let s = String::from("hello");
    s
}
```
  
### Slice Type

- Does not have ownership
- References a contiguous sequence of elements in a collection (like Go slice)

```shell
let s = String::from("hello world");
let s1 = s[..4];  // index 0 to index 3
let s2 = s[3..7]; // index 3 to index 6
let s3 = s[4..];  // index 4 to the end
let s4 = s[..];   // slice of entire string
```

### Structs

- Named custom data type that groups related values (fields) and behaviour 
  (methods)
- Field labels add more meaning to data  

```rust
struct User {
 username: String,
 email: String,
 full_name: String,
}

let u1 = User {
  username: String::from("miked"),
  email: String::from("miked"),
  full_name: String::from("Michael Donnici"),
}
```

- Field values are retrieved and set using dot notation
- To mutate a struct value the entire instance must be mutable - cannot just 
  specify some fields as mutable. 
  
#### Field  Init Shorthand

- Can use _field init shorthand_ to initialise field values:

```rust
fn new_user(username: String, email: String, full_name: String) -> User {
  // Field Init Shorthand: because field names match param names
  User {
    username,
    email,
    full_name,
  }
}
```

#### Struct Update Syntax

- Create a new instance and copy remaining values from another instance

```rust

struct CountThings {
  foo: uint32,
  bar: uint32,
  baz: uint32,
  bing: uint32,
}

let things1 = CountThings {
    foo: 8,
    bar: 8,
    baz: 10,
    bing: 10,
}

let things2 = CountThings {
    foo: 10,
    bar: 10,
    ..things1
}
```

#### Tuple structs

- Structs without field names
- Each is still a distinct type:

```rust
fn main() {
    struct Color(i32, i32, i32);
    struct Point(i32, i32, i32);

    let black = Color(0, 0, 0);
    let origin = Point(0, 0, 0);
}
```

#### Unit-Like Structs without Fields

- Used when need to implement a trait on a type but don't need any data

#### Printing struct values using `println!`

- Default behaviour for `println!` using `{}` is to to use `Display`
- Primitives implement `Display` by default but structs do not
- `{:?}` or `{:#?}` (pretty print) specifies print format using the `Debug` 
- Adding the annotation `#[derive(Debug)]` will implement the `Debug` trait:

```rust
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };

    println!("rect1 is {:?}", rect1);
}
// -> rect1 is Rectangle { width: 30, height: 50 }
```

### Methods

- Functions defined within the context of a struct
- First arg is always `self` - the instance
- Define an `impl` (implementation block) ... like a receiver (?)
- Can have multiple methods in an `impl` block and can also have multiple 
  `impl` blocks for the same receiver.

```rust
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    println!("rect1 area {}", rect1.area());
}
```

#### Associated Functions

- Functions defined in an `impl` block that _don't_ take self as a parameter
- Are not methods as they don't need an instance - so like _static_ functions
- Invoked using `::`, eg `String::from("hello")`
- Often used as a constructor, ie to return an instance of the type

```rust
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn square(size: u32) -> Rectangle {
        Rectangle {
            width: size,
            height: size,
        }
    }
}

fn main() {
    let rect = Rectangle {
        width: 30,
        height: 50,
    };
    println!("rect area {}", rect.area());

    let squ = Rectangle::square(10);
    println!("squ = {:?}", squ);
}
```

### Enums and Pattern Matching

- Enumerations or _enums_ allow a type to be defined by enumerating its possible 
  variants
- Can define enums with, or without associated data types
- Enum values are namespaced by identifier and use `::`
- Provide a more compact way to do what could be done with a struct
- Can also create methods using a `impl` block

```rust
// These would be analogous to iota enums in Go
#[derive(Debug)]
enum Fruit {
  Apple,
  Orange,
  Banana,
}

impl Fruit {
  fn eat(&self) {
    println!("You ate the {:?}", self);
  }
}

// More enums, some with associated data types
enum Message {
  Quit,                        // no data type
  Move { x: i32, y: i32 },     // an anonymous struct
  Write(String),               // a String
  ChangeColor(i32, i32, i32),  // a three item Tuple of i32 
}

fn main() {
  let f1 = Fruit::Apple;
  f1.eat();
  let msg = Message::Write(String::from("hello"));
}
```

#### The Option Enum

- `Option` enum is defined by the standard library and provides a type that can 
   be used where a value could be something, or could be nothing
- As such, Rust does not have `null`

```rust
enum Option<T> {
    Some(T),
    None,
}
```

- `Option` is included in the _prelude_ so does not have to be brought into 
  scope - and its variants can be used without the `Option::` prefix.
- `<T>` is a generic type parameter so the `Some` variant of the enum `Option` 
  can hold data of any type:
  
```rust
fn main() {
  let some_number = Some(5);
  let some_string = Some("a string"); 
  let absent_number: Option<i32> = None;
}
```

- So a data type can really never be null, unless it is an `Option` and, when 
  this is the case, the possibility of a null value _must_ be handled explicitly 
  or there will be compiler errors:
  
```rust
fn main() {
  let x: i8 = 5;
  let y: Option<i8> = Some(5);

  let sum = x + y; // <- error!
}
```

- See the [`Option`](https://doc.rust-lang.org/std/option/enum.Option.html) docs 
  for more info. 

#### `match` control flow operator

- Used to execute code based on pattern matches 
- Used for handling `enum` variants
- Similar to a `switch`
- Form:

```
match [expression] { <-- expression can be any value
    [pattern] => [code; expression], <-- match arm
    [pattern] => [code; expression], <-- match arm
    [pattern] => [code; expression], <-- match arm
}
```

- First arm that matches will return that expression from the `match {}` block
- Can run multiple lines of code from a `match` arm:

```rust
enum Coin {
  Penny,
  Nickel,
  Dime,
  Quarter,
}

fn value_in_cents(coin: Coin) -> u8 {
  match coin {
    Coin::Penny => {
      println!("Lucky penny!");
      1
    }
    Coin::Nickel => 5,
    Coin::Dime => 10,
    Coin::Quarter => 25,
  }
}
```





  
