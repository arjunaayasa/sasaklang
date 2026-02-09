$ErrorActionPreference = 'Stop'

# Detect Arch
if ($env:PROCESSOR_ARCHITECTURE -eq "AMD64") {
    $ARCH = "amd64"
} else {
    $ARCH = "386"
}

$ASSET_NAME = "sasaklang-windows-$ARCH.exe"
$INSTALL_DIR = "$env:USERPROFILE\.sasaklang\bin"
$BIN_PATH = "$INSTALL_DIR\sasaklang.exe"

Write-Host "Downloading SasakLang for Windows ($ARCH)..."

try {
    # Get latest release info
    $latestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/arjunaayasa/sasaklang/releases/latest"
    $downloadUrl = ($latestRelease.assets | Where-Object { $_.name -eq $ASSET_NAME }).browser_download_url

    if (-not $downloadUrl) {
        Write-Error "Could not find release asset for $ASSET_NAME"
        exit 1
    }

    # Install
    if (-not (Test-Path $INSTALL_DIR)) {
        New-Item -ItemType Directory -Force -Path $INSTALL_DIR | Out-Null
    }

    Invoke-WebRequest -Uri $downloadUrl -OutFile $BIN_PATH
    Write-Host "Installed to $BIN_PATH"

    # Add to PATH
    $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if (-not ($userPath -split ';' -contains $INSTALL_DIR)) {
        Write-Host "Adding to PATH..."
        [Environment]::SetEnvironmentVariable("Path", "$userPath;$INSTALL_DIR", "User")
        Write-Host "PATH updated. Please restart your terminal."
    } else {
        Write-Host "PATH already configured."
    }

    Write-Host "✅ SasakLang installed successfully!" -ForegroundColor Green
} catch {
    Write-Error "Failed to install: $_"
    exit 1
}
