package decimal

import (
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

type decimal struct {
	val big.Int // v × 10ᵉ
	exp int
}

var ten = big.NewInt(10)

func (r *decimal) number() Number {
	var buf []byte
	buf = r.val.Append(buf, 10)

	exp := r.exp
	var zeros int
	for zeros = 0; zeros < len(buf) && buf[len(buf)-zeros-1] == '0'; zeros++ {
	}
	if exp != 0 || zeros > 2 {
		buf = buf[:len(buf)-zeros]
		exp += zeros
	}

	if exp != 0 {
		buf = append(buf, 'e')
		buf = strconv.AppendInt(buf, int64(exp), 10)
	}
	return Number(unsafe.String(&buf[0], len(buf)))
}

func (r *decimal) setNumber(n Number) bool {
	if !valid(n) {
		return false
	}

	if e := strings.IndexAny(string(n), "eE"); e >= 0 {
		i, err := strconv.Atoi(string(n[e+1:]))
		if err != nil {
			return false
		}
		n = n[:e]
		r.exp = i
	} else {
		r.exp = 0
	}

	num, frac, _ := strings.Cut(string(n), ".")
	if _, ok := r.val.SetString(num, 10); !ok {
		return false
	}

	if len(frac) > 0 {
		var tmp decimal
		tmp.exp = r.exp - len(frac)
		if _, ok := tmp.val.SetString(frac, 10); !ok {
			return false
		}
		if num[0] == '-' {
			r.sub(r, &tmp)
		} else {
			r.add(r, &tmp)
		}
	}

	return true
}

// Neg sets z to the product x*y.
func (z *decimal) neg(x *decimal) *decimal {
	z.val.Neg(&x.val)
	return z
}

// Mul sets z to the product x*y.
func (z *decimal) mul(x, y *decimal) *decimal {
	z.exp = x.exp + y.exp
	z.val.Mul(&x.val, &y.val)
	return z
}

// Add sets z to the sum x+y.
func (z *decimal) add(x, y *decimal) *decimal {
	rescale(x, y)
	z.exp = x.exp
	z.val.Add(&x.val, &y.val)
	return z
}

// Sub sets z to the difference x-y.
func (z *decimal) sub(x, y *decimal) *decimal {
	rescale(x, y)
	z.exp = x.exp
	z.val.Sub(&x.val, &y.val)
	return z
}

// Cmp compares x and y.
func (x *decimal) cmp(y *decimal) int {
	rescale(x, y)
	return x.val.Cmp(&y.val)
}

// Exp2 sets z to 2ⁿ.
func (z *decimal) exp2(n int) *decimal {
	if n >= 0 {
		z.exp = 0
		z.val.SetUint64(1)
		z.val.Lsh(&z.val, uint(n))
		return z
	}

	var tmp decimal // 5 × 10⁻¹
	tmp.val.SetUint64(5)
	tmp.exp = -1

	n = ^n // n = -n - 1
	z.val.Set(&tmp.val)
	z.exp = tmp.exp

	for {
		if n&1 != 0 {
			z.mul(z, &tmp)
		}
		n >>= 1
		if n == 0 {
			return z
		}
		tmp.mul(&tmp, &tmp)
	}
}

// Rescale arguments to have same exponent.
// This may modify one of its arguments,
// but that argument will retain its exact value,
// so this should still be safe.
func rescale(x, y *decimal) {
	exp := x.exp - y.exp
	var zero big.Int
	switch {
	case exp == 0:
		return
	case x.val.CmpAbs(&zero) == 0:
		x.exp = y.exp
		return
	case y.val.CmpAbs(&zero) == 0:
		y.exp = x.exp
		return
	}

	r := x
	if exp < 0 {
		r = y
		exp = -exp
	}

	var tmp big.Int
	tmp.Exp(ten, tmp.SetUint64(uint64(exp)), nil)
	r.val.Mul(&r.val, &tmp) // v × 10ᵉ
	r.exp -= exp
}
