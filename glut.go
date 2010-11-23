package glut

// #ifdef __Darwin
// # include <GLUT/glut.h>
// #else
// # include <GL/glut.h>
// #endif
// #include <stdlib.h>
//
// // DEFINE_FUNCS is needed because we cannot pass C function pointers in Go.
// #define DEFINE_FUNCS(x, y, ...) \
// extern void go_##y(__VA_ARGS__); \
// void set##x##Func() { glut##x##Func(go_##y); } \
// void clear##x##Func() { glut##x##Func(NULL); } \
// 
// DEFINE_FUNCS(Display, a)
// DEFINE_FUNCS(OverlayDisplay, b)
// DEFINE_FUNCS(Reshape, c, int width, int height)
// DEFINE_FUNCS(Keyboard, d, unsigned char key, int x, int y)
// DEFINE_FUNCS(Mouse, e, int button, int state, int x, int y)
// DEFINE_FUNCS(Motion, f, int x, int y)
// DEFINE_FUNCS(PassiveMotion, g, int x, int y)
// DEFINE_FUNCS(Visibility, h, int state)
// DEFINE_FUNCS(Entry, i, int state)
// DEFINE_FUNCS(Special, j, int key, int x, int y)
// DEFINE_FUNCS(SpaceballMotion, k, int x, int y, int z)
// DEFINE_FUNCS(SpaceballRotate, l, int x, int y, int z)
// DEFINE_FUNCS(SpaceballButton, m, int button, int state)
// /* just in case you're on an SGI box.. :) */
// DEFINE_FUNCS(ButtonBox, n, int button, int state)
// DEFINE_FUNCS(Dials, o, int dial, int value)
// DEFINE_FUNCS(TabletMotion, p, int x, int y)
// DEFINE_FUNCS(TabletButton, q, int button, int state, int x, int y)
// DEFINE_FUNCS(MenuStatus, r, int status, int x, int y)
// DEFINE_FUNCS(Idle, s)
// // timer's an odd duck - we ignore it for now.
import "C"

//import "gl"
import "os"
import "unsafe"

type (
	Window int

	windowFuncs struct {
		display         func()
		overlayDisplay  func()
		reshape         func(width, height int)
		keyboard        func(key byte, x, y int)
		mouse           func(button, state, x, y int)
		motion          func(x, y int)
		passiveMotion   func(x, y int)
		visibility      func(state int)
		entry           func(state int)
		special         func(key, x, y int)
		spaceballMotion func(x, y, z int)
		spaceballRotate func(x, y, z int)
		spaceballButton func(button, state int)
		buttonBox       func(button, state int)
		dials           func(dial, value int)
		tabletMotion    func(x, y int)
		tabletButton    func(button, state, x, y int)
		menuStatus      func(status, x, y int)
		idle            func()
	}
)

