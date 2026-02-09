#!/bin/bash
# SasakLang Installer for Linux/macOS
# Usage: curl -fsSL https://raw.githubusercontent.com/arjunaayasa/sasaklang/main/scripts/install.sh | bash

set -e

REPO="arjunaayasa/sasaklang"
INSTALL_DIR="$HOME/.sasaklang/bin"
BINARY_NAME="sasaklang"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}"
echo "   _____                 __   __                   "
echo "  / ___/____ _________ _/ /__/ /   ____ _____  ____ _"
echo "  \\__ \\/ __ \`/ ___/ __ \`/ //_/ /   / __ \`/ __ \\/ __ \`/"
echo " ___/ / /_/ (__  ) /_/ / ,< / /___/ /_/ / / / / /_/ / "
echo "/____/\\__,_/____/\\__,_/_/|_/_____/\\__,_/_/ /_/\\__, /  "
echo "                                             /____/   "
echo -e "${NC}"
echo "SasakLang Installer"
echo "==================="
echo ""

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}Error: Arsitektur $ARCH tidak didukung${NC}"
        exit 1
        ;;
esac

case "$OS" in
    linux)
        PLATFORM="linux"
        ;;
    darwin)
        PLATFORM="darwin"
        ;;
    *)
        echo -e "${RED}Error: OS $OS tidak didukung${NC}"
        exit 1
        ;;
esac

echo "Terdeteksi: $PLATFORM-$ARCH"

# Get latest release version
echo -e "${YELLOW}Mengambil versi terbaru...${NC}"
LATEST_VERSION=$(curl -sL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/' || echo "v1.0.0")

if [ -z "$LATEST_VERSION" ]; then
    LATEST_VERSION="v1.0.0"
fi

echo "Versi: $LATEST_VERSION"

# Download URL
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_VERSION/${BINARY_NAME}_${PLATFORM}_${ARCH}"

# Create install directory
echo -e "${YELLOW}Membuat direktori instalasi...${NC}"
mkdir -p "$INSTALL_DIR"

# Download binary
echo -e "${YELLOW}Mengunduh $BINARY_NAME...${NC}"
if ! curl -fsSL "$DOWNLOAD_URL" -o "$INSTALL_DIR/$BINARY_NAME" 2>/dev/null; then
    echo -e "${YELLOW}Download dari release gagal, mencoba build dari source...${NC}"
    
    # Check if Go is installed
    if ! command -v go &> /dev/null; then
        echo -e "${RED}Error: Go tidak terinstall. Silakan install Go terlebih dahulu.${NC}"
        echo "Kunjungi: https://golang.org/dl/"
        exit 1
    fi
    
    # Build from source
    TEMP_DIR=$(mktemp -d)
    echo "Cloning repository..."
    git clone "https://github.com/$REPO.git" "$TEMP_DIR/sasaklang" 2>/dev/null || {
        echo -e "${RED}Error: Gagal clone repository${NC}"
        exit 1
    }
    
    cd "$TEMP_DIR/sasaklang"
    echo "Building..."
    go build -o "$INSTALL_DIR/$BINARY_NAME" ./cmd/sasaklang
    cd -
    rm -rf "$TEMP_DIR"
fi

# Make executable
chmod +x "$INSTALL_DIR/$BINARY_NAME"

# Add to PATH
SHELL_NAME=$(basename "$SHELL")
PROFILE_FILE=""

case "$SHELL_NAME" in
    bash)
        PROFILE_FILE="$HOME/.bashrc"
        ;;
    zsh)
        PROFILE_FILE="$HOME/.zshrc"
        ;;
    fish)
        PROFILE_FILE="$HOME/.config/fish/config.fish"
        ;;
    *)
        PROFILE_FILE="$HOME/.profile"
        ;;
esac

PATH_EXPORT="export PATH=\"\$HOME/.sasaklang/bin:\$PATH\""

if ! grep -q ".sasaklang/bin" "$PROFILE_FILE" 2>/dev/null; then
    echo "" >> "$PROFILE_FILE"
    echo "# SasakLang" >> "$PROFILE_FILE"
    echo "$PATH_EXPORT" >> "$PROFILE_FILE"
    echo -e "${GREEN}PATH ditambahkan ke $PROFILE_FILE${NC}"
fi

echo ""
echo -e "${GREEN}✓ SasakLang berhasil diinstall!${NC}"
echo ""
echo "Untuk mulai menggunakan, jalankan:"
echo -e "  ${YELLOW}source $PROFILE_FILE${NC}"
echo ""
echo "Atau buka terminal baru, lalu coba:"
echo -e "  ${YELLOW}sasaklang version${NC}"
echo -e "  ${YELLOW}sasaklang${NC}  (untuk REPL)"
echo ""
echo "Dokumentasi: https://github.com/$REPO"
