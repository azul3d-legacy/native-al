// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
#include "AL/al.h"
*/
import "C"

const (
	// "no distance model" or "no buffer"
	NONE int32 = C.AL_NONE

	// Boolean False.
	FALSE int32 = C.AL_FALSE

	// Boolean True.
	TRUE int32 = C.AL_TRUE

	// Relative source.
	// Type:	uint8
	// Range:   [TRUE, FALSE]
	// Default: FALSE
	//
	// Specifies if the Source has relative coordinates.
	SOURCE_RELATIVE int32 = C.AL_SOURCE_RELATIVE

	// Inner cone angle, in degrees.
	// Type:	int32, float32
	// Range:   [0 - 360]
	// Default: 360
	//
	// The angle covered by the inner cone, where the source will not attenuate.
	CONE_INNER_ANGLE int32 = C.AL_CONE_INNER_ANGLE

	// Outer cone angle, in degrees.
	// Range:   [0 - 360]
	// Default: 360
	//
	// The angle covered by the outer cone, where the source will be fully
	// attenuated.
	CONE_OUTER_ANGLE int32 = C.AL_CONE_OUTER_ANGLE

	// Source pitch.
	// Type:	float32
	// Range:   [0.5 - 2.0]
	// Default: 1.0
	//
	// A multiplier for the frequency (sample rate) of the source's buffer.
	PITCH int32 = C.AL_PITCH

	// Source or listener position.
	// Type:	float32[3], int32[3]
	// Default: {0, 0, 0}
	//
	// The source or listener location in three dimensional space.
	//
	// OpenAL, like OpenGL, uses a right handed coordinate system, where in a
	// frontal default view X (thumb) points right, Y points up (index finger), and
	// Z points towards the viewer/camera (middle finger).
	//
	// To switch from a left handed coordinate system, flip the sign on the Z
	// coordinate.
	POSITION int32 = C.AL_POSITION

	// Source direction.
	// Type:	float32[3], int32[3]
	// Default: {0, 0, 0}
	//
	// Specifies the current direction in local space.
	// A zero-length vector specifies an omni-directional source (cone is ignored).
	DIRECTION int32 = C.AL_DIRECTION

	// Source or listener velocity.
	// Type:	float32[3], int32[3]
	// Default: {0, 0, 0}
	//
	// Specifies the current velocity in local space.
	VELOCITY int32 = C.AL_VELOCITY

	// Source looping.
	// Type:	uint8
	// Range:   [TRUE, FALSE]
	// Default: FALSE
	//
	// Specifies whether source is looping.
	LOOPING int32 = C.AL_LOOPING

	// Source buffer.
	// Type:  uint32
	// Range: any valid Buffer.
	//
	// Specifies the buffer to provide sound samples.
	BUFFER int32 = C.AL_BUFFER

	// Source or listener gain.
	// Type:  float32
	// Range: [0.0 - ]
	//
	// A value of 1.0 means unattenuated. Each division by 2 equals an attenuation
	// of about -6dB. Each multiplicaton by 2 equals an amplification of about
	// +6dB.
	//
	// A value of 0.0 is meaningless with respect to a logarithmic scale; it is
	// silent.
	GAIN int32 = C.AL_GAIN

	// Minimum source gain.
	// Type:  float32
	// Range: [0.0 - 1.0]
	//
	// The minimum gain allowed for a source, after distance and cone attenation is
	// applied (if applicable).
	MIN_GAIN int32 = C.AL_MIN_GAIN

	// Maximum source gain.
	// Type:  float32
	// Range: [0.0 - 1.0]
	//
	// The maximum gain allowed for a source, after distance and cone attenation is
	// applied (if applicable).
	MAX_GAIN int32 = C.AL_MAX_GAIN

	// Listener orientation.
	// Type: float32[6]
	// Default: {0.0, 0.0, -1.0, 0.0, 1.0, 0.0}
	//
	// Effectively two three dimensional vectors. The first vector is the front (or
	// "at") and the second is the top (or "up").
	//
	// Both vectors are in local space.
	ORIENTATION int32 = C.AL_ORIENTATION

	// Source state (query only).
	// Type:  int32
	// Range: [INITIAL, PLAYING, PAUSED, STOPPED]
	SOURCE_STATE int32 = C.AL_SOURCE_STATE

	// Source state value.
	INITIAL int32 = C.AL_INITIAL
	PLAYING int32 = C.AL_PLAYING
	PAUSED  int32 = C.AL_PAUSED
	STOPPED int32 = C.AL_STOPPED

	// Source Buffer Queue size (query only).
	// Type: int32
	//
	// The number of buffers queued using alSourceQueueBuffers, minus the buffers
	// removed with alSourceUnqueueBuffers.
	BUFFERS_QUEUED int32 = C.AL_BUFFERS_QUEUED

	// Source Buffer Queue processed count (query only).
	// Type: int32
	//
	// The number of queued buffers that have been fully processed, and can be
	// removed with alSourceUnqueueBuffers.
	//
	// Looping sources will never fully process buffers because they will be set to
	// play again for when the source loops.
	BUFFERS_PROCESSED int32 = C.AL_BUFFERS_PROCESSED

	// Source reference distance.
	// Type:	float32
	// Range:   [0.0 - ]
	// Default: 1.0
	//
	// The distance in units that no attenuation occurs.
	//
	// At 0.0, no distance attenuation ever occurs on non-linear attenuation models.
	REFERENCE_DISTANCE int32 = C.AL_REFERENCE_DISTANCE

	// Source rolloff factor.
	// Type:	float32
	// Range:   [0.0 - ]
	// Default: 1.0
	//
	// Multiplier to exaggerate or diminish distance attenuation.
	//
	// At 0.0, no distance attenuation ever occurs.
	ROLLOFF_FACTOR int32 = C.AL_ROLLOFF_FACTOR

	// Outer cone gain.
	// Type:	float32
	// Range:   [0.0 - 1.0]
	// Default: 0.0
	//
	// The gain attenuation applied when the listener is outside of the source's
	// outer cone.
	CONE_OUTER_GAIN int32 = C.AL_CONE_OUTER_GAIN

	// Source maximum distance.
	// Type:	float32
	// Range:   [0.0 - ]
	// Default: +inf
	//
	// The distance above which the source is not attenuated any further with a
	// clamped distance model, or where attenuation reaches 0.0 gain for linear
	// distance models with a default rolloff factor.
	MAX_DISTANCE int32 = C.AL_MAX_DISTANCE

	// Source buffer position, in seconds
	SEC_OFFSET int32 = C.AL_SEC_OFFSET

	// Source buffer position, in sample frames
	SAMPLE_OFFSET int32 = C.AL_SAMPLE_OFFSET

	// Source buffer position, in bytes
	BYTE_OFFSET int32 = C.AL_BYTE_OFFSET

	// Source type (query only).
	// Type:  int32
	// Range: [STATIC, STREAMING, UNDETERMINED]
	//
	// A Source is Static if a Buffer has been attached using BUFFER.
	//
	// A Source is Streaming if one or more Buffers have been attached using
	// alSourceQueueBuffers.
	//
	// A Source is Undetermined when it has the NULL buffer attached using
	// BUFFER.
	SOURCE_TYPE int32 = C.AL_SOURCE_TYPE

	// Source type value.
	STATIC       int32 = C.AL_STATIC
	STREAMING    int32 = C.AL_STREAMING
	UNDETERMINED int32 = C.AL_UNDETERMINED

	// Buffer format specifier.
	FORMAT_MONO8    int32 = C.AL_FORMAT_MONO8
	FORMAT_MONO16   int32 = C.AL_FORMAT_MONO16
	FORMAT_STEREO8  int32 = C.AL_FORMAT_STEREO8
	FORMAT_STEREO16 int32 = C.AL_FORMAT_STEREO16

	// Buffer frequency (query only).
	FREQUENCY int32 = C.AL_FREQUENCY

	// Buffer bits per sample (query only).
	BITS int32 = C.AL_BITS

	// Buffer channel count (query only).
	CHANNELS int32 = C.AL_CHANNELS

	// Buffer data size (query only).
	SIZE int32 = C.AL_SIZE

	// Context string: Vendor ID.
	VENDOR int32 = C.AL_VENDOR

	// Context string: Version.
	VERSION int32 = C.AL_VERSION

	// Context string: Renderer ID.
	RENDERER int32 = C.AL_RENDERER

	// Context string: Space-separated extension list.
	EXTENSIONS int32 = C.AL_EXTENSIONS

	// Doppler scale.
	// Type:	float32
	// Range:   [0.0 - ]
	// Default: 1.0
	//
	// Scale for source and listener velocities.
	DOPPLER_FACTOR int32 = C.AL_DOPPLER_FACTOR

	// Doppler velocity (deprecated).
	//
	// A multiplier applied to the Speed of Sound.
	DOPPLER_VELOCITY int32 = C.AL_DOPPLER_VELOCITY

	// Speed of Sound, in units per second.
	// Type:	float32
	// Range:   [0.0001 - ]
	// Default: 343.3
	//
	// The speed at which sound waves are assumed to travel, when calculating the
	// doppler effect.
	SPEED_OF_SOUND int32 = C.AL_SPEED_OF_SOUND

	// Distance attenuation model.
	// Type:	int32
	// Range:   [NONE, INVERSE_DISTANCE, INVERSE_DISTANCE_CLAMPED,
	//		   LINEAR_DISTANCE, LINEAR_DISTANCE_CLAMPED,
	//		   EXPONENT_DISTANCE, EXPONENT_DISTANCE_CLAMPED]
	// Default: INVERSE_DISTANCE_CLAMPED
	//
	// The model by which sources attenuate with distance.
	//
	// None	 - No distance attenuation.
	// Inverse  - Doubling the distance halves the source gain.
	// Linear   - Linear gain scaling between the reference and max distances.
	// Exponent - Exponential gain dropoff.
	//
	// Clamped variations work like the non-clamped counterparts, except the
	// distance calculated is clamped between the reference and max distances.
	DISTANCE_MODEL int32 = C.AL_DISTANCE_MODEL

	// Distance model value.
	INVERSE_DISTANCE          int32 = C.AL_INVERSE_DISTANCE
	INVERSE_DISTANCE_CLAMPED  int32 = C.AL_INVERSE_DISTANCE_CLAMPED
	LINEAR_DISTANCE           int32 = C.AL_LINEAR_DISTANCE
	LINEAR_DISTANCE_CLAMPED   int32 = C.AL_LINEAR_DISTANCE_CLAMPED
	EXPONENT_DISTANCE         int32 = C.AL_EXPONENT_DISTANCE
	EXPONENT_DISTANCE_CLAMPED int32 = C.AL_EXPONENT_DISTANCE_CLAMPED
)
