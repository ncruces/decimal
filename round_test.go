package decimal_test

import (
	"testing"

	"github.com/ncruces/decimal"
)

func TestTrunc(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		unit decimal.Number
		want decimal.Number
	}{
		{"0", "1", "0"},
		{"1", "1", "1"},
		{"-1", "1", "-1"},
		{"0.1", "1", "0"},
		{"0.5", "1", "0"},
		{"0.9", "1", "0"},
		{"1.1", "1", "1"},
		{"1.5", "1", "1"},
		{"1.9", "1", "1"},
		{"0.01", "1", "0"},
		{"0.05", "1", "0"},
		{"0.09", "1", "0"},
		{"1.01", "1", "1"},
		{"1.05", "1", "1"},
		{"1.09", "1", "1"},
		{"-0.1", "1", "0"},
		{"-0.5", "1", "0"},
		{"-0.9", "1", "0"},
		{"-1.1", "1", "-1"},
		{"-1.5", "1", "-1"},
		{"-1.9", "1", "-1"},
		{"-0.01", "1", "0"},
		{"-0.05", "1", "0"},
		{"-0.09", "1", "0"},
		{"-1.01", "1", "-1"},
		{"-1.05", "1", "-1"},
		{"-1.09", "1", "-1"},
		{"0.1", "0.5", "0"},
		{"0.5", "0.5", "0.5"},
		{"0.9", "0.5", "0.5"},
		{"1.1", "0.5", "1"},
		{"1.5", "0.5", "1.5"},
		{"1.9", "0.5", "1.5"},
		{"-0.1", "0.5", "0"},
		{"-0.5", "0.5", "-0.5"},
		{"-0.9", "0.5", "-0.5"},
		{"-1.1", "0.5", "-1"},
		{"-1.5", "0.5", "-1.5"},
		{"-1.9", "0.5", "-1.5"},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			if got := decimal.Trunc(tt.x, tt.unit); got != tt.want {
				t.Errorf("Trunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		unit decimal.Number
		want decimal.Number
	}{
		{"0", "1", "0"},
		{"1", "1", "1"},
		{"-1", "1", "-1"},
		{"0.1", "1", "0"},
		{"0.5", "1", "0"},
		{"0.9", "1", "0"},
		{"1.1", "1", "1"},
		{"1.5", "1", "1"},
		{"1.9", "1", "1"},
		{"0.01", "1", "0"},
		{"0.05", "1", "0"},
		{"0.09", "1", "0"},
		{"1.01", "1", "1"},
		{"1.05", "1", "1"},
		{"1.09", "1", "1"},
		{"-0.1", "1", "-1"},
		{"-0.5", "1", "-1"},
		{"-0.9", "1", "-1"},
		{"-1.1", "1", "-2"},
		{"-1.5", "1", "-2"},
		{"-1.9", "1", "-2"},
		{"-0.01", "1", "-1"},
		{"-0.05", "1", "-1"},
		{"-0.09", "1", "-1"},
		{"-1.01", "1", "-2"},
		{"-1.05", "1", "-2"},
		{"-1.09", "1", "-2"},
		{"0.1", "0.5", "0"},
		{"0.5", "0.5", "0.5"},
		{"0.9", "0.5", "0.5"},
		{"1.1", "0.5", "1"},
		{"1.5", "0.5", "1.5"},
		{"1.9", "0.5", "1.5"},
		{"-0.1", "0.5", "-0.5"},
		{"-0.5", "0.5", "-0.5"},
		{"-0.9", "0.5", "-1"},
		{"-1.1", "0.5", "-1.5"},
		{"-1.5", "0.5", "-1.5"},
		{"-1.9", "0.5", "-2"},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			if got := decimal.Floor(tt.x, tt.unit); got != tt.want {
				t.Errorf("Floor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		unit decimal.Number
		want decimal.Number
	}{
		{"0", "1", "0"},
		{"1", "1", "1"},
		{"-1", "1", "-1"},
		{"0.1", "1", "1"},
		{"0.5", "1", "1"},
		{"0.9", "1", "1"},
		{"1.1", "1", "2"},
		{"1.5", "1", "2"},
		{"1.9", "1", "2"},
		{"0.01", "1", "1"},
		{"0.05", "1", "1"},
		{"0.09", "1", "1"},
		{"1.01", "1", "2"},
		{"1.05", "1", "2"},
		{"1.09", "1", "2"},
		{"-0.1", "1", "0"},
		{"-0.5", "1", "0"},
		{"-0.9", "1", "0"},
		{"-1.1", "1", "-1"},
		{"-1.5", "1", "-1"},
		{"-1.9", "1", "-1"},
		{"-0.01", "1", "0"},
		{"-0.05", "1", "0"},
		{"-0.09", "1", "0"},
		{"-1.01", "1", "-1"},
		{"-1.05", "1", "-1"},
		{"-1.09", "1", "-1"},
		{"0.1", "0.5", "0.5"},
		{"0.5", "0.5", "0.5"},
		{"0.9", "0.5", "1"},
		{"1.1", "0.5", "1.5"},
		{"1.5", "0.5", "1.5"},
		{"1.9", "0.5", "2"},
		{"-0.1", "0.5", "0"},
		{"-0.5", "0.5", "-0.5"},
		{"-0.9", "0.5", "-0.5"},
		{"-1.1", "0.5", "-1"},
		{"-1.5", "0.5", "-1.5"},
		{"-1.9", "0.5", "-1.5"},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			if got := decimal.Ceil(tt.x, tt.unit); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		unit decimal.Number
		want decimal.Number
	}{
		{"0", "1", "0"},
		{"1", "1", "1"},
		{"-1", "1", "-1"},
		{"0.1", "1", "0"},
		{"0.5", "1", "1"},
		{"0.9", "1", "1"},
		{"1.1", "1", "1"},
		{"1.5", "1", "2"},
		{"1.9", "1", "2"},
		{"0.01", "1", "0"},
		{"0.05", "1", "0"},
		{"0.09", "1", "0"},
		{"1.01", "1", "1"},
		{"1.05", "1", "1"},
		{"1.09", "1", "1"},
		{"-0.1", "1", "0"},
		{"-0.5", "1", "-1"},
		{"-0.9", "1", "-1"},
		{"-1.1", "1", "-1"},
		{"-1.5", "1", "-2"},
		{"-1.9", "1", "-2"},
		{"-0.01", "1", "0"},
		{"-0.05", "1", "0"},
		{"-0.09", "1", "0"},
		{"-1.01", "1", "-1"},
		{"-1.05", "1", "-1"},
		{"-1.09", "1", "-1"},
		{"0.1", "0.5", "0"},
		{"0.5", "0.5", "0.5"},
		{"0.9", "0.5", "1"},
		{"1.1", "0.5", "1"},
		{"1.5", "0.5", "1.5"},
		{"1.9", "0.5", "2"},
		{"-0.1", "0.5", "0"},
		{"-0.5", "0.5", "-0.5"},
		{"-0.9", "0.5", "-1"},
		{"-1.1", "0.5", "-1"},
		{"-1.5", "0.5", "-1.5"},
		{"-1.9", "0.5", "-2"},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			if got := decimal.Round(tt.x, tt.unit); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundToEven(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		unit decimal.Number
		want decimal.Number
	}{
		{"0", "1", "0"},
		{"1", "1", "1"},
		{"-1", "1", "-1"},
		{"0.1", "1", "0"},
		{"0.5", "1", "0"},
		{"0.9", "1", "1"},
		{"1.1", "1", "1"},
		{"1.5", "1", "2"},
		{"1.9", "1", "2"},
		{"0.01", "1", "0"},
		{"0.05", "1", "0"},
		{"0.09", "1", "0"},
		{"1.01", "1", "1"},
		{"1.05", "1", "1"},
		{"1.09", "1", "1"},
		{"-0.1", "1", "0"},
		{"-0.5", "1", "0"},
		{"-0.9", "1", "-1"},
		{"-1.1", "1", "-1"},
		{"-1.5", "1", "-2"},
		{"-1.9", "1", "-2"},
		{"-0.01", "1", "0"},
		{"-0.05", "1", "0"},
		{"-0.09", "1", "0"},
		{"-1.01", "1", "-1"},
		{"-1.05", "1", "-1"},
		{"-1.09", "1", "-1"},
		{"0.1", "0.5", "0"},
		{"0.5", "0.5", "0.5"},
		{"0.9", "0.5", "1"},
		{"1.1", "0.5", "1"},
		{"1.5", "0.5", "1.5"},
		{"1.9", "0.5", "2"},
		{"-0.1", "0.5", "0"},
		{"-0.5", "0.5", "-0.5"},
		{"-0.9", "0.5", "-1"},
		{"-1.1", "0.5", "-1"},
		{"-1.5", "0.5", "-1.5"},
		{"-1.9", "0.5", "-2"},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			if got := decimal.RoundToEven(tt.x, tt.unit); got != tt.want {
				t.Errorf("RoundToEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
