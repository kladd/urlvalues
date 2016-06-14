package urlvalues

import "testing"

type S1 struct {
	F01 int     `url:"f01" alttag:"t01"`
	F02 int     `url:"-" alttag:"-"`
	F03 string  `url:"f03" alttag:"t03"`
	F04 string  `url:"f04,omitempty" alttag:"t04,omitempty"`
	F05 bool    `url:"f05" alttag:"t05"`
	F06 bool    `url:"f06" alttag:"t06"`
	F07 *string `url:"f07" alttag:"t07"`
	F08 *int8   `url:"f08" alttag:"t08"`
	F09 float64 `url:"f09" alttag:"t09"`
	F10 S2
}

type S2 struct {
	F01 int `url:"f201" alttag:"t201"`
}

func testFilled(t *testing.T, useAltTag bool) {
	f07 := "seven"
	var f08 int8 = 8
	s2 := S2{10}
	s := &S1{
		F01: 1,
		F02: 2,
		F03: "three",
		F04: "four",
		F05: true,
		F06: false,
		F07: &f07,
		F08: &f08,
		F09: 1.618,
		F10: s2,
	}

	vals := make(map[string][]string)
	encoder := NewEncoder()
	tagPrefix := "f"

	if useAltTag {
		encoder.SetAliasTag("alttag")
		tagPrefix = "t"
	}
	_ = encoder.Encode(s, vals)

	valExists(t, tagPrefix+"01", "1", vals)
	valNotExists(t, tagPrefix+"02", vals)
	valExists(t, tagPrefix+"03", "three", vals)
	valExists(t, tagPrefix+"05", "1", vals)
	valExists(t, tagPrefix+"06", "0", vals)
	valExists(t, tagPrefix+"07", "seven", vals)
	valExists(t, tagPrefix+"08", "8", vals)
	valExists(t, tagPrefix+"09", "1.618000", vals)
	valExists(t, tagPrefix+"201", "10", vals)
}

func TestFilledWithDefaultTag(t *testing.T) {
	testFilled(t, false)
}

func TestFilledWithCustomTag(t *testing.T) {
	testFilled(t, true)
}

func testEmpty(t *testing.T, useAltTag bool) {
	s := &S1{
		F01: 1,
		F02: 2,
		F03: "three",
	}

	vals := make(map[string][]string)
	encoder := NewEncoder()
	tagPrefix := "f"

	if useAltTag {
		encoder.SetAliasTag("alttag")
		tagPrefix = "t"
	}
	_ = encoder.Encode(s, vals)

	valExists(t, tagPrefix+"03", "three", vals)
	valNotExists(t, tagPrefix+"04", vals)
}

func TestEmptyWithDefaultTag(t *testing.T) {
	testEmpty(t, false)
}

func TestEmptyWithCustomTag(t *testing.T) {
	testEmpty(t, true)
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
