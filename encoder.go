package urlvalues

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// Encoder encodes values from a struct into url.Values.
type Encoder struct {
	TagID string
}

// NewEncoder returns a new Encoder with defaults.
func NewEncoder() *Encoder {
	return &Encoder{TagID: "url"}
}

// Encode encodes a struct into map[string][]string.
//
// Intended for use with url.Values.
func (e *Encoder) Encode(src interface{}, dst map[string][]string) error {
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("urlutil: interface must be a pointer to a struct")
	}
	v = v.Elem()
	t := v.Type()

	var opts string
	var value string
	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag.Get(e.TagID)
		name := tag
		if idx := strings.Index(tag, ","); idx != -1 {
			name = tag[:idx]
			opts = tag[idx+1:]
		}

		if name == "-" {
			continue
		}

		switch v.Field(i).Type().Kind() {
		case reflect.String:
			value = v.Field(i).String()
		case reflect.Int:
			value = strconv.Itoa(int(v.Field(i).Int()))
		}

		if value == "" && strings.Contains(opts, "omitempty") {
			continue
		}

		dst[name] = []string{value}
	}

	return nil
}
