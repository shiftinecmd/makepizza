@echo off
set CGO_ENABLED=1
set CC=clang
set CXX=clang++

for /f "tokens=*" %%i in ('go env GOHOSTOS') do set GOOS=%%i
for /f "tokens=*" %%i in ('go env GOHOSTARCH') do set GOARCH=%%i

echo Compiling DOS (`mkpizza`) to dist\mkpizza.exe...
go build -o dist\mkpizza.exe .\dos
echo Compiling POSIX (`makepizza`) to dist\makepizza.exe...
go build -o dist\makepizza.exe .\posix
echo Compiling PowerShell (`New-Pizza`) to dist\new-pizza.exe...
go build -o dist\new-pizza.exe .\powershell
