package main

import (
	"fmt"
	"path/filepath"

	"go-sciter-v6-demo/api"
	"go-sciter-v6-demo/utils/panicutil"

	"github.com/Mr-ShiHuaYu/go-sciter"
	"github.com/Mr-ShiHuaYu/go-sciter/window"
)

var win *window.Window

func main() {
	w, err := window.New(sciter.DefaultWindowCreateFlag,
		&sciter.Rect{Left: 400, Top: 150, Right: 1600, Bottom: 950})
	if err != nil {
		fmt.Println("Create Window Error: ", err)
	}
	fullpath, err := filepath.Abs("resources/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	win = w
	// 设置运行时特性，允许调试器
	w.SetOption(sciter.SCITER_SET_SCRIPT_RUNTIME_FEATURES, sciter.ALLOW_FILE_IO|sciter.ALLOW_SOCKET_IO|sciter.ALLOW_EVAL|sciter.ALLOW_SYSINFO)
	// 设置调试模式
	w.SetOption(sciter.SCITER_SET_DEBUG_MODE, 1)

	w.SetTitle("go-sciter-demo")
	w.LoadFile(fullpath)
	api.MainWin = w
	setEventHandler(w)
	w.Show()
	w.Run()
}

func setEventHandler(w *window.Window) {

	w.DefineFunction("loadFile", func(args ...*sciter.Value) *sciter.Value {
		panicutil.TryPanic(func(err interface{}) {
			fmt.Println("[PANIC][loadFile]", "catch panic:", err)
		})
		// 传入参数读取
		resFile := args[0].String()
		fullpath, err := filepath.Abs("resources/" + resFile)
		if err != nil {
			fmt.Println(err)
			return sciter.NewValue(0)
		}
		win.LoadFile(fullpath)    // 重新加载界面
		return sciter.NewValue(1) // 返回单个值
	})

	w.DefineFunction("execCmd", api.ExecCmd)
	w.DefineFunction("execCmdAsync", api.ExecCmdAsync)

	w.DefineFunction("execNativeFunc", api.ExecNativeFunc)
	w.DefineFunction("execJsCallback", api.ExecJsCallback)
	w.DefineFunction("execReturnNativeCallback", api.ExecReturnNativeCallback)

	w.DefineFunction("downloadFile", api.DownloadFile)

}
