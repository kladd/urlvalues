package urlvalues

import "testing"

type S1 struct {
	F01 int    `url:"f01"`
	F02 int    `url:"-"`
	F03 string `url:"f03"`
	F04 string `url:"-"`
	F05 string `url:"f05,omitempty"`
}

func TestFilled(t *testing.T) {
	s := &S1{
		F01: 1,
		F02: 2,
		F03: "three",
		F04: "four",
		F05: "five",
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
			t.Error("expected 'four' got ", val[0])
		}
	}
	if val, ok := vals["f04"]; ok {
		t.Error("Expected 'four' to be omitted, got ", val[0])
	}
	if val, ok := vals["f05"]; ok {
		if val[0] != "five" {
			t.Error("expected 'five' got ", val[0])
		}
	}
}

func TestEmpty(t *testing.T) {
	s := &S1{
		F01: 1,
		F02: 2,
		F04: "four",
	}

	vals := make(map[string][]string)
	_ = NewEncoder().Encode(s, vals)

	if _, ok := vals["f03"]; !ok {
		t.Error("omitempty expected not empty, got empty")
	}
	if val, ok := vals["f05"]; ok {
		t.Error("omitempty expected empty, got ", val)
	}
}
