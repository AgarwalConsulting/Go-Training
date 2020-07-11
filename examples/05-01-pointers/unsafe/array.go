// Source: https://copyninja.info/blog/workaround-gotypesystems.html
package main

/*
#cgo CFLAGS: -std=c99

#include "bytetest.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func ReadArray() unsafe.Pointer {
       var outArray = unsafe.Pointer (C.calloc(20,1))
       C.ArrayReadFunc((*C.UBYTE)(outArray))

       return outArray
}

func WriteArray(inArray unsafe.Pointer) {
       C.ArrayWriteFunc((*C.UBYTE)(inArray))
}

func CArrayToByteSlice(array unsafe.Pointer, size int) []byte {
       var arrayptr = uintptr(array)
       var byteSlice = make([]byte, size)

       for i := 0; i < len(byteSlice); i ++ {
               byteSlice[i] = byte(*(*C.UBYTE)(unsafe.Pointer(arrayptr)))
               arrayptr ++
       }

       return byteSlice
}

func ByteSliceToCArray (byteSlice []byte) unsafe.Pointer {
       var array = unsafe.Pointer(C.calloc(C.size_t(len(byteSlice)), 1))
       var arrayptr = uintptr(array)

       for i := 0; i < len(byteSlice); i ++ {
              *(*C.UBYTE)(unsafe.Pointer(arrayptr)) = C.UBYTE(byteSlice[i])
              arrayptr ++
       }

       return array
}

func main(){
        carray := ReadArray()
        defer C.free(carray)

        carraybytes := CArrayToByteSlice(carray, 20)

        fmt.Println("C array converted to byte slice:")
        for i := 0; i < len(carraybytes); i ++ {
                fmt.Printf("%d ", carraybytes[i])
        }

        fmt.Println()

        gobytes := []byte{21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
                31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
        gobytesarray := ByteSliceToCArray(gobytes)
        defer C.free(gobytesarray)

        WriteArray(gobytesarray)
}
