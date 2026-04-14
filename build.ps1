# Dynamically check whether this is Windows or not
# PowerShell versions less than 6 will be treated as Windows

if ($Env:GOOS -eq "windows") {
    Write-Output "Compiling DOS (``mkpizza``) to dist\mkpizza.exe..."
    go build -o dist\bin\mkpizza.exe .\dos
    Write-Output "Compiling POSIX (``makepizza``) to dist\makepizza.exe..."
    go build -o dist\bin\makepizza.exe .\posix
    Write-Output "Compiling PowerShell (``New-Pizza``) to dist\new-pizza.exe..."
    go build -o dist\bin\new-pizza.exe .\powershell
} else {
    Write-Output "Compiling DOS (``mkpizza``) to dist/mkpizza..."
    go build -o dist/bin/mkpizza ./dos
    Write-Output "Compiling POSIX (``makepizza``) to dist/makepizza..."
    go build -o dist/bin/makepizza ./posix
    Write-Output "Compiling PowerShell (``New-Pizza``) to dist/new-pizza..."
    go build -o dist/bin/new-pizza ./powershell
}
