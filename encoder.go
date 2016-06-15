package urlvalues

import (
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

// SetAliasTag changes the tag used to locate urlvalues annotations. The default
// value is "url"
func (e *Encoder) SetAliasTag(tag string) *Encoder {
	e.TagID = tag
	return e
}

// Encode encodes a struct into map[string][]string.
//
// Intended for use with url.Values.
func (e *Encoder) Encode(src interface{}, dst map[string][]string) error {
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		v = v.Elem()
	}

	return e.encode(v, dst)
}

func (e *Encoder) encode(v reflect.Value, dst map[string][]string) error {
	var opts string
	var value string
	t := v.Type()

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

		encFunc, recurse := encoder(v.Field(i).Type())
		if recurse {
			e.encode(v.Field(i), dst)
			continue
		}

		value = encFunc(v.Field(i))

		if value == "" && strings.Contains(opts, "omitempty") {
			continue
		}

		dst[name] = []string{value}
	}

	return nil
}

func encoder(t reflect.Type) (func(v reflect.Value) string, bool) {
	switch t.Kind() {
	case reflect.Bool:
		return boolEncoder, false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intEncoder, false
	case reflect.Float32:
		return float32Encoder, false
	case reflect.Float64:
		return float64Encoder, false
	case reflect.Ptr:
		f, recurse := ptrEncoder(t)
		return f, recurse
	case reflect.String:
		return stringEncoder, false
	case reflect.Struct:
		return unsupportedEncoder, true
	default:
		return unsupportedEncoder, false
	}
}

func boolEncoder(v reflect.Value) string {
	if v.Bool() {
		return "1"
	}
	return "0"
}

func intEncoder(v reflect.Value) string {
	return strconv.Itoa(int(v.Int()))
}

func float32Encoder(v reflect.Value) string {
	return strconv.FormatFloat(v.Float(), 'f', 6, 32)
}

func float64Encoder(v reflect.Value) string {
	return strconv.FormatFloat(v.Float(), 'f', 6, 64)
}

func ptrEncoder(t reflect.Type) (func(v reflect.Value) string, bool) {
	f, recurse := encoder(t.Elem())

	return func(v reflect.Value) string {
		if v.IsNil() {
			return "null"
		}
		return f(v.Elem())
	}, recurse
}

func stringEncoder(v reflect.Value) string {
	return v.String()
}

func unsupportedEncoder(v reflect.Value) string {
	return ""
}