const (
	RGB         = C.GLUT_RGB
	RGBA        = C.GLUT_RGBA
	INDEX       = C.GLUT_INDEX
	SINGLE      = C.GLUT_SINGLE
	DOUBLE      = C.GLUT_DOUBLE
	ACCUM       = C.GLUT_ACCUM
	ALPHA       = C.GLUT_ALPHA
	DEPTH       = C.GLUT_DEPTH
	STENCIL     = C.GLUT_STENCIL
	MULTISAMPLE = C.GLUT_MULTISAMPLE
	STEREO      = C.GLUT_STEREO
	LUMINANCE   = C.GLUT_LUMINANCE

	LEFT_BUTTON   = C.GLUT_LEFT_BUTTON
	MIDDLE_BUTTON = C.GLUT_MIDDLE_BUTTON
	RIGHT_BUTTON  = C.GLUT_RIGHT_BUTTON

	DOWN = C.GLUT_DOWN
	UP   = C.GLUT_UP

	KEY_F1  = C.GLUT_KEY_F1
	KEY_F2  = C.GLUT_KEY_F2
	KEY_F3  = C.GLUT_KEY_F3
	KEY_F4  = C.GLUT_KEY_F4
	KEY_F5  = C.GLUT_KEY_F5
	KEY_F6  = C.GLUT_KEY_F6
	KEY_F7  = C.GLUT_KEY_F7
	KEY_F8  = C.GLUT_KEY_F8
	KEY_F9  = C.GLUT_KEY_F9
	KEY_F10 = C.GLUT_KEY_F10
	KEY_F11 = C.GLUT_KEY_F11
	KEY_F12 = C.GLUT_KEY_F12

	KEY_LEFT      = C.GLUT_KEY_LEFT
	KEY_UP        = C.GLUT_KEY_UP
	KEY_RIGHT     = C.GLUT_KEY_RIGHT
	KEY_DOWN      = C.GLUT_KEY_DOWN
	KEY_PAGE_UP   = C.GLUT_KEY_PAGE_UP
	KEY_PAGE_DOWN = C.GLUT_KEY_PAGE_DOWN
	KEY_HOME      = C.GLUT_KEY_HOME
	KEY_END       = C.GLUT_KEY_END
	KEY_INSERT    = C.GLUT_KEY_INSERT

	LEFT    = C.GLUT_LEFT
	ENTERED = C.GLUT_ENTERED

	MENU_NOT_IN_USE = C.GLUT_MENU_NOT_IN_USE
	MENU_IN_USE     = C.GLUT_MENU_IN_USE

	NOT_VISIBLE = C.GLUT_NOT_VISIBLE
	VISIBLE     = C.GLUT_VISIBLE

	HIDDEN             = C.GLUT_HIDDEN
	FULLY_RETAINED     = C.GLUT_FULLY_RETAINED
	PARTIALLY_RETAINED = C.GLUT_PARTIALLY_RETAINED
	FULLY_COVERED      = C.GLUT_FULLY_COVERED

	RED   = C.GLUT_RED
	GREEN = C.GLUT_GREEN
	BLUE  = C.GLUT_BLUE

	NORMAL  = C.GLUT_NORMAL
	OVERLAY = C.GLUT_OVERLAY

	// STROKE_ROMAN          = C.GLUT_STROKE_ROMAN
	// STROKE_MONO_ROMAN     = C.GLUT_STROKE_MONO_ROMAN
	// BITMAP_9_BY_15        = C.GLUT_BITMAP_9_BY_15
	// BITMAP_8_BY_13        = C.GLUT_BITMAP_8_BY_13
	// BITMAP_TIMES_ROMAN_10 = C.GLUT_BITMAP_TIMES_ROMAN_10
	// BITMAP_TIMES_ROMAN_24 = C.GLUT_BITMAP_TIMES_ROMAN_24
	// BITMAP_HELVETICA_10   = C.GLUT_BITMAP_HELVETICA_10
	// BITMAP_HELVETICA_12   = C.GLUT_BITMAP_HELVETICA_12
	// BITMAP_HELVETICA_18   = C.GLUT_BITMAP_HELVETICA_18

	WINDOW_X                = C.GLUT_WINDOW_X
	WINDOW_Y                = C.GLUT_WINDOW_Y
	WINDOW_WIDTH            = C.GLUT_WINDOW_WIDTH
	WINDOW_HEIGHT           = C.GLUT_WINDOW_HEIGHT
	WINDOW_BUFFER_SIZE      = C.GLUT_WINDOW_BUFFER_SIZE
	WINDOW_STENCIL_SIZE     = C.GLUT_WINDOW_STENCIL_SIZE
	WINDOW_DEPTH_SIZE       = C.GLUT_WINDOW_DEPTH_SIZE
	WINDOW_RED_SIZE         = C.GLUT_WINDOW_RED_SIZE
	WINDOW_GREEN_SIZE       = C.GLUT_WINDOW_GREEN_SIZE
	WINDOW_BLUE_SIZE        = C.GLUT_WINDOW_BLUE_SIZE
	WINDOW_ALPHA_SIZE       = C.GLUT_WINDOW_ALPHA_SIZE
	WINDOW_ACCUM_RED_SIZE   = C.GLUT_WINDOW_ACCUM_RED_SIZE
	WINDOW_ACCUM_GREEN_SIZE = C.GLUT_WINDOW_ACCUM_GREEN_SIZE
	WINDOW_ACCUM_BLUE_SIZE  = C.GLUT_WINDOW_ACCUM_BLUE_SIZE
	WINDOW_ACCUM_ALPHA_SIZE = C.GLUT_WINDOW_ACCUM_ALPHA_SIZE
	WINDOW_DOUBLEBUFFER     = C.GLUT_WINDOW_DOUBLEBUFFER
	WINDOW_RGBA             = C.GLUT_WINDOW_RGBA
	WINDOW_PARENT           = C.GLUT_WINDOW_PARENT
	WINDOW_NUM_CHILDREN     = C.GLUT_WINDOW_NUM_CHILDREN
	WINDOW_COLORMAP_SIZE    = C.GLUT_WINDOW_COLORMAP_SIZE
	WINDOW_NUM_SAMPLES      = C.GLUT_WINDOW_NUM_SAMPLES
	WINDOW_STEREO           = C.GLUT_WINDOW_STEREO
	WINDOW_CURSOR           = C.GLUT_WINDOW_CURSOR
	SCREEN_WIDTH            = C.GLUT_SCREEN_WIDTH
	SCREEN_HEIGHT           = C.GLUT_SCREEN_HEIGHT
	SCREEN_WIDTH_MM         = C.GLUT_SCREEN_WIDTH_MM
	SCREEN_HEIGHT_MM        = C.GLUT_SCREEN_HEIGHT_MM
	MENU_NUM_ITEMS          = C.GLUT_MENU_NUM_ITEMS
	DISPLAY_MODE_POSSIBLE   = C.GLUT_DISPLAY_MODE_POSSIBLE
	INIT_WINDOW_X           = C.GLUT_INIT_WINDOW_X
	INIT_WINDOW_Y           = C.GLUT_INIT_WINDOW_Y
	INIT_WINDOW_WIDTH       = C.GLUT_INIT_WINDOW_WIDTH
	INIT_WINDOW_HEIGHT      = C.GLUT_INIT_WINDOW_HEIGHT
	INIT_DISPLAY_MODE       = C.GLUT_INIT_DISPLAY_MODE
	ELAPSED_TIME            = C.GLUT_ELAPSED_TIME
	WINDOW_FORMAT_ID        = C.GLUT_WINDOW_FORMAT_ID

	HAS_KEYBOARD             = C.GLUT_HAS_KEYBOARD
	HAS_MOUSE                = C.GLUT_HAS_MOUSE
	HAS_SPACEBALL            = C.GLUT_HAS_SPACEBALL
	HAS_DIAL_AND_BUTTON_BOX  = C.GLUT_HAS_DIAL_AND_BUTTON_BOX
	HAS_TABLET               = C.GLUT_HAS_TABLET
	NUM_MOUSE_BUTTONS        = C.GLUT_NUM_MOUSE_BUTTONS
	NUM_SPACEBALL_BUTTONS    = C.GLUT_NUM_SPACEBALL_BUTTONS
	NUM_BUTTON_BOX_BUTTONS   = C.GLUT_NUM_BUTTON_BOX_BUTTONS
	NUM_DIALS                = C.GLUT_NUM_DIALS
	NUM_TABLET_BUTTONS       = C.GLUT_NUM_TABLET_BUTTONS
	DEVICE_IGNORE_KEY_REPEAT = C.GLUT_DEVICE_IGNORE_KEY_REPEAT
	DEVICE_KEY_REPEAT        = C.GLUT_DEVICE_KEY_REPEAT
	HAS_JOYSTICK             = C.GLUT_HAS_JOYSTICK
	OWNS_JOYSTICK            = C.GLUT_OWNS_JOYSTICK
	JOYSTICK_BUTTONS         = C.GLUT_JOYSTICK_BUTTONS
	JOYSTICK_AXES            = C.GLUT_JOYSTICK_AXES
	JOYSTICK_POLL_RATE       = C.GLUT_JOYSTICK_POLL_RATE

	OVERLAY_POSSIBLE  = C.GLUT_OVERLAY_POSSIBLE
	LAYER_IN_USE      = C.GLUT_LAYER_IN_USE
	HAS_OVERLAY       = C.GLUT_HAS_OVERLAY
	TRANSPARENT_INDEX = C.GLUT_TRANSPARENT_INDEX
	NORMAL_DAMAGED    = C.GLUT_NORMAL_DAMAGED
	OVERLAY_DAMAGED   = C.GLUT_OVERLAY_DAMAGED

	VIDEO_RESIZE_POSSIBLE     = C.GLUT_VIDEO_RESIZE_POSSIBLE
	VIDEO_RESIZE_IN_USE       = C.GLUT_VIDEO_RESIZE_IN_USE
	VIDEO_RESIZE_X_DELTA      = C.GLUT_VIDEO_RESIZE_X_DELTA
	VIDEO_RESIZE_Y_DELTA      = C.GLUT_VIDEO_RESIZE_Y_DELTA
	VIDEO_RESIZE_WIDTH_DELTA  = C.GLUT_VIDEO_RESIZE_WIDTH_DELTA
	VIDEO_RESIZE_HEIGHT_DELTA = C.GLUT_VIDEO_RESIZE_HEIGHT_DELTA
	VIDEO_RESIZE_X            = C.GLUT_VIDEO_RESIZE_X
	VIDEO_RESIZE_Y            = C.GLUT_VIDEO_RESIZE_Y
	VIDEO_RESIZE_WIDTH        = C.GLUT_VIDEO_RESIZE_WIDTH
	VIDEO_RESIZE_HEIGHT       = C.GLUT_VIDEO_RESIZE_HEIGHT

	ACTIVE_SHIFT = C.GLUT_ACTIVE_SHIFT
	ACTIVE_CTRL  = C.GLUT_ACTIVE_CTRL
	ACTIVE_ALT   = C.GLUT_ACTIVE_ALT

	CURSOR_RIGHT_ARROW         = C.GLUT_CURSOR_RIGHT_ARROW
	CURSOR_LEFT_ARROW          = C.GLUT_CURSOR_LEFT_ARROW
	CURSOR_INFO                = C.GLUT_CURSOR_INFO
	CURSOR_DESTROY             = C.GLUT_CURSOR_DESTROY
	CURSOR_HELP                = C.GLUT_CURSOR_HELP
	CURSOR_CYCLE               = C.GLUT_CURSOR_CYCLE
	CURSOR_SPRAY               = C.GLUT_CURSOR_SPRAY
	CURSOR_WAIT                = C.GLUT_CURSOR_WAIT
	CURSOR_TEXT                = C.GLUT_CURSOR_TEXT
	CURSOR_CROSSHAIR           = C.GLUT_CURSOR_CROSSHAIR
	CURSOR_UP_DOWN             = C.GLUT_CURSOR_UP_DOWN
	CURSOR_LEFT_RIGHT          = C.GLUT_CURSOR_LEFT_RIGHT
	CURSOR_TOP_SIDE            = C.GLUT_CURSOR_TOP_SIDE
	CURSOR_BOTTOM_SIDE         = C.GLUT_CURSOR_BOTTOM_SIDE
	CURSOR_LEFT_SIDE           = C.GLUT_CURSOR_LEFT_SIDE
	CURSOR_RIGHT_SIDE          = C.GLUT_CURSOR_RIGHT_SIDE
	CURSOR_TOP_LEFT_CORNER     = C.GLUT_CURSOR_TOP_LEFT_CORNER
	CURSOR_TOP_RIGHT_CORNER    = C.GLUT_CURSOR_TOP_RIGHT_CORNER
	CURSOR_BOTTOM_RIGHT_CORNER = C.GLUT_CURSOR_BOTTOM_RIGHT_CORNER
	CURSOR_BOTTOM_LEFT_CORNER  = C.GLUT_CURSOR_BOTTOM_LEFT_CORNER
	CURSOR_INHERIT             = C.GLUT_CURSOR_INHERIT
	CURSOR_NONE                = C.GLUT_CURSOR_NONE
	CURSOR_FULL_CROSSHAIR      = C.GLUT_CURSOR_FULL_CROSSHAIR
)

