package api

import (
	"fmt"
	"os/exec"
	"strings"

	"go-sciter-v6-demo/utils/panicutil"

	"github.com/Mr-ShiHuaYu/go-sciter"
)

func ExecCmd(args ...*sciter.Value) *sciter.Value {
	panicutil.TryPanic(func(err interface{}) {
		fmt.Println("[PANIC][execCmd]", "catch panic:", err)
	})
	arg0 := args[0].String()
	arg1 := args[1].String()
	list := strings.Split(arg1, ",")
	c := exec.Command(arg0, list...)
	output, err := c.CombinedOutput()
	fmt.Println(arg0, list)
	fmt.Println("result=", string(output))
	result := sciter.NewValue()
	if err == nil {
		result.Set("status", 200)
	} else {
		result.Set("status", 400)
		result.Set("error", err.Error())
	}
	result.Set("result", string(output))
	return result
}

func ExecCmdAsync(args ...*sciter.Value) *sciter.Value {
	panicutil.TryPanic(func(err interface{}) {
		fmt.Println("[PANIC][ExecCmdAsync]", "catch panic:", err)
	})
	arg0 := args[0].String()
	arg1 := args[1].String()
	go func() {
		list := strings.Split(arg1, ",")
		c := exec.Command(arg0, list...)
		output, err := c.CombinedOutput()
		fmt.Println(arg0, list)
		fmt.Println("result=", string(output))
		result := sciter.NewValue()
		if err == nil {
			result.Set("status", 200)
		} else {
			result.Set("status", 400)
			result.Set("error", err.Error())
		}
		result.Set("result", string(output))

		if res, err := MainWin.Call("postCustomEvent", sciter.NewValue("execCmdResult"), result); err != nil {
			fmt.Println("postCustomEvent resturn=", res, err)
		}
	}()
	return sciter.NullValue()
}
