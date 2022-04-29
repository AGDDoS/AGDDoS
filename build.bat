@echo off
cd /d %~dp0
echo.
echo ----------------------------What do you want to do?----------------------------
echo.
echo 1. Build With GUI
echo 2. Build Without GUI
echo.
echo.
echo What do you want to do? Enter the number.
set /p num=
if "%num%"=="1" (
go build -o AGDDoS.exe .\AGDDoS.go
)

if "%num%"=="2" (
go build -ldflags "-H=windowsgui" -o AGDDoS.exe .\AGDDoS.go
)
