package decimal

import "strings"

// Count sigificant digits in a valid JSON number.
func digits(x Number) int {
	if i := strings.IndexAny(string(x), "eE"); i >= 0 {
		x = x[:i]
	}
	s := strings.Trim(string(x), "-.0")
	if strings.Contains(s, ".") {
		return len(s) - 1
	}
	return len(s)
}
