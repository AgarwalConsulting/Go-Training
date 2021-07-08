// Use strings.Builder
// Builder: Design a html builder
// Builder Facet: Design a PersonBuilder, PersonJobBuilder, PersonAddressBuilder
// Builder Parameter: Design an EmailBuilder => func SendEmail(action func(b *EmailBuilder) {})
// Functional Builder: Design PersonBuilder combining Facet with Builder Parameter for lazy building

package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder // Initializes to it's zero value
	// b := strings.Builder{} // Struct-literal syntax

	fmt.Fprintln(&b, "<html>")
	fmt.Fprintln(&b, "<body>")
	fmt.Fprintln(&b, "Hello, World!")
	fmt.Fprintln(&b, "</body>")
	fmt.Fprintln(&b, "</html>")

	fmt.Println(b.String())
}
