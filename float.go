package decimal

import (
	"math"
	"math/bits"
)

func floatNumber(f float64) Number {
	mant, exp2 := mantExp(f)
	if exp2 < 0 && exp2 == -exp2 {
		// not finite
		switch {
		case f != f:
			return "NaN"
		case f > 0:
			return "+Inf"
		case f < 0:
			return "-Inf"
		}
	}
	var a, b decimal
	a.exp = 0
	a.val.SetInt64(mant)
	return a.mul(&a, b.exp2(int(exp2))).number()
}

func mantExp(f float64) (int64, int) {
	n := math.Float64bits(f)

	mant := n << 12 >> 12
	exp2 := int(n>>52) & 0x7ff

	switch exp2 {
	case 0:
		if mant == 0 {
			return 0, 0 // zero
		}
		exp2++ // denormals
	case 0x7ff:
		// infinities, NaNs
		return 0, math.MinInt
	default:
		mant |= 1 << 52
	}

	if z := bits.TrailingZeros64(mant); z != 0 {
		mant >>= z
		exp2 += z
	}
	if int64(n) < 0 {
		mant = -mant
	}
	exp2 -= 1023 + 52
	return int64(mant), exp2
}
