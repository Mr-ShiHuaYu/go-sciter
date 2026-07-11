@echo off
chcp 65001

E:\go\gopath\bin\g.exe use 1.11.13-386
go version
set PATH=E:\soft\w64devkit-x86-2.0.0\bin;%PATH%;
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1

echo go build ...
@rm examples\callback\callback.exe
cp dll\windows.xp\x32\sciter.dll examples\callback
go build -x -o ./examples/callback/callback.exe ./examples/callback
set EC=%ERRORLEVEL%

echo.
echo BUILD_DONE exitcode=%EC%

if %EC% NEQ 0 (
  echo [ERROR] go build FAILED, exitcode=%EC% - will NOT start callback64.exe
  pause
  exit /b %EC%
)

cd /d examples\callback
echo [OK] Build succeeded, launching callback64.exe ...
start "" callback.exe
pause