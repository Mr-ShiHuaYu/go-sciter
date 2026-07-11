@echo off
REM ================================================================
REM   Generate libsciter64.a (MinGW x64 / 64-bit PE import library)
REM   For building 64-bit Go binaries (GOARCH=amd64) against 64-bit
REM   sciter.dll on Vista/7/8/10/11.
REM
REM   Notes:
REM     * 64-bit Microsoft x64 calling convention does NOT use @N
REM       stdcall decoration — all symbols are bare names.  We still
REM       keep the @0 alias line in the .def just in case any 32-bit
REM       header gets #included by mistake; it is harmless on x64.
REM     * Final PE Import Directory -> sciter.dll -> SciterAPI (bare)
REM ================================================================
setlocal
set PATH=E:\soft\w64devkit\bin;%PATH%
set DLLTOOL=dlltool
cd /d %~dp0

echo [STEP 1/3] Auto-generating sciter64.def (embedded into this bat) ...
(
  echo LIBRARY sciter.dll
  echo EXPORTS
  echo   SciterAPI@0 = SciterAPI
  echo   SciterAPI = SciterAPI
) > sciter64.def

echo [STEP 2/3] Auto-generating libsciter64.a (x64) via dlltool ...
if exist libsciter64.a del /F /Q libsciter64.a
REM --------------------------------------------------------------
REM  x64 PE / 64-bit MinGW dlltool flags — differences from 32-bit:
REM    -m i386:x86-64    -> produce PE32+ import stubs (x86-64 ABI)
REM                       (some builds accept -m x86-64 as well)
REM    --as-flags=--64   -> pass --64 to GNU assembler
REM    -k                -> still keep -k (--kill-at); although x64
REM                       never emits @N decorations, -k is a no-op
REM                       that guards against any future aliases.
REM --------------------------------------------------------------
%DLLTOOL% -d sciter64.def -D sciter.dll -l libsciter64.a -k -m i386:x86-64 --as-flags=--64
if errorlevel 1 (
  echo [ERROR] dlltool failed!  Cannot build x64 import library libsciter64.a.
  del /F /Q sciter64.def
  pause
  exit /b 1
)
del /F /Q sciter64.def

echo [STEP 3/3] Verification ...
echo ---- file format check (should say "pe-x86-64" / 64-bit PE) ----
for /f "tokens=*" %%F in ('ar t libsciter64.a ^| findstr /r /c:"_s[0-9]*\.o$"') do (
  echo   %%F  format:
  objdump -f libsciter64.a --archive-headers 2>nul | findstr /i "pe-x86-64 file format" || echo     (objdump not available or format mismatch)
  goto :format_check_done
)
:format_check_done
echo ---- symbols (x64: NO leading underscore `_` prefix, bare names) ----
nm -g libsciter64.a
echo ---- archive members (t.o + h.o + stub .o(s)) ----
ar t libsciter64.a
echo.
echo [OK] 64-bit libsciter64.a generated successfully.
echo     DLL-side lookup name: SciterAPI (bare, no @N)
echo     Target architecture : PE32+ / x86-64
echo.
echo Build command example:
echo   set GOARCH=amd64
echo   set CGO_ENABLED=1
echo   go build -o callback.exe .\examples\callback
endlocal
pause
exit /b 0
