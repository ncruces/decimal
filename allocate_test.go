package decimal_test

import (
	"reflect"
	"testing"

	"github.com/ncruces/decimal"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		x     decimal.Number
		unit  decimal.Number
		count uint
		want  []decimal.Number
	}{
		{
			"100", "1", 3,
			[]decimal.Number{"34", "33", "33"},
		},
		{
			"0.99", "0.01", 2,
			[]decimal.Number{"0.5", "0.49"},
		},
		{
			"-0.99", "0.01", 2,
			[]decimal.Number{"-0.49", "-0.5"},
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			got := decimal.Split(tt.x, tt.unit, tt.count)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocate(t *testing.T) {
	tests := []struct {
		x      decimal.Number
		unit   decimal.Number
		ratios []uint
		want   []decimal.Number
	}{
		{
			"100", "1", []uint{1, 2},
			[]decimal.Number{"34", "66"},
		},
		{
			"0.99", "0.01", []uint{1, 1},
			[]decimal.Number{"0.5", "0.49"},
		},
		{
			"-100", "1", []uint{1, 2, 3, 1},
			[]decimal.Number{"-14", "-28", "-43", "-15"},
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.x), func(t *testing.T) {
			got := decimal.Allocate(tt.x, tt.unit, tt.ratios...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Allocate() = %v, want %v", got, tt.want)
			}
		})
	}
}
