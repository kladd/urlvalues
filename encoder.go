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

		value = encoder(v.Field(i).Type())(v.Field(i))
		if value == "" && strings.Contains(opts, "omitempty") {
			continue
		}

		dst[name] = []string{value}
	}

	return nil
}

func encoder(t reflect.Type) func(v reflect.Value) string {
	switch t.Kind() {
	case reflect.String:
		return stringEncoder
	case reflect.Int:
		return intEncoder
	case reflect.Bool:
		return boolEncoder
	default:
		return unsupportedEncoder
	}
}

func stringEncoder(v reflect.Value) string {
	return v.String()
}

func intEncoder(v reflect.Value) string {
	return strconv.Itoa(int(v.Int()))
}

func boolEncoder(v reflect.Value) string {
	if v.Bool() {
		return "1"
	}
	return "0"
}

func unsupportedEncoder(v reflect.Value) string {
	return ""
}
