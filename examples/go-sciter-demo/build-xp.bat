@echo off
chcp 65001
set PATH=E:\go\resources\go-sciter\xp\w64devkit-x86-2.0.0\bin;%PATH%;
set GOPATH=E:\go\resources\go-sciter\xp\gopath
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1

echo go build
go build -x -o main.exe
set EC=%ERRORLEVEL%

echo.
echo BUILD_DONE exitcode=%EC%
pause