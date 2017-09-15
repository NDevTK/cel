@echo off
setlocal

set TARGET=C:\staging\bin
set CACHE=C:\staging\cache

if "%1"=="" (
  echo Usage: download-build [isolate-hash]
  exit /b 1
)

if not exist C:\staging mkdir C:\staging
if not errorlevel 0 exit /b 1

if not exist %CACHE% mkdir %CACHE%
if not errorlevel 0 exit /b 1

if exist %TARGET% (
  @rem Try to delete the contents of %TARGET% without deleting the directory itself.
  @rem This is imporatant since we want to be able to run download-build from the
  @rem target directory.

  del /f /q %TARGET%\*
  for /D %%f in ( %TARGET%\* ) do (
    rmdir /s /q %%f
    if not errorlevel 0 (
      echo Can't clean up %%f.
      exit /b 1
    )
  )
)

if not exist %TARGET% mkdir %TARGET%
if not errorlevel 0 exit /b 1

call isolateserver.cmd download --cache %CACHE% --target %TARGET% --isolated %1
