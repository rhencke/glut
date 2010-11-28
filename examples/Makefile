# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

all: fontdemo simple

clean:
	rm -f $(patsubst %.go,%.$(O),$(wildcard *.go)) $(patsubst %.go,%,$(wildcard *.go))

%: %.go
	$(GC) $<
	$(LD) -o $@ $@.$(O)

