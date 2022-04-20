@echo off
::切换到当前文件所在目录
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
::构建exe文件
go build -o unitst000.exe .\AGDDoS.go
)

if "%num%"=="2" (
::构建exe文件
go build -ldflags "-H=windowsgui" -o unitst000.exe .\AGDDoS.go
)