---
document type: cmdlet
external help file: MakePizza-makepizza-help.xml
HelpUri: https://makepizza.shiftine.sh
Locale: en-US
Module Name: MakePizza
schema: 2.0.0
title: makepizza
---

# makepizza

## SYNOPSIS

Dummy command-line tool to teach how to make pizza (POSIX convention version)

## SYNTAX

### DefaultParameterSet

**makepizza** <string>

### DetailedParameterSet

**makepizza** <string> [--name <string>] [--duration <int>] [--pineapple|--no-pineapple] [--check-superuser|--no-check-superuser] [--exit-code <int>]

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

## EXAMPLE

### Example 1: Create Johny, the Pizza

This command creates a new pizza named Johny, using the `-Name` parameter.

```ps1
makepizza --name Johny "Johny, the Pizza"
```

### Example 2: You forgor to remove the pineapple

Apparently, this program has the algorithmic affinity to put pineapples on pizza.

Oh, yuck! Make another one without it.

```ps1
makepizza --name Johny --no-pineapple "Johny, the Pizza"
```

### Example 3: Cook it a little bit more

Nice! You'd think cooking the pizza longer could mean better, crispier piece of bread!

Try to cook for 67 seconds while you put the fries in the bag.

```ps1
makepizza --name Johny --no-pineapple --duration 67 "Johny, the Pizza"
```

### Example 4: root vibes only

You'd think cooking a pizza for 67 seconds is insecure by default. Skibidi people could have tried to poison your pizza.

Therefore, to maintain supply-chain security, only allow `root`, `sudoer`, or Administrator users to make the pizza. And just to be secure, please return an exit code of `165` instead of `0`.

```ps1
makepizza --name Johny --no-pineapple --duration 67 --check-superuser --exit-code 165 "Johny, the Pizza"
```

## PARAMETERS

### -n, --name

Give your pizza some name! (optional)

```yaml
Type: System.String
DefaultValue: ''
SupportsWildcards: false
Aliases: []
ParameterSets:
- Name: (All)
  Position: Named
  IsRequired: false
  ValueFromPipeline: false
  ValueFromPipelineByPropertyName: false
  ValueFromRemainingArguments: false
DontShow: false
```

### -d, --duration

Delays the execution of this program, which should have been done in less than 1 second (default: 0)

```yaml
Type: System.Int32
DefaultValue: 0
SupportsWildcards: false
Aliases: []
ParameterSets:
- Name: (All)
  Position: Named
  IsRequired: false
  ValueFromPipeline: false
  ValueFromPipelineByPropertyName: false
  ValueFromRemainingArguments: false
DontShow: false
```

### -p, --pineapple

Add pineapple to the pizza (default)

```yaml
Type: System.Boolean
DefaultValue: ''
SupportsWildcards: false
Aliases: []
ParameterSets:
- Name: (All)
  Type: SwitchParameter
  DefaultValue: False
  IsRequired: false
  ValueFromPipeline: false
  ValueFromPipelineByPropertyName: false
  ValueFromRemainingArguments: false
DontShow: false
AcceptedValues: []
```

### -P, --no-pineapple

DON'T add pineapple to the pizza

```yaml
Type: System.Boolean
DefaultValue: ''
SupportsWildcards: false
Aliases: []
ParameterSets:
- Name: (All)
  Type: SwitchParameter
  DefaultValue: False
  IsRequired: false
  ValueFromPipeline: false
  ValueFromPipelineByPropertyName: false
  ValueFromRemainingArguments: false
DontShow: false
AcceptedValues: []
```

### -s, --check-superuser

Ensure that only root or sudoer can make this pizza

```yaml
Type: System.Boolean
DefaultValue: ''
SupportsWildcards: false
Aliases: []
ParameterSets:
- Name: (All)
  Type: SwitchParameter
  DefaultValue: False
  IsRequired: false
  ValueFromPipeline: false
  ValueFromPipelineByPropertyName: false
  ValueFromRemainingArguments: false
DontShow: false
AcceptedValues: []
```

### -S, --no-check-superuser

Anyone can make this pizza

```yaml
Type: System.Boolean
DefaultValue: ''
SupportsWildcards: false
Aliases: []
ParameterSets:
- Name: (All)
  Type: SwitchParameter
  DefaultValue: False
  IsRequired: false
  ValueFromPipeline: false
  ValueFromPipelineByPropertyName: false
  ValueFromRemainingArguments: false
DontShow: false
AcceptedValues: []
```

### -e, --exit-code

Change the exit code of this program, after the pizza is made (default: 0)

```yaml
Type: System.Int32
DefaultValue: 0
SupportsWildcards: false
Aliases: []
ParameterSets:
- Name: (All)
  Position: Named
  IsRequired: false
  ValueFromPipeline: false
  ValueFromPipelineByPropertyName: false
  ValueFromRemainingArguments: false
DontShow: false
```

### -?, -h, --help

Display help

```yaml
Type: System.Boolean
DefaultValue: ''
SupportsWildcards: false
Aliases: []
ParameterSets:
- Name: (All)
  Type: SwitchParameter
  DefaultValue: False
  IsRequired: false
  ValueFromPipeline: false
  ValueFromPipelineByPropertyName: false
  ValueFromRemainingArguments: false
DontShow: false
AcceptedValues: []
```

### -v, --version

Display version

```yaml
Type: System.Boolean
DefaultValue: ''
SupportsWildcards: false
Aliases: []
ParameterSets:
- Name: (All)
  Type: SwitchParameter
  DefaultValue: False
  IsRequired: false
  ValueFromPipeline: false
  ValueFromPipelineByPropertyName: false
  ValueFromRemainingArguments: false
DontShow: false
AcceptedValues: []
```

## INPUTS

### System.String

The content of the pizza. Will be printed on the output.

## OUTPUTS

### System.String

The resulting pizza that you've just made (>_ )!

## NOTES

### General Limitations

Even though the DOS, POSIX, and PowerShell variants of this programs are made available for many operating systems and shell environments, some advanced command-line features might not be consistent between them.

For example, you want to call PowerShell-based `New-Pizza -Pineapple $false -Name "Great Pizza"` in Bash. But in Bash, `$false` will instead recall the value of the shell session variable `false`.

Since `false` is clearly unset, Bash will return `$false` as null (or no value), and `New-Pizza` will instead receive your command as `New-Pizza -Pineapple -Name "Great Pizza"`, which means the pineapple will still be added to your pizza. Oh, no!

> To solve this issue in POSIX-based shells: use escape characters `\`, so the command becomes `New-Pizza -Pineapple \$false -Name "Great Pizza"`

`mkpizza`, `makepizza`, and `New-Pizza` will try to interpret the remaining command-line input as literal strings, so they won't be able to parse advanced shell data structures like lists/arrays, and PowerShell's `HashTable`s. Many shell environments allow you to execute shell subcommands and carry their results to the final `makepizza` command, but the uniqueness or differences between these shells are out of the scope of this software.

## RELATED LINKS

- [Source Code](https://github.com/shiftinecmd/makepizza)
