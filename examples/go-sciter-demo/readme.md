此demo来源于:
https://gitee.com/mylofter/go-sciter-demo

只是原来的demo,版本是v5的,没有更新到v6的版本
此demo,已经更新到了v6的版本

## 此demo兼容XP系统

## 版本说明
此demo,已经更新到了v6的版本
根目录下的sciter.dll是v6.0.4.5的,截止2026-07-05最新版本

## 常规编译
常规编译:
先下载gcc 64位版本,设置环境变量PATH
go mod tidy
set PATH=D:\5CPP\5.C\mingw64\bin;%PATH%;
go build -o go-sciter-demo.exe

## 编译XP系统

需要先下载 go 1.11.13-386 32位版本
下载gcc 32位,推荐w64devkit-x86-2.0.0,支持最新的XP系统的GCC
修改 build-xp.bat 中的gcc 目录,以及gopath,双击运行 build-xp.bat 进行编译

