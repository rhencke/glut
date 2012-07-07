// Ported from GLUT's samples.  Original copyright below applies.

/* Copyright (c) Mark J. Kilgard, 1997, 1998. */

/* This program is freely distributable without licensing fees
   and is provided without guarantee or warrantee expressed or
   implied. This program is -not- in the public domain. */

package main

import (
	"math"
	"os"

	"github.com/banthar/gl"
	"github.com/rhencke/glut"
)

var (
	angle                         float32
	left, right                   bool
	leftTime, rightTime           int
	thrust                        bool
	thrustTime                    int
	joyThrust, joyLeft, joyRight  bool
	x                             float32 = 20
	y                             float32 = 20
	xv, yv, v                     float32
	shield, joyShield             bool
	cursor                        bool = true
	lastTime                      int
	paused                        bool
	resuming                      bool = true
	originalWindow, currentWindow glut.Window
)

type Bullet struct {
	inuse           bool
	x, y, v, xv, yv float32
	expire          int
}

const MAX_BULLETS = 10

var bullet [MAX_BULLETS]Bullet

func allocBullet() int {
	for i := 0; i < MAX_BULLETS; i++ {
		if !bullet[i].inuse {
			return i
		}
	}
	return -1
}

func initBullet(i, time int) {
	c := float32(math.Cos(float64(angle) * math.Pi / 180.0))
	s := float32(math.Sin(float64(angle) * math.Pi / 180.0))

	bullet[i].inuse = true
	bullet[i].x = x + 2*c
	bullet[i].y = y + 2*s
	bullet[i].v = 0.025
	bullet[i].xv = xv + c*bullet[i].v
	bullet[i].yv = yv + s*bullet[i].v
	bullet[i].expire = time + 1000
}

func advanceBullets(delta, time int) {
	for i := 0; i < MAX_BULLETS; i++ {
		if bullet[i].inuse {
			if time > bullet[i].expire {
				bullet[i].inuse = false
				continue
			}
			x := bullet[i].x + bullet[i].xv*float32(delta)
			y := bullet[i].y + bullet[i].yv*float32(delta)
			x = x / 40.0
			bullet[i].x = (x - float32(math.Floor(float64(x)))) * 40.0
			y = y / 40.0
			bullet[i].y = (y - float32(math.Floor(float64(x)))) * 40.0
		}
	}
}

func shotBullet() {
	entry := allocBullet()
	if entry >= 0 {
		initBullet(entry, glut.Get(glut.ELAPSED_TIME))
	}
}

func drawBullets() {
	gl.Begin(gl.POINTS)
	gl.Color3f(1.0, 0.0, 1.0)
	for i := 0; i < MAX_BULLETS; i++ {
		if bullet[i].inuse {
			gl.Vertex2f(bullet[i].x, bullet[i].y)
		}
	}
	gl.End()
}

func drawShip(angle float32) {
	gl.PushMatrix()
	gl.Translatef(x, y, 0.0)
	gl.Rotatef(angle, 0.0, 0.0, 1.0)
	if thrust {
		gl.Color3f(1.0, 0.0, 0.0)
		gl.Begin(gl.LINE_STRIP)
		gl.Vertex2f(-0.75, -0.5)
		gl.Vertex2f(-1.75, 0)
		gl.Vertex2f(-0.75, 0.5)
		gl.End()
	}
	gl.Color3f(1.0, 1.0, 0.0)
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(2.0, 0.0)
	gl.Vertex2f(-1.0, -1.0)
	gl.Vertex2f(-0.5, 0.0)
	gl.Vertex2f(-1.0, 1.0)
	gl.Vertex2f(2.0, 0.0)
	gl.End()
	if shield {
		gl.Color3f(0.1, 0.1, 1.0)
		gl.Begin(gl.LINE_LOOP)
		for rad := 0.0; rad < 12.0; rad += 1.0 {
			gl.Vertex2f(
				float32(2.3*math.Cos(2*float64(rad)/math.Pi)+0.2),
				float32(2.0*math.Sin(2*float64(rad)/math.Pi)))
		}
		gl.End()
	}
	gl.PopMatrix()
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	drawShip(angle)
	drawBullets()
	glut.SwapBuffers()
}

func idle() {
	var delta int

	time := glut.Get(glut.ELAPSED_TIME)
	if resuming {
		lastTime = time
		resuming = false
	}
	if left {
		delta = time - leftTime
		angle = angle + float32(delta)*float32(0.4)
		leftTime = time
	}
	if right {
		delta = time - rightTime
		angle = angle - float32(delta)*float32(0.4)
		rightTime = time
	}
	if thrust {
		delta = time - thrustTime
		v = float32(delta) * 0.00004
		xv = xv + float32(math.Cos(float64(angle)*math.Pi/180.0))*v
		yv = yv + float32(math.Sin(float64(angle)*math.Pi/180.0))*v
		thrustTime = time
	}
	delta = time - lastTime
	x = x + xv*float32(delta)
	y = y + yv*float32(delta)
	x = x / 40.0
	x = (x - float32(math.Floor(float64(x)))) * 40.0
	y = y / 40.0
	y = (y - float32(math.Floor(float64(y)))) * 40.0
	lastTime = time
	advanceBullets(delta, time)
	currentWindow.PostRedisplay()
}

