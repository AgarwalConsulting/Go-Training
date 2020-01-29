package geo

import "fmt"

func (loc Location) Print() {
	fmt.Printf("%f : %f\n", loc.Lat, loc.Long)
}

func Print(loc Location) {
	fmt.Printf("%f : %f\n", loc.Lat, loc.Long)
}
