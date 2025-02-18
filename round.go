package decimal

import "math/big"

func Trunc(x Number) Number {
	checkValid(x)
	var rx big.Rat
	rx.SetString(string(x))
	if rx.IsInt() {
		return x
	}

	ix := rx.Num()
	ix.Quo(ix, rx.Denom())
	return Number(ix.String())
}

func Floor(x Number) Number {
	checkValid(x)
	var rx big.Rat
	rx.SetString(string(x))
	if rx.IsInt() {
		return x
	}

	ix := rx.Num()
	ix.Div(ix, rx.Denom())
	return Number(ix.String())
}

func Ceil(x Number) Number {
	checkValid(x)
	var rx big.Rat
	rx.SetString(string(x))
	if rx.IsInt() {
		return x
	}

	ix := rx.Num()
	ix.Div(ix.Neg(ix), rx.Denom())
	return Number(ix.Neg(ix).String())
}

func Round(x Number) Number {
	return round(x, false)
}

func RoundToEven(x Number) Number {
	return round(x, true)
}

func round(x Number, toEven bool) Number {
	checkValid(x)
	var rx big.Rat
	rx.SetString(string(x))
	if rx.IsInt() {
		return x
	}

	ix := rx.Num()
	dx := rx.Denom()
	neg := ix.Sign() < 0
	ix.Abs(ix)

	var mx big.Int
	ix.QuoRem(ix, dx, &mx)
	if cmp := mx.Lsh(&mx, 1).CmpAbs(dx); cmp > 0 ||
		cmp == 0 && (!toEven || ix.Bit(0) != 0) {
		ix.Add(ix, mx.SetUint64(1))
	}
	if neg {
		ix.Neg(ix)
	}
	return Number(ix.String())
}
