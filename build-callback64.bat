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

REM --------------------------------------------------------------
REM  Fix for w64devkit GCC 14+ producing COFF BIGOBJ -> Go 1.20 cgo cannot parse pe-bigobj-x86-64.
REM  We route CC/CXX through a wrapper: gcc -c step writes temp bigobj,
REM  then wrapper calls objcopy to rewrite every ./*.o as classic pe-x86-64.
REM  Non-compile steps (link, -E, -S etc.) pass through 100% verbatim.
REM --------------------------------------------------------------
REM set "CC=%~dp0gcc-amd64-nobigobj.bat"
REM set "CXX=%~dp0gcc-amd64-nobigobj.bat"

echo go build ...
@rm examples\callback\callback64.exe
cp dll\windows\x64\sciter.dll examples\callback
go build -x -a -o ./examples/callback/callback64.exe ./examples/callback
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
start "" callback64.exe
pause