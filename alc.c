// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "AL/alc.h"

ALCcontext* goal_alcCreateContext(void* proc, ALCdevice* a, const ALCint* b) {
	LPALCCREATECONTEXT fn = (LPALCCREATECONTEXT)proc;
	return fn(a, b);
}

ALCboolean goal_alcMakeContextCurrent(void* proc, ALCcontext* a) {
	LPALCMAKECONTEXTCURRENT fn = (LPALCMAKECONTEXTCURRENT)proc;
	return fn(a);
}

void goal_alcProcessContext(void* proc, ALCcontext* a) {
	LPALCPROCESSCONTEXT fn = (LPALCPROCESSCONTEXT)proc;
	fn(a);
}

void goal_alcSuspendContext(void* proc, ALCcontext* a) {
	LPALCSUSPENDCONTEXT fn = (LPALCSUSPENDCONTEXT)proc;
	fn(a);
}

void goal_alcDestroyContext(void* proc, ALCcontext* a) {
	LPALCDESTROYCONTEXT fn = (LPALCDESTROYCONTEXT)proc;
	fn(a);
}

ALCcontext* goal_alcGetCurrentContext(void* proc) {
	LPALCGETCURRENTCONTEXT fn = (LPALCGETCURRENTCONTEXT)proc;
	return fn();
}

ALCdevice* goal_alcGetContextsDevice(void* proc, ALCcontext* a) {
	LPALCGETCONTEXTSDEVICE fn = (LPALCGETCONTEXTSDEVICE)proc;
	return fn(a);
}

ALCdevice* goal_alcOpenDevice(void* proc, const ALCchar* a) {
	LPALCOPENDEVICE fn = (LPALCOPENDEVICE)proc;
	return fn(a);
}

ALCboolean goal_alcCloseDevice(void* proc, ALCdevice* a) {
	LPALCCLOSEDEVICE fn = (LPALCCLOSEDEVICE)proc;
	return fn(a);
}

ALCenum goal_alcGetError(void* proc, ALCdevice* a) {
	LPALCGETERROR fn = (LPALCGETERROR)proc;
	return fn(a);
}

ALCboolean goal_alcIsExtensionPresent(void* proc, ALCdevice* a, const ALCchar* b) {
	LPALCISEXTENSIONPRESENT fn = (LPALCISEXTENSIONPRESENT)proc;
	return fn(a, b);
}

void* goal_alcGetProcAddress(void* proc, ALCdevice* a, const ALCchar* b) {
	LPALCGETPROCADDRESS fn = (LPALCGETPROCADDRESS)proc;
	return fn(a, b);
}

ALCenum goal_alcGetEnumValue(void* proc, ALCdevice* a, const ALCchar* b) {
	LPALCGETENUMVALUE fn = (LPALCGETENUMVALUE)proc;
	return fn(a, b);
}

const ALCchar* goal_alcGetString(void* proc, ALCdevice* a, ALCenum b) {
	LPALCGETSTRING fn = (LPALCGETSTRING)proc;
	return fn(a, b);
}

void goal_alcGetIntegerv(void* proc, ALCdevice* a, ALCenum b, ALCsizei c, ALCint* d) {
	LPALCGETINTEGERV fn = (LPALCGETINTEGERV)proc;
	fn(a, b, c, d);
}

ALCdevice* goal_alcCaptureOpenDevice(void* proc, const ALCchar* a, ALCuint b, ALCenum c, ALCsizei d) {
	LPALCCAPTUREOPENDEVICE fn = (LPALCCAPTUREOPENDEVICE)proc;
	return fn(a, b, c, d);
}

ALCboolean goal_alcCaptureCloseDevice(void* proc, ALCdevice* a) {
	LPALCCAPTURECLOSEDEVICE fn = (LPALCCAPTURECLOSEDEVICE)proc;
	return fn(a);
}

void goal_alcCaptureStart(void* proc, ALCdevice* a) {
	LPALCCAPTURESTART fn = (LPALCCAPTURESTART)proc;
	fn(a);
}

void goal_alcCaptureStop(void* proc, ALCdevice* a) {
	LPALCCAPTURESTOP fn = (LPALCCAPTURESTOP)proc;
	fn(a);
}

void goal_alcCaptureSamples(void* proc, ALCdevice* a, ALCvoid* b, ALCsizei c) {
	LPALCCAPTURESAMPLES fn = (LPALCCAPTURESAMPLES)proc;
	fn(a, b, c);
}
