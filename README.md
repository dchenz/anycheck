# anycheck

Linter to detect inconsistent usage of `interface{}` and `any` in your Go code.


## Install

```sh
go get -u github.com/dchenz/anycheck/cmd/anycheck
```

## Motivation

In Go 1.18, the introduction of `any` as a type alias for `interface{}` allows for interchangeable usage between the two. However, this flexibility can lead to inconsistent coding styles when some parts of the codebase utilize either option.

```go
var a any
var b interface{}
```
