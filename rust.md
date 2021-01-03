# Rust

- [Dual licensing](https://github.com/sfackler/rust-postgres-macros/issues/19)
- [Procedural macros workshop](https://github.com/dtolnay/proc-macro-workshop)
- [Building distributed GraphQL backend using Rust and Apollo Federation](https://dev.to/rkudryashov/building-distributed-graphql-backend-using-rust-and-apollo-federation-50bm)
- [On Generics and Associated Types](https://blog.thomasheartman.com/posts/on-generics-and-associated-types)

## Books

- [The Rust Programming Language](https://doc.rust-lang.org/book/)
- [Rust Cookbook](https://rust-lang-nursery.github.io/rust-cookbook/)
- [Engineering Rust Web Applications](https://erwabook.com/)
- [Zero To Production In Rust](https://www.zero2prod.com/)

## Libraries

- [The Rust Standard Library](https://doc.rust-lang.org/std/)
- [Image - basic image processing functions and methods for converting to and from various image formats](https://github.com/image-rs/image)
- [clap - Command Line Argument Parser](https://github.com/clap-rs/clap)

## Concurrency and parallelism

- [Tokio - runtime for writing reliable, asynchronous and slim applications](https://github.com/tokio-rs/tokio)
- [Rayon - data-parallelism library](https://github.com/rayon-rs/rayon)
- [Crossbeam - set of tools for concurrent programming](https://github.com/crossbeam-rs/crossbeam)
- [Concurrency in modern programming languages: Rust 1/2](https://dev.to/deepu105/concurrency-in-modern-programming-languages-introduction-ckk)
- [Concurrency in modern programming languages: Rust 2/2](https://dev.to/deepu105/concurrency-in-modern-programming-languages-rust-19co)

## Tools

- [Rustup - installer for the systems programming language Rust](https://rustup.rs/)

## Container images

- [distroless](https://github.com/GoogleContainerTools/distroless/blob/master/examples/rust/Dockerfile)

## Trait bounds

```rust
fn notify(item: impl Summary) -> String {}

fn notify<T: Summary>(item: T) -> String {}

fn notify<T> Summary(item: T) -> String
    where T: Summary {}
```

## Macros

### matches!

```rust
fn is_first(data: Test) -> bool {
    match data {
            Test::FIRST => true,
            _ => false,
        }
}
```

```rust
fn is_first(data: Test) -> bool {
    matches!(data, Test::FIRST)
}
```

## Error handling

### Do not unwind stack on panic

```rust
[profile.dev]
panic = "abort"

[profile.release]
panic = "abort"
```

### Match error kind

```rust
let f = match File::open("file.txt") {
    Ok(file) => do_something_with(file),
    Err(e) => match e.kind() {
        ErrorKind::NotFound => match File::create("file.txt") {
            Ok(file2) => do_something_with(file2),
            Err(e2) => panic!("{:?}", e2)
        },
        _ => panic!("something weird happened")
    }
}
```

### Return errors in functions

```rust
fn read_file() => Result<String, io::Error> {
    let mut f = File::open("file.txt") {
        Ok(file) => do_something_with(file),
        Err(e) => return Err(e), // here we need return
    }

    let mut s = String::new();
    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),   // return not needed here because match is the last expression in the function
        Err(e) => Err(e), // return not needed here because match is the last expression in the function
    }
}
```

### Short way to return errors in functions

```rust
fn read_file() => Result<String, io::Error> {
    let mut f = File::open("file.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}
```

```rust
fn read_file() => Result<String, io::Error> {
    let mut s = String::new();
    File::open("file.txt")?.read_to_string(&mut s)?;
    Ok(s)
}
```

### Return error in main()

```rust
fn main() => Result<(), Box<dyn Error>> {
    let f File::open("file.txt")?;
    Ok(s)
}
```

## Iterator over struct fields

```rust
use std::iter::once;

struct Color {
    r: u8,
    g: u8,
    b: u8,
}

impl Color {
    fn iter(&self) -> impl Iterator<Item = u8> {
        once(self.r).chain(once(self.g)).chain(once(self.b))
    }
}

fn main() {
    let color = Color {
        r: 100,
        g: 0,
        b: 150,
    };
    for v in color.iter() {
        println!("{}", v);
    }
}
```

## Cast to std::any::Any

```rust
Box::new(Foo { data: 32 }) as Box<dyn Any>
```

## View create documentation skipping dependencies

```sh
cargo doc --no-deps --open
```
