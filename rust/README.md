# Rust

- [The Rust Book](https://doc.rust-lang.org/stable/book)
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






