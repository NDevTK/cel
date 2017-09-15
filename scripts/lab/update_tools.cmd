@if "%1"=="post-update" goto postupdate

@echo Verifying directories ...
if not exist C:\tools mkdir C:\tools
@if not errorlevel 0 exit /b 1

if not exist C:\tools\chrome-auth-lab mkdir C:\tools\chrome-auth-lab
@if not errorlevel 0 exit /b 1

if not exist C:\keys mkdir C:\keys
@if not errorlevel 0 exit /b 1

if not exist C:\staging mkdir C:\staging
@if not errorlevel 0 exit /b 1

@echo Updating chrome-auth-lab ...
call gsutil -m rsync -d gs://chrome-auth-lab-staging/tools C:\Tools\chrome-auth-lab
@if not errorlevel 0 exit /b 1

call C:\tools\chrome-auth-lab\update_tools.cmd post-update
exit /b 0

:postupdate

@echo Updating keys ...
call gsutil -m rsync -d gs://chrome-auth-lab-staging/keys C:\keys
@if not errorlevel 0 exit /b 1

@echo Updating depot_tools ...
call gclient --version

@echo Updating Luci-Py ...
cd C:\tools\luci-py
call git pull origin master

@echo Done
