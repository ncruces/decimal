package decimal

import "math/big"

func Trunc(x, unit Number) Number {
	return scale(x, unit, func(rx *big.Rat) {
		nx := rx.Num()
		dx := rx.Denom()
		nx.Quo(nx, dx)
		dx.SetUint64(1)
	})
}

func Floor(x, unit Number) Number {
	return scale(x, unit, func(rx *big.Rat) {
		nx := rx.Num()
		dx := rx.Denom()
		nx.Div(nx, dx)
		dx.SetUint64(1)
	})
}

func Ceil(x, unit Number) Number {
	return scale(x, unit, func(rx *big.Rat) {
		nx := rx.Num()
		dx := rx.Denom()
		nx.Neg(nx)
		nx.Div(nx, dx)
		nx.Neg(nx)
		dx.SetUint64(1)
	})
}

func Round(x, unit Number) Number {
	return scale(x, unit, tiesAway.round)
}

func RoundToEven(x, unit Number) Number {
	return scale(x, unit, tiesToEven.round)
}

type rounder int

const (
	tiesAway rounder = iota
	tiesToEven
)

func (r rounder) round(rx *big.Rat) {
	nx := rx.Num()
	dx := rx.Denom()
	neg := nx.Sign() < 0

	nx.Abs(nx)
	var mx big.Int
	nx.QuoRem(nx, dx, &mx)
	cmp := mx.Lsh(&mx, 1).CmpAbs(dx)
	dx.SetUint64(1)

	if cmp > 0 || cmp == 0 && (r == tiesAway || nx.Bit(0) != 0) {
		nx.Add(nx, dx)
	}
	if neg {
		nx.Neg(nx)
	}
}

func scale(x, unit Number, round func(*big.Rat)) Number {
	var rx, ru big.Rat
	rx.SetString(checkValid(x))

	if unit != "1" {
		ru.SetString(checkValid(unit))
		rx.Quo(&rx, &ru)
	}
	if rx.IsInt() {
		return x
	}

	round(&rx)
	if unit != "1" {
		return toNumber(rx.Mul(&rx, &ru))
	}
	return Number(rx.Num().String())
}
