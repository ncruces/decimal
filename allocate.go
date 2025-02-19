package decimal

import (
	"math/big"
)

// Split splits an integer amount of units over n parties.
// Leftover units are distributed round-robin,
// from left to right.
func Split(amount, unit Number, n uint) []Number {
	return allocate(amount, unit, n, nil)
}

// Allocate allocates an integer amount of units
// according to a list of ratios.
// Leftover units are distributed round-robin,
// from left to right.
func Allocate(amount, unit Number, ratios ...uint) []Number {
	return allocate(amount, unit, 0, ratios)
}

func allocate(amount, unit Number, n uint, ratios []uint) []Number {
	var ra, ru big.Rat
	fromNumber(&ra, amount)

	if unit != "1" {
		fromNumber(&ru, unit)
		if ru.Sign() <= 0 {
			panic("nonpositive unit")
		}
		ra.Quo(&ra, &ru)
	}
	if !ra.IsInt() {
		panic("noninteger amount")
	}

	var res []big.Int
	var sum, mod big.Int

	if ratios == nil {
		sum.SetUint64(uint64(n))
		mod.Set(ra.Num())

		res = make([]big.Int, n)
		for i := range n {
			res[i].Div(ra.Num(), &sum)
			mod.Sub(&mod, &res[i])
		}
	} else {
		for _, r := range ratios {
			sum.Add(&sum, mod.SetUint64(uint64(r)))
		}
		mod.Set(ra.Num())

		res = make([]big.Int, len(ratios))
		for i, r := range ratios {
			res[i].SetUint64(uint64(r))
			res[i].Mul(&res[i], ra.Num())
			res[i].Div(&res[i], &sum)
			mod.Sub(&mod, &res[i])
		}
	}

	sum.SetUint64(1)
	for i := range mod.Uint64() {
		res[i].Add(&res[i], &sum)
	}

	num := make([]Number, len(res))
	if unit == "1" {
		for i := range num {
			num[i] = Number(res[i].String())
		}
	} else {
		for i := range num {
			ra.SetInt(&res[i])
			num[i] = toNumber(ra.Mul(&ra, &ru))
		}
	}
	return num
}
