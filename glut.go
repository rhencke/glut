// The glut package provides a platform-independent window system for OpenGL, through the system's native GLUT or freeglut library.
// It also provides convenience functions for drawing teapots.

package glut

import (
	"os"
	"runtime"
	"unsafe"

	"github.com/go-gl/gl"
)

// #cgo darwin  LDFLAGS: -framework GLUT
// #cgo linux   LDFLAGS: -lglut
// #cgo windows LDFLAGS: -lglut32
// #ifdef __APPLE__
// # include <GLUT/glut.h>
// #else
// # include <GL/glut.h>
// #endif
// #include <stdlib.h>
// #include "support.h"
import "C"

type (
	Window     C.int
	BitmapFont C.int
	StrokeFont C.int
	Menu       C.int
)

type windowFuncs struct {
	display         func()
	overlayDisplay  func()
	reshape         func(width, height int)
	keyboard        func(key byte, x, y int)
	mouse           func(button, state, x, y int)
	motion          func(x, y int32)
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
	windowStatus    func(state int)
	keyboardUp      func(key byte, x, y int)
	specialUp       func(key, x, y int)
	joystick        func(buttonMask uint, x, y, z int)
}

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

	STROKE_ROMAN          StrokeFont = iota //C.GLUT_STROKE_ROMAN
	STROKE_MONO_ROMAN     StrokeFont = iota //C.GLUT_STROKE_MONO_ROMAN
	BITMAP_9_BY_15        BitmapFont = iota //C.GLUT_BITMAP_9_BY_15
	BITMAP_8_BY_13        BitmapFont = iota //C.GLUT_BITMAP_8_BY_13
	BITMAP_TIMES_ROMAN_10 BitmapFont = iota //C.GLUT_BITMAP_TIMES_ROMAN_10
	BITMAP_TIMES_ROMAN_24 BitmapFont = iota //C.GLUT_BITMAP_TIMES_ROMAN_24
	BITMAP_HELVETICA_10   BitmapFont = iota //C.GLUT_BITMAP_HELVETICA_10
	BITMAP_HELVETICA_12   BitmapFont = iota //C.GLUT_BITMAP_HELVETICA_12
	BITMAP_HELVETICA_18   BitmapFont = iota //C.GLUT_BITMAP_HELVETICA_18

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

	KEY_REPEAT_OFF     = C.GLUT_KEY_REPEAT_OFF
	KEY_REPEAT_ON      = C.GLUT_KEY_REPEAT_ON
	KEY_REPEAT_DEFAULT = C.GLUT_KEY_REPEAT_DEFAULT

	JOYSTICK_BUTTON_A = C.GLUT_JOYSTICK_BUTTON_A
	JOYSTICK_BUTTON_B = C.GLUT_JOYSTICK_BUTTON_B
	JOYSTICK_BUTTON_C = C.GLUT_JOYSTICK_BUTTON_C
	JOYSTICK_BUTTON_D = C.GLUT_JOYSTICK_BUTTON_D

	GAME_MODE_ACTIVE          = C.GLUT_GAME_MODE_ACTIVE
	GAME_MODE_POSSIBLE        = C.GLUT_GAME_MODE_POSSIBLE
	GAME_MODE_WIDTH           = C.GLUT_GAME_MODE_WIDTH
	GAME_MODE_HEIGHT          = C.GLUT_GAME_MODE_HEIGHT
	GAME_MODE_PIXEL_DEPTH     = C.GLUT_GAME_MODE_PIXEL_DEPTH
	GAME_MODE_REFRESH_RATE    = C.GLUT_GAME_MODE_REFRESH_RATE
	GAME_MODE_DISPLAY_CHANGED = C.GLUT_GAME_MODE_DISPLAY_CHANGED
)

var (
	idleFunc  func()
	winFuncs  = make(map[Window]*windowFuncs)
	menuFuncs = make(map[Menu]func(value int))
)

var gameWindow *Window

// - Initialization

