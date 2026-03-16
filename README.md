<img width="128" height="128" alt="stg-2601-t-l@512" src="https://github.com/user-attachments/assets/5a7aec05-e42b-44dd-8dc1-ecfb80e725b1" alt="(>_ )" />

# `makepizza` 🍕

> [!NOTE]
>
> `makepizza` is currently under development, and yes, before the program is here yet, we are ramping up the command-line specification here. (>_ );

`makepizza` is a dummy command-line program that educates people how to use the Command-Line Interface (CLI).

It is offerred in three major variants:

+ `mkpizza` for DOS or legacy Windows Batch command-line program conventions
  - Example: `mkpizza.exe /n="Hello Pizza (^$- )"`
+ `makepizza` for general, UNIX/POSIX command-line program conventions
  - Example: `makepizza --name="Hello Pizza (\$- )"`
+ `Make-Pizza` for PowerShell cmdlet conventions
  - Example: ``Make-Pizza -Name "Hello Pizza (`$- )"``"

`makepizza` is actually written and compiled in Go (instead of usual Bash/Batch/PowerShell scripts), so you can enjoy working with different command-line conventions without having to install their runtimes.

That means you can use `mkpizza` on your Mac without having to install Windows, `makepizza` on Windows without using Bash or WSL, `Make-Pizza` on Linux without PowerShell/.NET for Linux, and so on.

## Usage

```sh
mkpizza {DOS-style arguments} [Contents <string>]

makepizza {POSIX-style arguments} [Contents <string>]

Make-Pizza {PowerShell-style arguments} [Contents <string>]
```

The `Contents` of the pizza can be filled by providing the remaining string (as interpreted _outside_ of the context of the given arguments/flags), or by piping this program by the outputs of the previous program, like `Get-Content foo.txt | Make-Pizza` (PowerShell).

> [!NOTE]
>
> **Notice for case-sensitive OS or filesystems**
> 
> For convenience, the PowerShell `Make-Pizza` binary is actually stored under the lowercase `make-pizza` filename. This is to follow PowerShell conventions that commands (technically either a "cmdlet" or a "function") and its arguments must be case-insensitive.

## Command-line arguments

| DOS | POSIX | PowerShell | Description | Example Value |
|---|---|---|---|---|
| `[/n=<string>]` | `[-n <string>]`<br/>`[--name <string>]` | `[-n <string>]`<br/>`[-Name <string>]` | Give your pizza some name! (Optional) | `-n "MozzaTuna(TM)"` |
| `[/d=<int>]` | `[-d <int>]`<br/>`[--duration <int>]` | `[-d <int>]`<br/>`[-Duration <int>]` | Delays the execution of this program, which should have been done in less than 1 second. Might be useful to simulate multithreading. **Must be a positive integer.** Default is `0` (no delay). | `1200` (1200 seconds) |
| `[/p:ON \| /p:OFF]` | `[-p \| -P]`<br/>`[--pineapple \| --no-pineapple]` | `[-p <bool>]`<br/>`[-Pineapple <bool>]` | Should we add pineapple to the pizza? Default is `true` (yes). | `/p:OFF` (DOS)<br/>`-P` (POSIX: must use the uppercase P to turn off)<br/>`-p $false` (PowerShell) |
| `[/s:ON \| /s:OFF]` | `[-s \| -S]`<br/>`[--check-superuser \| --no-check-superuser]` | `[-s <bool>]`<br/>`[-CheckSuperuser <bool>]` | Should the pizza be made by a `root`, `sudoer`, or an `Administrator`? Default is `false` (no). | `/p:ON` (DOS)<br/>`-p` (POSIX: must use the lowercase P to turn on)<br/>`-p $true` (PowerShell) |
| `[/e=<int>]` | `[-e <int>]`<br/>`[--exit-code <int>]` | `[-e <int>]`<br/>`[-ExitCode <int>]` | Emulate the exit code of this program. **Must be a positive integer.** Default is `0` (OK). | `137` (POSIX `SIGKILL`) |
| `[/?]` | `[-h]`<br/>`[--help]` | `[-h]`<br/>`[-Help]` | Display help. Note that help is currently **not** accessible via shells' native `man`, `help`, or `Get-Help` commands. | `mkpizza /?` |
| `[/v]` | `[-v]`<br/>`[--version]` | `[-v]`<br/>`[-Version]` | Display program version info for each of `mkpizza`, `makepizza`, and `Make-Pizza`. | `mkpizza /v` |

## Limitations

Even though the DOS, POSIX, and PowerShell variants of this programs are made available for many operating systems and shell environments, some advanced command-line features might not be consistent between them.

For example, you want to call PowerShell-based `Make-Pizza -Pineapple $false -Name "Great Pizza"` in Bash. But in Bash, `$false` will instead recall the value of the shell session variable `false`.

Since `false` is clearly unset, Bash will return `$false` as null (or no value), and `Make-Pizza` will instead receive your command as `Make-Pizza -Pineapple -Name "Great Pizza"`, which means the pineapple will still be added to your pizza. Oh, no!

> To solve this issue in POSIX-based shells: use escape characters `\`, so the command becomes `Make-Pizza -Pineapple \$false -Name "Great Pizza"`

`mkpizza`, `makepizza`, and `Make-Pizza` will try to interpret the remaining command-line input as literal strings, so they won't be able to parse advanced shell data structures like lists/arrays, and PowerShell's `HashTable`s. Many shell environments allow you to execute shell subcommands and carry their results to the final `makepizza` command, but the uniqueness or differences between these shells are out of the scope of this software.

## Frequently-asked questions

### Why not name this thing `make-a-sandwich` so I can `sudo make-a-sandwich`?

*What? Make it yourself.*

### Why doesn't this program generate some ASCII pizza art (like in `cowsay`)?

Adding some ASCII pizza art is indeed a requested feature, but it might confuse people who learn to `makepizza` under OS-level multithreading, that the pizza artworks became another piece of art within the multiverse.

### Are the arguments/flags case-sensitive?

Flags are only case-sensitive for the POSIX program (`makepizza`).

### Can POSIX learners use group flags to `makepizza`?

Yes, and only in the POSIX-style program (`makepizza`), such as `sudo makepizza -sP` or `sudo makepizza -Ps`.

### Is the program help page accessible through `man` or PowerShell's `Get-Help`?

Not yet. We want to make `makepizza` stable enough to start writing proper documentation for each of the commands.
