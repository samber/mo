# mo - Monads

[![tag](https://img.shields.io/github/tag/samber/mo.svg)](https://github.com/samber/mo/releases)
[![GoDoc](https://godoc.org/github.com/samber/mo?status.svg)](https://pkg.go.dev/github.com/samber/mo)
![Build Status](https://github.com/samber/mo/actions/workflows/go.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/mo)](https://goreportcard.com/report/github.com/samber/mo)
[![codecov](https://codecov.io/gh/samber/mo/branch/master/graph/badge.svg)](https://codecov.io/gh/samber/mo)

🦄 **`samber/mo` brings monads and populars FP abstractions to Go projects. `samber/mo` uses the recent Go 1.18+ Generics.**

**Inspired by:**

- Scala
- Rust
- FP-TS

**See also:**

- [samber/lo](https://github.com/samber/lo): A Lodash-style Go library based on Go 1.18+ Generics
- [samber/do](https://github.com/samber/do): A dependency injection toolkit based on Go 1.18+ Generics

**Why this name?**

I love **short name** for such utility library. This name is similar to "Monad Go" and no Go package currently uses this name.

## 💡 Features

We currently support the following data types:

- `Option[T]` (Maybe)
- `Result[T]`
- `Either[A, B]`
- `Future[T]`
- `IO[T]`
- `IOEither[T]`
- `Task[T]`
- `TaskEither[T]`
- `State[S, A]`

## 🚀 Install

```sh
go get github.com/samber/mo@v1
```

This library is v1 and follows SemVer strictly.

No breaking changes will be made to exported APIs before v2.0.0.

## 💡 Quick start

You can import `mo` using:

```go
import (
    "github.com/samber/mo"
)
```

Then use one of the helpers below:

```go
option1 := mo.Some(42)
// Some(42)

option1.
    FlatMap(func (value int) Option[int] {
        return Some(value*2)
    }).
    FlatMap(func (value int) Option[int] {
        return Some(value%2)
    }).
    FlatMap(func (value int) Option[int] {
        return Some(value+21)
    }).
    OrElse(1234)
// 21

option2 := mo.None[int]()
// None

option2.OrElse(1234)
// 1234

option3 := option1.Match(
    func(i int) (int, bool) {
        // when value is present
        return i * 2, true
    },
    func() (int, bool) {
        // when value is absent
        return 0, false
    }
)
// Some(42)
```

More examples in [documentation](https://godoc.org/github.com/samber/mo).

## 🤠 Documentation and examples

[GoDoc: https://godoc.org/github.com/samber/mo](https://godoc.org/github.com/samber/mo)

### Option[T any]

`Option` is a container for an optional value of type `T`. If value exists, `Option` is of type `Some`. If the value is absent, `Option` is of type `None`.

Constructors:

- mo.Some()
- mo.None()
- mo.TupleToOption()

Methods:

- .IsPresent()
- .IsAbsent()
- .Size()
- .Get()
- .MustGet()
- .OrElse()
- .OrEmpty()
- .ForEach()
- .Match()
- .Map()
- .MapNone()
- .FlatMap()

### Result[T any]

`Result` respresent a result of an action having one of the following output: success or failure. An instance of `Result` is an instance of either `Ok` or `Err`. It could be compared to `Either[error, T]`.

Constructors:

- mo.Ok()
- mo.Err()
- mo.TupleToResult()

Methods:

- .IsOk()
- .IsError()
- .Error()
- .Get()
- .MustGet()
- .OrElse()
- .OrEmpty()
- .ToEither()
- .ForEach()
- .Match()
- .Map()
- .MapErr()
- .FlatMap()

### Either[L any, R any]

`Either` respresents a value of 2 possible types. An instance of `Either` is an instance of either `A` or `B`.

Constructors:

- mo.Left()
- mo.Right()

Methods:

- .IsLeft()
- .IsRight()
- .Left()
- .Right()
- .MustLeft()
- .MustRight()
- .LeftOrElse()
- .RightOrElse()
- .LeftOrEmpty()
- .RightOrEmpty()
- .Swap()
- .ForEach()
- .Match()
- .MapLeft()
- .MapRight()

### Future[T any]

`Future` represents a value which may or may not currently be available, but will be available at some point, or an exception if that value could not be made available.

Constructors:

- mo.NewFuture()

Methods:

- .Then()
- .Catch()
- .Finally()
- .Collect()
- .Result()
- .Cancel()

### IO[T any]

`IO` represents a non-deterministic synchronous computation that can cause side effects, yields a value of type `R` and never fails.

Constructors:

- mo.NewIO()
- mo.NewIO1()
- mo.NewIO2()
- mo.NewIO3()
- mo.NewIO4()
- mo.NewIO5()

Methods:

- .Run()

### IOEither[T any]

`IO` represents a non-deterministic synchronous computation that can cause side effects, yields a value of type `R` and can fail.

Constructors:

- mo.NewIOEither()
- mo.NewIOEither1()
- mo.NewIOEither2()
- mo.NewIOEither3()
- mo.NewIOEither4()
- mo.NewIOEither5()

Methods:

- .Run()

### Task[T any]

`Task` represents a non-deterministic asynchronous computation that can cause side effects, yields a value of type `R` and never fails.

Constructors:

- mo.NewTask()
- mo.NewTask1()
- mo.NewTask2()
- mo.NewTask3()
- mo.NewTask4()
- mo.NewTask5()
- mo.NewTaskFromIO()
- mo.NewTaskFromIO1()
- mo.NewTaskFromIO2()
- mo.NewTaskFromIO3()
- mo.NewTaskFromIO4()
- mo.NewTaskFromIO5()

Methods:

- .Run()

### TaskEither[T any]

`TaskEither` represents a non-deterministic asynchronous computation that can cause side effects, yields a value of type `R` and can fail.

Constructors:

- mo.NewTaskEither()
- mo.NewTaskEitherFromIOEither()

Methods:

- .Run()
- .OrElse()
- .Match()
- .TryCatch()
- .ToTask()
- .ToEither()

### State[S any, A any]

`State` represents a function `(S) -> (A, S)`, where `S` is state, `A` is result.

Constructors:

- mo.NewState()
- mo.ReturnState()

Methods:

- .Run()
- .Get()
- .Modify()
- .Put()

## 🛩 Benchmark

// @TODO

This library does not use `reflect` package. We don't expect overhead.

## 🤝 Contributing

- Ping me on twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/samber/mo)
- Fix [open issues](https://github.com/samber/mo/issues) or request new features

Don't hesitate ;)

### With Docker

```bash
docker-compose run --rm dev
```

### Without Docker

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Authors

- Samuel Berthe

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![support us](https://c5.patreon.com/external/logo/become_a_patron_button.png)](https://www.patreon.com/samber)

## 📝 License

Copyright © 2022 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
