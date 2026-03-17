#!/bin/sh

case $(uname | tr '[:upper:]' '[:lower:]') in
    msys*|cygwin*|mingw*|nt|win*)
        echo "Compiling DOS (\`mkpizza\`) to dist\\mkpizza.exe..."
        go build -o dist/mkpizza.exe ./dos
        echo "Compiling POSIX (\`makepizza\`) to dist\\makepizza.exe..."
        go build -o dist/makepizza.exe ./posix
        echo "Compiling PowerShell (\`New-Pizza\`) to dist\\new-pizza.exe..."
        go build -o dist/new-pizza.exe ./powershell
        ;;
    *)
        echo "Compiling DOS (\`mkpizza\`) to dist/mkpizza..."
        go build -o dist/mkpizza ./dos
        echo "Compiling POSIX (\`makepizza\`) to dist/makepizza..."
        go build -o dist/makepizza ./posix
        echo "Compiling PowerShell (\`New-Pizza\`) to dist/new-pizza..."
        go build -o dist/new-pizza ./powershell
        ;;
esac
