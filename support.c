// DEFINE_FUNCS is needed because I'm unsure how to pass C function pointers in Go.
#define DEFINE_FUNCS(x, y, ...) \
extern void go_##y(__VA_ARGS__); \
void set##x##Func() { glut##x##Func(go_##y); } \
void clear##x##Func() { glut##x##Func(NULL); } \

DEFINE_FUNCS(Display, a)
DEFINE_FUNCS(OverlayDisplay, b)
DEFINE_FUNCS(Reshape, c, int width, int height)
DEFINE_FUNCS(Keyboard, d, unsigned char key, int x, int y)
DEFINE_FUNCS(Mouse, e, int button, int state, int x, int y)
DEFINE_FUNCS(Motion, f, int x, int y)
DEFINE_FUNCS(PassiveMotion, g, int x, int y)
DEFINE_FUNCS(Visibility, h, int state)
DEFINE_FUNCS(Entry, i, int state)
DEFINE_FUNCS(Special, j, int key, int x, int y)
DEFINE_FUNCS(SpaceballMotion, k, int x, int y, int z)
DEFINE_FUNCS(SpaceballRotate, l, int x, int y, int z)
DEFINE_FUNCS(SpaceballButton, m, int button, int state)
/* just in case you're on an SGI box.. :) */
DEFINE_FUNCS(ButtonBox, n, int button, int state)
DEFINE_FUNCS(Dials, o, int dial, int value)
DEFINE_FUNCS(TabletMotion, p, int x, int y)
DEFINE_FUNCS(TabletButton, q, int button, int state, int x, int y)
DEFINE_FUNCS(MenuStatus, r, int status, int x, int y)
DEFINE_FUNCS(Idle, s)
DEFINE_FUNCS(KeyboardUp, u, unsigned char key, int x, int y)
DEFINE_FUNCS(SpecialUp, v, int key, int x, int y)

// glutCreateMenu callback
extern void go_t(int value); // 
int goCreateMenu() { return glutCreateMenu(go_t); }
int goCreateMenuWithoutCallback() { return glutCreateMenu(NULL); }

// glutJoystickFunc callback
extern void go_w(unsigned int buttonMask, int x, int y, int z);
void setJoystickFunc(int pollInterval) { glutJoystickFunc(go_w, pollInterval); }
void clearJoystickFunc(int pollInterval) { glutJoystickFunc(NULL, pollInterval); }

// cgo does not correctly interpret the GLUT font constants, so we try a different approach.
#define DEFINE_FONT(x) void* go_##x() { return x; }
DEFINE_FONT(GLUT_STROKE_ROMAN)
DEFINE_FONT(GLUT_STROKE_MONO_ROMAN)
DEFINE_FONT(GLUT_BITMAP_9_BY_15)
DEFINE_FONT(GLUT_BITMAP_8_BY_13)
DEFINE_FONT(GLUT_BITMAP_TIMES_ROMAN_10)
DEFINE_FONT(GLUT_BITMAP_TIMES_ROMAN_24)
DEFINE_FONT(GLUT_BITMAP_HELVETICA_10)
DEFINE_FONT(GLUT_BITMAP_HELVETICA_12)
DEFINE_FONT(GLUT_BITMAP_HELVETICA_18)