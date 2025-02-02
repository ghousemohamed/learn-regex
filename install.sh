#!/bin/bash

green='\033[0;32m'
red='\033[0;31m'
yellow='\033[0;33m'
cyan='\033[0;36m'
bright_green='\033[1;32m'
bright_yellow='\033[1;33m'
bright_cyan='\033[1;36m'
nc='\033[0m' # No Color

REPO="ghousemohamed/learn-regex"
BIN_NAME="learn-regex"
INSTALL_DIR="$HOME/.local/bin"
VERSION="v0.2.0"

mkdir -p "$INSTALL_DIR"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
    ARCH="amd64"
elif [[ "$ARCH" == "aarch64" || "$ARCH" == arm* ]]; then
    ARCH="arm64"
else
    echo -e "${red}❌ Unsupported architecture: ${ARCH}${nc}"
    exit 1
fi

case "$OS" in
    linux)   BINARY="$BIN_NAME-linux-$ARCH" ;;
    darwin)  BINARY="$BIN_NAME-darwin-$ARCH" ;;
    msys*|mingw*|cygwin*) BINARY="$BIN_NAME-windows-amd64.exe" ;;
    *)
        echo -e "${red}❌ Unsupported OS: ${OS}${nc}"
        exit 1
        ;;
esac

# Download the binary
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$VERSION/$BINARY"
echo -e "${bright_yellow}Downloading ${cyan}$BIN_NAME $VERSION for $OS ($ARCH)...${nc}"
curl -sL -o "$INSTALL_DIR/$BIN_NAME" "$DOWNLOAD_URL"

# Make executable (except on Windows)
if [[ "$OS" != "msys" && "$OS" != "mingw" && "$OS" != "cygwin" ]]; then
    chmod +x "$INSTALL_DIR/$BIN_NAME"
fi

# Add to PATH if not already present
if ! [[ ":$PATH:" == *":$INSTALL_DIR:"* ]]; then
    echo -e "${bright_yellow}Adding $INSTALL_DIR to PATH...${nc}"
    case $SHELL in
        */bash)
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
            echo -e "${bright_cyan}PATH added to ~/.bashrc${nc}" ;;
        */zsh)
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
            echo -e "${bright_cyan}PATH added to ~/.zshrc${nc}" ;;
        */fish)
            echo 'fish_add_path "$HOME/.local/bin"' >> ~/.config/fish/config.fish
            echo -e "${bright_cyan}PATH added to ~/.config/fish/config.fish${nc}" ;;
        */ksh)
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.kshrc
            echo -e "${bright_cyan}PATH added to ~/.kshrc${nc}" ;;
        */xonsh)
            echo '$PATH.prepend("$HOME/.local/bin")' >> ~/.xonshrc
            echo -e "${bright_cyan}PATH added to ~/.xonshrc${nc}" ;;
        */csh|*/tcsh)
            echo 'setenv PATH "$HOME/.local/bin:$PATH"' >> ~/.cshrc
            echo -e "${bright_cyan}PATH added to ~/.cshrc${nc}" ;;
        *)
            echo -e "${red}⚠ Unsupported shell: $SHELL. Please manually add ${cyan}$INSTALL_DIR${red} to your PATH.${nc}" ;;
    esac
    echo -e "${bright_cyan}Please restart your terminal or source your shell config file.${nc}"
fi

echo -e "🎉 ${bright_green}Installation complete! You can now run '${BIN_NAME}' from anywhere.${nc}"

