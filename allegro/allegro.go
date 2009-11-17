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
#include <alleggl.h>

void a_init(void) {
    allegro_init();
}

void a_install_keyboard(void) {
    install_keyboard();
}

int a_set_gfx_mode(void) {
    int depth = 24; //or 8, 16, 24
    set_color_depth(depth);
    if (set_gfx_mode(GFX_AUTODETECT_WINDOWED, 640, 400, 0, 0) != 0) {
        if (set_gfx_mode(GFX_SAFE, 640, 400, 0, 0) != 0) {
            set_gfx_mode(GFX_TEXT, 0, 0, 0, 0);
            allegro_message("Unable to set any graphic mode\n%s\n", allegro_error);
            return 1;
        }
    }
    return 0;
}

int a_set_opengl_mode(void) {
    install_allegro_gl();
	allegro_gl_set(AGL_COLOR_DEPTH, 24);

    if (set_gfx_mode(GFX_OPENGL_WINDOWED, 640, 400, 0, 0) != 0) {
        if (set_gfx_mode(GFX_SAFE, 640, 400, 0, 0) != 0) {
            set_gfx_mode(GFX_TEXT, 0, 0, 0, 0);
            allegro_message("Unable to set any graphic mode\n%s\n", allegro_error);
            return 1;
        }
    }
    
    glShadeModel (GL_FLAT);
	glEnable (GL_DEPTH_TEST);
	glCullFace (GL_BACK);
	glEnable (GL_CULL_FACE);
    glShadeModel(GL_SMOOTH);
    glClearColor(0.0, 0.0, 0.0, 0.0);
    return 0;
}

void a_glClear(void) {
    glClear(GL_COLOR_BUFFER_BIT);
}

void a_glBegin(void) {
    glBegin(GL_TRIANGLE_FAN);
}

void a_glColor3f(float r, float g, float b) {
    glColor3f(r, g, b);
}

void a_glVertex2f(float x, float y) {
    glVertex2f(x, y);
}

void a_glVertex3d(float x, float y, float z) {
    glVertex3d(x, y, z);
}

void a_glEnd(void) {
    glEnd();
}

void a_glFlush(void) {
    glFlush();
}

void a_allegro_gl_flip(void) {
    allegro_gl_flip();
}

void a_glRotatef(float a, float b, float c, float d) {
    glRotatef(a, b, c, d);
}

void a_glTranslatef(float a, float b, float c) {
    glTranslatef(a, b, c);
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

void a_END_OF_MAIN(void) {
    END_OF_MAIN();
}
*/
import "C"
//import "unsafe"

func Init()  {
	C.a_init();
}

func Install_keyboard() {
    C.a_install_keyboard();
}

func Set_gfx_mode() {
    C.a_set_gfx_mode();
}

func Set_opengl_mode() {
    C.a_set_opengl_mode();
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

func GLClear() {
    C.a_glClear();
}

func GLBegin() {
    C.a_glBegin();
}

func GLEnd() {
    C.a_glEnd();
}

func GLFlush() {
    C.a_glFlush();
}

func Allegro_gl_flip(){
    C.a_allegro_gl_flip();
}

func GLColor3f(r, g, b float) {
    C.a_glColor3f(_C_float(r), _C_float(g), _C_float(b));
}

func GLVertex2f(x, y float) {
    C.a_glVertex2f(_C_float(x), _C_float(y));
}

func GLVertex3d(x, y, z float) {
    C.a_glVertex3d(_C_float(x), _C_float(y), _C_float(z));
}

func GLRotatef(a, b, c, d float) {
    C.a_glRotatef(_C_float(a), _C_float(b), _C_float(c), _C_float(d));
}

func GLTranslatef(a, b, c float) {
    C.a_glTranslatef(_C_float(a), _C_float(b), _C_float(c));
}

func End() {
    C.a_END_OF_MAIN();
}
