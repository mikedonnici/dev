# Rust

- [The Rust Book](https://doc.rust-lang.org/stable/book)
- [Cargo Guide](https://doc.rust-lang.org/cargo/guide/)
- [Crate Register](https://crates.io/)

## Install and Update

```script
# Installation
$ curl https://sh.rustup.rs -sSf | sh

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
