# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=glut

CGOFILES:=glut.go

PLATFORM:=$(shell uname -s)

CGO_DEPS:=_cgo_export.o
ifeq ($(PLATFORM),Darwin)
CGO_LDFLAGS:= _cgo_export.o -framework GLUT
CGO_CFLAGS:=-D__Darwin
else
CGO_LDFLAGS:=-lGLUT _cgo_export.o 
CGO_CFLAGS:=-D__$(PLATFORM)
endif

include $(GOROOT)/src/Make.pkg
