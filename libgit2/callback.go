package main

/*
#include <git2.h>
*/
import "C"
import "fmt"

//export TreeWalker
func TreeWalker(entry *C.git_tree_entry) {
	cname := C.git_tree_entry_name(entry)
	filename := C.GoString(cname)
	fmt.Println(filename)
}
