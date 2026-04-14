#!/bin/sh

case $(uname | tr '[:upper:]' '[:lower:]') in
    msys*|cygwin*|mingw*|nt|win*)
        echo "Compiling DOS (\`mkpizza\`) to dist\\mkpizza.exe..."
        go build -o dist/mkpizza.exe ./dos
        echo "Compiling POSIX (\`makepizza\`) to dist\\makepizza.exe..."
        go build -o dist/makepizza.exe ./posix
        echo "Compiling PowerShell (\`New-Pizza\`) to dist\\new-pizza.exe..."
        go build -o dist/new-pizza.exe ./powershell
        echo "Compiling manpages..."
        find docs/system-help/man-pandoc -name "*.md" -type f -exec sh -c 'pandoc "$1" -s -t man -o "dist/man1/${1%.md}.1" && gzip -f "dist/man1/${1%.md}.1"' _ {} \;
        ;;
    *)
        echo "Compiling DOS (\`mkpizza\`) to dist/mkpizza..."
        go build -o dist/bin/mkpizza ./dos
        echo "Compiling POSIX (\`makepizza\`) to dist/makepizza..."
        go build -o dist/bin/makepizza ./posix
        echo "Compiling PowerShell (\`New-Pizza\`) to dist/new-pizza..."
        go build -o dist/bin/new-pizza ./powershell
        echo "Compiling manpages..."
        find docs/system-help/man-pandoc -name "*.md" -type f -exec sh -c 'pandoc "$1" -s -t man -o "dist/man1/${1%.md}.1" && gzip -f "dist/man1/${1%.md}.1"' _ {} \;
        ;;
esac
