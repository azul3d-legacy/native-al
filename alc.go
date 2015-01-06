// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
#include <stdlib.h>
#include "AL/alc.h"

#include "AL/alc.h"

#include "AL/al.h"
void* goal_GetProcAddress(void* proc, const ALchar* a);

ALCcontext* goal_alcCreateContext(void* proc, ALCdevice* a, const ALCint* b);
ALCboolean goal_alcMakeContextCurrent(void* proc, ALCcontext* a);
void goal_alcProcessContext(void* proc, ALCcontext* a);
void goal_alcSuspendContext(void* proc, ALCcontext* a);
void goal_alcDestroyContext(void* proc, ALCcontext* a);
ALCcontext* goal_alcGetCurrentContext(void* proc);
ALCdevice* goal_alcGetContextsDevice(void* proc, ALCcontext* a);
ALCdevice* goal_alcOpenDevice(void* proc, const ALCchar* a);
ALCboolean goal_alcCloseDevice(void* proc, ALCdevice* a);
ALCenum goal_alcGetError(void* proc, ALCdevice* a);
ALCboolean goal_alcIsExtensionPresent(void* proc, ALCdevice* a, const ALCchar* b);
void* goal_alcGetProcAddress(void* proc, ALCdevice* a, const ALCchar* b);
ALCenum goal_alcGetEnumValue(void* proc, ALCdevice* a, const ALCchar* b);
const ALCchar* goal_alcGetString(void* proc, ALCdevice* a, ALCenum b);
void goal_alcGetIntegerv(void* proc, ALCdevice* a, ALCenum b, ALCsizei c, ALCint* d);
ALCdevice* goal_alcCaptureOpenDevice(void* proc, const ALCchar* a, ALCuint b, ALCenum c, ALCsizei d);
ALCboolean goal_alcCaptureCloseDevice(void* proc, ALCdevice* a);
void goal_alcCaptureStart(void* proc, ALCdevice* a);
void goal_alcCaptureStop(void* proc, ALCdevice* a);
void goal_alcCaptureSamples(void* proc, ALCdevice* a, ALCvoid* b, ALCsizei c);
*/
import "C"

import (
	"unsafe"
)

var (
	p_alcCreateContext, p_alcMakeContextCurrent, p_alcProcessContext,
	p_alcSuspendContext, p_alcDestroyContext, p_alcGetCurrentContext,
	p_alcGetContextsDevice, p_alcOpenDevice, p_alcCloseDevice, p_alcGetError,
	p_alcIsExtensionPresent, p_alcGetProcAddress, p_alcGetEnumValue,
	p_alcGetString, p_alcGetIntegerv, p_alcCaptureOpenDevice,
	p_alcCaptureCloseDevice, p_alcCaptureStart, p_alcCaptureStop,
	p_alcCaptureSamples unsafe.Pointer
)

func alcInit() {
	p_alcCreateContext = symbol("alcCreateContext")
	p_alcMakeContextCurrent = symbol("alcMakeContextCurrent")
	p_alcProcessContext = symbol("alcProcessContext")
	p_alcSuspendContext = symbol("alcSuspendContext")
	p_alcDestroyContext = symbol("alcDestroyContext")
	p_alcGetCurrentContext = symbol("alcGetCurrentContext")
	p_alcGetContextsDevice = symbol("alcGetContextsDevice")
	p_alcOpenDevice = symbol("alcOpenDevice")
	p_alcCloseDevice = symbol("alcCloseDevice")
	p_alcGetError = symbol("alcGetError")
	p_alcIsExtensionPresent = symbol("alcIsExtensionPresent")
	p_alcGetProcAddress = symbol("alcGetProcAddress")
	p_alcGetEnumValue = symbol("alcGetEnumValue")
	p_alcGetString = symbol("alcGetString")
	p_alcGetIntegerv = symbol("alcGetIntegerv")
	p_alcCaptureOpenDevice = symbol("alcCaptureOpenDevice")
	p_alcCaptureCloseDevice = symbol("alcCaptureCloseDevice")
	p_alcCaptureStart = symbol("alcCaptureStart")
	p_alcCaptureStop = symbol("alcCaptureStop")
	p_alcCaptureSamples = symbol("alcCaptureSamples")
}

func (d *Device) alcCheckError() {
	cerr := C.goal_alcGetError(p_alcGetError, d.c)
	var err error
	switch cerr {
	case C.ALC_NO_ERROR:
		break
	case C.ALC_INVALID_ENUM:
		err = ErrInvalidEnum
	case C.ALC_INVALID_VALUE:
		err = ErrInvalidValue
	case C.ALC_OUT_OF_MEMORY:
		err = ErrOutOfMemory
	}
	if err != nil {
		go ErrorHandler()(err)
	}
	return
}

