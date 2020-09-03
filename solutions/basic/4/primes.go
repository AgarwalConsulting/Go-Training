package main

import "fmt"

func getProducts(number int) []int {
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return []int{i, number / i}
		}
	}
	return nil
}

func main() {
	for i := 0; i < 25; i++ {
		num := i + 1
		products := getProducts(num)

		if products == nil {
			fmt.Printf("%d is prime\n", num)
		} else {
			fmt.Printf("%d is a product of %d * %d\n", num, products[0], products[1])
		}
	}
}