var (
	idleFunc func()
	winFuncs = make(map[Window]*windowFuncs)
)

// Initialization

func Init() {
	argc := C.int(len(os.Args))
	argv := make([]*C.char, argc)
	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
	}

	C.glutInit(&argc, &argv[0])

	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
}

func InitWindowPosition(x, y int) {
	C.glutInitWindowPosition(C.int(x), C.int(y))
}

func InitWindowSize(width, height int) {
	C.glutInitWindowSize(C.int(width), C.int(height))
}

func InitDisplayMode(mode uint) {
	C.glutInitDisplayMode(C.uint(mode))
}

// Beginning Event Processing

func MainLoop() {
	C.glutMainLoop()
}

// Window Management

func CreateWindow(title string) (w Window) {
	ctitle := C.CString(title)
	w = Window(C.glutCreateWindow(ctitle))
	C.free(unsafe.Pointer(ctitle))

	winFuncs[w] = new(windowFuncs)

	return
}

func (w Window) CreateSubWindow(x, y, width, height int) Window {
	return Window(C.glutCreateSubWindow(C.int(w), C.int(x), C.int(y), C.int(width), C.int(height)))
}

func SetWindow(window Window) {
	C.glutSetWindow(C.int(window))
}

func GetWindow() Window {
	return Window(C.glutGetWindow())
}

