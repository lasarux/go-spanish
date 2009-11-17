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
package main

import "./allegro"
import "math"

func main() {
    // Init window
    allegro.Init();
    allegro.Install_keyboard();
    allegro.Set_gfx_mode();
    allegro.Set_window_title("Go rules!");
    
    // Show text centered
    allegro.Textout_centr_ex("Hello, world!");
    allegro.Readkey();
    
    // More text centered
    allegro.Textout_centr_ex("Hola, mundo!");
    allegro.Readkey();
    
    // Draw a mandelbrot set
    a := 0.;
    b := 0.;
    aa := 0.;
    bb := 0.;
    ca := 0.;
    cb := 0.;
    i := 0;
    max := 64;
    
    for y := 0; y<400; y++ {
        for x := 0; x<640; x++ {
            i = 0;
            ca = (float) (x - 420)/200;
			cb = (float) (y - 200)/200;
			a = 0.;
            b = 0.;
            for i < max && math.Sqrt(float64(a*a)+float64(b*b)) <= 2 { 
                aa = a*a - b*b + ca;
				bb = 2*a*b + cb;
				a = aa;
				b = bb;
				i++;
            }
            
            allegro.Putpixel(x, y, i);
        }
    }
    allegro.Readkey();
}
