// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
#include <stdlib.h>
#include "AL/al.h"
#include "AL/alext.h"

void goal_Enable(void* proc, ALenum a);
void goal_Disable(void* proc, ALenum a);
ALboolean goal_IsEnabled(void* proc, ALenum a);
const ALchar* goal_GetString(void* proc, ALenum a);
void goal_GetBooleanv(void* proc, ALenum a, ALboolean* b);
void goal_GetIntegerv(void* proc, ALenum a, ALint* b);
void goal_GetFloatv(void* proc, ALenum a, ALfloat* b);
void goal_GetDoublev(void* proc, ALenum a, ALdouble* b);
ALboolean goal_GetBoolean(void* proc, ALenum a);
ALint goal_GetInteger(void* proc, ALenum a);
ALfloat goal_GetFloat(void* proc, ALenum a);
ALdouble goal_GetDouble(void* proc, ALenum a);
ALenum goal_GetError(void* proc);
ALboolean goal_IsExtensionPresent(void* proc, const ALchar* a);
void* goal_GetProcAddress(void* proc, const ALchar* a);
ALenum goal_GetEnumValue(void* proc, const ALchar* a);
void goal_Listenerf(void* proc, ALenum a, ALfloat b);
void goal_Listener3f(void* proc, ALenum a, ALfloat b, ALfloat c, ALfloat d);
void goal_Listenerfv(void* proc, ALenum a, ALfloat* b);
void goal_Listeneri(void* proc, ALenum a, ALint b);
void goal_Listener3i(void* proc, ALenum a, ALint b, ALint c, ALint d);
void goal_Listeneriv(void* proc, ALenum a, ALint* b);
void goal_GetListenerf(void* proc, ALenum a, ALfloat* b);
void goal_GetListener3f(void* proc, ALenum a, ALfloat* b, ALfloat* c, ALfloat* d);
void goal_GetListenerfv(void* proc, ALenum a, ALfloat* b);
void goal_GetListeneri(void* proc, ALenum a, ALint* b);
void goal_GetListener3i(void* proc, ALenum a, ALint* b, ALint* c, ALint* d);
void goal_GetListeneriv(void* proc, ALenum a, ALint* b);
void goal_GenSources(void* proc, ALsizei a, ALuint* b);
void goal_DeleteSources(void* proc, ALsizei a, ALuint* b);
ALboolean goal_IsSource(void* proc, ALuint a);
void goal_Sourcef(void* proc, ALuint a, ALenum b, ALfloat c);
void goal_Source3f(void* proc, ALuint a, ALenum b, ALfloat c, ALfloat d, ALfloat e);
void goal_Sourcefv(void* proc, ALuint a, ALenum b, ALfloat* c);
void goal_Sourcei(void* proc, ALuint a, ALenum b, ALint c);
void goal_Source3i(void* proc, ALuint a, ALenum b, ALint c, ALint d, ALint e);
void goal_Sourceiv(void* proc, ALuint a, ALenum b, ALint* c);
void goal_GetSourcef(void* proc, ALuint a, ALenum b, ALfloat* c);
void goal_GetSource3f(void* proc, ALuint a, ALenum b, ALfloat* c, ALfloat* d, ALfloat* e);
void goal_GetSourcefv(void* proc, ALuint a, ALenum b, ALfloat* c);
void goal_GetSourcei(void* proc, ALuint a, ALenum b, ALint* c);
void goal_GetSource3i(void* proc, ALuint a, ALenum b, ALint* c, ALint* d, ALint* e);
void goal_GetSourceiv(void* proc, ALuint a, ALenum b, ALint* c);
void goal_SourcePlayv(void* proc, ALsizei a, ALuint* b);
void goal_SourceStopv(void* proc, ALsizei a, ALuint* b);
void goal_SourceRewindv(void* proc, ALsizei a, ALuint* b);
void goal_SourcePausev(void* proc, ALsizei a, ALuint* b);
void goal_SourcePlay(void* proc, ALuint a);
void goal_SourceStop(void* proc, ALuint a);
void goal_SourceRewind(void* proc, ALuint a);
void goal_SourcePause(void* proc, ALuint a);
void goal_SourceQueueBuffers(void* proc, ALuint a, ALsizei b, const ALuint* c);
void goal_SourceUnqueueBuffers(void* proc, ALuint a, ALsizei b, ALuint* c);
void goal_GenBuffers(void* proc, ALsizei a, ALuint* b);
void goal_DeleteBuffers(void* proc, ALsizei a, const ALuint* b);
ALboolean goal_IsBuffer(void* proc, ALuint a);
void goal_BufferData(void* proc, ALuint a, ALenum b, const ALvoid* c, ALsizei d, ALsizei e);
void goal_Bufferf(void* proc, ALuint a, ALenum b, ALfloat c);
void goal_Buffer3f(void* proc, ALuint a, ALenum b, ALfloat c, ALfloat d, ALfloat e);
void goal_Bufferfv(void* proc, ALuint a, ALenum b, ALfloat* c);
void goal_Bufferi(void* proc, ALuint a, ALenum b, ALint c);
void goal_Buffer3i(void* proc, ALuint a, ALenum b, ALint c, ALint d, ALint e);
void goal_Bufferiv(void* proc, ALuint a, ALenum b, ALint* c);
void goal_GetBufferf(void* proc, ALuint a, ALenum b, ALfloat* c);
void goal_GetBuffer3f(void* proc, ALuint a, ALenum b, ALfloat* c, ALfloat* d, ALfloat* e);
void goal_GetBufferfv(void* proc, ALuint a, ALenum b, ALfloat* c);
void goal_GetBufferi(void* proc, ALuint a, ALenum b, ALint* c);
void goal_GetBuffer3i(void* proc, ALuint a, ALenum b, ALint* c, ALint* d, ALint* e);
void goal_GetBufferiv(void* proc, ALuint a, ALenum b, ALint* c);
void goal_DopplerFactor(void* proc, ALfloat a);
void goal_DopplerVelocity(void* proc, ALfloat a);
void goal_SpeedOfSound(void* proc, ALfloat a);
void goal_DistanceModel(void* proc, ALenum a);



