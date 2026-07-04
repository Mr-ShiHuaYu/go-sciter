package sciter

/*
// ---------- Windows Build Flags — XP-SP3 COMPAT BY DEFAULT ----------
//
// -Iinclude                 : Sciter SDK headers
// -DWIN32_LEAN_AND_MEAN     : Skip rarely-used legacy Win32 headers
//
// XP version macros (_WIN32_WINNT=0x0501 etc.) live in include/sciter-x-types.h
// with #ifndef guards — they apply globally to every .c file, not just this TU.
//
// --- PE Import Table Fix (CRITICAL for XP stability) ------------------------
// We link libsciter.a (dlltool import lib) so sciter.dll appears in the final
// PE Import Table.  Windows loader then loads sciter.dll BEFORE the Go runtime
// spawns any threads (still single-threaded at EXE entry).
// Eliminates half-initialised TLS / singleton state that caused NULL+0x02 AVs
// inside SciterExec(APP_INIT) (XP and newer systems both benefited).
#cgo windows,386 CFLAGS: -Iinclude -DWIN32_LEAN_AND_MEAN
#cgo windows,amd64 CFLAGS: -Iinclude -DWIN32_LEAN_AND_MEAN
#cgo windows,386 LDFLAGS: -L${SRCDIR} -lsciter
#cgo windows,amd64 LDFLAGS: -L${SRCDIR} -lsciter
// Generic fallback CFLAGS for other cgo invocations.
#cgo CFLAGS: -Iinclude
#include "sciter-x.h"

// Forward-declaration of the helper implemented in sciter-x-api.c (Windows
// branch only).  It performs EXACTLY the same 13 steps that c_xp_test.exe
// proved works 100% on Windows XP (Sciter 6.0.4.5 DLL).  Using C rather than
// Go for this bootstrap avoids VTable-slot offset mistakes and __stdcall
// parameter ordering bugs that plagued the Go syscall-based attempts.
HWINDOW XPBootstrapAndCreateWindow(UINT creationFlags,
                                   const RECT* frame,
                                   HWINDOW parent);
*/
import "C"
import (
	"unsafe"

	"github.com/lxn/win"
)

// LRESULT  SciterProc (HWINDOW hwnd, UINT msg, WPARAM wParam, LPARAM lParam)
// LRESULT  SciterProcND (HWINDOW hwnd, UINT msg, WPARAM wParam, LPARAM lParam, BOOL* pbHandled)

func ProcND(hwnd win.HWND, msg uint, wParam uintptr, lParam uintptr) (ret int, handled bool) {
	var bHandled C.BOOL
	ret = int(C.SciterProcND(C.HWINDOW(unsafe.Pointer(hwnd)), C.UINT(msg), C.WPARAM(wParam), C.LPARAM(lParam), &bHandled))
	if bHandled == 0 {
		handled = false
	} else {
		handled = true
	}
	return
}

// HWINDOW  SciterCreateWindow ( UINT creationFlags,LPRECT frame, LPVOID, LPVOID, HWINDOW parent);
// Note: In Sciter API v10, the 3rd and 4th parameters (delegate/delegateParam) are reserved and MUST be NULL.

// Create sciter window.
//
//	On Windows returns HWND of either top-level or child window created.
//	On OS X returns NSView* of either top-level window or child view .
//
//	\param[in] creationFlags \b SCITER_CREATE_WINDOW_FLAGS, creation flags.
//	\param[in] frame \b LPRECT, window frame position and size.
//	\param[in] delegate \b uintptr, RESERVED in API v10 - must be 0.
//	\param[in] delegateParam \b uintptr, RESERVED in API v10 - must be 0.
//	\param[in] parent \b HWINDOW, optional parent window.
//
// rect is the display area
func CreateWindow(createFlags WindowCreationFlag, rect *Rect, delegate uintptr, delegateParam uintptr, parent C.HWINDOW) C.HWINDOW {
	// =========================================================================
	// WINDOWS XP BOOTSTRAP — DELEGATE 100% TO PURE C HELPER!
	// =========================================================================
	// Previous Go attempts failed because of (1) a VTable slot off-by-one
	// (SciterSetOption was slot 21 not 20) AND then (2) subtle Go
	// syscall.Syscall parameter-order vs __stdcall disagreements that could
	// not be diagnosed without a debugger.  Here we simply call a C helper
	// whose source is a VERBATIM copy of c_xp_test.c STEPS 4..13, so it
	// MUST produce byte-identical Sciter vtable calls compared to the
	// working pure-C test.
	// =========================================================================

	// Copy Rect to C.RECT on stack (if non-nil) because C helper needs a
	// stable pointer — the caller's Go Rect struct may be GC-movable or nil.
	// (C helper itself also handles NULL by using a default 800x600 stack RECT.)
	var pFrame unsafe.Pointer
	if rect != nil {
		r := C.RECT{
			left:   C.LONG(rect.Left),
			top:    C.LONG(rect.Top),
			right:  C.LONG(rect.Right),
			bottom: C.LONG(rect.Bottom),
		}
		pFrame = unsafe.Pointer(&r)
	}

	hwnd := C.XPBootstrapAndCreateWindow(
		C.UINT(createFlags),
		(*C.RECT)(pFrame),
		parent)

	// C helper returns BAD_HWINDOW (0) on failure.
	if int(uintptr(unsafe.Pointer(hwnd))) == 0 {
		return BAD_HWINDOW
	}
	return hwnd
}