func (w Window) Destroy() {
	C.glutDestroyWindow(C.int(w))

	winFuncs[w] = nil, false
}

func PostRedisplay() {
	C.glutPostRedisplay()
}

func SwapBuffers() {
	C.glutSwapBuffers()
}

func PositionWindow(x, y int) {
	C.glutPositionWindow(C.int(x), C.int(y))
}

func ReshapeWindow(width, height int) {
	C.glutReshapeWindow(C.int(width), C.int(height))
}

func FullScreen() {
	C.glutFullScreen()
}

func PopWindow() {
	C.glutPopWindow()
}

func PushWindow() {
	C.glutPushWindow()
}

func ShowWindow() {
	C.glutShowWindow()
}

func HideWindow() {
	C.glutHideWindow()
}

func IconifyWindow() {
	C.glutIconifyWindow()
}

func SetWindowTitle(name string) {
	cname := C.CString(name)
	C.glutSetWindowTitle(cname)
	C.free(unsafe.Pointer(cname))
}

func SetIconTitle(name string) {
	cname := C.CString(name)
	C.glutSetIconTitle(cname)
	C.free(unsafe.Pointer(cname))
}

func SetCursor(cursor int) {
	C.glutSetCursor(C.int(cursor))
}

// Overlay Management

// Menu Management