// Begin extensions -- do not place non-extension functions below here.



void goal_BufferDataStatic(void* proc, const ALuint a, ALenum b, ALvoid* c, ALsizei d, ALsizei e);
void goal_BufferSubDataSOFT(void* proc, ALuint a, ALenum b, const ALvoid* c, ALsizei d, ALsizei e);
void goal_BufferSamplesSOFT(void* proc, ALuint a, ALuint b, ALenum c, ALsizei d, ALenum e, ALenum f, const ALvoid* g);
void goal_BufferSubSamplesSOFT(void* proc, ALuint a, ALsizei b, ALsizei c, ALenum d, ALenum e, const ALvoid* f);
void goal_GetBufferSamplesSOFT(void* proc, ALuint a, ALsizei b, ALsizei c, ALenum d, ALenum e, ALvoid* f);
ALboolean goal_IsBufferFormatSupportedSOFT(void* proc, ALenum a);
void goal_SourcedSOFT(void* proc, ALuint a, ALenum b, ALdouble c);
void goal_Source3dSOFT(void* proc, ALuint a, ALenum b, ALdouble c, ALdouble d, ALdouble e);
void goal_SourcedvSOFT(void* proc, ALuint a, ALenum b, const ALdouble* c);
void goal_GetSourcedSOFT(void* proc, ALuint a, ALenum b, ALdouble* c);
void goal_GetSource3dSOFT(void* proc, ALuint a, ALenum b, ALdouble* c, ALdouble* d, ALdouble* e);
void goal_GetSourcedvSOFT(void* proc, ALuint a, ALenum b, ALdouble* c);
void goal_Sourcei64SOFT(void* proc, ALuint a, ALenum b, ALint64SOFT c);
void goal_Source3i64SOFT(void* proc, ALuint a, ALenum b, ALint64SOFT c, ALint64SOFT d, ALint64SOFT e);
void goal_Sourcei64vSOFT(void* proc, ALuint a, ALenum b, const ALint64SOFT* c);
void goal_GetSourcei64SOFT(void* proc, ALuint a, ALenum b, ALint64SOFT* c);
void goal_GetSource3i64SOFT(void* proc, ALuint a, ALenum b, ALint64SOFT* c, ALint64SOFT* d, ALint64SOFT* e);
void goal_GetSourcei64vSOFT(void* proc, ALuint a, ALenum b, ALint64SOFT* c);
*/
import "C"

import (
	"errors"
	"log"
	"sync"
	"unsafe"
)

var (
	ErrInvalidName      = errors.New("openal: invalid name parameter")
	ErrInvalidEnum      = errors.New("openal: invalid enum parameter")
	ErrInvalidValue     = errors.New("openal: invalid value parameter")
	ErrInvalidOperation = errors.New("openal: invalid operation")
	ErrOutOfMemory      = errors.New("openal: out of memory")

	ErrCantOpenDevice         = errors.New("openal: cant open device")
	ErrCantCreateContext      = errors.New("openal: cant create context")
	ErrCantMakeContextCurrent = errors.New("openal: cant make context current")
)

var (
	errorHandlerAccess sync.RWMutex
	errorHandler       = func(e error) {
		log.Println(e)
	}
)

func SetErrorHandler(f func(error)) {
	errorHandlerAccess.Lock()
	defer errorHandlerAccess.Unlock()

	errorHandler = f
}

func ErrorHandler() func(error) {
	errorHandlerAccess.RLock()
	defer errorHandlerAccess.RUnlock()

	return errorHandler
}

func checkError() {
	cerr := C.goal_GetError(p_GetError)
	var err error
	switch cerr {
	case C.AL_NO_ERROR:
		break
	case C.AL_INVALID_NAME:
		err = ErrInvalidName
	case C.AL_INVALID_ENUM:
		err = ErrInvalidEnum
	case C.AL_INVALID_VALUE:
		err = ErrInvalidValue
	case C.AL_INVALID_OPERATION:
		err = ErrInvalidOperation
	case C.AL_OUT_OF_MEMORY:
		err = ErrOutOfMemory
	}
	if err != nil {
		ErrorHandler()(err)
	}
	return
}

