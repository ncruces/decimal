package decimal

import "strconv"

func integerDigits(x string) int {
	exp := 0
	dot := -1
	lead := -1

loop:
	for i, b := range []byte(x) {
		switch b {
		case '.':
			if dot < 0 {
				dot = i
			}
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if lead < 0 {
				lead = i
			}
		case 'e', 'E':
			exp, _ = strconv.Atoi(x[i+1:])
			x = x[:i]
			break loop
		}
	}

	switch {
	case lead < 0:
		return 0
	case dot < 0:
		dot = len(x)
	case dot < lead:
		dot++
	}

	return exp + dot - lead
}

func significantDigits(x string) int {
	zeros, digits := 0, 0
	for _, b := range []byte(x) {
		switch b {
		case '0':
			if digits > 0 {
				zeros++
			}
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			digits += zeros + 1
			zeros = 0
		case 'e', 'E':
			return digits
		}
	}
	return digits
}
