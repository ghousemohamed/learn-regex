#!/bin/bash

REPO="ghousemohamed/learn-regex"
BIN_NAME="learn-regex"
INSTALL_DIR="$HOME/.local/bin"
VERSION="0.1.0"

mkdir -p "$INSTALL_DIR"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$OS" == "linux" ]]; then
    if [[ "$ARCH" == "x86_64" ]]; then
        BINARY="learn-regex_Linux_x86_64.tar.gz"
        EXTRACT_DIR="learn-regex_Linux_x86_64"
    elif [[ "$ARCH" == "aarch64" ]]; then
        BINARY="learn-regex_Linux_arm64.tar.gz"
        EXTRACT_DIR="learn-regex_Linux_arm64"
    elif [[ "$ARCH" == "i386" || "$ARCH" == "i686" ]]; then
        BINARY="learn-regex_Linux_i386.tar.gz"
        EXTRACT_DIR="learn-regex_Linux_i386"
    else
        echo "Unsupported architecture: $ARCH"
        exit 1
    fi
elif [[ "$OS" == "darwin" ]]; then
    if [[ "$ARCH" == "x86_64" ]]; then
        BINARY="learn-regex_Darwin_x86_64.tar.gz"
        EXTRACT_DIR="learn-regex_Darwin_x86_64"
    elif [[ "$ARCH" == "arm64" ]]; then
        BINARY="learn-regex_Darwin_arm64.tar.gz"
        EXTRACT_DIR="learn-regex_Darwin_arm64"
    else
        echo "Unsupported architecture: $ARCH"
        exit 1
    fi
elif [[ "$OS" == "mingw32nt" || "$OS" == "cygwin" || "$OS" == "msys" || "$OS" == "windows_nt" ]]; then
    if [[ "$ARCH" == "x86_64" ]]; then
        BINARY="learn-regex_Windows_x86_64.zip"
        EXTRACT_DIR="learn-regex_Windows_x86_64"
    elif [[ "$ARCH" == "aarch64" ]]; then
        BINARY="learn-regex_Windows_arm64.zip"
        EXTRACT_DIR="learn-regex_Windows_arm64"
    elif [[ "$ARCH" == "i386" || "$ARCH" == "i686" ]]; then
        BINARY="learn-regex_Windows_i386.zip"
        EXTRACT_DIR="learn-regex_Windows_i386"
    else
        echo "Unsupported architecture: $ARCH"
        exit 1
    fi
else
    echo "Unsupported OS: $OS"
    exit 1
fi

DOWNLOAD_URL="https://github.com/$REPO/releases/download/$VERSION/$BINARY"
echo "Downloading $BINARY from $DOWNLOAD_URL..."

TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

if [[ "$OS" == "mingw32nt" || "$OS" == "cygwin" || "$OS" == "msys" || "$OS" == "windows_nt" ]]; then
    curl -L -o "$BINARY" "$DOWNLOAD_URL"
    unzip "$BINARY"
    mv "$EXTRACT_DIR/learn-regex.exe" "$INSTALL_DIR/$BIN_NAME.exe"
else
    curl -L -o "$BINARY" "$DOWNLOAD_URL"
    tar xzf "$BINARY"
    mv "$EXTRACT_DIR/learn-regex" "$INSTALL_DIR/$BIN_NAME"
    chmod +x "$INSTALL_DIR/$BIN_NAME"
fi

cd - > /dev/null
rm -rf "$TMP_DIR"

if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
    echo "Adding $INSTALL_DIR to PATH"
    echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$HOME/.bashrc"
    if [[ -f "$HOME/.zshrc" ]]; then
        echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$HOME/.zshrc"
    fi
    echo "Please restart your terminal or run 'source ~/.bashrc' to update your PATH."
fi

echo "Installation complete. Happy Learning!"