package decimal

import "testing"

func Test_integerDigits(t *testing.T) {
	// The result is either:
	//	- the number of digits to the left of the decimal point, or
	//	- the number of zeros immediately to the right (negated)
	tests := []struct {
		arg  string
		want int
	}{
		{"0", 0},
		{"1", 1},
		{"10", 2},
		{"100", 3},
		{"0.1", 0},
		{"0.01", -1},
		{"0e+10", 0},
		{"0e-10", 0},
		{"1e+10", 11},
		{"1e-10", -9},
		{"10e+10", 12},
		{"10e-10", -8},
		{"100e+10", 13},
		{"100e-10", -7},
		{"0.1e+10", 10},
		{"0.1e-10", -10},
		{"0.01e+10", 9},
		{"0.01e-10", -11},
	}
	for _, tt := range tests {
		t.Run(string(tt.arg), func(t *testing.T) {
			if got := integerDigits(tt.arg); got != tt.want {
				t.Errorf("integerDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_significantDigits(t *testing.T) {
	// The result is the number of digits between
	// the first and last non zero digit.

	tests := []struct {
		arg  string
		want int
	}{
		{"0", 0},
		{"1", 1},
		{"10", 1},
		{"11", 2},
		{"100", 1},
		{"101", 3},
		{"0.1", 1},
		{"1.1", 2},
		{"0.01", 1},
		{"1.01", 3},
	}
	for _, tt := range tests {
		t.Run(string(tt.arg), func(t *testing.T) {
			if got := significantDigits(tt.arg); got != tt.want {
				t.Errorf("significantDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
