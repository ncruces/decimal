package decimal

import (
	"regexp"
	"testing"
)

func Fuzz_valid(f *testing.F) {
	re := regexp.MustCompile(`^-?(?:0|[1-9]\d*)(?:\.\d+)?(?:[eE][+-]?\d+)?$`)

	f.Add("0")
	f.Add("-1")
	f.Add("0e0")
	f.Add("1.0")
	f.Add("1e-1")
	f.Add("1E+10")
	f.Add("1.0e0")
	f.Add("01")
	f.Add("1+")
	f.Add("0E.")
	f.Add("0E+-")
	f.Add("0E0-")
	f.Add("0.E0")
	f.Add("0.0.0")
	f.Add("+Inf")
	f.Add("null")
	f.Add("NaN")
	f.Add(`"0"`)
	f.Add(`""`)
	f.Add("")

	f.Fuzz(func(t *testing.T, str string) {
		if ok := valid(Number(str)); ok != re.MatchString(str) {
			t.Fatalf("%q: ok=%v", str, ok)
		}
	})
}
