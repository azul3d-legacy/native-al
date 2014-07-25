// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
#include <stdlib.h>
#include <dlfcn.h>
#cgo LDFLAGS: -ldl

*/
import "C"

import (
	"errors"
	"unsafe"
)

var (
	lib unsafe.Pointer
)

func loadLibrary(name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	lib = C.dlopen(cname, C.RTLD_LAZY|C.RTLD_GLOBAL)
	cerr := C.dlerror()
	if lib == nil || cerr != nil {
		err := C.GoString(cerr)
		return errors.New(err)
	}
	return nil
}

func symbol(name string) unsafe.Pointer {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return C.dlsym(lib, cname)
}
