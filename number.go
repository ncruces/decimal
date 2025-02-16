// Package decimal implements arbitrary-precision decimal arithmetic.
package decimal

import (
	"encoding/json"
	"strconv"
	"strings"
)

// A Number is an arbitrary precision decimal number, stored as JSON text.
type Number = json.Number

// Int64 converts i into a decimal number.
func Int64(i int64) Number {
	return Number(strconv.FormatInt(i, 10))
}

// Float64 converts f into a decimal number.
func Float64(f float64) Number {
	var rf rat
	if rf.SetFloat64(f) == nil {
		switch {
		case f != f:
			return "NaN"
		case f > 0:
			return "+Inf"
		case f < 0:
			return "-Inf"
		}
	}
	return rf.toNumber()
}

// Neg returns x with its sign negated.
func Neg(x Number) Number {
	if !valid(x) {
		return "Neg(" + x + ")"
	}
	if len(x) > 1 && x[0] == '-' {
		return x[1:]
	}
	return "-" + x
}

// Add returns the sum x + y.
func Add(x, y Number) Number {
	var rx, ry rat
	if rx.fromNumber(x) && ry.fromNumber(y) {
		return rx.add(&rx, &ry).toNumber()
	}
	return "Add(" + x + "," + y + ")"
}

// Sub returns the difference x - y.
func Sub(x, y Number) Number {
	var rx, ry rat
	if rx.fromNumber(x) && ry.fromNumber(y) {
		return rx.sub(&rx, &ry).toNumber()
	}
	return "Sub(" + x + "," + y + ")"
}

// Mul returns the product x * y.
func Mul(x, y Number) Number {
	var rx, ry rat
	if rx.fromNumber(x) && ry.fromNumber(y) {
		return rx.mul(&rx, &ry).toNumber()
	}
	return "Mul(" + x + "," + y + ")"
}

// Cmp compares a and b, like [cmp.Compare].
func Cmp(x, y Number) int {
	var rx, ry rat
	if rx.fromNumber(x) && ry.fromNumber(y) {
		return rx.cmp(&ry)
	}
	return strings.Compare(x.String(), y.String())
}

// Sum returns the sum of all n.
func Sum(n ...Number) Number {
	var rs, rn rat
	for _, n := range n {
		if rn.fromNumber(n) {
			rs.add(&rs, &rn)
		} else {
			return "Sum(...," + n + ",...)"
		}
	}
	return rs.toNumber()
}
