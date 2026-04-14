---
title: new-pizza
section: 1
header: makepizza manual
footer:
date: March 26, 2026
---

# NAME

New-Pizza - dummy command-line tool to teach how to make pizza (PowerShell convention version)

# SYNOPSIS

**New-Pizza** [OPTION]... PIZZA_CONTENTS

# DESCRIPTION

`makepizza` is a dummy command-line program that educates people how to use the Command-Line Interface (CLI).

It is offerred in three major variants:

+ `mkpizza` for DOS, CP/M, or legacy Windows Batch command-line program conventions
  - Example: `mkpizza.exe /n="Hello Pizza (^$- )"`
+ `makepizza` for general, UNIX/POSIX command-line program conventions
  - Example: `makepizza --name="Hello Pizza (\$- )"`
+ `New-Pizza` for PowerShell cmdlet conventions
  - Example: ``New-Pizza -Name "Hello Pizza (`$- )"``"

`makepizza` is actually written and compiled in Go (instead of usual Bash/Batch/PowerShell scripts), so you can enjoy working with different command-line conventions without having to install their runtimes.

That means you can use `mkpizza` on your Mac without having to install Windows, `makepizza` on Windows without using Bash or WSL, `New-Pizza` on Linux without PowerShell/.NET for Linux, and so on.

# OPTIONS

`-n, -Name NAME`

:   Give your pizza some name! (optional)

`-d, -Duration DURATION`

:   Delays the execution of this program, which should have been done in less than 1 second (default: 0)

`-p [true|false], -Pineapple [true|false]`

:   Add pineapple to the pizza (default: true)

`-s [true|false], -CheckSuperuser [true|false]`

:   Ensure that only root or sudoer can make this pizza (default: false)

`-e, -ExitCode`

:   Change the exit code of this program, after the pizza is made (default: 0)

`-?, -h, -Help`

:   Display help

`-v, -Version`

:   Display version

# EXAMPLES

This command creates a new pizza named Johny, using the `-Name` parameter.

`New-Pizza -Name Johny "Johny, the Pizza"`

Apparently, this program has the algorithmic affinity to put pineapples on pizza.

Oh, yuck! Make another one without it.

`New-Pizza -Name Johny -Pineapple false "Johny, the Pizza"`

Nice! You'd think cooking the pizza longer could mean better, crispier piece of bread!

Try to cook for 67 seconds while you put the fries in the bag.

`New-Pizza -Name Johny -Pineapple false -Duration 67 "Johny, the Pizza"`

You'd think cooking a pizza for 67 seconds is insecure by default. Skibidi people could have tried to poison your pizza.

Therefore, to maintain supply-chain security, only allow `root`, `sudoer`, or Administrator users to make the pizza. And just to be secure, please return an exit code of `165` instead of `0`.

`New-Pizza -Name Johny -Pineapple false -Duration 67 -CheckSuperuser -ExitCode 165 "Johny, the Pizza"`

# NOTES

This command also accepts PowerShell-style Boolean values `$true` and `$false`, but the dollar sign may need to be escaped in some shells (like `bash`) to avoid confusing it with shell variables.

The `PIZZA_CONTENTS` argument can also be explicitly set using `-Contents PIZZA_CONTENTS`, to follow common PowerShell cmdlet conventions. If multiple parts of pizza contents are defined, the program only selects the first string detected after the first invocation of `-Contents`.

# LIMITATIONS

Even though the DOS, POSIX, and PowerShell variants of this programs are made available for many operating systems and shell environments, some advanced command-line features might not be consistent between them.

For example, you want to call PowerShell-based `New-Pizza -Pineapple $false -Name "Great Pizza"` in Bash. But in Bash, `$false` will instead recall the value of the shell session variable `false`.

Since `false` is clearly unset, Bash will return `$false` as null (or no value), and `New-Pizza` will instead receive your command as `New-Pizza -Pineapple -Name "Great Pizza"`, which means the pineapple will still be added to your pizza. Oh, no!

> To solve this issue in POSIX-based shells: use escape characters `\`, so the command becomes `New-Pizza -Pineapple \$false -Name "Great Pizza"`

`mkpizza`, `makepizza`, and `New-Pizza` will try to interpret the remaining command-line input as literal strings, so they won't be able to parse advanced shell data structures like lists/arrays, and PowerShell's `HashTable`s. Many shell environments allow you to execute shell subcommands and carry their results to the final `makepizza` command, but the uniqueness or differences between these shells are out of the scope of this software.
