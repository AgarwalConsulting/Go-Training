package main

import (
	"fmt"
	"unsafe"
)

// type Coordinates struct {
// 	Lat  float64
// 	Long float64
// }

type Place struct {
	Name string
	// Coordinates Coordinates
	Coordinates struct {
		Lat  float64
		Long float64
	}
}

func main() {
	// place := Place{
	// 	Name: "Global Edge",
	// 	Coordinates: Coordinates{
	// 		Lat:  42.1245,
	// 		Long: 12.124,
	// 	},
	// }

	place := Place{"Global Edge", struct {
		Lat  float64
		Long float64
	}{42.1245, 12.124}}

	var placePtr *Place

	fmt.Println(unsafe.Sizeof(placePtr))

	placePtr = new(Place)
	fmt.Println(unsafe.Sizeof(placePtr))
	// placePtr = &(Place{})

	fmt.Println(placePtr)

	placePtr.Name = "Google"

	fmt.Println(place)
}
