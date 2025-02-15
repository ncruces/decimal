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
	return floatNumber(f)
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
	var da, db decimal
	if da.setNumber(a) && db.setNumber(b) {
		return da.add(&da, &db).number()
	}
	return "Add(" + a + "," + b + ")"
}

func Sub(a, b Number) Number {
	var da, db decimal
	if da.setNumber(a) && db.setNumber(b) {
		return da.sub(&da, &db).number()
	}
	return "Sub(" + a + "," + b + ")"
}

func Mul(a, b Number) Number {
	var da, db decimal
	if da.setNumber(a) && db.setNumber(b) {
		return da.mul(&da, &db).number()
	}
	return "Mul(" + a + "," + b + ")"
}

func Cmp(a, b Number) int {
	var da, db decimal
	if da.setNumber(a) && db.setNumber(b) {
		return da.cmp(&db)
	}
	return strings.Compare(a.String(), b.String())
}

func Sum(n ...Number) Number {
	var da, db decimal
	for _, n := range n {
		if db.setNumber(n) {
			da.add(&da, &db)
		} else {
			return "Sum(...," + n + ",...)"
		}
	}
	return da.number()
}
