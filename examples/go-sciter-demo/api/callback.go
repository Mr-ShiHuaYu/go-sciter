package api

import (
	"fmt"

	"github.com/Mr-ShiHuaYu/go-sciter"
)

func ExecNativeFunc(args ...*sciter.Value) *sciter.Value {
	fmt.Println("js call ExecNativeFunc，args=", args)
	return sciter.NewValue("OK")
}

func ExecJsCallback(args ...*sciter.Value) *sciter.Value {
	fn := args[0] // fn.IsObjectFunction()=TRUE
	// go func() { // 如果是用fn是callback，无法在异步线程中执行
	// 异步回调参考,go暂未实现：https://sciter.com/forums/topic/check-if-native-callback-is-valid/
	fmt.Println(fn.Invoke(sciter.NullValue(), "[Native Script]", sciter.NewValue("ABCBDBA")))
	return sciter.NullValue()
}

func ExecReturnNativeCallback(args ...*sciter.Value) *sciter.Value {
	result := sciter.NewValue()
	fn := func(args ...*sciter.Value) *sciter.Value {
		fmt.Println("args:", len(args), args)
		return sciter.NewValue("native functor called")
	}
	result.Set("f", fn) // 返回回调函数供js调用
	return result
}
