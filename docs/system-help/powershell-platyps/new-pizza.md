---
document type: cmdlet
external help file: en-US\New-Pizza-Help.xml
HelpUri: https://makepizza.shiftine.sh
Locale: en-US
Module Name: MakePizza.New-Pizza
schema: 2.0.0
title: New-Pizza
---

# New-Pizza

## SYNOPSIS
Dummy command-line tool to teach how to make pizza (PowerShell convention version)

## SYNTAX

```
**New-Pizza** [-Contents] <string> [-Name <string>] [-Duration <int>] [-Pineapple <bool>] [-CheckSuperuser <bool>] [-ExitCode <int>]
```

## DESCRIPTION
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

## EXAMPLES

### Example 1: Create Johny, the Pizza
```powershell
New-Pizza -Name Johny "Johny, the Pizza"
```

This command creates a new pizza named Johny, using the `-Name` parameter.

### Example 2: You forgor to remove the pineapple
```powershell
New-Pizza -Name Johny -Pineapple $false "Johny, the Pizza"
```

Apparently, this program has the algorithmic affinity to put pineapples on pizza.

Oh, yuck! Make another one without it.

### Example 3: Cook it a little bit more
```powershell
New-Pizza -Name Johny -Pineapple $false -Duration 67 "Johny, the Pizza"
```

Nice! You'd think cooking the pizza longer could mean better, crispier piece of bread!

Try to cook for 67 seconds while you put the fries in the bag.

### Example 4: root vibes only
```powershell
New-Pizza -Name Johny -Pineapple $false -Duration 67 -CheckSuperuser -ExitCode 165 "Johny, the Pizza"
```

You'd think cooking a pizza for 67 seconds is insecure by default. Skibidi people could have tried to poison your pizza.

Therefore, to maintain supply-chain security, only allow `root`, `sudoer`, or Administrator users to make the pizza. And just to be secure, please return an exit code of `165` instead of `0`.

## PARAMETERS

### -Name
Give your pizza some name! (optional)

```yaml
Type: System.String
Default value: ''
Accept wildcard characters: false
Aliases: n
Parameter Sets: (All)
Position: Named
Required: false
Accept pipeline input: false
```

### -Contents
Add some content to your pizza!

```yaml
Type: System.String
Default value: ''
Accept wildcard characters: false
Parameter Sets: (All)
Position: Named
Required: true
Accept pipeline input: true
```

### -Duration
Delays the execution of this program, which should have been done in less than 1 second (default: 0)

```yaml
Type: System.Int32
Default value: 0
Accept wildcard characters: false
Aliases: d
Parameter Sets: (All)
Position: Named
Required: false
Accept pipeline input: false
```

### -Pineapple
Add pineapple to the pizza (default: true)

```yaml
Type: System.Boolean
Default value: ''
Accept wildcard characters: false
Aliases: p
Parameter Sets: (All)
Type: SwitchParameter
Default value: False
Required: false
Accept pipeline input: false
```

### -CheckSuperuser
Ensure that only `root`, `sudoer`, or Administrators can make this pizza (default: false)

```yaml
Type: System.Boolean
Default value: ''
Accept wildcard characters: false
Aliases: s
Parameter Sets: (All)
Type: SwitchParameter
Default value: False
Required: false
Accept pipeline input: false
```

### -ExitCode

Change the exit code of this program, after the pizza is made (default: 0)

```yaml
Type: System.Int32
Default value: 0
Accept wildcard characters: false
Aliases: e
Parameter Sets: (All)
Position: Named
Required: false
Accept pipeline input: false
```

### -Help
Display help

```yaml
Type: ActionPreference
Default value: ''
Accept wildcard characters: false
Aliases: ["?", "h"]
Parameter Sets: (All)
Type: SwitchParameter
Default value: False
Required: false
Accept pipeline input: false
```

### -Version

Display version

```yaml
Type: ActionPreference
Default value: ''
Accept wildcard characters: false
Aliases: v
Parameter Sets: (All)
Type: SwitchParameter
Default value: False
Required: false
Accept pipeline input: false
```

## INPUTS

### System.String
The content of the pizza. Will be printed on the output.

## OUTPUTS

### System.String
The resulting pizza that you've just made (>_ )!

## NOTES
The `PIZZA_CONTENTS` argument can also be explicitly set using `-Contents PIZZA_CONTENTS`, to follow common PowerShell cmdlet conventions. If multiple parts of pizza contents are defined, the program only selects the first string detected after the first invocation of `-Contents`.

Even though the DOS, POSIX, and PowerShell variants of this programs are made available for many operating systems and shell environments, some advanced command-line features might not be consistent between them.

For example, you want to call PowerShell-based `New-Pizza -Pineapple $false -Name "Great Pizza"` in Bash. But in Bash, `$false` will instead recall the value of the shell session variable `false`.

Since `false` is clearly unset, Bash will return `$false` as null (or no value), and `New-Pizza` will instead receive your command as `New-Pizza -Pineapple -Name "Great Pizza"`, which means the pineapple will still be added to your pizza. Oh, no!

> To solve this issue in POSIX-based shells: use escape characters `\`, so the command becomes `New-Pizza -Pineapple \$false -Name "Great Pizza"`

`mkpizza`, `makepizza`, and `New-Pizza` will try to interpret the remaining command-line input as literal strings, so they won't be able to parse advanced shell data structures like lists/arrays, and PowerShell's `HashTable`s. Many shell environments allow you to execute shell subcommands and carry their results to the final `makepizza` command, but the uniqueness or differences between these shells are out of the scope of this software.

## RELATED LINKS
[Source Code](https://github.com/shiftinecmd/makepizza)
