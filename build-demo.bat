@echo off
chcp 65001

E:\go\gopath\bin\g.exe use 1.11.13-386
go version
set PATH=E:\soft\w64devkit-x86-2.0.0\bin;%PATH%;
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1

echo go build ...
@rm examples\go-sciter-demo\go-sciter-demo.exe
cp dll\windows.xp\x32\sciter.dll examples\go-sciter-demo
go build -x -o ./examples/go-sciter-demo/go-sciter-demo.exe ./examples/go-sciter-demo
set EC=%ERRORLEVEL%

echo.
echo BUILD_DONE exitcode=%EC%
cd /d examples\go-sciter-demo
start "" go-sciter-demo.exe
pause