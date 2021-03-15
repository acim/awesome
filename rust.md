# Rust

- [Awesome Rust](https://github.com/rust-unofficial/awesome-rust)
- [Dual licensing](https://github.com/sfackler/rust-postgres-macros/issues/19)
- [Procedural macros workshop](https://github.com/dtolnay/proc-macro-workshop)
- [Building distributed GraphQL backend using Rust and Apollo Federation](https://dev.to/rkudryashov/building-distributed-graphql-backend-using-rust-and-apollo-federation-50bm)
- [On Generics and Associated Types](https://blog.thomasheartman.com/posts/on-generics-and-associated-types)
- [Allowed and denied lints](https://doc.rust-lang.org/rustc/lints/listing/allowed-by-default.html)
- [Building a runtime reflection system for Rust 1/3: dyn Class](https://www.osohq.com/post/rust-reflection-pt-1)
- [Building a runtime reflection system for Rust 2/3: dyn Attribute](https://www.osohq.com/post/runtime-reflection-pt-2)
- [Building a runtime reflection system for Rust 2/3: dyn Method](https://www.osohq.com/post/runtime-reflection-pt-3)
- [Closures in Rust: Capture by Immutable Borrow, by Mutable Borrow, by Move, by Unique Immutable Borrow](https://zhauniarovich.com/post/2020/2020-12-closures-in-rust)
- [Recursive iterators in Rust](https://fasterthanli.me/articles/recursive-iterators-rust)
- [Rust Traits and Trait Objects](https://joshleeb.com/posts/rust-traits-trait-objects.html)
- [Rust Language Cheat Sheet](https://cheats.rs/)
- [Programming idioms - Go vs Rust](https://programming-idioms.org/cheatsheet/Go/Rust)
- [rust-learning - bunch of links to blog posts, articles, videos, etc for learning Rust](https://github.com/ctjhoa/rust-learning)
- [Rosetta Code](http://rosettacode.org/wiki/Category:Rust)
- [Rust Programming Techniques](https://www.youtube.com/watch?v=vqavdUGKeb4)
- [From Go to Rust - JSON and YAML](http://technosophos.com/2018/06/12/from-go-to-rust-json-and-yaml.html)
- [From Go to Rust - Unit Testing](http://technosophos.com/2018/07/07/from-go-to-rust-testing.html)
- [From Go To Rust - Advanced Testing](http://technosophos.com/2018/07/25/from-go-to-rust-advanced-testing.html)
- [Procedural Macros](https://www.youtube.com/watch?v=9IbYBl48eTQ)
- [Generic `impl` blocks are kinda like macros](https://dev.to/somedood/generic-impl-blocks-are-kinda-like-macros-1aa0)
- [Hello Hyper!](https://www.fpcomplete.com/blog/captures-closures-async/)

## YouTube channels

- [David Pedersen](https://www.youtube.com/channel/UCDmSWx6SK0zCU2NqPJ0VmDQ/videos)

## Blogs

- [Sven Assmann](https://www.d34dl0ck.me/)

## Books

- [The Rust Programming Language](https://doc.rust-lang.org/book/)
- [Rust Cookbook](https://rust-lang-nursery.github.io/rust-cookbook/)
- [Engineering Rust Web Applications](https://erwabook.com/)
- [Zero To Production In Rust](https://www.zero2prod.com/)
- [Rust Design Patterns](https://rust-unofficial.github.io/patterns/)
- [The Little Book of Rust Books](https://lborb.github.io/book/title-page.html)
- [Learn Rust by writing Entirely Too Many Linked Lists](https://rust-unofficial.github.io/too-many-lists)
- [Rust API Guidelines](https://rust-lang.github.io/api-guidelines/)

## Traits

- [3 Things to Try When You Can't Make a Trait Object](https://www.possiblerust.com/pattern/3-things-to-try-when-you-can-t-make-a-trait-object)
- [Fn - used for closures](https://doc.rust-lang.org/std/ops/trait.Fn.html)
- [FnMut - used for closures](https://doc.rust-lang.org/std/ops/trait.FnMut.html)
- [FnOnce - used for closures](https://doc.rust-lang.org/std/ops/trait.FnOnce.html)
- Forms of syntax when used as type:

```rust
&dyn Trait
&Trait
Box<Trait>
Box<dyn Trait>
```

## Libraries

- [The Rust Standard Library](https://doc.rust-lang.org/std/)
- [Image - basic image processing functions and methods for converting to and from various image formats](https://github.com/image-rs/image)
- [clap - Command Line Argument Parser](https://github.com/clap-rs/clap)
- [kadabra - tool that makes data sharing between terminal windows easy](https://crates.io/crates/kadabra)
- [mdBook - utility to create modern online books from Markdown files](https://crates.io/crates/mdbook)
- [Yew - modern Rust framework for creating multi-threaded front-end web apps with WebAssembly](https://github.com/yewstack/yew)
- [nom - parser combinators library written in Rust.](https://github.com/Geal/nom)
- [libp2p - protocols, specifications and libraries for peer-to-peer network applications](https://github.com/libp2p/rust-libp2p)
- [SQLx - async SQL crate featuring compile-time checked queries](https://github.com/launchbadge/sqlx)

## Concurrency and parallelism

- [Tokio - runtime for writing reliable, asynchronous and slim applications](https://github.com/tokio-rs/tokio)
- [Rayon - data-parallelism library](https://github.com/rayon-rs/rayon)
- [Crossbeam - set of tools for concurrent programming](https://github.com/crossbeam-rs/crossbeam)
- [Concurrency in modern programming languages: Rust 1/2](https://dev.to/deepu105/concurrency-in-modern-programming-languages-introduction-ckk)
- [Concurrency in modern programming languages: Rust 2/2](https://dev.to/deepu105/concurrency-in-modern-programming-languages-rust-19co)
- [Actors with Tokio](https://ryhl.io/blog/actors-with-tokio/)

## Kubernetes

- [Writing a Kubernetes CRD Controller in Rust](http://technosophos.com/2019/08/07/writing-a-kubernetes-controller-in-rust.html)
- [Evolution of kube](https://clux.github.io/probes/post/2021-02-28-kube-evolution/)

## GUI

- [Creating a GTK/Glade based GUI for a Rust application](https://dev.to/henrybarreto/creating-a-gui-for-a-rust-application-2h1g)

## HTTP server

- [tide example app - fully fledged backend application built with Rust and tide including CRUD operations, authentication, routing, pagination, and more](https://github.com/colinbankier/realworld-tide)
- [Validating JSON input in Rust web services](https://dev.to/vinted/validating-json-input-in-rust-web-services-5gp0)
- [Building Powerful GraphQL Servers with Rust](https://dev.to/open-graphql/building-powerful-graphql-servers-with-rust-3gla)

## Tools

- [Rustup - installer for the systems programming language Rust](https://rustup.rs/)
- [Rust Search Browser Extension](https://rust.extension.sh/)

## Performance

- [Automatic Flamegraphs for Benchmarks in Rust](https://www.jibbow.com/posts/criterion-flamegraphs/)

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

- [Macros in Rust: A tutorial with examples](https://blog.logrocket.com/macros-in-rust-a-tutorial-with-examples/)

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

### unimplemented!

When a function is not implemented, use unimplemented! macro inside the body.

## Error handling

- [Wrapping errors in Rust](https://edgarluque.com/blog/wrapping-errors-in-rust)

### [failure crate - handle multiple error types, chain errors](https://docs.rs/failure/0.1.8/failure/)

```rust
#[derive(Fail)]
pub enum MyError {
    Server(u8)
    User(String)
    Connection,
}
```

### Mutex poison to custom error conversion

```rust
#[derive(Debug)]
enum MyError {
    Lock,
}

impl<T> From<PoisonError<T>> for MyError {
    fn from(_: PoisonError<T>) -> MyError {
        MyError::Lock
    }
}

fn main() -> Result<(), MyError> {}
```

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

## Iterators

- [Iteration patterns for Result & Option](http://xion.io/post/code/rust-iter-patterns.html)
- [Returning Rust Iterators](https://depth-first.com/articles/2020/06/22/returning-rust-iterators/)
- [Creating an Iterator in Rust](https://aloso.github.io/2021/03/09/creating-an-iterator)

### Iterator over struct fields

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

## Stricter compiler lints

```rust
#![warn(missing_debug_implementations, rust_2018_idioms, missing_docs)]
```

## List all compiler lints

```sh
rustc -W help
```

## Update Rust toolchain

```sh
rustup update stable
```

## _use_ can be used inside functions

```rust
fn print_power_action(state: PowerState) {
    use PowerState::*;
    match state {
        Off,
        Sleep,
        Reboot,
        Shutdown,
        Hibernate,
    }
}
```

## PostgreSQL mTLS connection

```rust
extern crate openssl;
extern crate postgres;

use postgres::{Connection, TlsMode};

use openssl::ssl::{SslConnectorBuilder, SslMethod, SslVerifyMode};
use openssl::x509;

fn main() {
    let mut connector = SslConnectorBuilder::new(SslMethod::tls()).unwrap();
    connector.set_ca_file("root.crt").unwrap();
    connector
        .set_certificate_file("postgresql.crt", x509::X509_FILETYPE_PEM)
        .unwrap();

    connector
        .set_private_key_file("postgresql.key", x509::X509_FILETYPE_PEM)
        .unwrap();

    // openssl::ssl::SslVerfifyMode constant in not defined yet in openssl 0.9.23 which is rust-postgres dependency
    // disable certificate hostname check
    let mode = SslVerifyMode::empty();
    connector.set_verify(mode);

    let negotiator = postgres::tls::openssl::OpenSsl::from(connector.build());

    let conn = Connection::connect(
        "postgres://postgres@localhost:5432",
        TlsMode::Require(&negotiator),
    ).unwrap();
    let res = conn.query("SELECT 1+1 as foo", &[]).unwrap();
    for row in &res {
        let foo: i32 = row.get(0);
        println!("{}", foo);
    }
}
```
