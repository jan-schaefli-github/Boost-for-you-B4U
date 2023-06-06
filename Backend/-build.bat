@echo off

REM Create "builds" directory if it doesn't exist
if not exist builds (
    mkdir builds
)

REM Clean previous builds
del /q builds\* 2>nul

REM Build for Linux
set GOOS=linux
set GOARCH=amd64
go build -o builds/b4u_backend_linux

REM Build for Windows
set GOOS=windows
set GOARCH=amd64
go build -o builds/b4u_backend_windows.exe

REM Reset environment variables
set GOOS=
set GOARCH=

echo Build completed!