type Device struct {
	destroyed, capture bool
	name               string
	ctx                *C.ALCcontext
	c, captureDevice   *C.ALCdevice

	p_BufferDataStatic, p_BufferSubDataSOFT, p_BufferSamplesSOFT,
	p_BufferSubSamplesSOFT, p_GetBufferSamplesSOFT,
	p_IsBufferFormatSupportedSOFT, p_SourcedSOFT, p_Source3dSOFT,
	p_SourcedvSOFT, p_GetSourcedSOFT, p_GetSource3dSOFT, p_GetSourcedvSOFT,
	p_Sourcei64SOFT, p_Source3i64SOFT, p_Sourcei64vSOFT, p_GetSourcei64SOFT,
	p_GetSource3i64SOFT, p_GetSourcei64vSOFT unsafe.Pointer
}

func (d *Device) mustBeCaptureDevice() {
	if !d.capture {
		panic("openal: illegal call (device is not a capture device).")
	}
}

func (d *Device) panicIfDestroyed() {
	if d.destroyed {
		panic("openal: illegal call (device is already closed).")
	}
}

var (
	activeDevice *Device
)

func (d *Device) doMakeCurrent() {
	if activeDevice != d {
		activeDevice = d
		C.goal_alcMakeContextCurrent(p_alcMakeContextCurrent, d.ctx)
	}
}

func (d *Device) alDispatch(f func()) {
	// Make our context current
	d.doMakeCurrent()

	// Execute function
	f()

	// Check for errors
	checkError()
}

func (d *Device) alcDispatch(f func()) {
	// Make our context current
	d.doMakeCurrent()

	// Execute function
	f()

	// Check for errors
	d.alcCheckError()
}

func (d *Device) alExtensionInit(ptr *unsafe.Pointer, name string) unsafe.Pointer {
	if *ptr == nil {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		*ptr = C.goal_GetProcAddress(p_GetProcAddress, (*C.ALchar)(unsafe.Pointer(cname)))
	}
	return *ptr
}

func (d *Device) AlcIsExtensionPresent(name string) bool {
	var ret bool
	d.alcDispatch(func() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		ret = int32(C.goal_alcIsExtensionPresent(p_alcIsExtensionPresent, d.c, (*C.ALCchar)(unsafe.Pointer(cname)))) == TRUE
	})
	return ret
}

func (d *Device) AlcGetProcAddress(name string) unsafe.Pointer {
	var ret unsafe.Pointer
	d.alcDispatch(func() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		ret = C.goal_alcGetProcAddress(p_alcGetProcAddress, d.c, (*C.ALCchar)(unsafe.Pointer(cname)))
	})
	return ret
}

func (d *Device) AlcGetEnumValue(name string) int32 {
	var ret int32
	d.alcDispatch(func() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		ret = int32(C.goal_alcGetEnumValue(p_GetEnumValue, d.c, (*C.ALCchar)(unsafe.Pointer(cname))))
	})
	return ret
}

func AlcGetRawString(d *Device, param int32) uintptr {
	var cdevice *C.ALCdevice
	if d != nil {
		cdevice = d.c
	}
	return uintptr(unsafe.Pointer(C.goal_alcGetString(p_alcGetString, cdevice, C.ALCenum(param))))
}

func AlcGetString(d *Device, param int32) string {
	raw := AlcGetRawString(d, param)
	return C.GoString((*C.char)(unsafe.Pointer(raw)))
}

func (d *Device) AlcGetRawString(param int32) uintptr {
	var ret uintptr
	d.alcDispatch(func() {
		ret = uintptr(unsafe.Pointer(C.goal_alcGetString(p_alcGetString, d.c, C.ALCenum(param))))
	})
	return ret
}

func (d *Device) AlcGetString(param int32) string {
	raw := d.AlcGetRawString(param)
	return C.GoString((*C.char)(unsafe.Pointer(raw)))
}

func AlcGetIntegerv(d *Device, param, size int32, values *int32) {
	var cdevice *C.ALCdevice
	if d != nil {
		cdevice = d.c
	}
	C.goal_alcGetIntegerv(
		p_alcGetIntegerv,
		cdevice,
		C.ALCenum(param),
		C.ALCsizei(size),
		(*C.ALCint)(unsafe.Pointer(values)),
	)
	return
}

func (d *Device) AlcGetIntegerv(param int32, size int32, values *int32) {
	d.alcDispatch(func() {
		C.goal_alcGetIntegerv(
			p_alcGetIntegerv,
			d.c,
			C.ALCenum(param),
			C.ALCsizei(size),
			(*C.ALCint)(unsafe.Pointer(values)),
		)
	})
	return
}

// StartCapture starts capturing on this device.
//
// This function will panic if the device is not yet initialized for capturing
// audo samples via the InitCapture() method.
//
// This function will panic if the device was already closed via the Close()
// method.
func (d *Device) StartCapture() {
	d.panicIfDestroyed()
	d.mustBeCaptureDevice()
	d.alcDispatch(func() {
		C.goal_alcCaptureStart(p_alcCaptureStart, d.c)
	})
}

