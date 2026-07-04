package sciter

/*
#cgo CFLAGS: -Iinclude
#include "sciter-x.h"
*/
import "C"
import (
	"unsafe"
)

// HWINDOW  SciterCreateWindow ( UINT creationFlags,LPRECT frame, LPVOID, LPVOID, HWINDOW parent);
// Note: In Sciter API v10, the 3rd and 4th parameters (delegate/delegateParam) are reserved and MUST be NULL.

// rect is the display area
func CreateWindow(createFlags WindowCreationFlag, rect *Rect, delegate uintptr, delegateParam uintptr, parent C.HWINDOW) C.HWINDOW {
	// set default size
	if rect == nil {
		rect = DefaultRect
	}
	// create window
	// NOTE: In Sciter API v10, delegate and delegateParam MUST be NULL (0)
	hwnd := C.SciterCreateWindow(
		C.UINT(createFlags),
		(*C.RECT)(unsafe.Pointer(rect)),
		nil,
		nil,
		parent)
	// in case of NULL
	if int(uintptr(unsafe.Pointer(hwnd))) == 0 {
		return BAD_HWINDOW
	}
	return hwnd
}
