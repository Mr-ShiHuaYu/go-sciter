@echo off
chcp 65001

E:\go\gopath\bin\g.exe use 1.11.13-386
go version
set PATH=E:\go\resources\go-sciter\xp\w64devkit-x86-2.0.0\bin;%PATH%;
set GOPATH=E:\go\resources\go-sciter\xp\gopath
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1

echo go build ...
rm examples\callback\callback.exe
cp dll\windows.xp\x32\sciter.dll examples\callback
go build -x -o ./examples/callback/callback.exe ./examples/callback
set EC=%ERRORLEVEL%

echo.
echo BUILD_DONE exitcode=%EC%
cd /d examples\callback
start "" callback.exe
pause