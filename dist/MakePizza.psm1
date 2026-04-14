# DOS
function mkpizza {
    [CmdletBinding()]
    param(
        [string]$Argument
    )
    & "$PSScriptRoot\bin\mkpizza.exe" $Argument
}

# POSIX
function makepizza {
    [CmdletBinding()]
    param(
        [string]$Argument
    )
    & "$PSScriptRoot\bin\makepizza.exe" $Argument
}

# PowerShell
function New-Pizza {
    [CmdletBinding()]
    param(
        [string]$Argument
    )
    & "$PSScriptRoot\bin\new-pizza.exe" $Argument
}
