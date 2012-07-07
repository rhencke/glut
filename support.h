// DEFINE_FUNCS is needed because I'm unsure how to pass C function pointers in Go.
#define DECLARE_FUNCS(x, ...) \
extern void set##x##Func(); \
extern void clear##x##Func();

DECLARE_FUNCS(Display)
DECLARE_FUNCS(OverlayDisplay)
DECLARE_FUNCS(Reshape, int width, int height)
DECLARE_FUNCS(Keyboard, unsigned char key, int x, int y)
DECLARE_FUNCS(Mouse, int button, int state, int x, int y)
DECLARE_FUNCS(Motion, int x, int y)
DECLARE_FUNCS(PassiveMotion, int x, int y)
DECLARE_FUNCS(Visibility, int state)
DECLARE_FUNCS(Entry, int state)
DECLARE_FUNCS(Special, int key, int x, int y)
DECLARE_FUNCS(SpaceballMotion, int x, int y, int z)
DECLARE_FUNCS(SpaceballRotate, int x, int y, int z)
DECLARE_FUNCS(SpaceballButton, int button, int state)
/* just in case you're on an SGI box.. :) */
DECLARE_FUNCS(ButtonBox, int button, int state)
DECLARE_FUNCS(Dials, int dial, int value)
DECLARE_FUNCS(TabletMotion, int x, int y)
DECLARE_FUNCS(TabletButton, int button, int state, int x, int y)
DECLARE_FUNCS(MenuStatus, int status, int x, int y)
DECLARE_FUNCS(Idle)
DECLARE_FUNCS(KeyboardUp, unsigned char key, int x, int y)
DECLARE_FUNCS(SpecialUp, int key, int x, int y)

// glutCreateMenu callback
extern int goCreateMenu();
extern int goCreateMenuWithoutCallback();

// glutJoystickFunc callback
extern void setJoystickFunc(int pollInterval);
extern void clearJoystickFunc(int pollInterval);

// cgo does not correctly interpret the GLUT font constants, so we try a different approach.
#define DECLARE_FONT(x) extern void* go_##x();

DECLARE_FONT(GLUT_STROKE_MONO_ROMAN)
DECLARE_FONT(GLUT_STROKE_ROMAN)
DECLARE_FONT(GLUT_BITMAP_9_BY_15)
DECLARE_FONT(GLUT_BITMAP_8_BY_13)
DECLARE_FONT(GLUT_BITMAP_TIMES_ROMAN_10)
DECLARE_FONT(GLUT_BITMAP_TIMES_ROMAN_24)
DECLARE_FONT(GLUT_BITMAP_HELVETICA_10)
DECLARE_FONT(GLUT_BITMAP_HELVETICA_12)
DECLARE_FONT(GLUT_BITMAP_HELVETICA_18)