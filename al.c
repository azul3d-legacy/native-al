// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "AL/al.h"
#include "AL/alext.h"

void goal_Enable(void* proc, ALenum a) {
	LPALENABLE fn = (LPALENABLE)proc;
	fn(a);
}

void goal_Disable(void* proc, ALenum a) {
	LPALDISABLE fn = (LPALDISABLE)proc;
	fn(a);
}

ALboolean goal_IsEnabled(void* proc, ALenum a) {
	LPALISENABLED fn = (LPALISENABLED)proc;
	return fn(a);
}

const ALchar* goal_GetString(void* proc, ALenum a) {
	LPALGETSTRING fn = (LPALGETSTRING)proc;
	return fn(a);
}

void goal_GetBooleanv(void* proc, ALenum a, ALboolean* b) {
	LPALGETBOOLEANV fn = (LPALGETBOOLEANV)proc;
	fn(a, b);
}

void goal_GetIntegerv(void* proc, ALenum a, ALint* b) {
	LPALGETINTEGERV fn = (LPALGETINTEGERV)proc;
	fn(a, b);
}

void goal_GetFloatv(void* proc, ALenum a, ALfloat* b) {
	LPALGETFLOATV fn = (LPALGETFLOATV)proc;
	fn(a, b);
}

void goal_GetDoublev(void* proc, ALenum a, ALdouble* b) {
	LPALGETDOUBLEV fn = (LPALGETDOUBLEV)proc;
	fn(a, b);
}

ALboolean goal_GetBoolean(void* proc, ALenum a) {
	LPALGETBOOLEAN fn = (LPALGETBOOLEAN)proc;
	return fn(a);
}

ALint goal_GetInteger(void* proc, ALenum a) {
	LPALGETINTEGER fn = (LPALGETINTEGER)proc;
	return fn(a);
}

ALfloat goal_GetFloat(void* proc, ALenum a) {
	LPALGETFLOAT fn = (LPALGETFLOAT)proc;
	return fn(a);
}

ALdouble goal_GetDouble(void* proc, ALenum a) {
	LPALGETDOUBLE fn = (LPALGETDOUBLE)proc;
	return fn(a);
}

ALenum goal_GetError(void* proc) {
	LPALGETERROR fn = (LPALGETERROR)proc;
	return fn();
}

ALboolean goal_IsExtensionPresent(void* proc, const ALchar* a) {
	LPALISEXTENSIONPRESENT fn = (LPALISEXTENSIONPRESENT)proc;
	return fn(a);
}

void* goal_GetProcAddress(void* proc, const ALchar* a) {
	LPALGETPROCADDRESS fn = (LPALGETPROCADDRESS)proc;
	return fn(a);
}

ALenum goal_GetEnumValue(void* proc, const ALchar* a) {
	LPALGETENUMVALUE fn = (LPALGETENUMVALUE)proc;
	return fn(a);
}

void goal_Listenerf(void* proc, ALenum a, ALfloat b) {
	LPALLISTENERF fn = (LPALLISTENERF)proc;
	fn(a, b);
}

void goal_Listener3f(void* proc, ALenum a, ALfloat b, ALfloat c, ALfloat d) {
	LPALLISTENER3F fn = (LPALLISTENER3F)proc;
	fn(a, b, c, d);
}

void goal_Listenerfv(void* proc, ALenum a, ALfloat* b) {
	LPALLISTENERFV fn = (LPALLISTENERFV)proc;
	fn(a, b);
}

void goal_Listeneri(void* proc, ALenum a, ALint b) {
	LPALLISTENERI fn = (LPALLISTENERI)proc;
	fn(a, b);
}

void goal_Listener3i(void* proc, ALenum a, ALint b, ALint c, ALint d) {
	LPALLISTENER3I fn = (LPALLISTENER3I)proc;
	fn(a, b, c, d);
}

void goal_Listeneriv(void* proc, ALenum a, ALint* b) {
	LPALLISTENERIV fn = (LPALLISTENERIV)proc;
	fn(a, b);
}

