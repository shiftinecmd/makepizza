#!/bin/bash
find docs/system-help/man-pandoc -name "*.md" -type f -exec sh -c 'pandoc "$1" -s -t man -o "dist/man1/${1%.md}.1" && gzip -f "dist/man1/${1%.md}.1"' _ {} \;
