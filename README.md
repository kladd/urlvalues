# urlvalues - Golang structs to form values

[![GoDoc](https://godoc.org/github.com/kladd/urlvalues?status.svg)](https://godoc.org/github.com/kladd/urlvalues) [![Build Status](https://travis-ci.org/kladd/urlvalues.png?branch=master)](https://travis-ci.org/kladd/urlvalues)

Package urlvalues fills form values with the contents of a struct.

This package is meant to be used in conjunction with the gorilla toolkit's [schema](https://github.com/gorilla/schema) library, which _decodes_ values into structs.

### Example

```go
type Person struct {
	Name  string `url:"name"`
	Phone string `url:"phone"`
}

func main() {
	jane := &Person{"Jane Doe", "(111) 555-5555"}
	vals := url.Values{}

	// Encode Person into url.Values
	encoder := urlvalues.NewEncoder()
	encoder.Encode(jane, vals)

	// Use url.Values.Encode() to output a query string
	// name=Jane+Doe&phone=%28111%29+555-5555
	fmt.Println(vals.Encode())
}
```

The supported field types so far:
* bool(&)
* int(8/16/32/64/&)
* float(32/64/&)
* string(&)
* struct


### License

MIT Licensed. See [LICENSE](./LICENSE).
