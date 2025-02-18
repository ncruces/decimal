package decimal

import (
	"testing"
)

func TestTrunc(t *testing.T) {
	tests := []struct {
		x    Number
		want Number
	}{
		{"0", "0"},
		{"1", "1"},
		{"-1", "-1"},
		{"0.1", "0"},
		{"0.5", "0"},
		{"0.9", "0"},
		{"1.1", "1"},
		{"1.5", "1"},
		{"1.9", "1"},
		{"0.01", "0"},
		{"0.05", "0"},
		{"0.09", "0"},
		{"1.01", "1"},
		{"1.05", "1"},
		{"1.09", "1"},
		{"-0.1", "0"},
		{"-0.5", "0"},
		{"-0.9", "0"},
		{"-1.1", "-1"},
		{"-1.5", "-1"},
		{"-1.9", "-1"},
		{"-0.01", "0"},
		{"-0.05", "0"},
		{"-0.09", "0"},
		{"-1.01", "-1"},
		{"-1.05", "-1"},
		{"-1.09", "-1"},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			if got := Trunc(tt.x); got != tt.want {
				t.Errorf("Trunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		x    Number
		want Number
	}{
		{"0", "0"},
		{"1", "1"},
		{"-1", "-1"},
		{"0.1", "0"},
		{"0.5", "0"},
		{"0.9", "0"},
		{"1.1", "1"},
		{"1.5", "1"},
		{"1.9", "1"},
		{"0.01", "0"},
		{"0.05", "0"},
		{"0.09", "0"},
		{"1.01", "1"},
		{"1.05", "1"},
		{"1.09", "1"},
		{"-0.1", "-1"},
		{"-0.5", "-1"},
		{"-0.9", "-1"},
		{"-1.1", "-2"},
		{"-1.5", "-2"},
		{"-1.9", "-2"},
		{"-0.01", "-1"},
		{"-0.05", "-1"},
		{"-0.09", "-1"},
		{"-1.01", "-2"},
		{"-1.05", "-2"},
		{"-1.09", "-2"},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			if got := Floor(tt.x); got != tt.want {
				t.Errorf("Floor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		x    Number
		want Number
	}{
		{"0", "0"},
		{"1", "1"},
		{"-1", "-1"},
		{"0.1", "1"},
		{"0.5", "1"},
		{"0.9", "1"},
		{"1.1", "2"},
		{"1.5", "2"},
		{"1.9", "2"},
		{"0.01", "1"},
		{"0.05", "1"},
		{"0.09", "1"},
		{"1.01", "2"},
		{"1.05", "2"},
		{"1.09", "2"},
		{"-0.1", "0"},
		{"-0.5", "0"},
		{"-0.9", "0"},
		{"-1.1", "-1"},
		{"-1.5", "-1"},
		{"-1.9", "-1"},
		{"-0.01", "0"},
		{"-0.05", "0"},
		{"-0.09", "0"},
		{"-1.01", "-1"},
		{"-1.05", "-1"},
		{"-1.09", "-1"},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			if got := Ceil(tt.x); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		x    Number
		want Number
	}{
		{"0", "0"},
		{"1", "1"},
		{"-1", "-1"},
		{"0.1", "0"},
		{"0.5", "1"},
		{"0.9", "1"},
		{"1.1", "1"},
		{"1.5", "2"},
		{"1.9", "2"},
		{"0.01", "0"},
		{"0.05", "0"},
		{"0.09", "0"},
		{"1.01", "1"},
		{"1.05", "1"},
		{"1.09", "1"},
		{"-0.1", "0"},
		{"-0.5", "-1"},
		{"-0.9", "-1"},
		{"-1.1", "-1"},
		{"-1.5", "-2"},
		{"-1.9", "-2"},
		{"-0.01", "0"},
		{"-0.05", "0"},
		{"-0.09", "0"},
		{"-1.01", "-1"},
		{"-1.05", "-1"},
		{"-1.09", "-1"},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			if got := Round(tt.x); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
