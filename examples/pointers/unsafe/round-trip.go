// Source: https://stackoverflow.com/questions/38001462/save-pointer-into-a-string-data-structure-in-golang
package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	//Given:
	data := "Hello"

	ptrData := &data // 824633778256

	fmt.Printf("%T, %#v, %d\n", ptrData, ptrData, ptrData)

	ptrString := fmt.Sprintf("%d", &data) // "824633778256"

	//Convert it to a uint64
	ptrInt, _ := strconv.ParseUint(ptrString, 10, 64) // 824633778256

	//They should match
	fmt.Printf("Address as String: %s as Int: %d\n", ptrString, ptrInt)

	//Convert the integer to a uintptr type
	ptrVal := uintptr(ptrInt) // Addr: 824633778256

	//Convert the uintptr to a Pointer type
	ptr := unsafe.Pointer(ptrVal)

	//Get the string pointer by address
	stringPtr := (*string)(ptr)

	//Get the value at that pointer
	newData := *stringPtr

	//Got it:
	fmt.Println(newData)

	//Test
	if stringPtr == &data && data == newData {
		fmt.Println("successful round trip!")
	} else {
		fmt.Println("uhoh! Something went wrong...")
	}
}
