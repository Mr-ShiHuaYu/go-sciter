@echo off
REM ================================================================
REM   Generate libsciter32.a (MinGW i386 import library)
REM   Matches the working behaviour of go-sciter v0.6.1 shipped .a:
REM   Final PE Import Directory -> sciter.dll -> SciterAPI (BARE name, no @0 suffix)
REM
REM   Strategy (based on user-provided recipe):
REM     * Step 1:  .def provides TWO symbol names on the LINKER side:
REM                 SciterAPI@0 = SciterAPI
REM                 SciterAPI   = SciterAPI
REM               so GCC / ld can resolve both plain and @stdcall names.
REM     * Step 2:  dlltool -k (--kill-at) strips @N from the DLL-side
REM               LOOKUP NAME written into every .idata$6 stub.  Even the
REM               @0 alias stub ends up with the bare "SciterAPI" string,
REM               so the EXE Import Table is always correct.
REM ================================================================
setlocal
set PATH=E:\soft\w64devkit-x86-2.0.0\bin;%PATH%
set DLLTOOL=dlltool
cd /d %~dp0

echo [STEP 1/3] Auto-generating sciter32.def (embedded into this bat) ...
(
  echo LIBRARY sciter.dll
  echo EXPORTS
  echo   SciterAPI@0 = SciterAPI
  echo   SciterAPI = SciterAPI
) > sciter32.def

echo [STEP 2/3] Auto-generating libsciter32.a via dlltool ...
if exist libsciter32.a del /F /Q libsciter32.a
%DLLTOOL% -d sciter32.def -D sciter.dll -l libsciter32.a -k -m i386 --as-flags=--32
if errorlevel 1 (
  echo [ERROR] dlltool failed!  Cannot build import library libsciter32.a.
  del /F /Q sciter32.def
  pause
  exit /b 1
)
del /F /Q sciter32.def

echo [STEP 3/3] Verification ...
echo ---- symbols (must be 4: @0 + bare for both T and I) ----
nm -g libsciter32.a
echo ---- members (t.o + h.o + 2 stub .o for SciterAPI / SciterAPI@0) ----
ar t libsciter32.a
echo.
echo [OK] libsciter32.a generated successfully. 
echo     DLL-side lookup name is ALWAYS "SciterAPI" (bare, no @0)
echo     because dlltool -k (--kill-at) strips @N from idata stubs.
echo Now run go-sciter-build-xp.bat to build.
endlocal
pause
exit /b 0
