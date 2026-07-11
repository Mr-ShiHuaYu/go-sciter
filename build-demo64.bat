@echo off
chcp 65001

E:\go\gopath\bin\g.exe use 1.20.14
go version
@REM 64位需要使用mingw64,不能使用 w64devkit, 因为w64devkit的gcc 14+ 会生产 COFF BIGOBJ -> Go 1.20 cgo 不能 parse pe-bigobj-x86-64.
@REM set PATH=E:\soft\w64devkit\bin;%PATH%;
set PATH=D:\5CPP\5.C\mingw64\bin;%PATH%;
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1

echo go build ...
@rm examples\go-sciter-demo\go-sciter-demo64.exe
cp dll\windows\x64\sciter.dll examples\go-sciter-demo
go build -x -o ./examples/go-sciter-demo/go-sciter-demo64.exe ./examples/go-sciter-demo
set EC=%ERRORLEVEL%

echo.
echo BUILD_DONE exitcode=%EC%
cd /d examples\go-sciter-demo
start "" go-sciter-demo64.exe
pause