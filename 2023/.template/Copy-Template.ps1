[CmdletBinding()]
param(
    [Parameter(
        Mandatory,
        Position=1)]
    [int]
    $Day
)

$dayText = $Day.ToString('00')

$sourcePath  = $PSScriptRoot
$sourcePy    = Join-Path $sourcePath 'day_nn.py'
$sourceTest  = Join-Path $sourcePath 'input.test.txt'
$sourceInput = Join-Path $sourcePath 'input.txt'
$destinationPath = Join-Path $sourcePath '..' $dayText
$destinationPy   = Join-Path $destinationPath "day_$dayText.py"

New-Item -Path $destinationPath -ItemType Directory | Out-Null
Copy-Item -Path $sourcePy    -Destination $destinationPy
Copy-Item -Path $sourceTest  -Destination $destinationPath
Copy-Item -Path $sourceInput -Destination $destinationPath

Push-Location $destinationPath
