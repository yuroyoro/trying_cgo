package main

import "C"
import "runtime"

//export GoVersion
func GoVersion() *C.char {
	version := runtime.Version()
	return C.CString(version)
}
