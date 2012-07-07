// Ported from GLUT's samples.  Original copyright below applies.

/* Copyright (c) Mark J. Kilgard, 1994. */

/* This program is freely distributable without licensing fees 
   and is provided without guarantee or warrantee expressed or 
   implied. This program is -not- in the public domain. */

package main

import (
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"github.com/rhencke/glut"
)

func bitmap_output(x, y float32, str string, font glut.BitmapFont) {
	gl.RasterPos2f(x, y)
	for _, ch := range str {
		font.Character(ch)
	}
}

func stroke_output(x, y float32, str string, font glut.StrokeFont) {
	gl.PushMatrix()
	gl.Translatef(x, y, 0)
	gl.Scalef(0.005, 0.005, 0.005)
	for _, ch := range str {
		font.Character(ch)
	}
	gl.PopMatrix()
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	bitmap_output(40, 35, "This is written in a GLUT bitmap font.", glut.BITMAP_TIMES_ROMAN_24)
	bitmap_output(30, 210, "More bitmap text is a fixed 9 by 15 font.", glut.BITMAP_9_BY_15)
	bitmap_output(70, 240, "                Helvetica is yet another bitmap font.", glut.BITMAP_HELVETICA_18)

	gl.MatrixMode(gl.PROJECTION)
	gl.PushMatrix()
	gl.LoadIdentity()
	glu.Perspective(40.0, 1.0, 0.1, 20.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.PushMatrix()
	glu.LookAt(0.0, 0.0, 4.0, /* eye is at (0,0,30) */
		0.0, 0.0, 0.0, /* center is at (0,0,0) */
		0.0, 1.0, 0.0) /* up is in postivie Y direction */
	gl.PushMatrix()
	gl.Translatef(0, 0, -4)
	gl.Rotatef(50, 0, 1, 0)
	stroke_output(-2.5, 1.1, "  This is written in a", glut.STROKE_ROMAN)
	stroke_output(-2.5, 0, " GLUT stroke font.", glut.STROKE_ROMAN)
	stroke_output(-2.5, -1.1, "using 3D perspective.", glut.STROKE_ROMAN)
	gl.PopMatrix()
	gl.MatrixMode(gl.MODELVIEW)
	gl.PopMatrix()
	gl.MatrixMode(gl.PROJECTION)
	gl.PopMatrix()
	gl.MatrixMode(gl.MODELVIEW)
	gl.Flush()
}

func reshape(w, h int) {
	gl.Viewport(0, 0, w, h)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
	gl.Scalef(1, -1, 1)
	gl.Translatef(0, float32(-h), 0)
	gl.MatrixMode(gl.MODELVIEW)
}

func main() {
	glut.InitDisplayMode(glut.SINGLE | glut.RGB)
	glut.InitWindowSize(465, 250)
	glut.CreateWindow("GLUT bitmap & stroke font example")
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Color3f(0, 0, 0)
	gl.LineWidth(3.0)
	glut.DisplayFunc(display)
	glut.ReshapeFunc(reshape)
	glut.MainLoop()
}