void goal_GetListenerf(void* proc, ALenum a, ALfloat* b) {
	LPALGETLISTENERF fn = (LPALGETLISTENERF)proc;
	fn(a, b);
}

void goal_GetListener3f(void* proc, ALenum a, ALfloat* b, ALfloat* c, ALfloat* d) {
	LPALGETLISTENER3F fn = (LPALGETLISTENER3F)proc;
	fn(a, b, c, d);
}

void goal_GetListenerfv(void* proc, ALenum a, ALfloat* b) {
	LPALGETLISTENERFV fn = (LPALGETLISTENERFV)proc;
	fn(a, b);
}

void goal_GetListeneri(void* proc, ALenum a, ALint* b) {
	LPALGETLISTENERI fn = (LPALGETLISTENERI)proc;
	fn(a, b);
}

void goal_GetListener3i(void* proc, ALenum a, ALint* b, ALint* c, ALint* d) {
	LPALGETLISTENER3I fn = (LPALGETLISTENER3I)proc;
	fn(a, b, c, d);
}

void goal_GetListeneriv(void* proc, ALenum a, ALint* b) {
	LPALGETLISTENERIV fn = (LPALGETLISTENERIV)proc;
	fn(a, b);
}

void goal_GenSources(void* proc, ALsizei a, ALuint* b) {
	LPALGENSOURCES fn = (LPALGENSOURCES)proc;
	fn(a, b);
}

void goal_DeleteSources(void* proc, ALsizei a, ALuint* b) {
	LPALDELETESOURCES fn = (LPALDELETESOURCES)proc;
	fn(a, b);
}

ALboolean goal_IsSource(void* proc, ALuint a) {
	LPALISSOURCE fn = (LPALISSOURCE)proc;
	return fn(a);
}

void goal_Sourcef(void* proc, ALuint a, ALenum b, ALfloat c) {
	LPALSOURCEF fn = (LPALSOURCEF)proc;
	fn(a, b, c);
}

void goal_Source3f(void* proc, ALuint a, ALenum b, ALfloat c, ALfloat d, ALfloat e) {
	LPALSOURCE3F fn = (LPALSOURCE3F)proc;
	fn(a, b, c, d, e);
}

void goal_Sourcefv(void* proc, ALuint a, ALenum b, ALfloat* c) {
	LPALSOURCEFV fn = (LPALSOURCEFV)proc;
	fn(a, b, c);
}

void goal_Sourcei(void* proc, ALuint a, ALenum b, ALint c) {
	LPALSOURCEI fn = (LPALSOURCEI)proc;
	fn(a, b, c);
}

void goal_Source3i(void* proc, ALuint a, ALenum b, ALint c, ALint d, ALint e) {
	LPALSOURCE3I fn = (LPALSOURCE3I)proc;
	fn(a, b, c, d, e);
}

void goal_Sourceiv(void* proc, ALuint a, ALenum b, ALint* c) {
	LPALSOURCEIV fn = (LPALSOURCEIV)proc;
	fn(a, b, c);
}

void goal_GetSourcef(void* proc, ALuint a, ALenum b, ALfloat* c) {
	LPALGETSOURCEF fn = (LPALGETSOURCEF)proc;
	fn(a, b, c);
}

void goal_GetSource3f(void* proc, ALuint a, ALenum b, ALfloat* c, ALfloat* d, ALfloat* e) {
	LPALGETSOURCE3F fn = (LPALGETSOURCE3F)proc;
	fn(a, b, c, d, e);
}

void goal_GetSourcefv(void* proc, ALuint a, ALenum b, ALfloat* c) {
	LPALGETSOURCEFV fn = (LPALGETSOURCEFV)proc;
	fn(a, b, c);
}

void goal_GetSourcei(void* proc, ALuint a, ALenum b, ALint* c) {
	LPALGETSOURCEI fn = (LPALGETSOURCEI)proc;
	fn(a, b, c);
}

void goal_GetSource3i(void* proc, ALuint a, ALenum b, ALint* c, ALint* d, ALint* e) {
	LPALGETSOURCE3I fn = (LPALGETSOURCE3I)proc;
	fn(a, b, c, d, e);
}

