package window

/*
#include <windows.h>
*/
import "C"
import (
	"fmt"
	"sync"
	"syscall"
	"unsafe"

	"github.com/Mr-ShiHuaYu/go-sciter"
	"github.com/lxn/win"
)

var (
	windowProcMap   = make(map[win.HWND]uintptr)
	windowProcMutex sync.Mutex
)

func New(creationFlags sciter.WindowCreationFlag, rect *sciter.Rect) (*Window, error) {
	w := new(Window)
	w.creationFlags = creationFlags

	// Initialize OLE for DnD and printing support
	win.OleInitialize()

	// create window
	// Note: In Sciter API v10, delegate/delegateParam MUST be 0 (NULL)
	hwnd := sciter.CreateWindow(
		creationFlags,
		rect,
		0,
		0,
		sciter.BAD_HWINDOW)

	if hwnd == sciter.BAD_HWINDOW {
		return nil, fmt.Errorf("Sciter CreateWindow failed [%d]", win.GetLastError())
	}

	w.Sciter = sciter.Wrap(hwnd)

	// Subclass the window to handle WM_DESTROY (since v10 delegate is not available)
	winHwnd := win.HWND(unsafe.Pointer(hwnd))
	subclassProc := syscall.NewCallback(windowSubclassProc)

	/* ================================================================
	 * CRITICAL (DEADLOCK PREVENTION #1 — New() subclassing order):
	 *   1. windowProcMutex is held ONLY around map read/write + GetWindowLongPtr
	 *   2. UNLOCK happens BEFORE SetWindowLongPtr(GWLP_WNDPROC)
	 * SetWindowLongPtr may send messages (e.g. WM_STYLECHANGING via USER32
	 * internals) that would re-enter the new subclass procedure on the same
	 * thread.  If we still held windowProcMutex at that moment the re-entrant
	 * windowSubclassProc would try to acquire the same NON-RECURSIVE sync.Mutex
	 * and the application would freeze forever in an unrecoverable deadlock.
	 * NEVER move windowProcMutex.Unlock() BELOW SetWindowLongPtr!
	 * ================================================================ */
	windowProcMutex.Lock()
	origProc := win.GetWindowLongPtr(winHwnd, win.GWLP_WNDPROC)
	windowProcMap[winHwnd] = origProc
	windowProcMutex.Unlock()
	win.SetWindowLongPtr(winHwnd, win.GWLP_WNDPROC, subclassProc)

	return w, nil
}

// windowSubclassProc handles window messages to intercept WM_DESTROY for quitting
func windowSubclassProc(hWnd win.HWND, message uint32, wParam uintptr, lParam uintptr) uintptr {
	/* ================================================================
	 * CRITICAL (DEADLOCK PREVENTION #2 — per-message lock discipline):
	 *   windowProcMutex is held ONLY for two TINY atomic map accesses:
	 *     a) map access (a):  read  origProc = windowProcMap[hWnd]   (below)
	 *     b) map access (b):  delete windowProcMap[hWnd]            (WM_DESTROY only)
	 *   AT ALL OTHER TIMES the mutex is 100 % UNLOCKED.
	 *
	 * In particular: the original window procedure (syscall.Syscall6 →
	 * origProc, or DefWindowProc) is ALWAYS invoked COMPLETELY OUTSIDE any
	 * critical section.  Many USER32 messages (WM_PAINT, WM_SIZE, WM_CTLCOLOR*,
	 * WM_SETTINGCHANGE etc.) call SendMessage internally and therefore re-enter
	 * windowSubclassProc recursively on the same thread.  That re-entrance is
	 * 100 % safe ONLY because we never hold the NON-RECURSIVE sync.Mutex
	 * across an origProc / DefWindowProc call.  If you ever add any lock /
	 * unlock around those calls you will get an unrecoverable thread deadlock
	 * (= window hangs forever, no crash, no log).
	 * ================================================================ */

	// ---------- map access (a): read origProc (lock + immediate unlock) ----------
	var origProc uintptr
	var hasOrig bool
	windowProcMutex.Lock()
	origProc, hasOrig = windowProcMap[hWnd]
	windowProcMutex.Unlock() // <-- UNLOCK BEFORE every outgoing call!

	var ret uintptr
	if hasOrig {
		/* NOTE: No lock held while calling the original wndproc.
		 * This comment line exists as a visual canary — if you ever add a
		 * windowProcMutex.Lock() in this function without matching Unlock()
		 * BEFORE reaching syscall.Syscall6, please delete your changes now. */
		ret, _, _ = syscall.Syscall6(
			uintptr(origProc),
			4,
			uintptr(hWnd),
			uintptr(message),
			wParam,
			lParam,
			0, 0)
	} else {
		ret = win.DefWindowProc(hWnd, message, wParam, lParam)
	}

	/* --- Post-processing (AFTER original wndproc has handled the msg): -----
	 *   WM_DESTROY: The window is now officially torn down from USER32's
	 *   perspective.  At this point two things must happen:
	 *   (1) PostQuitMessage(0) — so that the main GetMessage loop exits with
	 *       WM_QUIT and the process terminates cleanly.
	 *   (2) Remove this HWND from windowProcMap IMMEDIATELY.  The global
	 *       map is keyed by HWND; if we leave stale entries then each
	 *       create/destroy cycle leaks one (HWND → origProc) entry forever.
	 *       Also, if the kernel later reuses that HWND value for another
	 *       window (in long-running processes) we'd attach a stale/incorrect
	 *       subclass-proc pointer to an unrelated window and crash.
	 *----------------------------------------------------------------------- */
	if message == win.WM_DESTROY {
		win.PostQuitMessage(0)

		// ---------- map access (b): remove entry (lock + immediate unlock) ----------
		windowProcMutex.Lock()
		delete(windowProcMap, hWnd)
		windowProcMutex.Unlock()
	}

	return ret
}

func (s *Window) Show() {
	// message handling
	hwnd := win.HWND(unsafe.Pointer(s.GetHwnd()))
	win.ShowWindow(hwnd, win.SW_SHOW)
	win.UpdateWindow(hwnd)
}

func (s *Window) SetTitle(title string) {
	// message handling
	hwnd := C.HWND(unsafe.Pointer(s.GetHwnd()))
	C.SetWindowTextW(hwnd, (*C.WCHAR)(unsafe.Pointer(sciter.StringToWcharPtr(title))))
}

func (s *Window) AddQuitMenu() {
	// Define behaviour for windows
}

func (s *Window) Run() {
	// for system drag-n-drop
	// win.OleInitialize()
	// defer win.OleUninitialize()
	s.run()
	// start main gui message loop
	msg := (*win.MSG)(unsafe.Pointer(win.GlobalAlloc(0, unsafe.Sizeof(win.MSG{}))))
	defer win.GlobalFree(win.HGLOBAL(unsafe.Pointer(msg)))
	for win.GetMessage(msg, 0, 0, 0) > 0 {
		win.TranslateMessage(msg)
		win.DispatchMessage(msg)
	}
}
