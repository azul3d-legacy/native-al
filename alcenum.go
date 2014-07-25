// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
#include "AL/alc.h"
*/
import "C"

const (
	// Context attribute: <int32> Hz.
	ALC_FREQUENCY int32 = C.ALC_FREQUENCY

	// Context attribute: <int32> Hz.
	ALC_REFRESH int32 = C.ALC_REFRESH

	// Context attribute: TRUE or FALSE.
	ALC_SYNC int32 = C.ALC_SYNC

	// Context attribute: <int32> requested Mono (3D) Sources.
	ALC_MONO_SOURCES int32 = C.ALC_MONO_SOURCES

	// Context attribute: <int32> requested Stereo Sources.
	ALC_STEREO_SOURCES int32 = C.ALC_STEREO_SOURCES

	// Runtime ALC version.
	ALC_MAJOR_VERSION int32 = C.ALC_MAJOR_VERSION
	ALC_MINOR_VERSION int32 = C.ALC_MINOR_VERSION

	// Context attribute list properties.
	ALC_ATTRIBUTES_SIZE int32 = C.ALC_ATTRIBUTES_SIZE
	ALC_ALL_ATTRIBUTES  int32 = C.ALC_ALL_ATTRIBUTES

	// String for the default device specifier.
	ALC_DEFAULT_DEVICE_SPECIFIER int32 = C.ALC_DEFAULT_DEVICE_SPECIFIER

	// String for the given device's specifier.
	//
	// If device handle is NULL, it is instead a null-char separated list of
	// strings of known device specifiers (list ends with an empty string).
	ALC_DEVICE_SPECIFIER int32 = C.ALC_DEVICE_SPECIFIER

	// String for space-separated list of ALC extensions.
	ALC_EXTENSIONS int32 = C.ALC_EXTENSIONS

	// Capture extension
	ALC_EXT_CAPTURE int32 = C.ALC_EXT_CAPTURE

	// String for the given capture device's specifier.
	//
	// If device handle is NULL, it is instead a null-char separated list of
	// strings of known capture device specifiers (list ends with an empty string).
	ALC_CAPTURE_DEVICE_SPECIFIER int32 = C.ALC_CAPTURE_DEVICE_SPECIFIER

	// String for the default capture device specifier.
	ALC_CAPTURE_DEFAULT_DEVICE_SPECIFIER int32 = C.ALC_CAPTURE_DEFAULT_DEVICE_SPECIFIER

	// Number of sample frames available for capture.
	ALC_CAPTURE_SAMPLES int32 = C.ALC_CAPTURE_SAMPLES

	// Enumerate All extension
	ALC_ENUMERATE_ALL_EXT int32 = C.ALC_ENUMERATE_ALL_EXT

	// String for the default extended device specifier.
	ALC_DEFAULT_ALL_DEVICES_SPECIFIER int32 = C.ALC_DEFAULT_ALL_DEVICES_SPECIFIER

	// String for the given extended device's specifier.
	//
	// If device handle is NULL, it is instead a null-char separated list of
	// strings of known extended device specifiers (list ends with an empty string).
	ALC_ALL_DEVICES_SPECIFIER int32 = C.ALC_ALL_DEVICES_SPECIFIER
)
