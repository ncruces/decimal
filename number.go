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
	return toNumber(&rf)
}

// Abs returns |x| (the absolute value of x).
func Abs(x Number) Number {
	checkValid(x)
	if x[0] == '-' {
		return x[1:]
	}
	return x
}

// Neg returns -x (x with its sign negated).
func Neg(x Number) Number {
	checkValid(x)
	if x[0] == '-' {
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
	return toNumber(rx.Add(&rx, &ry))
}

// Sub returns the difference x - y.
func Sub(x, y Number) Number {
	checkValid(x)
	checkValid(y)
	var rx, ry big.Rat
	rx.SetString(string(x))
	ry.SetString(string(y))
	return toNumber(rx.Sub(&rx, &ry))
}

// Mul returns the product x * y.
func Mul(x, y Number) Number {
	checkValid(x)
	checkValid(y)
	var rx, ry big.Rat
	rx.SetString(string(x))
	ry.SetString(string(y))
	return toNumber(rx.Mul(&rx, &ry))
}

// Sum returns the sum of all n.
func Sum(n ...Number) Number {
	var rs, rn big.Rat
	for _, n := range n {
		checkValid(n)
		rn.SetString(string(n))
		rs.Add(&rs, &rn)
	}
	return toNumber(&rs)
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

// Fmt is a formatter for a decimal number.
type Fmt Number

// Format implements fmt.Formatter.
// It accepts the formats for decimal floating-point numbers: 'e', 'E', 'f', 'F', 'g', 'G'.
// The 'v' format is handled like 'g'.
func (x Fmt) Format(f fmt.State, v rune) {
	s := string(x)
	prec, ok := f.Precision()

	if !IsValid(Number(s)) {
		fmt.Fprintf(f, "%%!%c(decimal=%s)", v, s)
		return
	}

	switch v {
	default:
		fmt.Fprintf(f, "%%!%c(decimal=%s)", v, s)
		return

	case 'e', 'E':
		if !ok {
			prec = 6
		}
		prec += 1

	case 'f', 'F':
		if !ok {
			prec = 6
		}
		prec += integerDigits(s)

	case 'v', 'g', 'G':
		if !ok {
			prec = significantDigits(s)
		}
	}

	var fx big.Float
	prec = max(0, prec)
	// prec in digits, multiply by log₂(10) for bits
	fx.SetPrec((107*uint(prec) + 31) / 32)
	fx.Parse(s, 10)
	fx.Format(f, v)
}

func toNumber(x *big.Rat) Number {
	n, exact := x.FloatPrec()
	if exact {
		return Number(x.FloatString(n))
	}
	panic("inexact decimal")
}
