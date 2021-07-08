// Use strings.Builder
// Builder: Design a html builder
// Builder Facet: Design a PersonBuilder, PersonJobBuilder, PersonAddressBuilder
// Builder Parameter: Design an EmailBuilder => func SendEmail(action func(b *EmailBuilder) {})
// Functional Builder: Design PersonBuilder combining Facet with Builder Parameter for lazy building

package main

func SendEmail(action func(b *EmailBuilder)) {

}

func main() {
	SendEmail(func(eb *EmailBuilder) {
		eb.To("")
		eb.From("")
		e.Subject("")
	})
}
