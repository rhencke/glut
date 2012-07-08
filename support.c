#ifdef __APPLE__
# include <GLUT/glut.h>
#else
# include <GL/glut.h>
#endif
#include <stdlib.h>
#include "_cgo_export.h"

// DEFINE_FUNCS is needed because I'm unsure how to pass C function pointers in Go.
#define DEFINE_FUNCS(x, ...) \
void set##x##Func() { glut##x##Func(internal##x##Func); } \
void clear##x##Func() { glut##x##Func(0); } \

DEFINE_FUNCS(Display)
DEFINE_FUNCS(OverlayDisplay)
DEFINE_FUNCS(Reshape, int width, int height)
DEFINE_FUNCS(Keyboard, unsigned char key, int x, int y)
DEFINE_FUNCS(Mouse, int button, int state, int x, int y)
DEFINE_FUNCS(Motion, int x, int y)
DEFINE_FUNCS(PassiveMotion, int x, int y)
DEFINE_FUNCS(Visibility, int state)
DEFINE_FUNCS(Entry, int state)
DEFINE_FUNCS(Special, int key, int x, int y)
DEFINE_FUNCS(SpaceballMotion, int x, int y, int z)
DEFINE_FUNCS(SpaceballRotate, int x, int y, int z)
DEFINE_FUNCS(SpaceballButton, int button, int state)
/* just in case you're on an SGI box.. :) */
DEFINE_FUNCS(ButtonBox, int button, int state)
DEFINE_FUNCS(Dials, int dial, int value)
DEFINE_FUNCS(TabletMotion, int x, int y)
DEFINE_FUNCS(TabletButton, int button, int state, int x, int y)
DEFINE_FUNCS(MenuStatus, int status, int x, int y)
DEFINE_FUNCS(Idle)
DEFINE_FUNCS(KeyboardUp, unsigned char key, int x, int y)
DEFINE_FUNCS(SpecialUp, int key, int x, int y)

// glutCreateMenu callback
int goCreateMenu() { return glutCreateMenu(internalMenuFunc); }
int goCreateMenuWithoutCallback() { return glutCreateMenu(0); }

// glutJoystickFunc callback
void setJoystickFunc(int pollInterval) { glutJoystickFunc(internalJoystickFunc, pollInterval); }
void clearJoystickFunc(int pollInterval) { glutJoystickFunc(0, pollInterval); }

// cgo does not correctly interpret the GLUT font constants, so we try a different approach.
#define DEFINE_FONT(x) void* go_##x() { return x; }

DEFINE_FONT(GLUT_STROKE_MONO_ROMAN)
DEFINE_FONT(GLUT_STROKE_ROMAN)
DEFINE_FONT(GLUT_BITMAP_9_BY_15)
DEFINE_FONT(GLUT_BITMAP_8_BY_13)
DEFINE_FONT(GLUT_BITMAP_TIMES_ROMAN_10)
DEFINE_FONT(GLUT_BITMAP_TIMES_ROMAN_24)
DEFINE_FONT(GLUT_BITMAP_HELVETICA_10)
DEFINE_FONT(GLUT_BITMAP_HELVETICA_12)
DEFINE_FONT(GLUT_BITMAP_HELVETICA_18)