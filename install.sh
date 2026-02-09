#!/bin/bash
set -e

# Detect OS and Arch
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

if [ "$ARCH" == "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH" == "aarch64" ]; then
    ARCH="arm64"
fi

ASSET_NAME="sasaklang-${OS}-${ARCH}"
INSTALL_DIR="$HOME/.sasaklang/bin"
BIN_NAME="sasaklang"

echo "Downloading SasakLang for ${OS}-${ARCH}..."

# Get latest release download URL
DOWNLOAD_URL=$(curl -s https://api.github.com/repos/arjunaayasa/sasaklang/releases/latest | grep "browser_download_url.*${ASSET_NAME}" | cut -d : -f 2,3 | tr -d \")

if [ -z "$DOWNLOAD_URL" ]; then
    echo "Error: Could not find release asset for ${ASSET_NAME}."
    echo "Please ensure a release exists with name format: sasaklang-${OS}-${ARCH}"
    exit 1
fi

# Download and Install
mkdir -p "$INSTALL_DIR"
curl -L -o "$INSTALL_DIR/$BIN_NAME" "$DOWNLOAD_URL"
chmod +x "$INSTALL_DIR/$BIN_NAME"

echo "Installed to $INSTALL_DIR/$BIN_NAME"

# Update PATH
SHELL_CONFIG=""
if [ "$SHELL" == */zsh ]; then
    SHELL_CONFIG="$HOME/.zshrc"
elif [ "$SHELL" == */bash ]; then
    SHELL_CONFIG="$HOME/.bashrc"
fi

if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "Adding to PATH..."
    echo "" >> "$SHELL_CONFIG"
    echo "# SasakLang" >> "$SHELL_CONFIG"
    echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$SHELL_CONFIG"
    echo "Updated $SHELL_CONFIG. Please restart your terminal or run: source $SHELL_CONFIG"
else
    echo "PATH already configured."
fi

echo "✅ SasakLang installed successfully!"
