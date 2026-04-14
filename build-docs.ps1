Write-Output "Compiling PowerShell Get-Help pages..."
New-ExternalHelp -Path docs\system-help\powershell-platyps\* -OutputPath dist\ -Force