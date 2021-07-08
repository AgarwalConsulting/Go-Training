package main

import "fmt"

type Product struct {
	Name  string
	Color string
	Size  string
}

type Specification interface {
	Is(product Product) bool
}

type ProductFilter struct {
	products []Product
}

func (pf ProductFilter) Filter(spec Specification) []Product {
	var res []Product

	for _, product := range pf.products {
		if spec.Is(product) {
			res = append(res, product)
		}
	}

	return res
}

type ColorSpecification string

func (cs ColorSpecification) Is(product Product) bool {
	return product.Color == string(cs)
}

type SizeSpecification string

func (cs SizeSpecification) Is(product Product) bool {
	return product.Size == string(cs)
}

type AndSpecification struct {
	operand [2]Specification
}

func (and AndSpecification) Is(product Product) bool {
	return and.operand[0].Is(product) && and.operand[1].Is(product)
}

func (pf ProductFilter) FilterBySize(size string) []Product {
	var res []Product

	for _, product := range pf.products {
		if product.Size == size {
			res = append(res, product)
		}
	}

	return res
}

func (pf ProductFilter) FilterByColor(color string) []Product {
	var res []Product

	for _, product := range pf.products {
		if product.Color == color {
			res = append(res, product)
		}
	}

	return res
}

func main() {
	products := []Product{
		{"Apple", "red", "small"},
		{"Table", "brown", "medium"},
		{"House", "green", "large"},
		{"Printer", "grey", "medium"},
	}

	pf := ProductFilter{products}

	// res := pf.FilterBySize("medium")

	res := pf.Filter(AndSpecification{[2]Specification{ColorSpecification("grey"), SizeSpecification("medium")}})

	fmt.Println(res)
}
