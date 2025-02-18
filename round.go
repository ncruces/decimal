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
	checkValid(x)
	var rx, rm big.Rat
	rx.SetString(string(x))
	if rx.IsInt() {
		return x
	}

	ix := rx.Num()
	neg := ix.Sign() < 0
	if neg {
		ix.Neg(ix)
	}

	rm.SetInt(rx.Denom())
	ix.QuoRem(ix, rx.Denom(), rm.Denom())

	if rm.Cmp(big.NewRat(2, 1)) <= 0 {
		ix.Add(ix, big.NewInt(1))
	}
	if neg {
		ix.Neg(ix)
	}
	return Number(ix.String())
}

// func RoundToEven(x Number) Number
