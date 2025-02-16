package decimal

import (
	"math/big"
	"strconv"
)

type rat struct{ big.Rat }

func (z *rat) toNumber() Number {
	n, exact := z.FloatPrec()
	if !exact {
		panic(strconv.ErrRange)
	}
	return Number(z.FloatString(n))
}

func (z *rat) fromNumber(a Number) (ok bool) {
	if valid(a) {
		_, ok = z.SetString(string(a))
	}
	return ok
}

func (z *rat) add(x, y *rat) *rat {
	z.Add(&x.Rat, &y.Rat)
	return z
}

func (z *rat) sub(x, y *rat) *rat {
	z.Sub(&x.Rat, &y.Rat)
	return z
}

func (z *rat) mul(x, y *rat) *rat {
	z.Mul(&x.Rat, &y.Rat)
	return z
}

func (x *rat) cmp(y *rat) int {
	return x.Cmp(&y.Rat)
}
