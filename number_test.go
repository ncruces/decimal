package decimal_test

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/ncruces/decimal"
)

func TestInt64(t *testing.T) {
	tests := []struct {
		i64  int64
		want decimal.Number
	}{
		{0, "0"},
		{1, "1"},
		{-1, "-1"},
		{math.MaxInt64, "9223372036854775807"},
		{math.MinInt64, "-9223372036854775808"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Int64(tt.i64); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		f64  float64
		want decimal.Number
	}{
		{0, "0"},
		{1, "1"},
		{-0, "0"},
		{-1, "-1"},
		{47.49, "47.49000000000000198951966012828052043914794921875"},
		{math.E, "2.718281828459045090795598298427648842334747314453125"},
		{math.Pi, "3.141592653589793115997963468544185161590576171875"},
		{math.Phi, "1.6180339887498949025257388711906969547271728515625"},
		{math.Sqrt2, "1.4142135623730951454746218587388284504413604736328125"},
		{math.SqrtE, "1.64872127070012819416433558217249810695648193359375"},
		{math.SqrtPi, "1.7724538509055161039640324815991334617137908935546875"},
		{math.SqrtPhi, "1.272019649514068984075265689170919358730316162109375"},
		{math.Ln2, "0.69314718055994528622676398299518041312694549560546875"},
		{math.Ln10, "2.30258509299404590109361379290930926799774169921875"},
		{math.Log2E, "1.442695040888963387004650940070860087871551513671875"},
		{math.Log10E, "0.43429448190325181666793241674895398318767547607421875"},
		{math.MaxFloat32, "340282346638528859811704183484516925440"},
		{math.SmallestNonzeroFloat32, "0.00000000000000000000000000000000000000000000140129846432481707092372958328991613128026194187651577175706828388979108268586060148663818836212158203125"},
		{math.MaxFloat64, "179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368"},
		{math.SmallestNonzeroFloat64, "0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004940656458412465441765687928682213723650598026143247644255856825006755072702087518652998363616359923797965646954457177309266567103559397963987747960107818781263007131903114045278458171678489821036887186360569987307230500063874091535649843873124733972731696151400317153853980741262385655911710266585566867681870395603106249319452715914924553293054565444011274801297099995419319894090804165633245247571478690147267801593552386115501348035264934720193790268107107491703332226844753335720832431936092382893458368060106011506169809753078342277318329247904982524730776375927247874656084778203734469699533647017972677717585125660551199131504891101451037862738167250955837389733598993664809941164205702637090279242767544565229087538682506419718265533447265625"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Float64(tt.f64); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}

	defer func() { recover() }()
	decimal.Float64(math.NaN())
	t.Fatal("want panic")
}

func TestAbs(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		want decimal.Number
	}{
		{"0", "0"},
		{"1", "1"},
		{"-0", "0"},
		{"-1", "1"},
		{"9223372036854775807", "9223372036854775807"},
		{"-9223372036854775808", "9223372036854775808"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Abs(tt.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeg(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		want decimal.Number
	}{
		{"0", "-0"},
		{"-0", "0"},
		{"1", "-1"},
		{"-1", "1"},
		{"-9223372036854775808", "9223372036854775808"},
		{"9223372036854775807", "-9223372036854775807"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Neg(tt.x); got != tt.want {
				t.Errorf("Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		y    decimal.Number
		want decimal.Number
	}{
		{"0", "-0", "0"},
		{"-0", "0", "0"},
		{"0.1", "0.2", "0.3"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Add(tt.x, tt.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		y    decimal.Number
		want decimal.Number
	}{
		{"0", "-0", "0"},
		{"-0", "0", "0"},
		{"0.3", "0.2", "0.1"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Sub(tt.x, tt.y); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		y    decimal.Number
		want decimal.Number
	}{
		{"0", "-0", "0"},
		{"-0", "0", "0"},
		{"0.1", "3", "0.3"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Mul(tt.x, tt.y); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmp(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		y    decimal.Number
		want int
	}{
		{"0", "-0", 0},
		{"-0", "0", 0},
		{"0.1", "0.2", -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := decimal.Cmp(tt.x, tt.y); got != tt.want {
				t.Errorf("Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		want bool
	}{
		{"0", true},
		{"0.1", false},
		{"0.99999999999999999999999999999999", false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := decimal.IsInt(tt.x); got != tt.want {
				t.Errorf("IsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		n    []decimal.Number
		want decimal.Number
	}{
		{nil, "0"},
		{[]decimal.Number{"-0"}, "0"},
		{[]decimal.Number{"-1"}, "-1"},
		{[]decimal.Number{"0.1", "0.1", "0.1"}, "0.3"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Sum(tt.n...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProd(t *testing.T) {
	tests := []struct {
		n    []decimal.Number
		want decimal.Number
	}{
		{nil, "1"},
		{[]decimal.Number{"-0"}, "0"},
		{[]decimal.Number{"-1"}, "-1"},
		{[]decimal.Number{"0.1", "0.1", "0.1"}, "0.001"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Prod(tt.n...); got != tt.want {
				t.Errorf("Prod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		x    decimal.Number
		n    uint
		want decimal.Number
	}{
		{"0", 0, "1"},
		{"1", 0, "1"},
		{"1", 1, "1"},
		{"2", 1, "2"},
		{"1", 10, "1"},
		{"2", 10, "1024"},
		{"0.01", 10, "0.00000000000000000001"},
		{"0.02", 10, "0.00000000000000001024"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := decimal.Pow(tt.x, tt.n); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFmt(t *testing.T) {
	tests := []struct {
		fmt  string
		x    decimal.Number
		want string
	}{
		{"%g", "0.1", "0.1"},
		{"%g", "1e-100", "1e-100"},
		{"%f", "1e-1", "0.100000"},
		{"%f", "1e-10", "0.000000"},
		{"%.1e", "1e-1", "1.0e-01"},
		{"%e", "1e-1", "1.000000e-01"},
		{"%.0f", "1e150", "1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"%.100f", "0.1", "0.1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"%.100f", "1e-10", "0.0000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"%g", "3.14159265358979323846264338327950288419716939937510582097494459", "3.14159265358979323846264338327950288419716939937510582097494459"},
		{"%.5g", "3.14159265358979323846264338327950288419716939937510582097494459", "3.1416"},
		{"%d", "1.0", "%!d(decimal=1.0)"},
		{"%v", "NaN", "%!v(decimal=NaN)"},
	}
	for _, tt := range tests {
		t.Run(string(tt.want), func(t *testing.T) {
			if got := fmt.Sprintf(tt.fmt, decimal.Fmt(tt.x)); got != tt.want {
				t.Errorf("Fmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzSum(f *testing.F) {
	f.Add("")
	f.Add("0;0;0")
	f.Add("1;1;1")
	f.Add("10;10;10")
	f.Add("0.1;0.1;0.1")

	f.Fuzz(func(t *testing.T, str string) {
		v := strings.Split(str, ";")
		if len(v) == 0 || len(v) > 97 {
			return
		}

		sum := decimal.Number("0")
		all := make([]decimal.Number, len(v))
		for i, v := range v {
			d := decimal.Number(v)
			if !decimal.IsValid(d) {
				return
			}
			if f, err := strconv.ParseFloat(v, 64); f == 0 || err != nil {
				return
			}
			sum = decimal.Add(sum, d)
			all[i] = d
		}

		if got := decimal.Sum(all...); got != sum {
			t.Fatalf("Sum=%q Add=%q", got, sum)
		}
	})
}

func FuzzProd(f *testing.F) {
	f.Add("")
	f.Add("0;0;0")
	f.Add("1;1;1")
	f.Add("10;10;10")
	f.Add("0.1;0.1;0.1")

	f.Fuzz(func(t *testing.T, str string) {
		v := strings.Split(str, ";")
		if len(v) == 0 || len(v) > 97 {
			return
		}

		mul := decimal.Number("1")
		all := make([]decimal.Number, len(v))
		for i, v := range v {
			d := decimal.Number(v)
			if !decimal.IsValid(d) {
				return
			}
			if f, err := strconv.ParseFloat(v, 64); f == 0 || err != nil {
				return
			}
			mul = decimal.Mul(mul, d)
			all[i] = d
		}

		if got := decimal.Prod(all...); got != mul {
			t.Fatalf("Prod=%q Mul=%q", got, mul)
		}
	})
}
