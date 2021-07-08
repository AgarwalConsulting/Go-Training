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

type HtmlElement struct {
	TagName  string
	Text     string
	Children []HtmlElement
}

func (he *HtmlElement) Render() string {
	b := &strings.Builder{}

	fmt.Fprintln(b, fmt.Sprintf("<%s>", he.TagName))

	for _, child := range he.Children {
		fmt.Fprintln(b, child.Render())
	}

	fmt.Fprintln(b, he.Text)
	fmt.Fprintln(b, fmt.Sprintf("</%s>", he.TagName))

	return b.String()
}

type HtmlBuilder struct {
	Root HtmlElement
}

func (b *HtmlBuilder) AddRoot(tagName string) *HtmlBuilder {
	b.Root = HtmlElement{
		TagName: tagName,
	}

	return b
}

func (b *HtmlBuilder) AddChild(tagName string, childText string) *HtmlBuilder {
	e := HtmlElement{
		TagName:  tagName,
		Text:     childText,
		Children: []HtmlElement{},
	}

	b.Root.Children = append(b.Root.Children, e)
	return b
}

func main() {
	b := HtmlBuilder{}

	b.AddRoot("ul")
	b.AddChild("li", "Hrishikesh")
	b.AddChild("li", "Bipul")
	b.AddChild("li", "Krishnan")

	fmt.Println(b.Root.Render())

	// b := HtmlBuilder{}

	// b.AddRoot("html")
	// b.AddChild("body", "")
	// b.AddChild("li", "Bipul")
	// b.AddChild("li", "Krishnan")

	// fmt.Println(b.Root.Render())

	// html := HtmlElement{TagName: "html"}

	// html.Children = []HtmlElement{HtmlElement{TagName: "body", Children: []HtmlElement{{Text: "Hello, World!"}}}}

	// fmt.Println(html.Render())

	// var b strings.Builder // Initializes to it's zero value
	// // b := strings.Builder{} // Struct-literal syntax

	// fmt.Fprintln(&b, "<html>")
	// fmt.Fprintln(&b, "<body>")
	// fmt.Fprintln(&b, "Hello, World!")
	// fmt.Fprintln(&b, "</body>")
	// fmt.Fprintln(&b, "</html>")

	// fmt.Println(b.String())
}