var (
	p_Enable, p_Disable, p_IsEnabled, p_GetString, p_GetBooleanv,
	p_GetIntegerv, p_GetFloatv, p_GetDoublev, p_GetBoolean, p_GetInteger,
	p_GetFloat, p_GetDouble, p_IsExtensionPresent, p_GetProcAddress,
	p_GetEnumValue, p_Listenerf, p_Listener3f, p_Listenerfv, p_Listeneri,
	p_Listener3i, p_Listeneriv, p_GetListenerf, p_GetListener3f,
	p_GetListenerfv, p_GetListeneri, p_GetListener3i, p_GetListeneriv,
	p_GenSources, p_DeleteSources, p_IsSource, p_Sourcef, p_Source3f,
	p_Sourcefv, p_Sourcei, p_Source3i, p_Sourceiv, p_GetSourcef, p_GetSource3f,
	p_GetSourcefv, p_GetSourcei, p_GetSource3i, p_GetSourceiv, p_SourcePlayv,
	p_SourceStopv, p_SourceRewindv, p_SourcePausev, p_SourcePlay, p_SourceStop,
	p_SourceRewind, p_SourcePause, p_SourceQueueBuffers,
	p_SourceUnqueueBuffers, p_GenBuffers, p_DeleteBuffers, p_IsBuffer,
	p_BufferData, p_Bufferf, p_Buffer3f, p_Bufferfv, p_Bufferi, p_Buffer3i,
	p_Bufferiv, p_GetBufferf, p_GetBuffer3f, p_GetBufferfv, p_GetBufferi,
	p_GetBuffer3i, p_GetBufferiv, p_DopplerFactor, p_DopplerVelocity,
	p_SpeedOfSound, p_DistanceModel, p_GetError unsafe.Pointer
)

func alInit() {
	p_GetError = symbol("alGetError")
	p_Enable = symbol("alEnable")
	p_Disable = symbol("alDisable")
	p_IsEnabled = symbol("alIsEnabled")
	p_GetString = symbol("alGetString")
	p_GetBooleanv = symbol("alGetBooleanv")
	p_GetIntegerv = symbol("alGetIntegerv")
	p_GetFloatv = symbol("alGetFloatv")
	p_GetDoublev = symbol("alGetDoublev")
	p_GetBoolean = symbol("alGetBoolean")
	p_GetInteger = symbol("alGetInteger")
	p_GetFloat = symbol("alGetFloat")
	p_GetDouble = symbol("alGetDouble")
	p_IsExtensionPresent = symbol("alIsExtensionPresent")
	p_GetProcAddress = symbol("alGetProcAddress")
	p_GetEnumValue = symbol("alGetEnumValue")
	p_Listenerf = symbol("alListenerf")
	p_Listener3f = symbol("alListener3f")
	p_Listenerfv = symbol("alListenerfv")
	p_Listeneri = symbol("alListeneri")
	p_Listener3i = symbol("alListener3i")
	p_Listeneriv = symbol("alListeneriv")
	p_GetListenerf = symbol("alGetListenerf")
	p_GetListener3f = symbol("alGetListener3f")
	p_GetListenerfv = symbol("alGetListenerfv")
	p_GetListeneri = symbol("alGetListeneri")
	p_GetListener3i = symbol("alGetListener3i")
	p_GetListeneriv = symbol("alGetListeneriv")
	p_GenSources = symbol("alGenSources")
	p_DeleteSources = symbol("alDeleteSources")
	p_IsSource = symbol("alIsSource")
	p_Sourcef = symbol("alSourcef")
	p_Source3f = symbol("alSource3f")
	p_Sourcefv = symbol("alSourcefv")
	p_Sourcei = symbol("alSourcei")
	p_Source3i = symbol("alSource3i")
	p_Sourceiv = symbol("alSourceiv")
	p_GetSourcef = symbol("alGetSourcef")
	p_GetSource3f = symbol("alGetSource3f")
	p_GetSourcefv = symbol("alGetSourcefv")
	p_GetSourcei = symbol("alGetSourcei")
	p_GetSource3i = symbol("alGetSource3i")
	p_GetSourceiv = symbol("alGetSourceiv")
	p_SourcePlayv = symbol("alSourcePlayv")
	p_SourceStopv = symbol("alSourceStopv")
	p_SourceRewindv = symbol("alSourceRewindv")
	p_SourcePausev = symbol("alSourcePausev")
	p_SourcePlay = symbol("alSourcePlay")
	p_SourceStop = symbol("alSourceStop")
	p_SourceRewind = symbol("alSourceRewind")
	p_SourcePause = symbol("alSourcePause")
	p_SourceQueueBuffers = symbol("alSourceQueueBuffers")
	p_SourceUnqueueBuffers = symbol("alSourceUnqueueBuffers")
	p_GenBuffers = symbol("alGenBuffers")
	p_DeleteBuffers = symbol("alDeleteBuffers")
	p_IsBuffer = symbol("alIsBuffer")
	p_BufferData = symbol("alBufferData")
	p_Bufferf = symbol("alBufferf")
	p_Buffer3f = symbol("alBuffer3f")
	p_Bufferfv = symbol("alBufferfv")
	p_Bufferi = symbol("alBufferi")
	p_Buffer3i = symbol("alBuffer3i")
	p_Bufferiv = symbol("alBufferiv")
	p_GetBufferf = symbol("alGetBufferf")
	p_GetBuffer3f = symbol("alGetBuffer3f")
	p_GetBufferfv = symbol("alGetBufferfv")
	p_GetBufferi = symbol("alGetBufferi")
	p_GetBuffer3i = symbol("alGetBuffer3i")
	p_GetBufferiv = symbol("alGetBufferiv")
	p_DopplerFactor = symbol("alDopplerFactor")
	p_DopplerVelocity = symbol("alDopplerVelocity")
	p_SpeedOfSound = symbol("alSpeedOfSound")
	p_DistanceModel = symbol("alDistanceModel")
}