// Callback Registration

func DisplayFunc(display func()) {
	if display == nil {
		panic("nil display func") // glut forbids this
	}
	winFuncs[GetWindow()].display = display
	C.setDisplayFunc()
}

func OverlayDisplayFunc(overlayDisplay func()) {
	winFuncs[GetWindow()].overlayDisplay = overlayDisplay
	if overlayDisplay != nil {
		C.setOverlayDisplayFunc()
	} else {
		C.clearOverlayDisplayFunc()
	}
}

func ReshapeFunc(reshape func(width, height int)) {
	winFuncs[GetWindow()].reshape = reshape
	if reshape != nil {
		C.setReshapeFunc()
	} else {
		C.clearReshapeFunc()
	}
}

func KeyboardFunc(keyboard func(key byte, x, y int)) {
	winFuncs[GetWindow()].keyboard = keyboard
	if keyboard != nil {
		C.setKeyboardFunc()
	} else {
		C.clearKeyboardFunc()
	}
}

func MouseFunc(mouse func(button, state, x, y int)) {
	winFuncs[GetWindow()].mouse = mouse
	if mouse != nil {
		C.setMouseFunc()
	} else {
		C.clearMouseFunc()
	}
}

func MotionFunc(motion func(x, y int)) {
	winFuncs[GetWindow()].motion = motion
	if motion != nil {
		C.setMotionFunc()
	} else {
		C.clearMotionFunc()
	}
}

func PassiveMotionFunc(passiveMotion func(x, y int)) {
	winFuncs[GetWindow()].passiveMotion = passiveMotion
	if passiveMotion != nil {
		C.setPassiveMotionFunc()
	} else {
		C.clearPassiveMotionFunc()
	}
}

func VisibilityFunc(visibility func(state int)) {
	winFuncs[GetWindow()].visibility = visibility
	if visibility != nil {
		C.setVisibilityFunc()
	} else {
		C.clearVisibilityFunc()
	}
}

func EntryFunc(entry func(state int)) {
	winFuncs[GetWindow()].entry = entry
	if entry != nil {
		C.setEntryFunc()
	} else {
		C.clearEntryFunc()
	}
}

func SpecialFunc(special func(key, x, y int)) {
	winFuncs[GetWindow()].special = special
	if special != nil {
		C.setSpecialFunc()
	} else {
		C.clearSpecialFunc()
	}
}

func SpaceballMotionFunc(spaceballMotion func(x, y, z int)) {
	winFuncs[GetWindow()].spaceballMotion = spaceballMotion
	if spaceballMotion != nil {
		C.setSpaceballMotionFunc()
	} else {
		C.clearSpaceballMotionFunc()
	}
}

func SpaceballRotateFunc(spaceballRotate func(x, y, z int)) {
	winFuncs[GetWindow()].spaceballRotate = spaceballRotate
	if spaceballRotate != nil {
		C.setSpaceballRotateFunc()
	} else {
		C.clearSpaceballRotateFunc()
	}
}

