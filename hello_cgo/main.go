package main

/*
#include <stdio.h>
#include <stdlib.h>

void hello(char *str)
{
	printf("Hello, cgo : %s\n",  str);
}
*/
import "C"
import "unsafe"

func main() {
	str := C.CString("Gopher")
	defer C.free(unsafe.Pointer(str))

	C.hello(str)
}
