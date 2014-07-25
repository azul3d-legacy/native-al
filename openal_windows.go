// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

import (
	"syscall"
	"unsafe"
)

var (
	lib *syscall.DLL
)

func loadLibrary(name string) error {
	var err error
	lib, err = syscall.LoadDLL(name)
	if err != nil {
		return err
	}
	return nil
}

func symbol(name string) unsafe.Pointer {
	proc := lib.MustFindProc(name)
	return unsafe.Pointer(proc.Addr())
}
