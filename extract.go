// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
extern void* openal_data[] asm("_binary_openal_start");

#if defined(_WIN32) || defined(_WIN64)
	// This stub is to fix a compilation error with Go 1.2 on Windows.
	//
	// It is likely that we can remove this once the external linker is used on
	// Windows in Go 1.3, right now this is just kind of a hack to get it to
	// compile OK with Go 1.2.
	int _binary_openal_size;
	void* openal_data_size = (void*)1619557; // hard-coded file size in bytes
#else
	extern void* openal_data_size[] asm("_binary_openal_size");
#endif


void* get_openal_data(void) {
	return (void*)openal_data;
}

void* get_openal_data_size(void) {
	return (void*)openal_data_size;
}
*/
import "C"

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"unsafe"
)

var (
	blob        []byte
	libraryPath string
)

func init() {
	// Initialize blob slice using blob data
	fsz := C.get_openal_data_size()
	sz := *(*int)(unsafe.Pointer(&fsz))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&blob))
	sliceHeader.Data = uintptr(C.get_openal_data())
	sliceHeader.Len = sz
	sliceHeader.Cap = sz

	// Determine where the library should actually be placed on the system
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	libraryPath = filepath.Join(usr.HomeDir, ".azul3d")
	libraryPath = filepath.Join(libraryPath, blobFileName)

	// Check if the library already exists at that location -- if it does it
	// means we've already extracted it there or the user has placed their own
	// implementation of the library there (per the LGPL restrictions).
	_, err = os.Stat(libraryPath)
	if err != nil {
		err := os.MkdirAll(filepath.Dir(libraryPath), 0777)
		if err != nil {
			log.Fatal(err)
		}

		// There is no dynamic library at that location, we can extract our
		// copy of it then.
		err = ioutil.WriteFile(libraryPath, blob, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Load the library
	err = loadLibrary(libraryPath)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize pointers
	alInit()
	alcInit()
}
