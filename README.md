# Go bindings for Sciter
## 更新到Sciter v6.0.4.5
## 兼容XP
## sciter sdk下载:
https://gitlab.com/sciter-engine/sciter-js-sdk/-/releases

## 本仓库的代码兼容XP系统

![xp下运行截图](https://github.com/Mr-ShiHuaYu/go-sciter/blob/master/imgs/1.jpg?raw=true)

## 版本说明

已经更新到了v6的版本
根目录下的sciter.dll是v6.0.4.5的,截止2026-07-05最新版本

## 编译提前配置
### 64位
先下载gcc 64位版本,设置环境变量PATH
set PATH=D:\5CPP\5.C\mingw64\bin;%PATH%;

```shell
REM 64位
@echo off
chcp 65001
set PATH=E:\go\resources\go-sciter\xp\w64devkit-x86-2.0.0\bin;%PATH%;
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
```

### 32位XP
- 需要先下载 go 1.11.13-386 32位版本,这是测试过支持XP的最后版本,可以使用go module
- 下载gcc 32位,推荐w64devkit-x86-2.0.0,支持最新的XP系统的GCC

```shell
REM 32位XP
@echo off
chcp 65001
set PATH=E:\go\resources\go-sciter\xp\w64devkit-x86-2.0.0\bin;%PATH%;
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1
```

## 开始编译示例代码
> go-sciter-demo 里面是最全的,其他的示例还没修改成兼容,只有以下2个示例是可以运行的,go-sciter-demo其中包含了其他全部

- go build -x -o ./examples/callback/callback.exe ./examples/callback
- go build -x -o ./examples/go-sciter-demo/go-sciter-demo.exe ./examples/go-sciter-demo

> -x 表示显示中间的过程
> 
> -a 可以强制重新编译
> 
> -o 指定输出文件名
> 
> go-sciter-demo 的代码来自于:
https://github.com/Mr-ShiHuaYu/go-sciter
> 特别感谢

## 运行特别注意
- 运行时,需要复制项目dll目录下对应的操作系统的sciter库到程序运行目录
- 32位的非XP系统,复制`dll\windows\x32\sciter.dll`到程序运行目录
- 64位的系统,复制`dll\windows\x64\sciter.dll`到程序运行目录
- xp系统,复制`dll\windows.xp\x32\sciter.dll`到程序运行目录
- 其他系统类似
> 此dll目录是官方sdk中bin目录复制,并经过upx压缩后的,windows和linux的库都经过压缩了

## dll目录下的 inspector.exe 是运行调试工具,类似于浏览器的开发者工具


## 程序运行时去掉黑框(cmd窗口)(建议在发布时使用)
没有黑色窗口,log.println或fmt.println就都看不到了

在编译时加上参数 -ldflags="-H windowsgui"，如:

```bash
go build -ldflags="-H windowsgui"
```

## 自定义程序图标及版本信息
在没有使用winres时,go build的exe程序是没有图标及版本信息的.现在介绍使用方法(此方法只在windows下有效)
详细见: https://gitee.com/ying32/govcl/wikis/pages?sort_id=410058&doc_id=102420
### 补充说明:
- app.rc包含了app.manifest(清单文件,可控制是否要管理员权限),applogo.ico(图标),版本信息(自定义)
- 编译命令(将rc转.syso)
* x86
  windres.exe -i app.rc -o defaultRes_windows_386.syso -F pe-i386
* x64
  windres.exe -i app.rc -o defaultRes_windows_amd64.syso -F pe-x86-64
- 编译后,只需要将.syso(名称无所谓,只能有一个,会默认找.syso)存放到项目根目录,就可以了,go build编译时,需要指定整个目录 go build .

## 直接打包html等资源文件到程序exe内部,防止代码被修改

```bash
sciter-js-sdk-main-6.0.4.5\sciter-js-sdk-main\bin\windows\packfolder.exe resources res.go -v resource -go
```

> packfolder 使用方式:
>
> packfolder.exe -h
> usage: packfolder.exe 要打包的文件夹 输出的文件 [options]
>
> [-i "foo/*;..."]          include files or folders, if defined only matching items are included
> [-x "*.ex1;*.ex2;..."]    exclude files or folders, if defined matching items are excluded from archive
>
> [-v "varname"]            name of blob variable(导出的变量名)
> [-csharp]                 generates C# class with a byte[] field literal
> [-dlang]                  generates D ubyte[] literal
> [-go]                     generates Go []byte literal
> [-binary]                 generates binary file

> 上面的 -v resource:-v 后面跟要导出的变量名
>
> -go 表示导出的是go语言

会在当前目录下生成 res.go 文件,这个文件内部,就是所有的资源文件了,并且,res.go中有一个 resource 变量,导出的,便于我们在main.go中使用这个变量

- 然后再修改main.go中引用html的那部分代码

```go
    // fullpath, err := filepath.Abs("resources/index.html")
	// if err != nil {
		// fmt.Println(err)
		// return
	// }
	// w.LoadFile(fullpath)
    w.SetResourceArchive(resource)
    w.LoadFile("this://app/index.html")
```

> 使用 SetResourceArchive 传入 resource 导出的变量名
>
> 使用 LoadFile 时,需要添加 this://app/ 前缀,才能访问资源中的文件

这样编译后的exe,就可以不用带html了,只需要带一个sciter.dll就可以了

## 减少编译体积

* -s: 去掉符号信息。
* -w: 去掉DWARF调试信息。

```bash
go build -ldflags="-s -w"
```

**此操作大约可以减少50%左右的，但用一`-s`参数后会造成原来的错误之类的信息无法具体化**  

如果要搭配上去除黑窗口,则使用:

```bash
go build -ldflags="-s -w -H windowsgui"
```

> 建议发布release版本时使用

还可以使用upx对生成的exe和sciter.dll进行压缩,可进一步大幅度减少体积

```bash
upx -9 xx.exe
upx -9 sciter.dll
```

## 跨平台编译

可在任意一平台进行开发，当要发布对应平台时需要到目标平台进行编译。

因为linux与macOS下用到了cgo，所以需要到目标平台进行编译（windows平台除外，可以在linux或者macOS下编译出windows应用）。

> 貌似可以使用zig的编译器,在windows 下,也使用cgo,来编译出linux的程序,待测试,但不知道zig的编译器支不支持XP了

## XP系统下运行说明
### 解决XP中文乱码问题

在所有公共的css或者当前页面头上添加:

```css
html {
}

* {
    font-family: "SimSun", "宋体",
    "SimHei", "黑体",
    "Microsoft YaHei", "微软雅黑",
    "PingFang SC",
    Tahoma, Arial, sans-serif !important;
}
```

不知道为什么,必须在**html空标签后面设置**才会生效:

### sqlite(XP下不支持)
运行sqlite时,会出现 无法定位程序输入点 ReleaseSWLockExcluseive 于动态链接库 KERNEL32.DLL 上
即使使用官方的scapp.exe也是同样的错误,并不是go的原因
因为,在官方的 bin\windows.xp\x32 目录下,就没有 sciter-sqlite.dll,所以,XP不支持,要想支持只能买源码,手动编译

- 但,GO可以使用GO的sqlite库,不使用sciter的

## 总结,生产打包

```bash
set PATH=E:\go\resources\go-sciter\xp\w64devkit-x86-2.0.0\bin;%PATH%;
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1
REM 打包资源文件到res.go
packfolder.exe resources res.go -v resource -go
REM 修改main.go文件
go build -ldflags="-s -w -H windowsgui" -o main.exe
REM 减少体积,使用upx压缩,若已经压缩过,则不能压缩
upx -9 main.exe
upx -9 sciter.dll
```



## 以下是原readme

[![AppVeyor status](https://ci.appveyor.com/api/projects/status/rphv883klffw9em9/branch/master?svg=true)](https://ci.appveyor.com/project/pravic/go-sciter)
[![Travis Status](https://travis-ci.com/sciter-sdk/go-sciter.svg?branch=master)](https://travis-ci.com/sciter-sdk/go-sciter)
[![License](https://img.shields.io/github/license/sciter-sdk/go-sciter.svg)](https://github.com/Mr-ShiHuaYu/go-sciter)
[![Join the forums at https://sciter.com/forums](https://img.shields.io/badge/forum-sciter.com-orange.svg)](https://sciter.com/forums)

Check [this page](http://sciter.com/developers/sciter-sdk-bindings/) for other language bindings (Delphi / D / Go / .NET / Python / Rust).

----


# Attention

The ownership of project is transferred to this new organization.
Thus the `import path` for golang should now be `github.com/Mr-ShiHuaYu/go-sciter`, but the package name is still `sciter`.

# Introduction

This package provides a Golang bindings of [Sciter][] using cgo.
Using go sciter you must have the platform specified `sciter dynamic library`
downloaded from [sciter-sdk][], the library itself is rather small
 (under 5MB, less than 2MB when upxed) .

Most [Sciter][] API are supported, including:

 * Html string/file loading
 * DOM manipulation/callback/event handling
 * DOM state/attribute handling
 * Custom resource loading
 * Sciter Behavior
 * Sciter Options
 * Sciter Value support
 * NativeFunctor (used in sciter scripting)

And the API are organized in more or less a gopher friendly way.

Things that are not supported:

 * Sciter Node API
 * TIScript Engine API

# Getting Started

###  At the moment only **Go 1.10** or higher is supported (issue #136).

 1. Download the [sciter-sdk][]
 2. Extract the sciter runtime library from [sciter-sdk][] to system PATH

    The runtime libraries lives in `bin` `bin.lnx` `bin.osx` with suffix like `dll` `so` or `dylib`

    * Windows: simply copying `bin\64\sciter.dll` to `c:\windows\system32` is just enough
    * Linux:
      - `cd sciter-sdk/bin.lnx/x64`
      - `export LIBRARY_PATH=$PWD`
      - `echo $PWD >> libsciter.conf`
      - `sudo cp libsciter.conf /etc/ld.so.conf.d/`
      - `sudo ldconfig`
      - `ldconfig -p | grep sciter` should print libsciter-gtk.so location
    * OSX:
      - `cd sciter-sdk/bin.osx/`
      - `export DYLD_LIBRARY_PATH=$PWD`

 3. Set up GCC envrionmnet for CGO

    [mingw64-gcc][] (5.2.0 and 7.2.0 are tested) is recommended for Windows users.

    Under Linux gcc(4.8 or above) and gtk+-3.0 are needed.

 4. `go get -x github.com/Mr-ShiHuaYu/go-sciter`

 5. Run the example and enjoy :)

# Sciter Desktop UI Examples

![](http://sciter.com/screenshots/slide-wt5.png)

![](http://sciter.com/screenshots/slide-norton360.png)

![](http://sciter.com/screenshots/slide-norton-nis.png)

![](http://sciter.com/screenshots/slide-cardio.png)

![](http://sciter.com/screenshots/slide-surveillance.png)

![](http://sciter.com/screenshots/slide-technology.png)

![](http://sciter.com/screenshots/slide-sciter-ide.png)

![](http://sciter.com/screenshots/slide-sciter-osx.png)

![](http://sciter.com/screenshots/slide-sciter-gtk.png)


# Sciter Version Support
Currently supports [Sciter][] version `4.0.0.0` and higher.

[Sciter]: http://sciter.com/
[sciter-sdk]: http://sciter.com/download/

# About Sciter

[Sciter][] is an `Embeddable HTML/CSS/script engine for modern UI development, Web designers, and developers, can reuse their experience and expertise in creating modern looking desktop applications.`

In my opinion, [Sciter][] , though not open sourced, is an great
desktop UI development envrionment using the full stack of web technologies,
which is rather small (under 5MB) especially compared to [CEF][],[Node Webkit][nw] and [Atom Electron][electron]. :)

Finally, according to [Andrew Fedoniouk][author] the author and the Sciter
`END USER LICENSE AGREEMENT` , the binary form of the [Sciter][]
dynamic libraries are totally free to use for commercial or
non-commercial applications.

# The Tailored Sciter C Headers
This binding ueses a tailored version of the sciter C Headers, which lives in directory: `include`. The included c headers are a modified version of the
[sciter-sdk][] standard headers.

It seems [Sciter][] is developed using C++, and the included headers in the
[Sciter SDK][sciter-sdk] are a mixture of C and C++, which is not
quite suitable for an easy golang binding.

I'm not much fond of C++ since I started to use Golang, so I made this
modification and hope [Andrew Fedoniouk][author] the author would provide
pure C header files for Sciter. :)

[CEF]:https://bitbucket.org/chromiumembedded/cef
[nw]: https://github.com/nwjs/nw.js
[electron]:https://github.com/atom/electron

[author]: http://sciter.com/about/
[mingw64-gcc]: http://sourceforge.net/projects/mingw-w64/