// StopCapture stops capturing on this device.
//
// This function will panic if the device is not yet initialized for capturing
// audo samples via the InitCapture() method.
//
// This function will panic if the device was already closed via the Close()
// method.
func (d *Device) StopCapture() {
	d.panicIfDestroyed()
	d.mustBeCaptureDevice()
	d.alcDispatch(func() {
		C.goal_alcCaptureStop(p_alcCaptureStop, d.c)
	})
}

// CaptureSamples fills the buffer with n audio samples.
//
// This function will panic if the device is not yet initialized for capturing
// audo samples via the InitCapture() method.
//
// This function will panic if the device was already closed via the Close()
// method.
func (d *Device) CaptureSamples(buffer unsafe.Pointer, n int32) {
	d.panicIfDestroyed()
	d.mustBeCaptureDevice()
	d.alcDispatch(func() {
		C.goal_alcCaptureSamples(p_alcCaptureSamples, d.c, buffer, C.ALCsizei(n))
	})
}

// InitCapture initializes this device for capturing audio samples.
func (d *Device) InitCapture(frequency uint32, format, bufferSize int32) error {
	var cname *C.char
	if len(d.name) > 0 {
		cname := C.CString(d.name)
		defer C.free(unsafe.Pointer(cname))
	}
	var err error

	d.alcDispatch(func() {
		d.captureDevice = C.goal_alcCaptureOpenDevice(
			p_alcCaptureOpenDevice,
			(*C.ALCchar)(unsafe.Pointer(cname)),
			C.ALCuint(frequency),
			C.ALCenum(format),
			C.ALCsizei(bufferSize),
		)
		if d.captureDevice == nil {
			err = ErrCantOpenDevice
			return
		}
	})

	d.capture = true
	if err != nil {
		return err
	}
	return nil
}

// Close closes the device. This function will panic if all of the device's
// associated buffers have not been destroyed (indicating a programmer fault).
//
// This function will panic if the device was already closed via the Close()
// method.
func (d *Device) Close() {
	d.panicIfDestroyed()
	d.alcDispatch(func() {
		d.destroyed = true

		// Destroy the context
		C.goal_alcMakeContextCurrent(p_alcMakeContextCurrent, nil)
		C.goal_alcDestroyContext(p_alcDestroyContext, d.ctx)

		// Close the device
		if d.capture {
			if int32(C.goal_alcCaptureCloseDevice(p_alcCaptureCloseDevice, d.captureDevice)) == FALSE {
				panic("openal: attempted to close audio device without closing all of it's associated buffers first.")
			}
		}
		if int32(C.goal_alcCloseDevice(p_alcCloseDevice, d.c)) == FALSE {
			panic("openal: attempted to close audio device without closing all of it's associated buffers first.")
		}
	})
}

func (d *Device) makeContext(attribs []int32) (err error) {
	var cattrs *C.ALCint
	if len(attribs) > 0 {
		cattrs = (*C.ALCint)(unsafe.Pointer(&attribs[0]))
	}
	d.ctx = C.goal_alcCreateContext(p_alcCreateContext, d.c, cattrs)
	if d.ctx == nil {
		err = ErrCantCreateContext
		return
	}
	if int32(C.goal_alcMakeContextCurrent(p_alcMakeContextCurrent, d.ctx)) == FALSE {
		err = ErrCantMakeContextCurrent
	}
	return
}

// StringList splits a string list. It must be NUL seperated with two NULs at
// the end.
func StringList(raw uintptr) []string {
	var list []string
	// Split memory at raw+i by NUL bytes until two consecutive ones.
	var (
		i       = raw
		segment []byte
	)
	for {
		b := (*byte)(unsafe.Pointer(i))
		if *b == 0 {
			// Do we have a non-empty segment?
			if len(segment) > 0 {
				list = append(list, string(segment))
				segment = segment[:0]
			}

			// Is the next byte NUL too? If so we're done.
			b2 := (*byte)(unsafe.Pointer(i + 1))
			if *b2 == 0 {
				return list
			}
			i++
			continue
		}

		// Append to segment.
		segment = append(segment, *b)
		i++
	}
}

// OpenDevice opens the named device (or "" for the default device).
func OpenDevice(name string, ctxAttribs []int32) (*Device, error) {
	var cname *C.char
	if len(name) > 0 {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))
	}

	d := new(Device)
	d.capture = false
	d.name = name
	d.c = C.goal_alcOpenDevice(
		p_alcOpenDevice,
		(*C.ALCchar)(unsafe.Pointer(cname)),
	)
	if d.c == nil {
		return nil, ErrCantOpenDevice
	}
	err := d.makeContext(ctxAttribs)
	if err != nil {
		return nil, err
	}
	return d, nil
}