func (d *Device) Enable(capability int32) {
	d.alDispatch(func() {
		C.goal_Enable(p_Enable, C.ALenum(capability))
	})
	return
}

func (d *Device) Disable(capability int32) {
	d.alDispatch(func() {
		C.goal_Disable(p_Disable, C.ALenum(capability))
	})
	return
}

func (d *Device) IsEnabled(capability int32) bool {
	var ret bool
	d.alDispatch(func() {
		ret = int32(C.goal_IsEnabled(p_IsEnabled, C.ALenum(capability))) == TRUE
	})
	return ret
}

func (d *Device) GetString(param int32) string {
	var ret string
	d.alDispatch(func() {
		cstr := C.goal_GetString(p_GetString, C.ALenum(param))
		ret = C.GoString((*C.char)(unsafe.Pointer(cstr)))
	})
	return ret
}

func (d *Device) GetBooleanv(param int32, values *uint8) {
	d.alDispatch(func() {
		C.goal_GetBooleanv(
			p_GetBooleanv,
			C.ALenum(param),
			(*C.ALboolean)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetIntegerv(param int32, values *int32) {
	d.alDispatch(func() {
		C.goal_GetIntegerv(
			p_GetIntegerv,
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetFloatv(param int32, values *float32) {
	d.alDispatch(func() {
		C.goal_GetFloatv(
			p_GetFloatv,
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetDoublev(param int32, values *float64) {
	d.alDispatch(func() {
		C.goal_GetDoublev(
			p_GetDoublev,
			C.ALenum(param),
			(*C.ALdouble)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetBoolean(param int32) bool {
	var ret bool
	d.alDispatch(func() {
		ret = int32(C.goal_GetBoolean(p_GetBoolean, C.ALenum(param))) == TRUE
	})
	return ret
}

func (d *Device) GetInteger(param int32) int32 {
	var ret int32
	d.alDispatch(func() {
		ret = int32(C.goal_GetInteger(p_GetInteger, C.ALenum(param)))
	})
	return ret
}

func (d *Device) GetFloat(param int32) float32 {
	var ret float32
	d.alDispatch(func() {
		ret = float32(C.goal_GetFloat(p_GetFloat, C.ALenum(param)))
	})
	return ret
}

func (d *Device) GetDouble(param int32) float64 {
	var ret float64
	d.alDispatch(func() {
		ret = float64(C.goal_GetDouble(p_GetDouble, C.ALenum(param)))
	})
	return ret
}

func (d *Device) IsExtensionPresent(name string) bool {
	var ret bool
	d.alDispatch(func() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		ret = int32(C.goal_IsExtensionPresent(p_IsExtensionPresent, (*C.ALchar)(unsafe.Pointer(cname)))) == TRUE
	})
	return ret
}

func (d *Device) GetProcAddress(name string) unsafe.Pointer {
	var ret unsafe.Pointer
	d.alDispatch(func() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		ret = C.goal_GetProcAddress(p_GetProcAddress, (*C.ALchar)(unsafe.Pointer(cname)))
	})
	return ret
}

func (d *Device) GetEnumValue(name string) int32 {
	var ret int32
	d.alDispatch(func() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))

		ret = int32(C.goal_GetEnumValue(p_GetEnumValue, (*C.ALchar)(unsafe.Pointer(cname))))
	})
	return ret
}

func (d *Device) Listenerf(param int32, value float32) {
	d.alDispatch(func() {
		C.goal_Listenerf(
			p_Listenerf,
			C.ALenum(param),
			C.ALfloat(value),
		)
	})
	return
}

func (d *Device) Listener3f(param int32, value1, value2, value3 float32) {
	d.alDispatch(func() {
		C.goal_Listener3f(
			p_Listener3f,
			C.ALenum(param),
			C.ALfloat(value1),
			C.ALfloat(value2),
			C.ALfloat(value3),
		)
	})
	return
}

func (d *Device) Listenerfv(param int32, values *float32) {
	d.alDispatch(func() {
		C.goal_Listenerfv(
			p_Listenerfv,
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) Listeneri(param int32, value int32) {
	d.alDispatch(func() {
		C.goal_Listeneri(
			p_Listeneri,
			C.ALenum(param),
			C.ALint(value),
		)
	})
	return
}

func (d *Device) Listener3i(param int32, value1, value2, value3 int32) {
	d.alDispatch(func() {
		C.goal_Listener3i(
			p_Listener3i,
			C.ALenum(param),
			C.ALint(value1),
			C.ALint(value2),
			C.ALint(value3),
		)
	})
	return
}

func (d *Device) Listeneriv(param int32, values *int32) {
	d.alDispatch(func() {
		C.goal_Listeneriv(
			p_Listeneriv,
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetListenerf(param int32, value *float32) {
	d.alDispatch(func() {
		C.goal_GetListenerf(
			p_GetListenerf,
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(value)),
		)
	})
	return
}

func (d *Device) GetListener3f(param int32, value1, value2, value3 *float32) {
	d.alDispatch(func() {
		C.goal_GetListener3f(
			p_GetListener3f,
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(value1)),
			(*C.ALfloat)(unsafe.Pointer(value2)),
			(*C.ALfloat)(unsafe.Pointer(value3)),
		)
	})
	return
}

func (d *Device) GetListenerfv(param int32, values *float32) {
	d.alDispatch(func() {
		C.goal_GetListenerfv(
			p_GetListenerfv,
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetListeneri(param int32, value *int32) {
	d.alDispatch(func() {
		C.goal_GetListeneri(
			p_GetListeneri,
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(value)),
		)
	})
	return
}

func (d *Device) GetListener3i(param int32, value1, value2, value3 *int32) {
	d.alDispatch(func() {
		C.goal_GetListener3i(
			p_GetListener3i,
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(value1)),
			(*C.ALint)(unsafe.Pointer(value2)),
			(*C.ALint)(unsafe.Pointer(value3)),
		)
	})
	return
}

func (d *Device) GetListeneriv(param int32, values *int32) {
	d.alDispatch(func() {
		C.goal_GetListeneriv(
			p_GetListeneriv,
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GenSources(n int32, sources *uint32) {
	d.alDispatch(func() {
		C.goal_GenSources(
			p_GenSources,
			C.ALsizei(n),
			(*C.ALuint)(unsafe.Pointer(sources)),
		)
	})
	return
}

func (d *Device) DeleteSources(n int32, sources *uint32) {
	d.alDispatch(func() {
		C.goal_DeleteSources(
			p_DeleteSources,
			C.ALsizei(n),
			(*C.ALuint)(unsafe.Pointer(sources)),
		)
	})
	return
}

func (d *Device) IsSource(source uint32) bool {
	var ret bool
	d.alDispatch(func() {
		ret = int32(C.goal_IsSource(p_IsSource, C.ALuint(source))) == TRUE
	})
	return ret
}

func (d *Device) Sourcef(source uint32, param int32, value float32) {
	d.alDispatch(func() {
		C.goal_Sourcef(
			p_Sourcef,
			C.ALuint(source),
			C.ALenum(param),
			C.ALfloat(value),
		)
	})
	return
}

func (d *Device) Source3f(source uint32, param int32, value1, value2, value3 float32) {
	d.alDispatch(func() {
		C.goal_Source3f(
			p_Source3f,
			C.ALuint(source),
			C.ALenum(param),
			C.ALfloat(value1),
			C.ALfloat(value2),
			C.ALfloat(value3),
		)
	})
	return
}

func (d *Device) Sourcefv(source uint32, param int32, values *float32) {
	d.alDispatch(func() {
		C.goal_Sourcefv(
			p_Sourcefv,
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) Sourcei(source uint32, param int32, value int32) {
	d.alDispatch(func() {
		C.goal_Sourcei(
			p_Sourcei,
			C.ALuint(source),
			C.ALenum(param),
			C.ALint(value),
		)
	})
	return
}

func (d *Device) Source3i(source uint32, param int32, value1, value2, value3 int32) {
	d.alDispatch(func() {
		C.goal_Source3i(
			p_Source3i,
			C.ALuint(source),
			C.ALenum(param),
			C.ALint(value1),
			C.ALint(value2),
			C.ALint(value3),
		)
	})
	return
}

func (d *Device) Sourceiv(source uint32, param int32, values *int32) {
	d.alDispatch(func() {
		C.goal_Sourceiv(
			p_Sourceiv,
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetSourcef(source uint32, param int32, value *float32) {
	d.alDispatch(func() {
		C.goal_GetSourcef(
			p_GetSourcef,
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(value)),
		)
	})
	return
}

func (d *Device) GetSource3f(source uint32, param int32, value1, value2, value3 *float32) {
	d.alDispatch(func() {
		C.goal_GetSource3f(
			p_GetSource3f,
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(value1)),
			(*C.ALfloat)(unsafe.Pointer(value2)),
			(*C.ALfloat)(unsafe.Pointer(value3)),
		)
	})
	return
}

func (d *Device) GetSourcefv(source uint32, param int32, values *float32) {
	d.alDispatch(func() {
		C.goal_GetSourcefv(
			p_GetSourcefv,
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetSourcei(source uint32, param int32, value *int32) {
	d.alDispatch(func() {
		C.goal_GetSourcei(
			p_GetSourcei,
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(value)),
		)
	})
	return
}

func (d *Device) GetSource3i(source uint32, param int32, value1, value2, value3 *int32) {
	d.alDispatch(func() {
		C.goal_GetSource3i(
			p_GetSource3i,
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(value1)),
			(*C.ALint)(unsafe.Pointer(value2)),
			(*C.ALint)(unsafe.Pointer(value3)),
		)
	})
	return
}

func (d *Device) GetSourceiv(source uint32, param int32, values *int32) {
	d.alDispatch(func() {
		C.goal_GetSourceiv(
			p_GetSourceiv,
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) SourcePlayv(sources []uint32) {
	d.alDispatch(func() {
		C.goal_SourcePlayv(
			p_SourcePlayv,
			C.ALsizei(len(sources)),
			(*C.ALuint)(unsafe.Pointer(&sources[0])),
		)
	})
	return
}

func (d *Device) SourceStopv(sources []uint32) {
	d.alDispatch(func() {
		C.goal_SourceStopv(
			p_SourceStopv,
			C.ALsizei(len(sources)),
			(*C.ALuint)(unsafe.Pointer(&sources[0])),
		)
	})
	return
}

func (d *Device) SourceRewindv(sources []uint32) {
	d.alDispatch(func() {
		C.goal_SourceRewindv(
			p_SourceRewindv,
			C.ALsizei(len(sources)),
			(*C.ALuint)(unsafe.Pointer(&sources[0])),
		)
	})
	return
}

func (d *Device) SourcePausev(sources []uint32) {
	d.alDispatch(func() {
		C.goal_SourcePausev(
			p_SourcePausev,
			C.ALsizei(len(sources)),
			(*C.ALuint)(unsafe.Pointer(&sources[0])),
		)
	})
	return
}

func (d *Device) SourcePlay(source uint32) {
	d.alDispatch(func() {
		C.goal_SourcePlay(p_SourcePlay, C.ALuint(source))
	})
	return
}

func (d *Device) SourceStop(source uint32) {
	d.alDispatch(func() {
		C.goal_SourceStop(p_SourceStop, C.ALuint(source))
	})
	return
}

func (d *Device) SourceRewind(source uint32) {
	d.alDispatch(func() {
		C.goal_SourceRewind(p_SourceRewind, C.ALuint(source))
	})
	return
}

func (d *Device) SourcePause(source uint32) {
	d.alDispatch(func() {
		C.goal_SourcePause(p_SourcePause, C.ALuint(source))
	})
	return
}

func (d *Device) SourceQueueBuffers(source uint32, buffers []uint32) {
	d.alDispatch(func() {
		C.goal_SourceQueueBuffers(
			p_SourceQueueBuffers,
			C.ALuint(source),
			C.ALsizei(len(buffers)),
			(*C.ALuint)(unsafe.Pointer(&buffers[0])),
		)
	})
	return
}

func (d *Device) SourceUnqueueBuffers(source uint32, buffers []uint32) {
	d.alDispatch(func() {
		C.goal_SourceUnqueueBuffers(
			p_SourceUnqueueBuffers,
			C.ALuint(source),
			C.ALsizei(len(buffers)),
			(*C.ALuint)(unsafe.Pointer(&buffers[0])),
		)
	})
	return
}

func (d *Device) GenBuffers(n int32, buffers *uint32) {
	d.alDispatch(func() {
		C.goal_GenBuffers(
			p_GenBuffers,
			C.ALsizei(n),
			(*C.ALuint)(unsafe.Pointer(buffers)),
		)
	})
	return
}

func (d *Device) DeleteBuffers(n int32, buffers *uint32) {
	d.alDispatch(func() {
		C.goal_DeleteBuffers(
			p_DeleteBuffers,
			C.ALsizei(n),
			(*C.ALuint)(unsafe.Pointer(buffers)),
		)
	})
	return
}

func (d *Device) IsBuffer(buffer uint32) bool {
	var ret bool
	d.alDispatch(func() {
		ret = int32(C.goal_IsBuffer(p_IsBuffer, C.ALuint(buffer))) == TRUE
	})
	return ret
}

func (d *Device) BufferData(buffer uint32, format int32, data unsafe.Pointer, size, freq int32) {
	d.alDispatch(func() {
		C.goal_BufferData(
			p_BufferData,
			C.ALuint(buffer),
			C.ALenum(format),
			data,
			C.ALsizei(size),
			C.ALsizei(freq),
		)
	})
	return
}

func (d *Device) Bufferf(buffer uint32, param int32, value float32) {
	d.alDispatch(func() {
		C.goal_Bufferf(
			p_Bufferf,
			C.ALuint(buffer),
			C.ALenum(param),
			C.ALfloat(value),
		)
	})
	return
}

func (d *Device) Buffer3f(buffer uint32, param int32, value1, value2, value3 float32) {
	d.alDispatch(func() {
		C.goal_Buffer3f(
			p_Buffer3f,
			C.ALuint(buffer),
			C.ALenum(param),
			C.ALfloat(value1),
			C.ALfloat(value2),
			C.ALfloat(value3),
		)
	})
	return
}

func (d *Device) Bufferfv(buffer uint32, param int32, values *float32) {
	d.alDispatch(func() {
		C.goal_Bufferfv(
			p_Bufferfv,
			C.ALuint(buffer),
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) Bufferi(buffer uint32, param int32, value int32) {
	d.alDispatch(func() {
		C.goal_Bufferi(
			p_Bufferi,
			C.ALuint(buffer),
			C.ALenum(param),
			C.ALint(value),
		)
	})
	return
}

func (d *Device) Buffer3i(buffer uint32, param int32, value1, value2, value3 int32) {
	d.alDispatch(func() {
		C.goal_Buffer3i(
			p_Buffer3i,
			C.ALuint(buffer),
			C.ALenum(param),
			C.ALint(value1),
			C.ALint(value2),
			C.ALint(value3),
		)
	})
	return
}

func (d *Device) Bufferiv(buffer uint32, param int32, values *int32) {
	d.alDispatch(func() {
		C.goal_Bufferiv(
			p_Bufferiv,
			C.ALuint(buffer),
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetBufferf(buffer uint32, param int32, value *float32) {
	d.alDispatch(func() {
		C.goal_GetBufferf(
			p_GetBufferf,
			C.ALuint(buffer),
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(value)),
		)
	})
	return
}

func (d *Device) GetBuffer3f(buffer uint32, param int32, value1, value2, value3 *float32) {
	d.alDispatch(func() {
		C.goal_GetBuffer3f(
			p_GetBuffer3f,
			C.ALuint(buffer),
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(value1)),
			(*C.ALfloat)(unsafe.Pointer(value2)),
			(*C.ALfloat)(unsafe.Pointer(value3)),
		)
	})
	return
}

func (d *Device) GetBufferfv(buffer uint32, param int32, values *float32) {
	d.alDispatch(func() {
		C.goal_GetBufferfv(
			p_GetBufferfv,
			C.ALuint(buffer),
			C.ALenum(param),
			(*C.ALfloat)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) GetBufferi(buffer uint32, param int32, value *int32) {
	d.alDispatch(func() {
		C.goal_GetBufferi(
			p_GetBufferi,
			C.ALuint(buffer),
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(value)),
		)
	})
	return
}

func (d *Device) GetBuffer3i(buffer uint32, param int32, value1, value2, value3 *int32) {
	d.alDispatch(func() {
		C.goal_GetBuffer3i(
			p_GetBuffer3i,
			C.ALuint(buffer),
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(value1)),
			(*C.ALint)(unsafe.Pointer(value2)),
			(*C.ALint)(unsafe.Pointer(value3)),
		)
	})
	return
}

func (d *Device) GetBufferiv(buffer uint32, param int32, values *int32) {
	d.alDispatch(func() {
		C.goal_GetBufferiv(
			p_GetBufferiv,
			C.ALuint(buffer),
			C.ALenum(param),
			(*C.ALint)(unsafe.Pointer(values)),
		)
	})
	return
}

func (d *Device) DopplerFactor(value float32) {
	d.alDispatch(func() {
		C.goal_DopplerFactor(p_DopplerFactor, C.ALfloat(value))
	})
	return
}

func (d *Device) DopplerVelocity(value float32) {
	d.alDispatch(func() {
		C.goal_DopplerVelocity(p_DopplerVelocity, C.ALfloat(value))
	})
	return
}

func (d *Device) SpeedOfSound(value float32) {
	d.alDispatch(func() {
		C.goal_SpeedOfSound(p_SpeedOfSound, C.ALfloat(value))
	})
	return
}

func (d *Device) DistanceModel(distanceModel int32) {
	d.alDispatch(func() {
		C.goal_DistanceModel(p_DistanceModel, C.ALenum(distanceModel))
	})
	return
}

// Begin extensions -- do not place non-extension functions below here.

func (d *Device) BufferDataStatic(buffer uint32, format int32, data unsafe.Pointer, length, freq int32) {
	d.alDispatch(func() {
		C.goal_BufferDataStatic(
			d.alExtensionInit(&d.p_BufferDataStatic, "alBufferDataStatic"),
			C.ALuint(buffer),
			C.ALenum(format),
			data,
			C.ALsizei(length),
			C.ALsizei(freq),
		)
	})
}

func (d *Device) BufferSubDataSOFT(buffer uint32, format int32, data unsafe.Pointer, offset, length int32) {
	d.alDispatch(func() {
		C.goal_BufferSubDataSOFT(
			d.alExtensionInit(&d.p_BufferSubDataSOFT, "alBufferSubDataSOFT"),
			C.ALuint(buffer),
			C.ALenum(format),
			data,
			C.ALsizei(offset),
			C.ALsizei(length),
		)
	})
}

func (d *Device) BufferSamplesSOFT(buffer uint32, sampleRate uint32, internalFormat, samples, channels, t int32, data unsafe.Pointer) {
	d.alDispatch(func() {
		C.goal_BufferSamplesSOFT(
			d.alExtensionInit(&d.p_BufferSamplesSOFT, "alBufferSamplesSOFT"),
			C.ALuint(buffer),
			C.ALuint(sampleRate),
			C.ALenum(internalFormat),
			C.ALsizei(samples),
			C.ALenum(channels),
			C.ALenum(t),
			data,
		)
	})
}

func (d *Device) BufferSubSamplesSOFT(buffer uint32, offset, samples, channels, t int32, data unsafe.Pointer) {
	d.alDispatch(func() {
		C.goal_BufferSubSamplesSOFT(
			d.alExtensionInit(&d.p_BufferSubSamplesSOFT, "alBufferSubSamplesSOFT"),
			C.ALuint(buffer),
			C.ALsizei(offset),
			C.ALsizei(samples),
			C.ALenum(channels),
			C.ALenum(t),
			data,
		)
	})
}

func (d *Device) GetBufferSamplesSOFT(buffer uint32, offset, samples, channels, t int32, data unsafe.Pointer) {
	d.alDispatch(func() {
		C.goal_GetBufferSamplesSOFT(
			d.alExtensionInit(&d.p_GetBufferSamplesSOFT, "alGetBufferSamplesSOFT"),
			C.ALuint(buffer),
			C.ALsizei(offset),
			C.ALsizei(samples),
			C.ALenum(channels),
			C.ALenum(t),
			data,
		)
	})
}

func (d *Device) IsBufferFormatSupportedSOFT(format int32) bool {
	var ret bool
	d.alDispatch(func() {
		ret = int32(C.goal_IsBufferFormatSupportedSOFT(
			d.alExtensionInit(&d.p_IsBufferFormatSupportedSOFT, "alIsBufferFormatSupportedSOFT"),
			C.ALenum(format),
		)) == TRUE
	})
	return ret
}

func (d *Device) SourcedSOFT(source uint32, param int32, value float64) {
	d.alDispatch(func() {
		C.goal_SourcedSOFT(
			d.alExtensionInit(&d.p_SourcedSOFT, "alSourcedSOFT"),
			C.ALuint(source),
			C.ALenum(param),
			C.ALdouble(value),
		)
	})
}

func (d *Device) Source3dSOFT(source uint32, param int32, value1, value2, value3 float64) {
	d.alDispatch(func() {
		C.goal_Source3dSOFT(
			d.alExtensionInit(&d.p_Source3dSOFT, "alSource3dSOFT"),
			C.ALuint(source),
			C.ALenum(param),
			C.ALdouble(value1),
			C.ALdouble(value2),
			C.ALdouble(value3),
		)
	})
}

func (d *Device) SourcedvSOFT(source uint32, param int32, values *float64) {
	d.alDispatch(func() {
		C.goal_SourcedvSOFT(
			d.alExtensionInit(&d.p_SourcedvSOFT, "alSourcedvSOFT"),
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALdouble)(unsafe.Pointer(values)),
		)
	})
}

func (d *Device) GetSourcedSOFT(source uint32, param int32, value *float64) {
	d.alDispatch(func() {
		C.goal_GetSourcedSOFT(
			d.alExtensionInit(&d.p_GetSourcedSOFT, "alGetSourcedSOFT"),
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALdouble)(unsafe.Pointer(value)),
		)
	})
}

func (d *Device) GetSource3dSOFT(source uint32, param int32, value1, value2, value3 *float64) {
	d.alDispatch(func() {
		C.goal_GetSource3dSOFT(
			d.alExtensionInit(&d.p_GetSource3dSOFT, "alGetSource3dSOFT"),
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALdouble)(unsafe.Pointer(value1)),
			(*C.ALdouble)(unsafe.Pointer(value2)),
			(*C.ALdouble)(unsafe.Pointer(value3)),
		)
	})
}

func (d *Device) GetSourcedvSOFT(source uint32, param int32, values *float64) {
	d.alDispatch(func() {
		C.goal_GetSourcedvSOFT(
			d.alExtensionInit(&d.p_GetSourcedvSOFT, "alGetSourcedvSOFT"),
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALdouble)(unsafe.Pointer(values)),
		)
	})
}

func (d *Device) Sourcei64SOFT(source uint32, param int32, value int64) {
	d.alDispatch(func() {
		C.goal_Sourcei64SOFT(
			d.alExtensionInit(&d.p_Sourcei64SOFT, "alSourcei64SOFT"),
			C.ALuint(source),
			C.ALenum(param),
			C.ALint64SOFT(value),
		)
	})
}

func (d *Device) Source3i64SOFT(source uint32, param int32, value1, value2, value3 int64) {
	d.alDispatch(func() {
		C.goal_Source3i64SOFT(
			d.alExtensionInit(&d.p_Source3i64SOFT, "Source3i64SOFT"),
			C.ALuint(source),
			C.ALenum(param),
			C.ALint64SOFT(value1),
			C.ALint64SOFT(value2),
			C.ALint64SOFT(value3),
		)
	})
}

func (d *Device) Sourcei64vSOFT(source uint32, param int32, values *int64) {
	d.alDispatch(func() {
		C.goal_Sourcei64vSOFT(
			d.alExtensionInit(&d.p_Sourcei64vSOFT, "Sourcei64vSOFT"),
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALint64SOFT)(unsafe.Pointer(values)),
		)
	})
}

func (d *Device) GetSourcei64SOFT(source uint32, param int32, value *int64) {
	d.alDispatch(func() {
		C.goal_GetSourcei64SOFT(
			d.alExtensionInit(&d.p_GetSourcei64SOFT, "GetSourcei64SOFT"),
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALint64SOFT)(unsafe.Pointer(value)),
		)
	})
}

func (d *Device) GetSource3i64SOFT(source uint32, param int32, value1, value2, value3 *int64) {
	d.alDispatch(func() {
		C.goal_GetSource3i64SOFT(
			d.alExtensionInit(&d.p_GetSource3i64SOFT, "GetSource3i64SOFT"),
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALint64SOFT)(unsafe.Pointer(value1)),
			(*C.ALint64SOFT)(unsafe.Pointer(value2)),
			(*C.ALint64SOFT)(unsafe.Pointer(value3)),
		)
	})
}

func (d *Device) GetSourcei64vSOFT(source uint32, param int32, values *int64) {
	d.alDispatch(func() {
		C.goal_GetSourcei64vSOFT(
			d.alExtensionInit(&d.p_GetSourcei64vSOFT, "GetSourcei64vSOFT"),
			C.ALuint(source),
			C.ALenum(param),
			(*C.ALint64SOFT)(unsafe.Pointer(values)),
		)
	})
}
