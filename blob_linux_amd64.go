// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
extern void* openal_data[] asm("_binary_openal_start");
extern void* openal_data_size[] asm("_binary_openal_size");

void* get_openal_data(void) {
  return (void*)openal_data;
}

void* get_openal_data_size(void) {
  return (void*)openal_data_size;
}
*/
import "C"

import (
  "reflect"
  "unsafe"
)

const blobFileName = "libopenal_soft.so.1.15.1"

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
