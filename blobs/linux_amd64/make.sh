# Copyright 2014 The Azul3D Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

objcopy -I binary -O elf64-x86-64 -B i386 openal ../../blob_linux_amd64.syso