func init() {
	runtime.LockOSThread()

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

func InitDisplayString(str string) {
	cstr := C.CString(str)
	C.glutInitDisplayString(cstr)
	C.free(unsafe.Pointer(cstr))
}

// - Beginning Event Processing

func MainLoop() {
	C.glutMainLoop()
}

// - Window Management

func registerWindow(w Window) {
	winFuncs[w] = new(windowFuncs)
}

func unregisterWindow(w Window) {
	delete(winFuncs, w)
}

func CreateWindow(title string) (w Window) {
	ctitle := C.CString(title)
	w = Window(C.glutCreateWindow(ctitle))
	C.free(unsafe.Pointer(ctitle))

	registerWindow(w)

	return
}

func (w Window) CreateSubWindow(x, y, width, height int) Window {
	neww := Window(C.glutCreateSubWindow(C.int(w), C.int(x), C.int(y), C.int(width), C.int(height)))

	registerWindow(neww)

	return neww
}

func SetWindow(window Window) {
	C.glutSetWindow(C.int(window))
}

func GetWindow() Window {
	return Window(C.glutGetWindow())
}

func (w Window) Destroy() {
	C.glutDestroyWindow(C.int(w))

	unregisterWindow(w)
}

func PostRedisplay() {
	C.glutPostRedisplay()
}

func (w Window) PostRedisplay() {
	C.glutPostWindowRedisplay(C.int(w))
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

func WarpPointer(x, y int) {
	C.glutWarpPointer(C.int(x), C.int(y))
}

// - Overlay Management

func EstablishOverlay() {
	C.glutEstablishOverlay()
}

func UseLayer(layer gl.GLenum) {
	C.glutUseLayer(C.GLenum(layer))
}

func RemoveOverlay() {
	C.glutRemoveOverlay()
}

func PostOverlayRedisplay() {
	C.glutPostOverlayRedisplay()
}

func (w Window) PostOverlayRedisplay() {
	C.glutPostWindowOverlayRedisplay(C.int(w))
}

func ShowOverlay() {
	C.glutShowOverlay()
}

func HideOverlay() {
	C.glutHideOverlay()
}

// - Menu Management

func CreateMenu(menu func(value int)) (m Menu) {
	if menu != nil {
		m = Menu(C.goCreateMenu())
	} else {
		m = Menu(C.goCreateMenuWithoutCallback())
	}

	menuFuncs[m] = menu

	return
}

func SetMenu(menu Menu) {
	C.glutSetMenu(C.int(menu))
}

func GetMenu() Menu {
	return Menu(C.glutGetMenu())
}

func (m Menu) Destroy() {
	C.glutDestroyMenu(C.int(m))
	delete(menuFuncs, m)
}

func AddMenuEntry(name string, value int) {
	cname := C.CString(name)
	C.glutAddMenuEntry(cname, C.int(value))
	C.free(unsafe.Pointer(cname))
}

func AddSubMenu(name string, value Menu) {
	cname := C.CString(name)
	C.glutAddSubMenu(cname, C.int(value))
	C.free(unsafe.Pointer(cname))
}

func ChangeToMenuEntry(entry int, name string, value int) {
	cname := C.CString(name)
	C.glutChangeToMenuEntry(C.int(entry), cname, C.int(value))
	C.free(unsafe.Pointer(cname))
}

func ChangeToSubMenu(entry int, name string, value int) {
	cname := C.CString(name)
	C.glutChangeToSubMenu(C.int(entry), cname, C.int(value))
	C.free(unsafe.Pointer(cname))
}

func RemoveMenuItem(entry int) {
	C.glutRemoveMenuItem(C.int(entry))
}

func AttachMenu(button int) {
	C.glutAttachMenu(C.int(button))
}

func DetachMenu(button int) {
	C.glutDetachMenu(C.int(button))
}

// - Callback Registration

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

func MotionFunc(motion func(x, y int32)) {
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

func KeyboardUpFunc(keyboardUp func(key byte, x, y int)) {
	winFuncs[GetWindow()].keyboardUp = keyboardUp
	if keyboardUp != nil {
		C.setKeyboardUpFunc()
	} else {
		C.clearKeyboardUpFunc()
	}
}

func SpecialUpFunc(specialUp func(key, x, y int)) {
	winFuncs[GetWindow()].specialUp = specialUp
	if specialUp != nil {
		C.setSpecialUpFunc()
	} else {
		C.clearSpecialUpFunc()
	}
}

func JoystickFunc(joystick func(buttonMask uint, x, y, z int), pollInterval int) {
	winFuncs[GetWindow()].joystick = joystick
	if joystick != nil {
		C.setJoystickFunc(C.int(pollInterval))
	} else {
		C.clearJoystickFunc(C.int(pollInterval))
	}
}

// - Color Index Colormap Management

func SetColor(cell int, red, green, blue gl.GLfloat) {
	C.glutSetColor(C.int(cell), C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
}

func GetColor(cell int) (red, green, blue gl.GLfloat) {
	ccell := C.int(cell)

	red = gl.GLfloat(C.glutGetColor(ccell, RED))
	green = gl.GLfloat(C.glutGetColor(ccell, GREEN))
	blue = gl.GLfloat(C.glutGetColor(ccell, BLUE))

	return
}

func CopyColormap(win Window) {
	C.glutCopyColormap(C.int(win))
}

// - State Retrieval

func Get(state gl.GLenum) int {
	return int(C.glutGet(C.GLenum(state)))
}

func LayerGet(info gl.GLenum) int {
	return int(C.glutLayerGet(C.GLenum(info)))
}

func DeviceGet(info gl.GLenum) int {
	return int(C.glutDeviceGet(C.GLenum(info)))
}

func GetModifiers() int {
	return int(C.glutGetModifiers())
}

func ExtensionSupported(extension string) (supported bool) {
	cextension := C.CString(extension)
	supported = C.glutExtensionSupported(cextension) != 0
	C.free(unsafe.Pointer(cextension))
	return
}

// - Font Rendering

func fontaddr(f int) unsafe.Pointer {
	switch f {
	case int(STROKE_ROMAN):
		return unsafe.Pointer(C.go_GLUT_STROKE_ROMAN())
	case int(STROKE_MONO_ROMAN):
		return unsafe.Pointer(C.go_GLUT_STROKE_MONO_ROMAN())
	case int(BITMAP_9_BY_15):
		return unsafe.Pointer(C.go_GLUT_BITMAP_9_BY_15())
	case int(BITMAP_8_BY_13):
		return unsafe.Pointer(C.go_GLUT_BITMAP_8_BY_13())
	case int(BITMAP_TIMES_ROMAN_10):
		return unsafe.Pointer(C.go_GLUT_BITMAP_TIMES_ROMAN_10())
	case int(BITMAP_TIMES_ROMAN_24):
		return unsafe.Pointer(C.go_GLUT_BITMAP_TIMES_ROMAN_24())
	case int(BITMAP_HELVETICA_10):
		return unsafe.Pointer(C.go_GLUT_BITMAP_HELVETICA_10())
	case int(BITMAP_HELVETICA_12):
		return unsafe.Pointer(C.go_GLUT_BITMAP_HELVETICA_12())
	case int(BITMAP_HELVETICA_18):
		return unsafe.Pointer(C.go_GLUT_BITMAP_HELVETICA_18())
	}
	panic("unknown font")
}

func (b BitmapFont) Character(character rune) {
	C.glutBitmapCharacter(fontaddr(int(b)), C.int(character))
}

func (b BitmapFont) Width(character rune) int {
	return int(C.glutBitmapWidth(fontaddr(int(b)), C.int(character)))
}

func (b BitmapFont) Length(str string) int {
	cstr := C.CString(str)
	strlen := C.glutBitmapLength(fontaddr(int(b)), (*C.uchar)(unsafe.Pointer(cstr)))
	C.free(unsafe.Pointer(cstr))
	return int(strlen)
}

func (s StrokeFont) Character(character rune) {
	C.glutStrokeCharacter(fontaddr(int(s)), C.int(character))
}

func (s StrokeFont) Width(character rune) int {
	return int(C.glutStrokeWidth(fontaddr(int(s)), C.int(character)))
}

func (s StrokeFont) Length(str string) int {
	cstr := C.CString(str)
	strlen := C.glutStrokeLength(fontaddr(int(s)), (*C.uchar)(unsafe.Pointer(cstr)))
	C.free(unsafe.Pointer(cstr))
	return int(strlen)
}

// - Geometric Object Rendering

func SolidSphere(radius gl.GLdouble, slices, stacks gl.GLint) {
	C.glutSolidSphere(C.GLdouble(radius), C.GLint(slices), C.GLint(stacks))
}

func WireSphere(radius gl.GLdouble, slices, stacks gl.GLint) {
	C.glutWireSphere(C.GLdouble(radius), C.GLint(slices), C.GLint(stacks))
}

func SolidCube(size gl.GLdouble) {
	C.glutSolidCube(C.GLdouble(size))
}

func WireCube(size gl.GLdouble) {
	C.glutWireCube(C.GLdouble(size))
}

func SolidCone(base, height gl.GLdouble, slices, stacks gl.GLint) {
	C.glutSolidCone(C.GLdouble(base), C.GLdouble(height), C.GLint(slices), C.GLint(stacks))
}

func WireCone(base, height gl.GLdouble, slices, stacks gl.GLint) {
	C.glutWireCone(C.GLdouble(base), C.GLdouble(height), C.GLint(slices), C.GLint(stacks))
}

func SolidTorus(innerRadius, outerRadius gl.GLdouble, nsides, rings gl.GLint) {
	C.glutSolidTorus(C.GLdouble(innerRadius), C.GLdouble(outerRadius), C.GLint(nsides), C.GLint(rings))
}

func WireTorus(innerRadius, outerRadius gl.GLdouble, nsides, rings gl.GLint) {
	C.glutWireTorus(C.GLdouble(innerRadius), C.GLdouble(outerRadius), C.GLint(nsides), C.GLint(rings))
}

func SolidDodecahedron() {
	C.glutSolidDodecahedron()
}

func WireDodecahedron() {
	C.glutWireDodecahedron()
}

func SolidOctahedron() {
	C.glutSolidOctahedron()
}

func WireOctahedron() {
	C.glutWireOctahedron()
}

func SolidTetrahedron() {
	C.glutSolidTetrahedron()
}

func WireTetrahedron() {
	C.glutWireTetrahedron()
}

func SolidIcosahedron() {
	C.glutSolidIcosahedron()
}

func WireIcosahedron() {
	C.glutWireIcosahedron()
}

// And, of course:
func SolidTeapot(size gl.GLdouble) {
	C.glutSolidTeapot(C.GLdouble(size))
}

func WireTeapot(size gl.GLdouble) {
	C.glutWireTeapot(C.GLdouble(size))
}

// - Video Resize

func VideoResizeGet(param gl.GLenum) int {
	return int(C.glutVideoResizeGet(C.GLenum(param)))
}

func SetupVideoResizing() {
	C.glutSetupVideoResizing()
}

func StopVideoResizing() {
	C.glutStopVideoResizing()
}

func VideoResize(x, y, width, height int) {
	C.glutVideoResize(C.int(x), C.int(y), C.int(width), C.int(height))
}

func VideoPan(x, y, width, height int) {
	C.glutVideoPan(C.int(x), C.int(y), C.int(width), C.int(height))
}

// - Debugging

func ReportErrors() {
	C.glutReportErrors()
}

// - Device Control

func IgnoreKeyRepeat(ignore int) {
	C.glutIgnoreKeyRepeat(C.int(ignore))
}

func SetKeyRepeat(repeatMode int) {
	C.glutSetKeyRepeat(C.int(repeatMode))
}

func ForceJoystickFunc() {
	C.glutForceJoystickFunc()
}

// - Game Mode

func GameModeString(str string) {
	cstr := C.CString(str)
	C.glutGameModeString(cstr)
	C.free(unsafe.Pointer(cstr))
}

func EnterGameMode() Window {
	w := Window(C.glutEnterGameMode())

	if gameWindow != nil {
		unregisterWindow(*gameWindow)
	}
	registerWindow(w)
	gameWindow = &w

	return w
}

func LeaveGameMode() {
	C.glutLeaveGameMode()

	unregisterWindow(*gameWindow)
	gameWindow = nil
}

func GameModeGet(mode gl.GLenum) int {
	return int(C.glutGameModeGet(C.GLenum(mode)))
}

// - Callbacks

//export internalButtonBoxFunc
func internalButtonBoxFunc(button, state int) {
	winFuncs[GetWindow()].buttonBox(button, state)
}

//export internalDialsFunc
func internalDialsFunc(dial, value int) {
	winFuncs[GetWindow()].dials(dial, value)
}

//export internalDisplayFunc
func internalDisplayFunc() {
	winFuncs[GetWindow()].display()
}

//export internalEntryFunc
func internalEntryFunc(state int) {
	winFuncs[GetWindow()].entry(state)
}

//export internalIdleFunc
func internalIdleFunc() {
	idleFunc()
}

//export internalJoystickFunc
func internalJoystickFunc(buttonMask uint, x, y, z int) {
	winFuncs[GetWindow()].joystick(buttonMask, x, y, z)
}

//export internalKeyboardFunc
func internalKeyboardFunc(key uint8, x, y int) {
	winFuncs[GetWindow()].keyboard(key, x, y)
}

//export internalKeyboardUpFunc
func internalKeyboardUpFunc(key uint8, x, y int) {
	winFuncs[GetWindow()].keyboardUp(key, x, y)
}

//export internalMenuFunc
func internalMenuFunc(state int) {
	menuFuncs[GetMenu()](state)
}

//export internalMenuStatusFunc
func internalMenuStatusFunc(status, x, y int) {
	winFuncs[GetWindow()].menuStatus(status, x, y)
}

//export internalMotionFunc
func internalMotionFunc(x, y int32) {
	winFuncs[GetWindow()].motion(x, y)
}

//export internalMouseFunc
func internalMouseFunc(button, state, x, y int) {
	winFuncs[GetWindow()].mouse(button, state, x, y)
}

//export internalOverlayDisplayFunc
func internalOverlayDisplayFunc() {
	winFuncs[GetWindow()].overlayDisplay()
}

//export internalPassiveMotionFunc
func internalPassiveMotionFunc(x, y int) {
	winFuncs[GetWindow()].passiveMotion(x, y)
}

//export internalReshapeFunc
func internalReshapeFunc(width, height int) {
	winFuncs[GetWindow()].reshape(width, height)
}

//export internalSpaceballButtonFunc
func internalSpaceballButtonFunc(button, state int) {
	winFuncs[GetWindow()].spaceballButton(button, state)
}

//export internalSpaceballMotionFunc
func internalSpaceballMotionFunc(x, y, z int) {
	winFuncs[GetWindow()].spaceballMotion(x, y, z)
}

//export internalSpaceballRotateFunc
func internalSpaceballRotateFunc(x, y, z int) {
	winFuncs[GetWindow()].spaceballRotate(x, y, z)
}

//export internalSpecialFunc
func internalSpecialFunc(key, x, y int) {
	winFuncs[GetWindow()].special(key, x, y)
}

//export internalSpecialUpFunc
func internalSpecialUpFunc(key, x, y int) {
	winFuncs[GetWindow()].specialUp(key, x, y)
}

//export internalTabletButtonFunc
func internalTabletButtonFunc(button, state, x, y int) {
	winFuncs[GetWindow()].tabletButton(button, state, x, y)
}

//export internalTabletMotionFunc
func internalTabletMotionFunc(x, y int) {
	winFuncs[GetWindow()].tabletMotion(x, y)
}

//export internalVisibilityFunc
func internalVisibilityFunc(state int) {
	winFuncs[GetWindow()].visibility(state)
}
