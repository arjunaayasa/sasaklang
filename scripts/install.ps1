# SasakLang Installer for Windows (PowerShell)
# Usage: irm https://raw.githubusercontent.com/arjunaayasa/sasaklang/main/scripts/install.ps1 | iex

$ErrorActionPreference = "Stop"

$Repo = "arjunaayasa/sasaklang"
$InstallDir = "$env:USERPROFILE\.sasaklang\bin"
$BinaryName = "sasaklang.exe"

# Logo
Write-Host @"
   _____                 __   __                   
  / ___/____ _________ _/ /__/ /   ____ _____  ____ _
  \__ \/ __  / ___/ __  / //_/ /   / __  / __ \/ __  /
 ___/ / /_/ (__  ) /_/ / ,< / /___/ /_/ / / / / /_/ / 
/____/\__,_/____/\__,_/_/|_/_____/\__,_/_/ /_/\__, /  
                                             /____/   
"@ -ForegroundColor Green

Write-Host "SasakLang Installer"
Write-Host "==================="
Write-Host ""

# Detect architecture
$Arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
Write-Host "Terdeteksi: windows-$Arch"

# Get latest release
Write-Host "Mengambil versi terbaru..." -ForegroundColor Yellow
try {
    $Release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest" -ErrorAction SilentlyContinue
    $LatestVersion = $Release.tag_name
} catch {
    $LatestVersion = "v1.0.0"
}

if (-not $LatestVersion) {
    $LatestVersion = "v1.0.0"
}

Write-Host "Versi: $LatestVersion"

# Download URL
$DownloadUrl = "https://github.com/$Repo/releases/download/$LatestVersion/sasaklang_windows_$Arch.exe"

# Create install directory
Write-Host "Membuat direktori instalasi..." -ForegroundColor Yellow
New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null

# Download binary
Write-Host "Mengunduh sasaklang..." -ForegroundColor Yellow
$BinaryPath = Join-Path $InstallDir $BinaryName

try {
    Invoke-WebRequest -Uri $DownloadUrl -OutFile $BinaryPath -ErrorAction Stop
} catch {
    Write-Host "Download dari release gagal, mencoba build dari source..." -ForegroundColor Yellow
    
    # Check if Go is installed
    $GoInstalled = Get-Command go -ErrorAction SilentlyContinue
    if (-not $GoInstalled) {
        Write-Host "Error: Go tidak terinstall. Silakan install Go terlebih dahulu." -ForegroundColor Red
        Write-Host "Kunjungi: https://golang.org/dl/"
        exit 1
    }
    
    # Build from source
    $TempDir = New-TemporaryFile | ForEach-Object { Remove-Item $_; New-Item -ItemType Directory -Path $_ }
    
    Write-Host "Cloning repository..."
    git clone "https://github.com/$Repo.git" "$TempDir\sasaklang" 2>$null
    
    if (-not $?) {
        Write-Host "Error: Gagal clone repository" -ForegroundColor Red
        exit 1
    }
    
    Push-Location "$TempDir\sasaklang"
    Write-Host "Building..."
    go build -o $BinaryPath ./cmd/sasaklang
    Pop-Location
    
    Remove-Item -Recurse -Force $TempDir
}

# Add to PATH
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    $NewPath = "$InstallDir;$UserPath"
    [Environment]::SetEnvironmentVariable("Path", $NewPath, "User")
    Write-Host "PATH diperbarui" -ForegroundColor Green
}

# Update current session PATH
$env:Path = "$InstallDir;$env:Path"

Write-Host ""
Write-Host "SasakLang berhasil diinstall!" -ForegroundColor Green
Write-Host ""
Write-Host "Buka PowerShell baru, lalu coba:"
Write-Host "  sasaklang version" -ForegroundColor Yellow
Write-Host "  sasaklang" -ForegroundColor Yellow
Write-Host "  (untuk REPL)"
Write-Host ""
Write-Host "Dokumentasi: https://github.com/$Repo"
