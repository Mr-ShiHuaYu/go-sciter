package window

import (
	"runtime"

	"github.com/Mr-ShiHuaYu/go-sciter"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func (w *Window) run() {
	// runtime.LockOSThread()
}

// https://github.com/golang/go/wiki/LockOSThread
// https://github.com/Mr-ShiHuaYu/go-sciter/issues/201
func init() {
	runtime.LockOSThread()
}
