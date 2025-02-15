package decimal

import (
	"encoding/json"
	"regexp"
	"testing"
)

func Fuzz_valid_regex(f *testing.F) {
	re := regexp.MustCompile(`^-?(?:0|[1-9]\d*)(?:\.\d+)?(?:[eE][+-]?\d+)?$`)

	f.Add("0")
	f.Add("-1")
	f.Add("1.0")
	f.Add("1e-1")
	f.Add("1E+1")
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

func Fuzz_valid_json(f *testing.F) {
	f.Add("0")
	f.Add("-1")
	f.Add("1.0")
	f.Add("1e-1")
	f.Add("1E+1")
	f.Add("0.E0")
	f.Add("0.0.0")
	f.Add("+Inf")
	f.Add("null")
	f.Add("NaN")
	f.Add(`"0"`)
	f.Add(`""`)
	f.Add("")

	f.Fuzz(func(t *testing.T, str string) {
		var num Number
		err := json.Unmarshal([]byte(str), &num)

		if err != nil {
			num = Number(str)
		}
		if num == "" {
			return
		}
		if ok := valid(num); ok != (err == nil) {
			t.Fatalf("%q: num=%v ok=%v err=%v", str, num, ok, err)
		}
	})
}