func SpaceballButtonFunc(spaceballButton func(button, state int)) {
	winFuncs[GetWindow()].spaceballButton = spaceballButton
	if spaceballButton != nil {
		C.setSpaceballButtonFunc()
	} else {
		C.clearSpaceballButtonFunc()
	}
}

func ButtonBoxFunc(buttonBox func(button, state int)) {
	winFuncs[GetWindow()].buttonBox = buttonBox
	if buttonBox != nil {
		C.setButtonBoxFunc()
	} else {
		C.clearButtonBoxFunc()
	}
}

func DialsFunc(dials func(dial, value int)) {
	winFuncs[GetWindow()].dials = dials
	if dials != nil {
		C.setDialsFunc()
	} else {
		C.clearDialsFunc()
	}
}

func TabletMotionFunc(tabletMotion func(x, y int)) {
	winFuncs[GetWindow()].tabletMotion = tabletMotion
	if tabletMotion != nil {
		C.setTabletMotionFunc()
	} else {
		C.clearTabletMotionFunc()
	}
}

func TabletButtonFunc(tabletButton func(button, state, x, y int)) {
	winFuncs[GetWindow()].tabletButton = tabletButton
	if tabletButton != nil {
		C.setTabletButtonFunc()
	} else {
		C.clearTabletButtonFunc()
	}
}

func MenuStatusFunc(menuStatus func(status, x, y int)) {
	winFuncs[GetWindow()].menuStatus = menuStatus
	if menuStatus != nil {
		C.setMenuStatusFunc()
	} else {
		C.clearMenuStatusFunc()
	}
}

func IdleFunc(idle func()) {
	idleFunc = idle
	if idle != nil {
		C.setIdleFunc()
	} else {
		C.clearIdleFunc()
	}
}

// Color Index Colormap Management

// State Retrieval

// Font Rendering

// Geometric Object Rendering

// Callbacks

// cgo does not allow callbacks to arbitrary functions, so we must handle this 
// ourselves.

// odd export names work around go bug on mac os x

//export go_a
func InternalDisplayFunc() {
	winFuncs[GetWindow()].display()
}

//export go_b
func InternalOverlayDisplayFunc() {
	winFuncs[GetWindow()].overlayDisplay()
}

//export go_c
func InternalReshapeFunc(width, height int) {
	winFuncs[GetWindow()].reshape(width, height)
}

//export go_d
func InternalKeyboardFunc(key uint8, x, y int) {
	winFuncs[GetWindow()].keyboard(key, x, y)
}

//export go_e
func InternalMouseFunc(button, state, x, y int) {
	winFuncs[GetWindow()].mouse(button, state, x, y)
}

//export go_f
func InternalMotionFunc(x, y int) {
	winFuncs[GetWindow()].motion(x, y)
}

//export go_g
func InternalPassiveMotionFunc(x, y int) {
	winFuncs[GetWindow()].passiveMotion(x, y)
}

//export go_h
func InternalVisibilityFunc(state int) {
	winFuncs[GetWindow()].visibility(state)
}

//export go_i
func InternalEntryFunc(state int) {
	winFuncs[GetWindow()].entry(state)
}

//export go_j
func InternalSpecialFunc(key, x, y int) {
	winFuncs[GetWindow()].special(key, x, y)
}

//export go_k
func InternalSpaceballMotionFunc(x, y, z int) {
	winFuncs[GetWindow()].spaceballMotion(x, y, z)
}

//export go_l
func InternalSpaceballRotateFunc(x, y, z int) {
	winFuncs[GetWindow()].spaceballRotate(x, y, z)
}

//export go_m
func InternalSpaceballButtonFunc(button, state int) {
	winFuncs[GetWindow()].spaceballButton(button, state)
}

//export go_n
func InternalButtonBoxFunc(button, state int) {
	winFuncs[GetWindow()].buttonBox(button, state)
}

//export go_o
func InternalDialsFunc(dial, value int) {
	winFuncs[GetWindow()].dials(dial, value)
}

//export go_p
func InternalTabletMotionFunc(x, y int) {
	winFuncs[GetWindow()].tabletMotion(x, y)
}

//export go_q
func InternalTabletButtonFunc(button, state, x, y int) {
	winFuncs[GetWindow()].tabletButton(button, state, x, y)
}

//export go_r
func InternalMenuStatusFunc(status, x, y int) {
	winFuncs[GetWindow()].menuStatus(status, x, y)
}

//export go_s
func InternalIdleFunc() {
	idleFunc()
}