void goal_GetSourceiv(void* proc, ALuint a, ALenum b, ALint* c) {
	LPALGETSOURCEIV fn = (LPALGETSOURCEIV)proc;
	fn(a, b, c);
}

void goal_SourcePlayv(void* proc, ALsizei a, ALuint* b) {
	LPALSOURCEPLAYV fn = (LPALSOURCEPLAYV)proc;
	fn(a, b);
}

void goal_SourceStopv(void* proc, ALsizei a, ALuint* b) {
	LPALSOURCESTOPV fn = (LPALSOURCESTOPV)proc;
	fn(a, b);
}

void goal_SourceRewindv(void* proc, ALsizei a, ALuint* b) {
	LPALSOURCEREWINDV fn = (LPALSOURCEREWINDV)proc;
	fn(a, b);
}

void goal_SourcePausev(void* proc, ALsizei a, ALuint* b) {
	LPALSOURCEPAUSEV fn = (LPALSOURCEPAUSEV)proc;
	fn(a, b);
}

void goal_SourcePlay(void* proc, ALuint a) {
	LPALSOURCEPLAY fn = (LPALSOURCEPLAY)proc;
	fn(a);
}

void goal_SourceStop(void* proc, ALuint a) {
	LPALSOURCESTOP fn = (LPALSOURCESTOP)proc;
	fn(a);
}

void goal_SourceRewind(void* proc, ALuint a) {
	LPALSOURCEREWIND fn = (LPALSOURCEREWIND)proc;
	fn(a);
}

void goal_SourcePause(void* proc, ALuint a) {
	LPALSOURCEPAUSE fn = (LPALSOURCEPAUSE)proc;
	fn(a);
}

void goal_SourceQueueBuffers(void* proc, ALuint a, ALsizei b, const ALuint* c) {
	LPALSOURCEQUEUEBUFFERS fn = (LPALSOURCEQUEUEBUFFERS)proc;
	fn(a, b, c);
}

void goal_SourceUnqueueBuffers(void* proc, ALuint a, ALsizei b, ALuint* c) {
	LPALSOURCEUNQUEUEBUFFERS fn = (LPALSOURCEUNQUEUEBUFFERS)proc;
	fn(a, b, c);
}

void goal_GenBuffers(void* proc, ALsizei a, ALuint* b) {
	LPALGENBUFFERS fn = (LPALGENBUFFERS)proc;
	fn(a, b);
}

void goal_DeleteBuffers(void* proc, ALsizei a, const ALuint* b) {
	LPALDELETEBUFFERS fn = (LPALDELETEBUFFERS)proc;
	fn(a, b);
}

ALboolean goal_IsBuffer(void* proc, ALuint a) {
	LPALISBUFFER fn = (LPALISBUFFER)proc;
	return fn(a);
}

void goal_BufferData(void* proc, ALuint a, ALenum b, const ALvoid* c, ALsizei d, ALsizei e) {
	LPALBUFFERDATA fn = (LPALBUFFERDATA)proc;
	fn(a, b, c, d, e);
}

void goal_Bufferf(void* proc, ALuint a, ALenum b, ALfloat c) {
	LPALBUFFERF fn = (LPALBUFFERF)proc;
	fn(a, b, c);
}

void goal_Buffer3f(void* proc, ALuint a, ALenum b, ALfloat c, ALfloat d, ALfloat e) {
	LPALBUFFER3F fn = (LPALBUFFER3F)proc;
	fn(a, b, c, d, e);
}

void goal_Bufferfv(void* proc, ALuint a, ALenum b, ALfloat* c) {
	LPALBUFFERFV fn = (LPALBUFFERFV)proc;
	fn(a, b, c);
}

void goal_Bufferi(void* proc, ALuint a, ALenum b, ALint c) {
	LPALBUFFERI fn = (LPALBUFFERI)proc;
	fn(a, b, c);
}

