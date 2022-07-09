# Rust

- [The Rust Book](https://doc.rust-lang.org/stable/book)
- [Rust By Example](https://doc.rust-lang.org/stable/rust-by-example)
- [Cargo Guide](https://doc.rust-lang.org/cargo/guide/)
- [Crate Register](https://crates.io/)

## Install and Update

```script
# Installation
curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh

# rustup toolchain manager
rustup --version

# rust compiler
rustc --version

# update
rustup update

# uninstall
rustup self uninstall
```

## Local Docs

```bash
rustup doc
```

## Build & Run

With the compiler:

```script
$ rustc main.rs
./main
```

With cargo:

```script
# create a new package
cargo new pkg_name
cd pkg_name

# build - also downloads depedencies in Cargo.toml
cargo build

# run
./target/debug/pkg_name

# build and run - from project root
# compiles binary with debug symbols by default -> debug/ folder
cargo run

# --release build without debug symbols, slower to compile -> release/ folder
cargo run --release 



# compilation check only (no build)
cargo check

# Update depedencies - ignored Cargo.lock
cargo update 
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
- Convention is UPPERCASE_UNDERSCORES (screaming snakecase)
- Type annotation always required
- Can only set with a _constant expression_, not func call or computed value
- Can be module scope (global) and are inlined at compile time so very fast

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
- Fixed length once declared (arity)
- Limited to size of 12
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
- Limited to size of 32, after which us a Vec
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
    // this is called a 'tail expression'
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

- to `break` or `continue` nested loops label a single tick identifier:

```rust
fn main() {
    'l1 loop {
        loop {
            loop {
                break 'l1;
            }
        }
    }
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
    // range inclusive of last item
    for n in 1..=4 {
        println!("{}", n);
    }
}
```

## Ownership

ref: <https://doc.rust-lang.org/book/ch04-01-what-is-ownership.html>

- Some languages require explicit memory allocation and release, eg C
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

- Deep copies _can_ be made using `.clone()`, eg:

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
let r1 = & mut s;
let r2 = & mut s; // < -- not allowed
```

- Can use curly brackets to create a new scope and allow for multiple, mutable
  references:

```rust
let mut s = String::from("hello");
{
let r1 = & mut s;
} // r1 is now out of scope so new ref is allowed
let r2 = & mut s;
```

- Multiple immutable references are ok, but cannot mix immutable and mutable
  references and this could effectively alter an immutable reference.

```rust
let s = String::from("hello");
let r1 = & s; // ok
let r2 = & s; // ok
let r3 = & mut s; // NOT ok
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

```bash
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
    Quit,
    // no data type
    Move { x: i32, y: i32 },
    // an anonymous struct
    Write(String),
    // a String
    ChangeColor(i32, i32, i32),  // a three item Tuple of i32 
}

fn main() {
    let f1 = Fruit::Apple;
    f1.eat();
    let msg = Message::Write(String::from("hello"));
}
```

#### Option Enum

- [`Option`](https://doc.rust-lang.org/std/option/enum.Option.html) enum is
  defined by the standard library and provides a type that can be used where a
  value could be something, or could be nothing
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

#### Result Enum

- [`Result`](https://doc.rust-lang.org/std/result/enum.Result.html) enum is when
  return may be a value or an error
- `must_use` annotation creates a compiler warning if all possible variants of
  the Result are not considered in some way
- `Result` is generic over separate types for return value (T) and an error (U)

```rust
#[must_use]
enum Result<T, U> {
    Ok(T),
    Err(U),
}
```

```rust
fn main() {
    let res = do_thing(true);
    match res {
        Ok(c) => println!("Result ok, code = {}", c),
        Err(e) => println!("Error, code = {}", e),
    }
}

fn do_thing(ret_err: bool) -> Result<i32, i32> {
    if ret_err {
        return Err(1);
    }
    Ok(0)
}
```

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

##### `match` patterns that bind to values

- Can match enums that hold additional values, for example:

```rust
enum Genus {
    Dionaea,
    Drosera,
    Sarracenia(Species),
}

#[derive(Debug)]
enum Species {
    Flava,
    Minor,
    Rubra,
}

fn main() {
    let p1 = Genus::Sarracenia(Species::Flava);

    match p1 {
        Genus::Dionaea => {
            println!("Fly trap");
        }
        Genus::Drosera => {
            println!("Sundew");
        }
        Genus::Sarracenia(species) => {
            println!("Sarracenia {:?}", species);
        }
    }
}
```

##### Matching with `Option<T>`

- Getting the value from an `Option`, here `None` is assigned a value of 0:

```rust
fn main() {
    println!("{}", opt_value(Some(5)));
    println!("{}", opt_value(Some(0)));
    println!("{}", opt_value(None));
}

fn opt_value(opt: Option<u8>) -> u8 {
    match opt {
        Some(i) => i,
        None => 0,
    }
}
```

- Matches are exhaustive, for the code to compile all possible values must be
  accounted for

- The `_` placeholder can be used for all remaining values:

```rust
fn main() {
    let n: u8 = 43;
    match n {
        1 => {
            println!("{}", "One");
        }
        2 => {
            println!("{}", "Two");
        }
        3 => {
            println!("{}", "Three");
        }
        _ => () // unit value - nada
    }
}
```

- `if let` can be used in the case where you want to match only one value:

```rust
fn main() {
    let some_u8_val = Some(0u8);

    // So this...
    match some_u8_val {
        Some(3) => println!("Three"),
        _ => (),
    }

    // ...can be written thus:
    if let Some(3) = some_u8_val {
        println!("Three");
    }

    // Can also include an else block:
    if let Some(3) = some_u8_val {
        println!("Three");
    } else {
        println!("Not three");
    }
}
```

## Code organisation - the module system

Ref: <https://doc.rust-lang.org/cargo/guide/project-layout.html>

### Packages and Crates

- A _crate_ is a binary or library
- The `crate root` is a source file that the compiler starts from and makes up
  the root module of the crate.
- A _package_ is one or more crates that provides a set of functionality
- A _package_ contains a `Cargo.toml` file that describes how to build the
  crates.
- A _package_ must contain:
    - at least one crate
    - zero or one library crates
    - any number of binary crates

Example:

```bash
cargo new my-pkg

my-pkg     # the pkg name
├── Cargo.toml
└── src/
    ├── lib.rs   # a library crate
    ├── main.rs  # the module root for binary crate named my-pkg
    └── foo.rs   # a second binary crate
```

### Modules - controlling scope and privacy

- Modules are used to organise code within a crate for readability and re-use
- Modules control _public_ / _private_ access to items (privacy boundary)

Example:

```bash
cargo new --lib restaurant 
```

```bash
restaurant
├── Cargo.toml
└── src/
    └── lib.rs
```

`lib.rs` with some nested (parent / child) modules defined with `mod` keyword

```rust
mod front_of_house {
    mod hosting {
        fn add_to_waitlist() {}

        fn seat_at_table() {}
    }

    mod serving {
        fn take_order() {}

        fn server_order() {}

        fn take_payment() {}
    }
}
```

This equates to the following module _tree_:

```bash
crate
└── front_of_house
    ├── hosting
    │   ├── add_to_waitlist
    │   └── seat_at_table
    └── serving
        ├── take_order
        ├── serve_order
        └── take_payment
```

### Paths - referencing items in a module tree

- _Absolute path_ starts from crate root or literal crate
- _Relative path_ starts from current module with `self`, `super` or an
  identifier in the current module.
- Both types are followed by one or more identifiers separated by `::`

Example:

```rust
mod front_of_house {
    mod hosting {
        fn add_to_waitlist() {}
    }
}

pub fn eat_at_restaurant() {

    // Absolute path
    crate::front_of_house::hosting::add_to_waitlist();

    // Relative path
    front_of_house::hosting::add_to_waitlist();
}
```

- This would not compile because child modules are not accessible by their
  parent modules (see `pub` section below).
- Child modules _can_ access their ancestors because they are defined _within_
  the scope of their ancestors (like closures).
- use the `pub` keyword to expose paths, eg to expose the `add_to_waitlist()`
  function above:

```rust
mod front_of_house {
    pub mod hosting {
        pub fn add_to_waitlist() {}
    }
}
```

- can also use `super` at the start of the module path to refer to the _parent_
  module - analogous to `..` in a file system path

## Privacy for structs and enums

- Structs can be declared `pub` but the struct fields will remain private by
  default.
- Each field must be explicitly declared `pub`, as required
- In the case where some struct fields remain private, a _public_, _associated_
  function is required to create an instance of the struct.

```rust
mod some_mod_name {
    pub struct Foo {
        pub a: string,
        b: string,
    }

    impl Foo {
        pub fn new(a_val: string) -> Foo {
            Foo {
                a: a_val,
                b: "default b",
            }
        }
    }
}
```

- If an `enum` is made public, all of its variants are public

```rust
mod some_mod_name {
    pub enum Colour {
        Red,
        Green,
        Blue,
    }
}

pub fn do_thing() {
    let red = some_mod_name::Colour::Red;
    let green = some_mod_name::Colour::Green;
}
```

## `use` - bringing modules into scope

- Can access module items with the full path each time:

```rust
fn some_func() {
    let res1 = crate::parent_mod::child_mod::some_fn();
    let res2 = crate::parent_mod::child_mod::some_fn();
}
```

- Easier with `use` paths:

```rust
use crate::parent_mod::child_mod;

fn some_func() {
    let res1 = child_mod::some_fn();
    let res2 = child_mod::some_fn();
}
```

- Note that the above is the _idiomatic_ way to implement `user` - that is, to
  import the parent module and call the function by referencing the parent
  module. This makes it clear that the function is _not_ locally defined while
  still reducing the path clutter.
- It is possible to do it like this as well, but not idiomatic:

```rust
use crate::parent_mod::child_mod::some_fn;

fn some_func() {
    let res1 = some_fn();
    let res2 = some_fn();
}
```

- Conversely, it _is_ idiomatic to use the full path when importing structs and
  enums:

```rust
use crate:: mod::SomeStruct;
use crate:: mod::SomeEnum;
```

- If there is a name clash can `use ... as ...`:

```rust
use crate::mod_one::FooStruct as FooOne;
use crate::mod_two::FooStruct as FooTwo;
```

- Items brought into scope are private to the local scope, however they can be _
  re-exported_ and thus made available to calling code:

```rust
pub use crate::mod1::mod1::item;

```

- [External packages](https://crates.io) must first be added to `Cargo.toml`
- The standard library `std` is also _external_ but is shipped with Rust so does
  not need to be added to `Cargo.toml`, but still need to bring into scope, eg:

```rust
use std::collections::HashMap;
```

- _Nested paths_ can be used to clean up large `use` lists
- Paths can be nested from any level in the path

```rust
use std: io;
use std::cmp::Ordering;
```

can be:

```rust
use std::{io, cmp::Ordering};
```

- _Glob_ operator `*` can also be used to bring _all_ items
- Can make it hard to know where names came from or how they came into scope

```rust
use std::collections::*;
```

## Separating modules into different files

Excellent
article: [Rust modules and project structure](https://medium.com/codex/rust-modules-and-project-structure-832404a33e2e)

- Rust requires you to explicitly include code in your application
- All files and folders and _modules_
- `mod` imports code from a module to the location where the `mod` statement is
  used
- Importing a module automatically creates a namespace (from the file name) to
  avoid name conflicts
- `use` provides a convenience to map fully qualified type names to something
  shorter

Example:

```
crate
└── lib.rs
├── main.rs
│   ├── add_to_waitlist
│   └── seat_at_table
└── serving
├── take_order
├── serve_order
└── take_payment
```

```rust
// lib.rs


```

```rust
// main.rs
fn main() {}
```

Rust 2015 used `dir/mod.rs` to define sub modules but in

See `modules` projects herein, and
also [project layout](https://doc.rust-lang.org/cargo/guide/project-layout.html)
.

## Collections

- [Collections](https://doc.rust-lang.org/std/collections/) are data structures
  for storing a dynamic number of values
- As size is dynamic, collections are stored on the _heap_
- Three commonly used types
  are [vectors](https://doc.rust-lang.org/std/vec/struct.Vec.html),
  [strings](https://doc.rust-lang.org/std/string/)
  and [hash maps](https://doc.rust-lang.org/std/collections/hash_map/struct.HashMap.html)

### Vectors - Vec<T>

- Stores multiple values of same type
- Implemented with generics
- Provided with `std` library
- Create a new, empty Vec

```rust
let v1<i32> = Vec::new();
```

- Infer type with initial values using `vec!` macro

```rust
let v2 = vec![1, 2, 3, 4];
```

- Update a vector
  ([play](https://play.rust-lang.org/?version=stable&mode=debug&edition=2021&gist=da21d80b31211078d9c92cfa45c61078))

```rust
fn main() {
    {
        let mut v = vec![1, 2, 3];
        v.push(4);
        v.push(5);
        println!("{:?}", v);
    } // <- v out of scope and is freed
    println!("{:?}", v); // <- error
}
```

- 2 ways to read elements of a vector, `&[]` and `get()`
  ([play](https://play.rust-lang.org/?version=stable&mode=debug&edition=2021&gist=f61b5b0b85311c72e785311756402d5d)):

```rust
fn main() {
    let v = [1, 2, 3, 4, 5];

    // Use indexing - pointer to value
    let a: &i32 = &v[2];
    println!("v[{}] = {}", 2, a);

    // Panics when the index is out of range
    // let b: &i32 = &v[99]; // <-- this will panic
    // println!("v[{}] = {}", 99, b);

    // Using .get(n) returns an Option enum so is best option if don't want 
    // a panic when the index is out of range.
    match v.get(2) { // returns an Option<T>
        Some(n) => println!("3rd element is {}", n),
        None => println!("There is no 3rd element."),
    }
    match v.get(99) { // returns an Option<T>
        Some(n) => println!("100th element is {}", n),
        None => println!("There is no 100th element."),
    }
}
```

- Iterating over a
  Vec [play](https://play.rust-lang.org/?version=stable&mode=debug&edition=2021&gist=7d9507be9d6a22b60e67df9f51b59742)

```rust
fn main() {
    let mut v = [1, 2, 3, 4, 5];
    for i in &v {
        println!("{}", i)
    }

    // loop with mutable references to change items in the vec
    for i in &mut v {
        *i += 10;
    }

    for i in &v {
        println!("{}", i)
    }
}
```

## Closures

- Closures are anonymous, inline functions, like a `lamda` in Python.
- format is param list followed by a block, eg  `|x, y| {x + y}`

```rust
fn main() {
    let p = |x| { x * 2 };
    let nn = vec![1, 2, 3, 4, 5];
    for n in nn {
        println!("2 x {}= {}", n, p(n));
    }
}
```

- a closure will borrow a reference to a variable in its scope which is fine
  provided the closure does not outlive the variable it is referencing

So, this is fine:

```rust
fn main() {
    let s = "dog".to_string();
    let p = || {
        println!("The string is '{}'", s)
    };
    p();
}
```

But this throws an error:

```rust
fn main() {
    {
        let s = "dog".to_string();
    } // <-- s is out of scope   
    let p = || {
        println!("The string is '{}'", s)
    };
    p();
}
```

- Can use `move` semantics to force the closure to take ownership of a variable:

```rust
fn main() {
    let s = "dog".to_string();
    let p = move || {
        println!("The string is '{}'", s);
    };
    p();
}
```

## Threads

- Like go routines
- Portable, so works on various platforms
- Require overhead for context switching so best for situations where can take
  advantage of multiple cpu cores
- For situations where delay is related to network or disk, `async await` is a
  better choice

```rust
use std::thread;

fn main() {

    // pass in a closure with no arguments which runs the main fn of the thread 
    let handle = thread::spawn(move || {
        run();
    });

    // do some other stuff in main thread
    println!("running main()");

    // wait until the 'run' thread has exited
    handle.join().unwrap();
}

fn run() {
  println!("running run()");
}
```