func visible(vis int) {
	if vis == glut.VISIBLE {
		if !paused {
			glut.IdleFunc(idle)
		}
	} else {
		glut.IdleFunc(nil)
	}
}

func key(key byte, px, py int) {
	switch key {
	case 27:
		os.Exit(0)
	case 'A':
	case 'a':
		thrust = true
		thrustTime = glut.Get(glut.ELAPSED_TIME)
	case 'S':
	case 's':
		shield = true
	case 'C':
	case 'c':
		cursor = !cursor
		if cursor {
			glut.SetCursor(glut.CURSOR_INHERIT)
		} else {
			glut.SetCursor(glut.CURSOR_NONE)
		}
	case 'z':
	case 'Z':
		x = 20
		y = 20
		xv = 0
		yv = 0
	case 'f':
		glut.GameModeString("640x480:32@60")
		glut.EnterGameMode()
		initWindow()
	case 'g':
		glut.GameModeString("800x600:32@60")
		glut.EnterGameMode()
		initWindow()
	case 'l':
		if originalWindow != 0 && currentWindow != originalWindow {
			glut.LeaveGameMode()
			currentWindow = originalWindow
		}
	case 'P':
	case 'p':
		paused = !paused
		if paused {
			glut.IdleFunc(nil)
		} else {
			glut.IdleFunc(idle)
			resuming = true
		}
	case 'Q':
	case 'q':
	case ' ':
		shotBullet()
	}
}

func keyup(key byte, x, y int) {
	switch key {
	case 'A':
	case 'a':
		thrust = false
	case 'S':
	case 's':
		shield = false
	}
}

func special(key, x, y int) {
	switch key {
	case glut.KEY_F1:
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		gl.Enable(gl.BLEND)
		gl.Enable(gl.LINE_SMOOTH)
		gl.Enable(gl.POINT_SMOOTH)
	case glut.KEY_F2:
		gl.Disable(gl.BLEND)
		gl.Disable(gl.LINE_SMOOTH)
		gl.Disable(gl.POINT_SMOOTH)
	case glut.KEY_UP:
		thrust = true
		thrustTime = glut.Get(glut.ELAPSED_TIME)
	case glut.KEY_LEFT:
		left = true
		leftTime = glut.Get(glut.ELAPSED_TIME)
	case glut.KEY_RIGHT:
		right = true
		rightTime = glut.Get(glut.ELAPSED_TIME)
	}
}

func specialup(key, x, y int) {
	switch key {
	case glut.KEY_UP:
		thrust = false
	case glut.KEY_LEFT:
		left = false
	case glut.KEY_RIGHT:
		right = false
	}
}

func joystick(buttons uint, x, y, z int) {
	if buttons&0x1 > 0 {
		thrust = true
		thrustTime = glut.Get(glut.ELAPSED_TIME)
		joyThrust = true
	} else {
		if joyThrust {
			thrust = false
			joyThrust = false
		}
	}
	if buttons&0x2 > 0 {
		shotBullet()
	}
	if buttons&0x4 > 0 {
		shield = true
		joyShield = true
	} else {
		if joyShield {
			shield = false
			joyShield = false
		}
	}
	if x < -300 {
		left = true
		leftTime = glut.Get(glut.ELAPSED_TIME)
		joyLeft = true
	} else {
		/* joyLeft helps avoid "joystick in neutral"
		   from continually stopping rotation. */
		if joyLeft {
			left = false
			joyLeft = false
		}
	}
	if x > 300 {
		right = true
		rightTime = glut.Get(glut.ELAPSED_TIME)
		joyRight = true
	} else {
		/* joyRight helps avoid "joystick in neutral"
		   from continually stopping rotation. */
		if joyRight {
			right = false
			joyRight = false
		}
	}
}

func initWindow() {
	glut.IgnoreKeyRepeat(1)

	glut.DisplayFunc(display)
	glut.VisibilityFunc(visible)
	glut.KeyboardFunc(key)
	glut.KeyboardUpFunc(keyup)
	glut.SpecialFunc(special)
	glut.SpecialUpFunc(specialup)
	glut.JoystickFunc(joystick, 100)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, 40, 0, 40, 0, 40)
	gl.MatrixMode(gl.MODELVIEW)
	gl.PointSize(3.0)

	currentWindow = glut.GetWindow()
}

func main() {
	glut.InitDisplayMode(glut.DOUBLE | glut.RGB)

	// FIXME: gamemode causes crash
	//if (argc > 1 && !strcmp(argv[1], "-fullscreen")) {
	//glut.GameModeString("800x600:32@60")
	//glut.EnterGameMode()
	// } else {
	originalWindow = glut.CreateWindow("asteroids")
	//}

	initWindow()

	glut.MainLoop()

}
