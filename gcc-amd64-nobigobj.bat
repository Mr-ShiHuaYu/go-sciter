@echo off
setlocal
set RG=E:\soft\w64devkit\bin\gcc.exe
set OC=E:\soft\w64devkit\bin\objcopy.exe
if not exist "%RG%" set RG=gcc.exe
if not exist "%OC%" set OC=objcopy.exe
set IC=0
set OUTLIST=
set NEXTO=0
:parse
if "%~1"=="" goto parse_done
set A=%~1
if /i "%A%"=="-c" set IC=1
if "%NEXTO%"=="1" (
  set NEXTO=0
  call :addout "%A%"
  goto next
)
set P=%A:~0,2%
if "%P%"=="-o" (
  if NOT "%A:~2%"=="" (
    call :addout "%A:~2%"
  ) else (
    set NEXTO=1
  )
  goto next
)
:next
shift
goto parse
:parse_done
if "%IC%"=="0" (
  "%RG%" %*
  exit /b %ERRORLEVEL%
)
"%RG%" %*
set E=%ERRORLEVEL%
if %E% NEQ 0 exit /b %E%
set AF=0
if defined OUTLIST for %%F in (%OUTLIST%) do call :REWRITE "%%~F"
exit /b %AF%
goto :EOF
:addout
set F=%~1
if "%F%"=="" goto :EOF
set OUTLIST=%OUTLIST% "%F%"
goto :EOF
:REWRITE
set S=%~1
if "%S%"=="" goto :EOF
if not exist "%S%" goto :EOF
set X=%S:~-2%
if /i NOT "%X%"==".o" goto :EOF
:mktmp
set T=%S%.__O%RANDOM%%%.tmp
if exist "%T%" goto mktmp
"%OC%" -O pe-x86-64 -- "%S%" "%T%" >nul 2>nul
set E2=%ERRORLEVEL%
if %E2% NEQ 0 (
  if exist "%T%" del /F /Q "%T%"
  set AF=1
  goto :EOF
)
if not exist "%T%" ( set AF=1& goto :EOF )
move /Y "%T%" "%S%" >nul 2>nul
if errorlevel 1 (
  copy /B /Y "%T%" "%S%" >nul 2>nul
  if exist "%T%" del /F /Q "%T%"
)
goto :EOF