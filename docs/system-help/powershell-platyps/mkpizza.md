---
document type: cmdlet
external help file: en-US\mkpizza-Help.xml
HelpUri: https://makepizza.shiftine.sh
Locale: en-US
Module Name: MakePizza.mkpizza
schema: 2.0.0
title: mkpizza
---

# mkpizza

## SYNOPSIS
Dummy command-line tool to teach how to make pizza (DOS convention version)

## SYNTAX

### DefaultParameterSet

```
**mkpizza** <string> [/N <string>] [/D <int>] [/P|/P:YES|/-P|/P:NO] [/S|/S:YES|/-S|/S:NO] [/E <int>]
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
mkpizza /N=Johny "Johny, the Pizza"
```

This command creates a new pizza named Johny, using the `/N` parameter.

### Example 2: You forgor to remove the pineapple
```powershell
mkpizza /N=Johny /P:OFF "Johny, the Pizza"
```

Apparently, this program has the algorithmic affinity to put pineapples on pizza.

Oh, yuck! Make another one without it.

### Example 3: Cook it a little bit more
```powershell
mkpizza /N=Johny /P:OFF /D:67 "Johny, the Pizza"
```

Nice! You'd think cooking the pizza longer could mean better, crispier piece of bread!

Try to cook for 67 seconds while you put the fries in the bag.

### Example 4: root vibes only
```powershell
mkpizza /N=Johny /P:OFF /D:67 /S:ON /E:165 "Johny, the Pizza"
```

You'd think cooking a pizza for 67 seconds is insecure by default. Skibidi people could have tried to poison your pizza.

Therefore, to maintain supply-chain security, only allow `root`, `sudoer`, or Administrator users to make the pizza. And just to be secure, please return an exit code of `165` instead of `0`.

## PARAMETERS

### /N
Give your pizza some name! (optional)

```yaml
Type: System.String
Default value: ''
Accept wildcard characters: false
Parameter Sets: (All)
Position: Named
Required: false
Accept pipeline input: false
```

### /D
Delays the execution of this program, which should have been done in less than 1 second (default: 0)

```yaml
Type: System.Int32
Default value: 0
Accept wildcard characters: false
Parameter Sets: (All)
Position: Named
Required: false
Accept pipeline input: false
```

### /P, /P:ON
Add pineapple to the pizza (default)

```yaml
Type: ActionPreference
Default value: ''
Accept wildcard characters: false
Parameter Sets: (All)
Type: SwitchParameter
Default value: False
Required: false
Accept pipeline input: false
```

### /-P, /P:OFF
DON'T add pineapple to the pizza

```yaml
Type: ActionPreference
Default value: ''
Accept wildcard characters: false
Parameter Sets: (All)
Type: SwitchParameter
Default value: False
Required: false
Accept pipeline input: false
```

### /S, /S:ON
Ensure that only `root`, `sudoer`, or Administrators can make this pizza

```yaml
Type: ActionPreference
Default value: ''
Accept wildcard characters: false
Parameter Sets: (All)
Type: SwitchParameter
Default value: False
Required: false
Accept pipeline input: false
```

### /-S, /S:OFF
Anyone can make this pizza

```yaml
Type: ActionPreference
Default value: ''
Accept wildcard characters: false
Parameter Sets: (All)
Type: SwitchParameter
Default value: False
Required: false
Accept pipeline input: false
```

### /E
Change the exit code of this program, after the pizza is made (default: 0)

Change the exit code of this program, after the pizza is made (default: 0)

```yaml
Type: System.Int32
Default value: 0
Accept wildcard characters: false
Parameter Sets: (All)
Position: Named
Required: false
Accept pipeline input: false
```

### /?
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

### /V
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
You are required to use either colon (`:`) or equal (`=`) sign to assign the value into each parameter. For example, `/D=10` or `/D:10`.

Even though the DOS, POSIX, and PowerShell variants of this programs are made available for many operating systems and shell environments, some advanced command-line features might not be consistent between them.

For example, you want to call PowerShell-based `New-Pizza -Pineapple $false -Name "Great Pizza"` in Bash. But in Bash, `$false` will instead recall the value of the shell session variable `false`.

Since `false` is clearly unset, Bash will return `$false` as null (or no value), and `New-Pizza` will instead receive your command as `New-Pizza -Pineapple -Name "Great Pizza"`, which means the pineapple will still be added to your pizza. Oh, no!

> To solve this issue in POSIX-based shells: use escape characters `\`, so the command becomes `New-Pizza -Pineapple \$false -Name "Great Pizza"`

`mkpizza`, `makepizza`, and `New-Pizza` will try to interpret the remaining command-line input as literal strings, so they won't be able to parse advanced shell data structures like lists/arrays, and PowerShell's `HashTable`s. Many shell environments allow you to execute shell subcommands and carry their results to the final `makepizza` command, but the uniqueness or differences between these shells are out of the scope of this software.

## RELATED LINKS
[Source Code](https://github.com/shiftinecmd/makepizza)
