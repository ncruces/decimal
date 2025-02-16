package decimal

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Number = json.Number

func Int64(i int64) Number {
	return Number(strconv.FormatInt(i, 10))
}

func Float64(f float64) Number {
	var tmp rat
	if tmp.SetFloat64(f) == nil {
		switch {
		case f != f:
			return "NaN"
		case f > 0:
			return "+Inf"
		case f < 0:
			return "-Inf"
		}
	}
	return tmp.toNumber()
}

func Neg(a Number) Number {
	if !valid(a) {
		return "Neg(" + a + ")"
	}
	if len(a) > 1 && a[0] == '-' {
		return a[1:]
	}
	return "-" + a
}

func Add(a, b Number) Number {
	var ra, rb rat
	if ra.fromNumber(a) && rb.fromNumber(b) {
		return ra.add(&ra, &rb).toNumber()
	}
	return "Add(" + a + "," + b + ")"
}

func Sub(a, b Number) Number {
	var ra, rb rat
	if ra.fromNumber(a) && rb.fromNumber(b) {
		return ra.sub(&ra, &rb).toNumber()
	}
	return "Sub(" + a + "," + b + ")"
}

func Mul(a, b Number) Number {
	var ra, rb rat
	if ra.fromNumber(a) && rb.fromNumber(b) {
		return ra.mul(&ra, &rb).toNumber()
	}
	return "Mul(" + a + "," + b + ")"
}

func Cmp(a, b Number) int {
	var ra, rb rat
	if ra.fromNumber(a) && rb.fromNumber(b) {
		return ra.cmp(&rb)
	}
	return strings.Compare(a.String(), b.String())
}

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
