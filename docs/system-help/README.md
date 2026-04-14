This directory contains the official help files for `makepizza`, `mkpizza`, and `New-Pizza` in the following formats:

| Directory | Input Format | Output Format | Output Target |
|---|---|---|---|
| fdos-html | `.HTM` | `.HTM`, `.TXT` | DOS-style help pages formatted using the FreeDOS `HTMLHELP` format. Also exported into `.TXT` with CRLF line endings. |
| man-pandoc | `.md` | `.tr` | Markdown files specially formatted for `pandoc` conversion to `tr` manpage files (for `man`). |
| powershell-platyps | `.md` | `.xml` (MAML) | Markdown files specially formatted for PlatyPS PowerShell utility, to generate MAML files for `Get-Help`. |
