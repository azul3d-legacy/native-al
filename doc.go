// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package al implements Go bindings to OpenAL.
//
// To avoid the requirement of users having to download OpenAL (or for
// developers to have to distribute dynamic libraries with their applications),
// this package works as a sort-of automatic 'installer'
//
// The dynamic OpenAL-Soft library is packed as a binary blob within this
// package automatically for you.
//
// When an application imports this package the installer will automatically
// run. It is not visible to users and only writes a single file to the hard
// drive *if it does not exist*.
//
// Because OpenAL-Soft is LGPL licensed it is required by the license that
// users be able to run the application with their own version of the library,
// we are not lawyers and you should consult a lawyer to know in full if you
// are obeying the LGPL license of OpenAL-Soft when using this library, but:
//
// You should simply inform your users where they can place their own
// OpenAL-Soft dynamic link library for their specific platform, file named as
// shown below, into their home directory under the ".azul3d" directory created
// on their operating system's file system:
//
//  linux/amd64: libopenal_soft.so.1.15.1
//
// And it will be dynamically linked and loaded instead of ours. Should the
// user wish to restore our dynamic link library, they can simply delete theirs
// and it will be automatically placed there once again.
//
package al // import "azul3d.org/native/al.v1-unstable"

