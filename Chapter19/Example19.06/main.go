package main

/*
#include <stdio.h>
#include <stdlib.h>

static void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C" // cgo
import (
	"fmt"
	"unsafe"
)

func main() {
	s := "Hello C!"

	cString := C.CString(s)
	defer C.free(unsafe.Pointer(cString))
	C.myprint(cString)

	b := C.GoBytes(unsafe.Pointer(cString), C.int(len(s)))
	fmt.Println(string(b))
}