void goal_Buffer3i(void* proc, ALuint a, ALenum b, ALint c, ALint d, ALint e) {
	LPALBUFFER3I fn = (LPALBUFFER3I)proc;
	fn(a, b, c, d, e);
}

void goal_Bufferiv(void* proc, ALuint a, ALenum b, ALint* c) {
	LPALBUFFERIV fn = (LPALBUFFERIV)proc;
	fn(a, b, c);
}

void goal_GetBufferf(void* proc, ALuint a, ALenum b, ALfloat* c) {
	LPALGETBUFFERF fn = (LPALGETBUFFERF)proc;
	fn(a, b, c);
}

void goal_GetBuffer3f(void* proc, ALuint a, ALenum b, ALfloat* c, ALfloat* d, ALfloat* e) {
	LPALGETBUFFER3F fn = (LPALGETBUFFER3F)proc;
	fn(a, b, c, d, e);
}

void goal_GetBufferfv(void* proc, ALuint a, ALenum b, ALfloat* c) {
	LPALGETBUFFERFV fn = (LPALGETBUFFERFV)proc;
	fn(a, b, c);
}

void goal_GetBufferi(void* proc, ALuint a, ALenum b, ALint* c) {
	LPALGETBUFFERI fn = (LPALGETBUFFERI)proc;
	fn(a, b, c);
}

void goal_GetBuffer3i(void* proc, ALuint a, ALenum b, ALint* c, ALint* d, ALint* e) {
	LPALGETBUFFER3I fn = (LPALGETBUFFER3I)proc;
	fn(a, b, c, d, e);
}

void goal_GetBufferiv(void* proc, ALuint a, ALenum b, ALint* c) {
	LPALGETBUFFERIV fn = (LPALGETBUFFERIV)proc;
	fn(a, b, c);
}

void goal_DopplerFactor(void* proc, ALfloat a) {
	LPALDOPPLERFACTOR fn = (LPALDOPPLERFACTOR)proc;
	fn(a);
}

void goal_DopplerVelocity(void* proc, ALfloat a) {
	LPALDOPPLERVELOCITY fn = (LPALDOPPLERVELOCITY)proc;
	fn(a);
}

void goal_SpeedOfSound(void* proc, ALfloat a) {
	LPALSPEEDOFSOUND fn = (LPALSPEEDOFSOUND)proc;
	fn(a);
}

void goal_DistanceModel(void* proc, ALenum a) {
	LPALDISTANCEMODEL fn = (LPALDISTANCEMODEL)proc;
	fn(a);
}



// Begin extensions -- do not place non-extension functions below here.



void goal_BufferDataStatic(void* proc, const ALuint a, ALenum b, ALvoid* c, ALsizei d, ALsizei e) {
	PFNALBUFFERDATASTATICPROC fn = (PFNALBUFFERDATASTATICPROC)proc;
	fn(a, b, c, d, e);
}

void goal_BufferSubDataSOFT(void* proc, ALuint a, ALenum b, const ALvoid* c, ALsizei d, ALsizei e) {
	PFNALBUFFERSUBDATASOFTPROC fn = (PFNALBUFFERSUBDATASOFTPROC)proc;
	fn(a, b, c, d, e);
}

void goal_BufferSamplesSOFT(void* proc, ALuint a, ALuint b, ALenum c, ALsizei d, ALenum e, ALenum f, const ALvoid* g) {
	LPALBUFFERSAMPLESSOFT fn = (LPALBUFFERSAMPLESSOFT)proc;
	fn(a, b, c, d, e, f, g);
}

void goal_BufferSubSamplesSOFT(void* proc, ALuint a, ALsizei b, ALsizei c, ALenum d, ALenum e, const ALvoid* f) {
	LPALBUFFERSUBSAMPLESSOFT fn = (LPALBUFFERSUBSAMPLESSOFT)proc;
	fn(a, b, c, d, e, f);
}

void goal_GetBufferSamplesSOFT(void* proc, ALuint a, ALsizei b, ALsizei c, ALenum d, ALenum e, ALvoid* f) {
	LPALGETBUFFERSAMPLESSOFT fn = (LPALGETBUFFERSAMPLESSOFT)proc;
	fn(a, b, c, d, e, f);
}

