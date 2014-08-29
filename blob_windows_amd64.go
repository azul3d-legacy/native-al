// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
extern void* openal_data[] asm("_binary_openal_start");

// This stub is to fix a compilation error with Go 1.2 on Windows.
//
// It is likely that we can remove this once the external linker is used on
// Windows in Go, right now this is just kind of a hack to get it to
// compile OK with Go 1.2.
int _binary_openal_size;
void* openal_data_size = (void*)1619557; // hard-coded file size in bytes

void* get_openal_data(void) {
  return (void*)openal_data;
}

void* get_openal_data_size(void) {
  return (void*)openal_data_size;
}
*/
import "C"

import(
  "reflect"
  "unsafe"
)

const blobFileName = "openal_soft.dll"

var blob []byte

func initBlob() {
  // Initialize blob slice using blob data
  fsz := C.get_openal_data_size()
  sz := *(*int)(unsafe.Pointer(&fsz))
  sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&blob))
  sliceHeader.Data = uintptr(C.get_openal_data())
  sliceHeader.Len = sz
  sliceHeader.Cap = sz
}
