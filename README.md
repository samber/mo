# mo - Monads

[![tag](https://img.shields.io/github/tag/samber/mo.svg)](https://github.com/samber/mo/releases)
[![GoDoc](https://godoc.org/github.com/samber/mo?status.svg)](https://pkg.go.dev/github.com/samber/mo)
![Build Status](https://github.com/samber/mo/actions/workflows/go.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/mo)](https://goreportcard.com/report/github.com/samber/mo)
[![codecov](https://codecov.io/gh/samber/mo/branch/master/graph/badge.svg)](https://codecov.io/gh/samber/mo)

🦄 **`samber/mo` brings monads and popular FP abstractions to Go projects. `samber/mo` uses the recent Go 1.18+ Generics.**

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
- `EitherX[T1, ..., TX]` (With X between 3 and 5) 
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
    },
)
// Some(42)
```

More examples in [documentation](https://godoc.org/github.com/samber/mo).

## 🤠 Documentation and examples

[GoDoc: https://godoc.org/github.com/samber/mo](https://godoc.org/github.com/samber/mo)

### Option[T any]

`Option` is a container for an optional value of type `T`. If value exists, `Option` is of type `Some`. If the value is absent, `Option` is of type `None`.

Constructors:

- `mo.Some()` [doc](https://pkg.go.dev/github.com/samber/mo#Some)
- `mo.None()` [doc](https://pkg.go.dev/github.com/samber/mo#None)
- `mo.TupleToOption()` [doc](https://pkg.go.dev/github.com/samber/mo#TupleToOption)

Methods:

- `.IsPresent()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.IsPresent)
- `.IsAbsent()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.IsAbsent)
- `.Size()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.Size)
- `.Get()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.Get)
- `.MustGet()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.MustGet)
- `.OrElse()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.OrElse)
- `.OrEmpty()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.OrEmpty)
- `.ForEach()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.ForEach)
- `.Match()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.Match)
- `.Map()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.Map)
- `.MapNone()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.MapNone)
- `.FlatMap()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.FlatMap)
- `.MarshalJSON()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.MarshalJSON)
- `.UnmarshalJSON()` [doc](https://pkg.go.dev/github.com/samber/mo#Option.UnmarshalJSON)

### Result[T any]

`Result` respresent a result of an action having one of the following output: success or failure. An instance of `Result` is an instance of either `Ok` or `Err`. It could be compared to `Either[error, T]`.

Constructors:

- `mo.Ok()` [doc](https://pkg.go.dev/github.com/samber/mo#Ok)
- `mo.Err()` [doc](https://pkg.go.dev/github.com/samber/mo#Err)
- `mo.TupleToResult()` [doc](https://pkg.go.dev/github.com/samber/mo#TupleToResult)
- `mo.Try()` [doc](https://pkg.go.dev/github.com/samber/mo#Try)

Methods:

- `.IsOk()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.IsOk)
- `.IsError()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.IsError)
- `.Error()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.Error)
- `.Get()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.Get)
- `.MustGet()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.MustGet)
- `.OrElse()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.OrElse)
- `.OrEmpty()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.OrEmpty)
- `.ToEither()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.ToEither)
- `.ForEach()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.ForEach)
- `.Match()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.Match)
- `.Map()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.Map)
- `.MapErr()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.MapErr)
- `.FlatMap()` [doc](https://pkg.go.dev/github.com/samber/mo#Result.FlatMap)

### Either[L any, R any]

`Either` respresents a value of 2 possible types. An instance of `Either` is an instance of either `A` or `B`.

Constructors:

- `mo.Left()` [doc](https://pkg.go.dev/github.com/samber/mo#Left)
- `mo.Right()` [doc](https://pkg.go.dev/github.com/samber/mo#Right)

Methods:

- `.IsLeft()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.IsLeft)
- `.IsRight()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.IsRight)
- `.Left()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.Left)
- `.Right()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.Right)
- `.MustLeft()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.MustLeft)
- `.MustRight()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.MustRight)
- `.LeftOrElse()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.LeftOrElse)
- `.RightOrElse()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.RightOrElse)
- `.LeftOrEmpty()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.LeftOrEmpty)
- `.RightOrEmpty()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.RightOrEmpty)
- `.Swap()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.Swap)
- `.ForEach()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.ForEach)
- `.Match()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.Match)
- `.MapLeft()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.MapLeft)
- `.MapRight()` [doc](https://pkg.go.dev/github.com/samber/mo#Either.MapRight)

### EitherX[T1, ..., TX] (With X between 3 and 5)

`EitherX` respresents a value of X possible types. For example, an `Either3` value is either `T1`, `T2` or `T3`.

Constructors:

- `mo.NewEitherXArgY()` [doc](https://pkg.go.dev/github.com/samber/mo#NewEither5Arg1). Eg:
  - `mo.NewEither3Arg1[A, B, C](A)`
  - `mo.NewEither3Arg2[A, B, C](B)`
  - `mo.NewEither3Arg3[A, B, C](C)`
  - `mo.NewEither4Arg1[A, B, C, D](A)`
  - `mo.NewEither4Arg2[A, B, C, D](B)`
  - ...

Methods:

- `.IsArgX()` [doc](https://pkg.go.dev/github.com/samber/mo#Either5.IsArg1)
- `.ArgX()` [doc](https://pkg.go.dev/github.com/samber/mo#Either5.Arg1)
- `.MustArgX()` [doc](https://pkg.go.dev/github.com/samber/mo#Either5.MustArg1)
- `.ArgXOrElse()` [doc](https://pkg.go.dev/github.com/samber/mo#Either5.Arg1OrElse)
- `.ArgXOrEmpty()` [doc](https://pkg.go.dev/github.com/samber/mo#Either5.Arg1OrEmpty)
- `.ForEach()` [doc](https://pkg.go.dev/github.com/samber/mo#Either5.ForEach)
- `.Match()` [doc](https://pkg.go.dev/github.com/samber/mo#Either5.Match)
- `.MapArgX()` [doc](https://pkg.go.dev/github.com/samber/mo#Either5.MapArg1)

### Future[T any]

`Future` represents a value which may or may not currently be available, but will be available at some point, or an exception if that value could not be made available.

Constructors:

- `mo.NewFuture()` [doc](https://pkg.go.dev/github.com/samber/mo#NewFuture)

Methods:

- `.Then()` [doc](https://pkg.go.dev/github.com/samber/mo#Future.Then)
- `.Catch()` [doc](https://pkg.go.dev/github.com/samber/mo#Future.Catch)
- `.Finally()` [doc](https://pkg.go.dev/github.com/samber/mo#Future.Finally)
- `.Collect()` [doc](https://pkg.go.dev/github.com/samber/mo#Future.Collect)
- `.Result()` [doc](https://pkg.go.dev/github.com/samber/mo#Future.Result)
- `.Cancel()` [doc](https://pkg.go.dev/github.com/samber/mo#Future.Cancel)

### IO[T any]

`IO` represents a non-deterministic synchronous computation that can cause side effects, yields a value of type `R` and never fails.

Constructors:

- `mo.NewIO()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIO)
- `mo.NewIO1()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIO1)
- `mo.NewIO2()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIO2)
- `mo.NewIO3()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIO3)
- `mo.NewIO4()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIO4)
- `mo.NewIO5()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIO5)

Methods:

- `.Run()` [doc](https://pkg.go.dev/github.com/samber/mo#Future.Run)

### IOEither[T any]

`IO` represents a non-deterministic synchronous computation that can cause side effects, yields a value of type `R` and can fail.

Constructors:

- `mo.NewIOEither()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIOEither)
- `mo.NewIOEither1()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIOEither1)
- `mo.NewIOEither2()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIOEither2)
- `mo.NewIOEither3()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIOEither3)
- `mo.NewIOEither4()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIOEither4)
- `mo.NewIOEither5()` [doc](https://pkg.go.dev/github.com/samber/mo#NewIOEither5)

Methods:

- `.Run()` [doc](https://pkg.go.dev/github.com/samber/mo#IOEither.Run)

### Task[T any]

`Task` represents a non-deterministic asynchronous computation that can cause side effects, yields a value of type `R` and never fails.

Constructors:

- `mo.NewTask()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTask)
- `mo.NewTask1()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTask1)
- `mo.NewTask2()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTask2)
- `mo.NewTask3()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTask3)
- `mo.NewTask4()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTask4)
- `mo.NewTask5()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTask5)
- `mo.NewTaskFromIO()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTaskFromIO)
- `mo.NewTaskFromIO1()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTaskFromIO1)
- `mo.NewTaskFromIO2()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTaskFromIO2)
- `mo.NewTaskFromIO3()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTaskFromIO3)
- `mo.NewTaskFromIO4()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTaskFromIO4)
- `mo.NewTaskFromIO5()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTaskFromIO5)

Methods:

- `.Run()` [doc](https://pkg.go.dev/github.com/samber/mo#Task.Run)

### TaskEither[T any]

`TaskEither` represents a non-deterministic asynchronous computation that can cause side effects, yields a value of type `R` and can fail.

Constructors:

- `mo.NewTaskEither()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTaskEither)
- `mo.NewTaskEitherFromIOEither()` [doc](https://pkg.go.dev/github.com/samber/mo#NewTaskEitherFromIOEither)

Methods:

- `.Run()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.Run)
- `.OrElse()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.OrElse)
- `.Match()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.Match)
- `.TryCatch()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.TryCatch)
- `.ToTask()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.ToTask)
- `.ToEither()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.ToEither)

### State[S any, A any]

`State` represents a function `(S) -> (A, S)`, where `S` is state, `A` is result.

Constructors:

- `mo.NewState()` [doc](https://pkg.go.dev/github.com/samber/mo#NewState)
- `mo.ReturnState()` [doc](https://pkg.go.dev/github.com/samber/mo#ReturnState)

Methods:

- `.Run()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.Run)
- `.Get()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.Get)
- `.Modify()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.Modify)
- `.Put()` [doc](https://pkg.go.dev/github.com/samber/mo#TaskEither.Put)

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
