package urlvalues

import "testing"

type S1 struct {
	F01 int     `url:"f01"`
	F02 int     `url:"-"`
	F03 string  `url:"f03"`
	F04 string  `url:"f04,omitempty"`
	F05 bool    `url:"f05"`
	F06 bool    `url:"f06"`
	F07 *string `url:"f07"`
	F08 *int8   `url:"f08"`
}

func TestFilled(t *testing.T) {
	f07 := "seven"
	var f08 int8 = 8
	s := &S1{
		F01: 1,
		F02: 2,
		F03: "three",
		F04: "four",
		F05: true,
		F06: false,
		F07: &f07,
		F08: &f08,
	}

	vals := make(map[string][]string)
	_ = NewEncoder().Encode(s, vals)

	valExists(t, "f01", "1", vals)
	valNotExists(t, "f02", vals)
	valExists(t, "f03", "three", vals)
	valExists(t, "f05", "1", vals)
	valExists(t, "f06", "0", vals)
	valExists(t, "f07", "seven", vals)
	valExists(t, "f08", "8", vals)
}

func TestEmpty(t *testing.T) {
	s := &S1{
		F01: 1,
		F02: 2,
		F03: "three",
	}

	vals := make(map[string][]string)
	_ = NewEncoder().Encode(s, vals)

	valExists(t, "f03", "three", vals)
	valNotExists(t, "f04", vals)
}

func valExists(t *testing.T, key string, expect string, result map[string][]string) {
	if val, ok := result[key]; !ok {
		t.Error("Key not found. Expected: " + expect)
	} else if val[0] != expect {
		t.Error("Unexpected value. Expected: " + expect + "; got: " + val[0] + ".")
	}
}

func valNotExists(t *testing.T, key string, result map[string][]string) {
	if val, ok := result[key]; ok {
		t.Error("Key not ommited. Expected: empty; got: " + val[0] + ".")
	}
}