ALboolean goal_IsBufferFormatSupportedSOFT(void* proc, ALenum a) {
	LPALISBUFFERFORMATSUPPORTEDSOFT fn = (LPALISBUFFERFORMATSUPPORTEDSOFT)proc;
	return fn(a);
}

void goal_SourcedSOFT(void* proc, ALuint a, ALenum b, ALdouble c) {
	LPALSOURCEDSOFT fn = (LPALSOURCEDSOFT)proc;
	fn(a, b, c);
}

void goal_Source3dSOFT(void* proc, ALuint a, ALenum b, ALdouble c, ALdouble d, ALdouble e) {
	LPALSOURCE3DSOFT fn = (LPALSOURCE3DSOFT)proc;
	fn(a, b, c, d, e);
}

void goal_SourcedvSOFT(void* proc, ALuint a, ALenum b, ALdouble* c) {
	LPALSOURCEDVSOFT fn = (LPALSOURCEDVSOFT)proc;
	fn(a, b, c);
}

void goal_GetSourcedSOFT(void* proc, ALuint a, ALenum b, ALdouble* c) {
	LPALGETSOURCEDSOFT fn = (LPALGETSOURCEDSOFT)proc;
	fn(a, b, c);
}

void goal_GetSource3dSOFT(void* proc, ALuint a, ALenum b, ALdouble* c, ALdouble* d, ALdouble* e) {
	LPALGETSOURCE3DSOFT fn = (LPALGETSOURCE3DSOFT)proc;
	fn(a, b, c, d, e);
}

void goal_GetSourcedvSOFT(void* proc, ALuint a, ALenum b, ALdouble* c) {
	LPALGETSOURCEDVSOFT fn = (LPALGETSOURCEDVSOFT)proc;
	fn(a, b, c);
}

void goal_Sourcei64SOFT(void* proc, ALuint a, ALenum b, ALint64SOFT c) {
	LPALSOURCEI64SOFT fn = (LPALSOURCEI64SOFT)proc;
	fn(a, b, c);
}

void goal_Source3i64SOFT(void* proc, ALuint a, ALenum b, ALint64SOFT c, ALint64SOFT d, ALint64SOFT e) {
	LPALSOURCE3I64SOFT fn = (LPALSOURCE3I64SOFT)proc;
	fn(a, b, c, d, e);
}

void goal_Sourcei64vSOFT(void* proc, ALuint a, ALenum b, const ALint64SOFT* c) {
	LPALSOURCEI64VSOFT fn = (LPALSOURCEI64VSOFT)proc;
	fn(a, b, c);
}

void goal_GetSourcei64SOFT(void* proc, ALuint a, ALenum b, ALint64SOFT* c) {
	LPALGETSOURCEI64SOFT fn = (LPALGETSOURCEI64SOFT)proc;
	fn(a, b, c);
}

void goal_GetSource3i64SOFT(void* proc, ALuint a, ALenum b, ALint64SOFT* c, ALint64SOFT* d, ALint64SOFT* e) {
	LPALGETSOURCE3I64SOFT fn = (LPALGETSOURCE3I64SOFT)proc;
	fn(a, b, c, d, e);
}

void goal_GetSourcei64vSOFT(void* proc, ALuint a, ALenum b, ALint64SOFT* c) {
	LPALGETSOURCEI64VSOFT fn = (LPALGETSOURCEI64VSOFT)proc;
	fn(a, b, c);
}

/*
typedef ALCdevice* (ALC_APIENTRY*LPALCLOOPBACKOPENDEVICESOFT)(const ALCchar*);
typedef ALCboolean (ALC_APIENTRY*LPALCISRENDERFORMATSUPPORTEDSOFT)(ALCdevice*,ALCsizei,ALCenum,ALCenum);
typedef void (ALC_APIENTRY*LPALCRENDERSAMPLESSOFT)(ALCdevice*,ALCvoid*,ALCsizei);

*/
