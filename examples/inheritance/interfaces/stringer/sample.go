package main

import "fmt"

type MyFloat float64

func (f MyFloat) String() string {
	real := int(f)
	fraction := f - MyFloat(real)

	return fmt.Sprintf("Real: %d, Fraction: %f", real, fraction)
}

// func PrintVerbose(f MyFloat) {
// 	real := int(f)
// 	fraction := f - MyFloat(real)

// 	fmt.Printf("Real: %d, Fraction: %f\n", real, fraction)
// }

func main() {
	f := MyFloat(64.2)

	fmt.Println(f)

	// PrintVerbose(f)
}
