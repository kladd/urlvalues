/*
Package urlvalues fills form values with the contents of a struct.

This package is meant to be used in conjunction with the gorilla toolkit's [schema](https://github.com/gorilla/schema) library, which _decodes_ values into structs.

Example

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
*/
package urlvalues
