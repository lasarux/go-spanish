##############################################################################
# Copyright (c) 2009, Pedro Gracia <lasarux@neuroomante.com>
# All rights reserved.
# 
# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions are met:
# 
# Redistributions of source code must retain the above copyright notice, this
# list of conditions and the following disclaimer.
# 
# Redistributions in binary form must reproduce the above copyright notice,
# this list of conditions and the following disclaimer in the documentation
# and/or other materials provided with the distribution.
# 
# Neither the name of Pedro Gracia nor the names of his contributors may
# be used to endorse or promote products derived from this software without
# specific prior written permission.
# 
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
# ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
# LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
# CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
# SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
# INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
# CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
# ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
# POSSIBILITY OF SUCH DAMAGE.
##############################################################################

# Makefile for allegro on 64-bit Linux. This needs modifications for other platforms.
# Makefile original of Kenneth Falck <kennu@iki.fi>

GOG	= 6g
GOL	= 6l
GOC	= 6c
GOCFL	= -I$(GOROOT)/src/pkg/runtime
CC	= gcc
CGO	= cgo
GOPACK	= gopack
OBJS	= allegro_test.6
BIN	= allegro_test
LIBS	= allegro.a
RM	= rm -f
RUNTIME_CFLAGS	= -D_64BIT
CGO_CFLAGS	= -m64
CGO_LDFLAGS	= -shared -lpthread -lm `allegro-config --libs` -lagl -lGL -lGLU #-dynamiclib -Wl,-undefined,dynamic_lookup


all:	$(BIN)

$(BIN):	$(LIBS) $(OBJS)
	$(GOL) -o $(BIN) $(OBJS)

allegro.a:
	$(CGO) allegro.go
	$(GOG) -o allegro.6 allegro.cgo1.go allegro.cgo2.go
	$(GOC) $(RUNTIME_CFLAGS) -o allegro_c.6 $(GOCFL) allegro.cgo3.c
	$(CC) -fPIC $(CGO_CFLAGS) -o allegro_so.o -c allegro.cgo4.c
	$(CC) -o $(GOROOT)/pkg/linux_amd64/allegro_allegro.so allegro_so.o $(CGO_LDFLAGS)
	$(RM) allegro.a
	$(GOPACK) grc allegro.a allegro.6 allegro_c.6

%.6:	%.go
	$(GOG) $<

clean:
	$(RM) *.o *.a *.6 *.so allegro.cgo* $(BIN)
