package main

// #cgo LDFLAGS: -lleveldb
// #include <stdlib.h>
// #include "leveldb/c.h"
import "C"

import (
	"errors"
	"unsafe"
)

// C Level pointer holder
type LevelDB struct {
	CLevelDB *C.leveldb_t
	Name     string
}

// Open LevelDB with given name
func OpenLevelDB(path string) (leveldb *LevelDB, err error) {

	cpath := C.CString(path) // convert path to c string
	defer C.leveldb_free(unsafe.Pointer(cpath))

	// allocate LevelDB Option struct to open
	opt := C.leveldb_options_create()
	defer C.leveldb_free(unsafe.Pointer(opt))

	// set open option
	C.leveldb_options_set_create_if_missing(opt, C.uchar(1))

	// open leveldb
	var cerr *C.char
	cleveldb := C.leveldb_open(opt, cpath, &cerr)

	if cerr != nil {
		defer C.leveldb_free(unsafe.Pointer(cerr))
		return nil, errors.New(C.GoString(cerr))
	}

	return &LevelDB{cleveldb, path}, nil
}

// Close the database
func (db *LevelDB) Close() {
	C.leveldb_close(db.CLevelDB)
}

// Put key, value to database
func (db *LevelDB) Put(key, value string) (err error) {

	opt := C.leveldb_writeoptions_create() // write option
	defer C.leveldb_free(unsafe.Pointer(opt))

	k := C.CString(key) // copy
	defer C.leveldb_free(unsafe.Pointer(k))

	v := C.CString(value)
	defer C.leveldb_free(unsafe.Pointer(v))

	var cerr *C.char
	C.leveldb_put(db.CLevelDB, opt, k, C.size_t(len(key)), v, C.size_t(len(value)), &cerr)

	if cerr != nil {
		defer C.leveldb_free(unsafe.Pointer(cerr))
		return errors.New(C.GoString(cerr))
	}

	return
}

func (db *LevelDB) Get(key string) (value string, err error) {

	opt := C.leveldb_readoptions_create() // write option
	defer C.leveldb_free(unsafe.Pointer(opt))

	k := C.CString(key) // copy
	defer C.leveldb_free(unsafe.Pointer(k))

	var vallen C.size_t
	var cerr *C.char
	cvalue := C.leveldb_get(db.CLevelDB, opt, k, C.size_t(len(key)), &vallen, &cerr)

	if cerr != nil {
		defer C.leveldb_free(unsafe.Pointer(cerr))
		return "", errors.New(C.GoString(cerr))
	}

	if cvalue == nil {
		return "", nil
	}

	defer C.leveldb_free(unsafe.Pointer(cvalue))
	return C.GoString(cvalue), nil
}

func (db *LevelDB) Delete(key string) (err error) {
	opt := C.leveldb_writeoptions_create() // write option
	defer C.leveldb_free(unsafe.Pointer(opt))

	k := C.CString(key) // copy
	defer C.leveldb_free(unsafe.Pointer(k))

	var cerr *C.char
	C.leveldb_delete(db.CLevelDB, opt, k, C.size_t(len(key)), &cerr)

	if cerr != nil {
		defer C.leveldb_free(unsafe.Pointer(cerr))
		return errors.New(C.GoString(cerr))
	}

	return
}
