# Decimal arithmetic

[![Go Reference](https://pkg.go.dev/badge/image)](https://pkg.go.dev/github.com/ncruces/decimal)
[![Go Report](https://goreportcard.com/badge/github.com/ncruces/decimal)](https://goreportcard.com/report/github.com/ncruces/decimal)
[![Go Coverage](https://github.com/ncruces/decimal/wiki/coverage.svg)](https://raw.githack.com/wiki/ncruces/decimal/coverage.html)

Arithmetic that works on a type alias of [`json.Number`](https://pkg.go.dev/encoding/json#Number),
inspired by [SQLite](https://www.sqlite.org/floatingpoint.html#the_decimal_c_extension).

Since `Number` is a string, literals are valid numbers, and you can do:

```go
decimal.Sum("0.1", "0.1", "0.1", "-0.3") // == "0"
```

All operations produce exact results, with arbitrary precision
(which is why we can't have `decimal.Div`).

Instead, we have `decimal.Split` (and `decimal.Allocate`)
to distribute an amount amongst parties without loosing units to rounding.

```go
decimal.Split("0.99", "0.01", 2) // == ["0.5", "0.49"]
decimal.Allocate("100", "0.05", 2, 3, 2) // == ["28.6", "48.85", "28.55"]
```

We also have operations to round amounts to multiples of a unit
under various rounding modes.

The library won't win any performance prizes,
but you may find the API more ergonomic than alternatives.