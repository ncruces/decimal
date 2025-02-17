// Package decimal implements arbitrary-precision decimal arithmetic.
package decimal

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

// A Number is an arbitrary precision decimal number, stored as JSON text.
type Number = json.Number

// Int64 converts i into a decimal number.
func Int64(i int64) Number {
	return Number(strconv.FormatInt(i, 10))
}

// Float64 converts f into a decimal number.
func Float64(f float64) Number {
	var rf big.Rat
	if rf.SetFloat64(f) == nil {
		switch {
		case f != f:
			panic("invalid decimal: NaN")
		case f > 0:
			panic("invalid decimal: +Inf")
		case f < 0:
			panic("invalid decimal: -Inf")
		}
	}
	return ratNumber(&rf)
}

// Neg returns x with its sign negated.
func Neg(x Number) Number {
	if !IsValid(x) {
		panic("invalid decimal: " + x)
	}
	if len(x) > 1 && x[0] == '-' {
		return x[1:]
	}
	return "-" + x
}

// Add returns the sum x + y.
func Add(x, y Number) Number {
	checkValid(x)
	checkValid(y)
	var rx, ry big.Rat
	rx.SetString(string(x))
	ry.SetString(string(y))
	return ratNumber(rx.Add(&rx, &ry))
}

// Sub returns the difference x - y.
func Sub(x, y Number) Number {
	checkValid(x)
	checkValid(y)
	var rx, ry big.Rat
	rx.SetString(string(x))
	ry.SetString(string(y))
	return ratNumber(rx.Sub(&rx, &ry))
}

// Mul returns the product x * y.
func Mul(x, y Number) Number {
	checkValid(x)
	checkValid(y)
	var rx, ry big.Rat
	rx.SetString(string(x))
	ry.SetString(string(y))
	return ratNumber(rx.Mul(&rx, &ry))
}

// Cmp compares x and y, like [cmp.Compare].
func Cmp(x, y Number) int {
	checkValid(x)
	checkValid(y)
	var rx, ry big.Rat
	rx.SetString(string(x))
	ry.SetString(string(y))
	return rx.Cmp(&ry)
}

// Sum returns the sum of all n.
func Sum(n ...Number) Number {
	var rs, rn big.Rat
	for _, n := range n {
		checkValid(n)
		rn.SetString(string(n))
		rs.Add(&rs, &rn)
	}
	return ratNumber(&rs)
}

// Fmt returns a formatter for x.
// The result will be accurate to at least
// 100 signifiant decimal digits more than
// the exact decimal representation of x.
func Fmt(x Number) fmt.Formatter {
	var fx big.Float
	if IsValid(x) {
		fx.SetPrec(333 + 107*uint(digits(x)/32))
	}
	f, _, _ := fx.Parse(string(x), 10)
	return f
}

func ratNumber(x *big.Rat) Number {
	n, exact := x.FloatPrec()
	if !exact {
		panic("inexact decimal")
	}
	return Number(x.FloatString(n))
}
