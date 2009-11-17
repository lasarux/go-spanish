/*****************************************************************************
 * Copyright (c) 2009, Pedro Gracia <lasarux@neuroomante.com>
 * All rights reserved.
 * 
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 * 
 * Redistributions of source code must retain the above copyright notice, this
 * list of conditions and the following disclaimer.
 * 
 * Redistributions in binary form must reproduce the above copyright notice,
 * this list of conditions and the following disclaimer in the documentation
 * and/or other materials provided with the distribution.
 * 
 * Neither the name of Pedro Gracia nor the names of his contributors may
 * be used to endorse or promote products derived from this software without
 * specific prior written permission.
 * 
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *****************************************************************************/
package allegro

/*
#include <allegro.h>

int a_init(void) {
    return allegro_init();
}

void a_install_keyboard(void) {
    install_keyboard(); 
}

int a_set_gfx_mode(void) {
    if (set_gfx_mode(GFX_AUTODETECT_WINDOWED, 640, 400, 0, 0) != 0) {
        if (set_gfx_mode(GFX_SAFE, 640, 400, 0, 0) != 0) {
            set_gfx_mode(GFX_TEXT, 0, 0, 0, 0);
            allegro_message("Unable to set any graphic mode\n%s\n", allegro_error);
            return 1;
        }
    }
    return 0;
}

void a_textout_centre_ex(char *text) {
    //set_palette(desktop_palette);
    clear_to_color(screen, makecol(255, 255, 255));
    acquire_screen();
    textout_centre_ex(screen, font, text, SCREEN_W/2, SCREEN_H/2, makecol(0,0,0), -1);
    release_screen();
}

void a_putpixel(int x, int y, int color) {
    putpixel(screen, x, y, color);
}

void a_readkey(void) {
    readkey();
}

void a_set_window_title(char *text) {
    set_window_title(text);
}

*/
import "C"
//import "unsafe"

func Init() {
	C.a_init();
}

func Install_keyboard() {
    C.a_install_keyboard();
}

func Set_gfx_mode() {
    C.a_set_gfx_mode();
}

func Readkey() {
    C.a_readkey();
}

func Textout_centr_ex(text string) {
    cs := C.CString(text);
    C.a_textout_centre_ex(cs);
}

func Set_window_title(text string) {
    cs := C.CString(text);
    C.a_set_window_title(cs);
}

func Putpixel(x, y, color int) {
    C.a_putpixel(_C_int(x), _C_int(y), _C_int(color));
}
