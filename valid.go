package decimal

// IsValid reports whether x is a valid JSON number literal.
func IsValid(x Number) bool {
	var state byte
	for _, d := range []byte(x) {
		switch state {
		case 0:
			switch d {
			case '0', '-':
				state = d
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = '1'
			default:
				return false
			}

		case '-':
			switch d {
			case '0':
				state = '0'
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = '1'
			default:
				return false
			}

		case '0':
			switch d {
			case '.':
				state = '.'
			case 'e', 'E':
				state = 'e'
			default:
				return false
			}

		case '1':
			switch d {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				continue
			case '.':
				state = '.'
			case 'e', 'E':
				state = 'e'
			default:
				return false
			}

		case '.':
			switch d {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = '2'
			default:
				return false
			}

		case '2':
			switch d {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				continue
			case 'e', 'E':
				state = 'e'
			default:
				return false
			}

		case 'e':
			switch d {
			case '-', '+':
				state = '+'
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = '3'
			default:
				return false
			}

		case '+':
			switch d {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = '3'
			default:
				return false
			}

		case '3':
			switch d {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				continue
			default:
				return false
			}
		}
	}
	switch state {
	case '0', '1', '2', '3':
		return true
	default:
		return false
	}
}
