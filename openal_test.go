// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

import (
	"testing"
)

func TestOpenAL(t *testing.T) {
	defaultDevice := AlcGetString(nil, ALC_DEFAULT_DEVICE_SPECIFIER)
	t.Log("Default Device:", defaultDevice)

	defaultCapture := AlcGetString(nil, ALC_CAPTURE_DEFAULT_DEVICE_SPECIFIER)
	t.Log("Default Capture Device:", defaultCapture)

	allDevices := StringList(AlcGetRawString(nil, ALC_DEVICE_SPECIFIER))
	for i, device := range allDevices {
		t.Logf("device %d. %q", i, device)
	}

	captureDevices := StringList(AlcGetRawString(nil, ALC_CAPTURE_DEVICE_SPECIFIER))
	for i, device := range captureDevices {
		t.Logf("capture device %d. %q", i, device)
	}

	extensions := AlcGetString(nil, ALC_EXTENSIONS)
	t.Log("Extensions:", extensions)

	device, err := OpenDevice("", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer device.Close()

	haveCapture := device.AlcIsExtensionPresent("ALC_EXT_CAPTURE")

	var maxSources int32
	device.AlcGetIntegerv(ALC_MONO_SOURCES, 1, &maxSources)
	t.Log("Maximum sources:", maxSources)

	if haveCapture && len(captureDevices) > 0 {
		t.Log("Have the ALC_EXT_CAPTURE extension.")
		err = device.InitCapture(44100, FORMAT_MONO16, 44100/2)
		if err != nil {
			t.Fatal(err)
		}
	}
}
