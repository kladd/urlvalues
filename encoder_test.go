package urlvalues

import "testing"

type S1 struct {
	F01 int    `url:"f01"`
	F02 int    `url:"-"`
	F03 string `url:"f03"`
	F04 string `url:"f04,omitempty"`
	F05 bool   `url:"f05"`
	F06 bool   `url:"f06"`
}

func TestFilled(t *testing.T) {
	s := &S1{
		F01: 1,
		F02: 2,
		F03: "three",
		F04: "four",
		F05: true,
		F06: false,
	}

	vals := make(map[string][]string)
	_ = NewEncoder().Encode(s, vals)

	if val, ok := vals["f01"]; ok {
		if val[0] != "1" {
			t.Error("expected '1' got ", val[0])
		}
	}
	if val, ok := vals["f02"]; ok {
		t.Error("Expected 'f02' to be omitted, got ", val[0])
	}
	if val, ok := vals["f03"]; ok {
		if val[0] != "three" {
			t.Error("expected 'three' got ", val[0])
		}
	}
	if val, ok := vals["f05"]; ok {
		if val[0] != "1" {
			t.Error("expected '1' got ", val[0])
		}
	}
	if val, ok := vals["f06"]; ok {
		if val[0] != "0" {
			t.Error("expected '0' got ", val[0])
		}
	}
}

func TestEmpty(t *testing.T) {
	s := &S1{
		F01: 1,
		F02: 2,
		F03: "three",
	}

	vals := make(map[string][]string)
	_ = NewEncoder().Encode(s, vals)

	if _, ok := vals["f03"]; !ok {
		t.Error("omitempty expected not empty, got empty")
	}
	if val, ok := vals["f04"]; ok {
		t.Error("omitempty expected empty, got ", val)
	}
}
